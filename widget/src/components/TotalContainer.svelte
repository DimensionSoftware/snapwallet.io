<script lang="ts">
  import { transactionStore } from '../stores/TransactionStore'
  import { TransactionIntents } from '../types'
  import { dropEndingZeros } from '../util'
  import { configStore } from '../stores/ConfigStore'

  $: isDonation = $configStore.intent === 'donate'
  $: ({ destinationCurrency, intent, destinationAmount } = $transactionStore)
  $: precision = intent === TransactionIntents.BUY ? 8 : 2
</script>

<div class="total-container">
  {#if isDonation}
    Donate
  {:else}
    You Get ≈
  {/if}
  {dropEndingZeros(destinationAmount.toFixed(precision))}
  {destinationCurrency.ticker}
</div>

<style lang="scss">
  .total-container {
    color: var(--theme-text-color);
    font-weight: normal;
  }
</style>
