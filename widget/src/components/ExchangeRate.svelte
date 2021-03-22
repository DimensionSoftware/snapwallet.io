<script lang="ts">
  import { transactionStore } from '../stores/TransactionStore'
  import { TransactionIntents } from '../types'
  import { formatLocaleCurrency } from '../util'

  export let exchangeRate: number
  export let fakePrice: number = 0
  export let isLoadingPrices: boolean = false

  $: ({ intent, destinationCurrency, sourceCurrency } = $transactionStore)
  $: isBuy = intent === TransactionIntents.BUY
  $: price = isLoadingPrices ? fakePrice : exchangeRate
  $: fiatTicker = isBuy ? sourceCurrency.ticker : destinationCurrency.ticker
  $: cryptoTicker = isBuy ? destinationCurrency.ticker : sourceCurrency.ticker
</script>

<div class="exchange-rate-container">
  1 {cryptoTicker} ≈ {formatLocaleCurrency(fiatTicker, price)}
</div>

<style lang="scss">
  .exchange-rate-container {
    height: 1.5rem;
    position: relative;
    z-index: 2;
    font-size: 0.9rem;
    color: var(--theme-text-color);
    font-weight: normal;
  }
</style>
