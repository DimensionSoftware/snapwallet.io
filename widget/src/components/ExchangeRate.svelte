<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { TransactionIntents } from '../types'
  import { formatLocaleCurrency } from '../util'

  export let exchangeRate: number
  export let fakePrice: number = 0
  export let isLoadingPrices: boolean = false

  const dispatch = createEventDispatcher()

  $: ({ intent, destinationCurrency, sourceCurrency } = $transactionStore)
  $: isBuy = intent === TransactionIntents.BUY
  $: price = isLoadingPrices ? fakePrice : exchangeRate
  $: fiatTicker = isBuy ? sourceCurrency.ticker : destinationCurrency.ticker
  $: cryptoTicker = isBuy ? destinationCurrency.ticker : sourceCurrency.ticker
</script>

<div
  title="Best Exchange Rate"
  class="exchange-rate-container"
  on:mousedown={() => dispatch('mousedown')}
>
  1 {cryptoTicker} ≈ {formatLocaleCurrency(fiatTicker, price)}
</div>

<style lang="scss">
  .exchange-rate-container {
    position: relative;
    text-align: left;
    top: -30px;
    left: 2.5rem;
    font-size: 0.75rem;
    z-index: 2;
    color: var(--theme-text-color);
    font-weight: normal;
  }
</style>
