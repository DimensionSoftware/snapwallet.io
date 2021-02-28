<script lang="ts">
  import { isValidNumber } from '../util'
  import { transactionStore } from '../stores/TransactionStore'

  export let isDestination: boolean
  export let rate: number = 0.0

  $: precision = isDestination ? 8 : 2

  $: ticker = isDestination
    ? $transactionStore.destinationCurrency.ticker
    : $transactionStore.sourceCurrency.ticker

  $: total = isValidNumber(rate) ? rate : 0
</script>

<div class="total-container">
  {total.toFixed(precision)}
  {ticker}
</div>

<style lang="scss">
  .total-container {
    display: flex;
    justify-content: flex-end;
    color: var(--theme-text-color);
    font-weight: 500;
    font-size: 0.9rem;
  }
</style>
