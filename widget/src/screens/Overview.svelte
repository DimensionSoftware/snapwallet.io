<script lang="ts">
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Button from '../components/Button.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { priceStore } from '../stores/PriceStore'
  import { CryptoIcons, isValidNumber } from '../util'
  import { onMount, afterUpdate } from 'svelte'

  $: Icon = CryptoIcons[$transactionStore.destinationCurrency.ticker]
  $: screenTitle = $transactionStore.intent === 'buy' ? 'Buying' : 'Selling'
  $: buttonText = $transactionStore.intent === 'buy' ? 'Buy Now' : 'Sell Now'

  // TODO: these prices will be removed once txn endpoint is wired up ;)
  // Price Data
  $: selectedDirection = `${$transactionStore.sourceCurrency.ticker}_${$transactionStore.destinationCurrency.ticker}`
  $: selectedPriceMap = $priceStore.prices[selectedDirection]
  $: destRate = selectedPriceMap[$transactionStore.destinationCurrency.ticker]
  $: exchangeRate = 1 / destRate
  $: destinationAmount = $transactionStore.sourceAmount * destRate

  // TODO: move to util
  const formatLocaleCurrency = (ticker: string, amount: number) => {
    amount = isValidNumber(amount) ? amount : 0
    const locale =
      (navigator?.languages || [])[0] || navigator?.language || 'en-US'
    return new Intl.NumberFormat(locale, {
      style: 'currency',
      currency: ticker,
    }).format(amount)
  }

  onMount(() => {
    priceStore.fetchPrices()
    const priceInterval = priceStore.pollPrices()
    return () => clearInterval(priceInterval)
  })
</script>

<ModalContent>
  <ModalBody>
    <ModalHeader>{screenTitle}</ModalHeader>
    <div class="checkout-item-box">
      <div class="checkout-item-icon">
        <Icon size="80" />
      </div>
      <div class="checkout-item-name">
        {(destinationAmount).toFixed(8)}
        {$transactionStore.destinationCurrency.ticker}
      </div>
    </div>
    <div class="line-items">
      <h4>Overview</h4>
      <div class="line-item muted">
        <div>Crypto Fee</div>
        <div>
          {(0).toFixed(8)}
          {$transactionStore.destinationCurrency.ticker}
        </div>
      </div>
      <div class="line-item muted">
        <div>Partner Fee</div>
        <div>
          {formatLocaleCurrency($transactionStore.sourceCurrency.ticker, 0)}
        </div>
      </div>
      <div class="line-item muted">
        <div>Exchange Rate</div>
        <div>
          {formatLocaleCurrency(
            $transactionStore.sourceCurrency.ticker,
            exchangeRate,
          )}
        </div>
      </div>     
      <div class="line dashed" />
      <div class="line-item">
        <div><b>Total</b></div>
        <div>
          {formatLocaleCurrency(
            $transactionStore.sourceCurrency.ticker,
            $transactionStore.sourceAmount,
          )}
        </div>
      </div>
      <div class="line dashed" /> 
      <div class="line-item muted">
        <div>Wallet</div>
        <div>3x2kdkdj...k34w</div>  
      </div> 
    </div>
  </ModalBody>
  <ModalFooter>
    <Button>{buttonText}</Button>
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
    margin-top: 0.8rem;
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
