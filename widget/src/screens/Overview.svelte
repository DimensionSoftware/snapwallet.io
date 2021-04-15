<script lang="ts">
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Button from '../components/Button.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { CryptoIcons, formatLocaleCurrency, dropEndingZeros } from '../util'
  import { TransactionIntents } from '../types'
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import { ParentMessenger } from '../util/parent_messenger'
  import { faClock, faLock } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import { onMount } from 'svelte'
  import { toaster } from '../stores/ToastStore'
  import { computeTransactionExpiration } from '../util/transactions'

  export let product

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
    expiresAt,
  } = wyrePreview)

  $: isBuy = intent === TransactionIntents.BUY
  $: cryptoTicker = isBuy ? destinationCurrency : sourceCurrency
  $: fiatTicker = isBuy ? sourceCurrency : destinationCurrency
  $: cryptoAmount = isBuy ? destinationAmount : sourceAmount
  $: Icon = CryptoIcons[cryptoTicker]
  $: screenTitle = $transactionStore.intent === 'buy' ? 'Buying' : 'Selling'
  // $: exchangeRate = isBuy ? 1 / txnExchangeRate : txnExchangeRate
  $: total = isBuy ? sourceAmount : destinationAmount

  $: cryptoPrecision = cryptoAmount % 1 === 0 ? 1 : 8
  $: isConfirmingTxn = false
  $: cryptoFee = isBuy
    ? fees[destinationCurrency] / txnExchangeRate
    : fees[sourceCurrency]
  $: trueSourceAmount = isBuy
    ? sourceAmount - cryptoFee - fees[sourceCurrency]
    : sourceAmount

  let buttonText
  $: {
    if (isBuy) {
      buttonText = isConfirmingTxn ? 'Buying' : 'Buy Now'
    } else {
      buttonText = isConfirmingTxn ? 'Selling' : 'Sell Now'
    }
  }

  let secondsUntilExpiration = computeTransactionExpiration(
    $transactionStore.wyrePreview?.expiresAt,
  )

  let formattedExpiration
  $: {
    const d = new Date(0)
    d.setSeconds(secondsUntilExpiration)
    const mins = d.getMinutes()
    if (mins > 1) {
      formattedExpiration = `${mins.toString()}m ${d.getSeconds().toString()}s`
    } else {
      formattedExpiration = `${d.getSeconds().toString()}s`
    }
  }

  const handleConfirmation = async () => {
    try {
      isConfirmingTxn = true
      const txn = await window.API.fluxWyreConfirmTransfer(txnId, {
        transferId: txnId,
      })
      ParentMessenger.success(txn.id)
      push(Routes.SUCCESS)
    } finally {
      isConfirmingTxn = false
    }
  }

  onMount(() => {
    const interval = setInterval(() => {
      secondsUntilExpiration = computeTransactionExpiration(
        $transactionStore.wyrePreview?.expiresAt,
      )
      if (secondsUntilExpiration <= 0) {
        toaster.pop({
          msg: 'Your preview has expired. Please create a new preview.',
          error: true,
        })
        push(Routes.ROOT)
      }
    }, 1000)
    return () => clearInterval(interval)
  })
</script>

<ModalContent>
  <ModalHeader>{screenTitle}</ModalHeader>
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
      {#if !product}
        <div style="width:30%;" class="checkout-item-icon">
          <Icon size="100%" height="100%" width="100%" viewBox="-4 0 40 40" />
        </div>
        <div class="checkout-item-name">
          {dropEndingZeros(cryptoAmount.toFixed(cryptoPrecision))}
          {cryptoTicker}
        </div>
      {:else}
        <div class="nft-title">
          {product.title}
        </div>
      {/if}
    </div>
    <div class="line-items" class:is-product={Boolean(product)}>
      {#if $transactionStore.selectedSourcePaymentMethod}
        <div class="line-item muted warning">
          <div>Price Expires</div>
          <div style="display:flex;justify-content:center;align-items:center;">
            <FaIcon data={faClock} />
            <div style="margin-right:0.35rem;" />
            <b>{formattedExpiration}</b>
          </div>
        </div>
        <div class="line dashed" />
        {#if isBuy}
          <div class="line-item muted">
            <div>From</div>
            <div>{$transactionStore.selectedSourcePaymentMethod?.name}</div>
          </div>
          <div class="line-item muted">
            <div>To</div>
            <div>
              {dest.substring(0, 6)}...{dest.substring(dest.length - 4)}
            </div>
          </div>
        {:else}
          <div class="line-item muted">
            <div>From</div>
            <div>
              {dest.substring(0, 6)}...{dest.substring(dest.length - 4)}
            </div>
          </div>
          <div class="line-item muted">
            <div>To</div>
            <div>{$transactionStore.selectedSourcePaymentMethod?.name}</div>
          </div>
        {/if}
        <div class="line dashed" />
      {/if}
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
      <div class="line-item">
        <div><b>Total</b></div>
        <div>
          <b>{formatLocaleCurrency(fiatTicker, total)}</b>
        </div>
      </div>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button isLoading={isConfirmingTxn} on:mousedown={handleConfirmation}>
      <div style="display:flex;justify-content:center;align-items:center;">
        <span style="margin-right:0.75rem;">
          {buttonText}
        </span>
        <FaIcon data={faLock} />
      </div>
    </Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';

  .checkout-item-box {
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
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
    height: 1px;
    max-height: 1px;
    width: 100%;
    border-bottom: 0.5px solid var(--theme-text-color);
    margin: 0.5rem 0 0.5rem 0;
    &.dashed {
      border-bottom: 0.7px dashed var(--theme-text-color);
    }
  }

  .line-items {
    width: 100%;
    align-self: center;
    margin-top: 2.5rem;
    padding: 0 0.7rem;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    &.is-product {
      margin-top: 1rem;
    }
    & > .line-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
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
</style>
