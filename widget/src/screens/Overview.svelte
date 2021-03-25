<script lang="ts">
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Button from '../components/Button.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { CryptoIcons, formatLocaleCurrency, dropEndingZeros } from '../util'
  import { TransactionIntents } from '../types'

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
  $: cryptoTicker = isBuy ? destinationCurrency : sourceCurrency
  $: fiatTicker = isBuy ? sourceCurrency : destinationCurrency
  $: cryptoAmount = isBuy ? destinationAmount : sourceAmount
  $: Icon = CryptoIcons[cryptoTicker]
  $: screenTitle = $transactionStore.intent === 'buy' ? 'Buying' : 'Selling'
  $: exchangeRate = isBuy ? 1 / txnExchangeRate : txnExchangeRate
  $: total = isBuy ? sourceAmount : destinationAmount

  $: cryptoPrecision = cryptoAmount % 1 === 0 ? 1 : 8
  $: isConfirmingTxn = false

  let buttonText
  $: {
    if (isBuy) {
      buttonText = isConfirmingTxn ? 'Buying' : 'Buy Now'
    } else {
      buttonText = isConfirmingTxn ? 'Selling' : 'Sell Now'
    }
  }

  const handleConfirmation = async () => {
    try {
      isConfirmingTxn = true
      await window.API.fluxWyreConfirmTransfer(txnId, { transferId: txnId })
    } finally {
      isConfirmingTxn = false
    }
  }
</script>

<ModalContent>
  <ModalBody>
    <ModalHeader>{screenTitle}</ModalHeader>
    <div class="checkout-item-box">
      <div class="checkout-item-icon">
        <Icon size="80" />
      </div>
      <div class="checkout-item-name">
        {dropEndingZeros(cryptoAmount.toFixed(cryptoPrecision))}
        {cryptoTicker}
      </div>
    </div>
    <div class="line-items">
      <h4>Overview</h4>
      {#if $transactionStore.selectedSourcePaymentMethod}
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
        <div>Crypto Fee</div>
        <div>
          {dropEndingZeros(fees[destinationCurrency].toFixed(cryptoPrecision))}
          {cryptoTicker}
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
      <div class="line-item muted">
        <div>Exchange Rate</div>
        <div>
          {formatLocaleCurrency(fiatTicker, exchangeRate)}
        </div>
      </div>
      <div class="line dashed" />
      <div class="line-item">
        <div><b>Total</b></div>
        <div>
          {formatLocaleCurrency(fiatTicker, total)}
        </div>
      </div>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button isLoading={isConfirmingTxn} on:click={handleConfirmation}
      >{buttonText}</Button
    >
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
      margin-top: 1.4rem;
      height: 100%;
      display: flex;
      flex-direction: column;
      justify-content: flex-start;
      font-weight: bold;
      font-size: 1.3rem;
    }
  }

  .line {
    height: 1px;
    max-height: 1px;
    width: 100%;
    border-bottom: 0.5px solid $textColor4;
    margin: 0.5rem 0 0.5rem 0;
    &.dashed {
      border-bottom: 0.7px dashed $textColor4;
    }
  }

  .line-items {
    width: 90%;
    align-self: center;
    margin-top: 1rem;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    & > .line-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
      &.muted {
        color: $textColor4;
        font-weight: 300;
      }
    }
  }
</style>
