<script lang="ts">
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Button from '../components/Button.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { CryptoIcons, formatLocaleCurrency, dropEndingZeros } from '../util'
  import { TransactionIntents, TransactionMediums } from '../types'
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import { ParentMessenger } from '../util/parent_messenger'
  import { faClock, faLock } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import { onMount } from 'svelte'
  import { toaster } from '../stores/ToastStore'
  import { formatExpiration } from '../util/transactions'
  import { configStore } from '../stores/ConfigStore'

  $: ({ product } = $configStore)

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
  } = wyrePreview)

  $: isBuy = intent === TransactionIntents.BUY
  $: isDebitCard = $transactionStore.inMedium === TransactionMediums.DEBIT_CARD
  $: cryptoTicker = isBuy ? destinationCurrency : sourceCurrency
  $: fiatTicker = isBuy ? sourceCurrency : destinationCurrency
  $: cryptoAmount = isBuy ? destinationAmount : sourceAmount
  $: Icon = CryptoIcons[cryptoTicker]
  $: screenTitle = screenTitleFrom($configStore.intent)
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

  const handleConfirmation = async () => {
    try {
      isConfirmingTxn = true
      if (isDebitCard) {
        return push(Routes.DEBIT_CARD)
      }
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
      if ($transactionStore.transactionExpirationSeconds <= 0) {
        toaster.pop({
          msg: 'Your preview has expired. Please create a new preview.',
          error: true,
        })
        push(Routes.ROOT)
      }
    }, 1000)
    return () => clearInterval(interval)
  })

  function screenTitleFrom(intent) {
    if (intent === 'donate') return $configStore.payee || 'Donation'
    return intent === 'buy' ? 'Buying' : 'Selling'
  }
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
      <div class="line-item muted warning">
        <div>Price Expires</div>
        <div style="display:flex;justify-content:center;align-items:center;">
          <FaIcon data={faClock} />
          <div style="margin-right:0.35rem;" />
          <b>{formattedExpiration}</b>
        </div>
      </div>
      <div class="line dashed" />
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
    <Button glow isLoading={isConfirmingTxn} on:mousedown={handleConfirmation}>
      <div style="display:flex;justify-content:center;align-items:center;">
        <span style="margin-right:0.75rem;">
          {buttonText}
        </span>
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
    line-height: 1.5rem;
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
      align-items: flex-start;
      & > div:first-child {
        margin-right: 1rem;
        font-weight: 400;
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
</style>
