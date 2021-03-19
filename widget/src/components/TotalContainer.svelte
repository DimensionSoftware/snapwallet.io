<script lang="ts">
  import { transactionStore } from '../stores/TransactionStore'
  import { TransactionIntents } from '../types'

  $: ({ destinationCurrency, intent, destinationAmount } = $transactionStore)
  $: precision = intent === TransactionIntents.BUY ? 8 : 2
</script>

<div class="total-container">
  {destinationAmount.toFixed(precision)}
  {destinationCurrency.ticker}
</div>

<style lang="scss">
  .total-container {
    color: var(--theme-text-color);
    font-weight: 500;
    font-size: 0.9rem;
  }
</style>
