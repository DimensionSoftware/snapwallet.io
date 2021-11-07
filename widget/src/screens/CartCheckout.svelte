<script lang="ts">
  import { onMount } from 'svelte'
  import { fly } from 'svelte/transition'
  import { push } from 'svelte-spa-router'
  import { faClock } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'

  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Button from '../components/Button.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { priceStore } from '../stores/PriceStore'
  import {
    CryptoIcons,
    formatLocaleCurrency,
    dropEndingZeros,
    resizeWidget,
    totalProducts,
  } from '../util'
  import { post } from '../util/api_2'
  import { TransactionIntents, TransactionMediums } from '../types'
  import { Routes } from '../constants'
  import { ParentMessenger } from '../util/parent_messenger'
  import { formatExpiration } from '../util/transactions'
  import { configStore } from '../stores/ConfigStore'

  $: ({ product, products } = $configStore)
  $: ({ intent, wyrePreview } = $transactionStore)

  $: ({
    id: txnId,
    dest,
    sourceAmount,
    sourceCurrency,
    destAmount: destinationAmount,
    destCurrency: destinationCurrency,
    exchangeRate: txnExchangeRate,
    fees,
  } = wyrePreview ?? {
    id: 0,
    dest: 0,
    sourceAmount: 0,
    sourceCurrency: 'ETH',
    destAmount: 0,
    destCurrency: 'ETH',
    fees: { ETH: 1, BTC: 0.00001 },
    exchangeRate: 1,
    txnExchangeRate: 1,
  })

  // TODO: move vars that can be consolidated into a single reactive block
  $: isBuy = intent !== TransactionIntents.SELL
  $: isDebitCard = $transactionStore.inMedium === TransactionMediums.DEBIT_CARD
  $: cryptoTicker = isBuy ? destinationCurrency : sourceCurrency
  $: fiatTicker = isBuy ? sourceCurrency : destinationCurrency
  $: Icon = CryptoIcons[cryptoTicker]
  // $: exchangeRate = isBuy ? 1 / txnExchangeRate : txnExchangeRate
  $: total = isBuy ? sourceAmount : destinationAmount

  $: isConfirmingTxn = false
  $: isPreviewing = false
  $: cryptoFee = isBuy
    ? fees[destinationCurrency] / txnExchangeRate || 0
    : fees[sourceCurrency] || 0
  // since sendwyre's amount has fees baked in, subtract out using absolute value
  $: trueSourceAmount = isBuy
    ? Math.abs(sourceAmount - cryptoFee - fees[sourceCurrency])
    : sourceAmount

  let buttonText
  $: {
    console.log(sourceAmount, cryptoFee, fees[sourceCurrency])
    if (isBuy && !isDebitCard) {
      buttonText = isConfirmingTxn ? 'Buying' : 'Buy Now'
    } else if (isBuy && isDebitCard) {
      buttonText = 'Continue'
    } else {
      buttonText = isConfirmingTxn ? 'Selling' : 'Sell Now'
    }
  }

  $: formattedExpiration = formatExpiration(
    $transactionStore.transactionExpirationSeconds,
  )

  $: hasManyProducts = products?.length > 0
  $: destinationTicker = hasManyProducts
    ? products[0].destinationTicker
    : product?.destinationTicker

  const handleConfirmation = async () => {
    try {
      isConfirmingTxn = true
      // if (isDebitCard) {
      //   return push(Routes.DEBIT_CARD)
      // }
      // const txn = await window.API.fluxWyreConfirmTransfer(txnId, {
      //   transferId: txnId,
      // })
      // ParentMessenger.success(txn.id)
      // push(Routes.SUCCESS)
      // push(Routes.AWAIT_PAYMENT)
    } finally {
      isConfirmingTxn = false
    }
  }

  onMount(async () => {
    // afford more space to lists of product
    if (hasManyProducts)
      resizeWidget({ height: 650, width: 500 }, $configStore.appName)
    // TODO generate wyrePreview
    try {
      isPreviewing = true
      // FIXME set total?
      total = hasManyProducts ? totalProducts(products) : destinationAmount
      const { preview, depositAddress } = await post('transfers', {
        sourceCurrency,
        sourceAmount: total,
        destCurrency: destinationCurrency,
      })
      console.log('preview', preview)
      transactionStore.setWyrePreview(preview)
    } finally {
      isPreviewing = false
    }
  })
</script>

<ModalContent animation="none">
  <ModalHeader hideBackButton>Checkout</ModalHeader>
  <ModalBody>
    {#if product}
      <div class="nft-container">
        {#if product.videoURL}
          <video loop playsinline autoplay muted class="nft-video">
            <source src={product.videoURL} />
          </video>
        {:else if product.imageURL && !product.videoURL}
          <img alt={product.title} class="nft-image" src={product.imageURL} />
        {/if}
      </div>
    {/if}
    <div class="checkout-item-box">
      {#if product}
        <div class="nft-title">
          {product.title}
        </div>
      {:else if hasManyProducts}
        {#each products as product, i}
          <div class="product" in:fly={{ y: 25, duration: 200 * (i + 1) }}>
            <img height="50" width="50" src={product.img} />
            <div class="title" title={product.subtitle || product.author}>
              {product.title}
              <small>x</small>
              <b>
                {product.qty || 1}
              </b>
            </div>
            <div class="right">
              <b>
                {formatLocaleCurrency(
                  product.destinationTicker,
                  product.destinationAmount,
                )}
              </b>
            </div>
          </div>
        {/each}
      {/if}
    </div>
    <div class="line-items" class:is-product={Boolean(product)}>
      <!-- ACH -->
      {#if !isDebitCard && $transactionStore.selectedSourcePaymentMethod}
        {#if isBuy}
          <div class="line-item muted">
            <div>From</div>
            <div>{$transactionStore.selectedSourcePaymentMethod?.name}</div>
          </div>
          <div class="line-item muted" title={dest}>
            <div>To</div>
            <div title={dest}>
              {dest.substring(0, 6)}...{dest.substring(dest.length - 6)}
            </div>
          </div>
        {:else}
          <div class="line-item muted" title={dest}>
            <div>From</div>
            <div>
              {dest.substring(0, 6)}...{dest.substring(dest.length - 6)}
            </div>
          </div>
          <div class="line-item muted">
            <div>To</div>
            <div>{$transactionStore.selectedSourcePaymentMethod?.name}</div>
          </div>
        {/if}
        <div class="line dashed" />
      {/if}
      <!-- Debit Card -->
      {#if isDebitCard}
        <div class="line-item muted">
          <div>From</div>
          <div>Debit Card</div>
        </div>
        <div class="line-item muted" title={dest}>
          <div>To</div>
          <div>
            {dest.substring(0, 6)}...{dest.substring(dest.length - 6)}
          </div>
        </div>
        <div class="line dashed" />
      {/if}
    </div>
  </ModalBody>
  <ModalFooter>
    <div
      class="line-items"
      style="margin: 1rem -1rem 0 -1rem"
      class:is-product={Boolean(product)}
    >
      <div class="line-item muted warning" style="padding-top: 1rem;">
        <div>Price Expires</div>
        <div style="display:flex;justify-content:center;align-items:center;">
          <FaIcon data={faClock} />
          <div style="margin-right:0.35rem;" />
          <b>{formattedExpiration}</b>
        </div>
      </div>
      <div class="line dashed" />
      <div class="line-item muted">
        <div>Subtotal</div>
        <div>
          {formatLocaleCurrency(sourceCurrency, trueSourceAmount)}
        </div>
      </div>
      <div class="line-item muted">
        <div>Crypto Fee</div>
        <div>
          {formatLocaleCurrency(sourceCurrency, cryptoFee)}
        </div>
      </div>
      {#if isBuy}
        <div class="line-item muted">
          <div>Service Fee</div>
          <div>
            {formatLocaleCurrency(sourceCurrency, fees[sourceCurrency])}
          </div>
        </div>
      {/if}
      <div class="line dashed" />
      <div class="line-item" style="margin-bottom: 2.15rem;">
        <div><b>Total</b></div>
        <div>
          <b class="total">{formatLocaleCurrency(fiatTicker, sourceAmount)}</b>
        </div>
      </div>
      <Button
        glow
        isLoading={isConfirmingTxn}
        on:mousedown={handleConfirmation}
      >
        <div style="display:flex;justify-content:center;align-items:center;">
          <span style="margin-right:0.75rem;">
            {buttonText}
          </span>
        </div>
      </Button>
    </div>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';

  .checkout-item-box {
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    .checkout-item-name {
      margin-top: 1rem;
      height: 100%;
      display: flex;
      flex-direction: column;
      justify-content: flex-start;
      font-weight: bold;
      font-size: 1.5rem;
    }
  }

  .line {
    position: relative;
    height: 1px;
    max-height: 1px;
    width: 100%;
    margin: 0.5rem 0 0.5rem 0;
    position: relative;
    &.dashed {
      &:after {
        content: '';
        position: absolute;
        background: linear-gradient(
          to right,
          transparent,
          var(--theme-text-color),
          var(--theme-text-color),
          transparent
        );
        opacity: 0.15;
        height: 1px;
        bottom: 0;
        left: 0;
        right: 0;
      }
    }
  }

  .line-items {
    width: 100%;
    line-height: 1.15rem;
    align-self: center;
    margin-top: 2.5rem;
    padding: 0;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    &.is-product {
      margin-top: 1rem;
    }
    & > .line-item {
      display: flex;
      justify-content: space-between;
      align-items: flex-start;
      & > div:first-child {
        margin-right: 1rem;
        font-weight: 400;
      }
      .total {
        font-size: 1.25rem;
      }
      &.muted {
        color: var(--theme-color-muted);
        font-weight: 300;
      }
    }
  }

  .nft-title {
    margin-top: 1rem;
    font-weight: bold;
  }

  .nft-container {
    height: 25%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin-bottom: 0.5rem;
    .nft-video {
      height: 100%;
    }
    .nft-image {
      height: 100%;
    }
  }

  .product {
    display: flex;
    width: 100%;
    padding: 0 0.5rem;
    margin-bottom: 0.5rem;
    justify-content: space-between;
    img {
      background: var(--theme-color-muted);
      height: 50px;
      width: 50px;
      border-radius: 8px;
      overflow: hidden;
      background-size: cover;
      margin-right: 1rem;
    }
    .title {
      flex: 2;
      margin: auto;
    }
    > div.right {
      text-align: right;
      margin: auto 0.5rem auto auto;
      flex: 0.75;
    }

    small {
      margin: 0.15rem 0 0.2rem 0;
      font-size: 0.75rem;
      opacity: 0.8;
    }
  }
</style>
