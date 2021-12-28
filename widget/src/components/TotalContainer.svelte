<script lang="ts">
  import { transactionStore } from '../stores/TransactionStore'
  import { TransactionIntents } from '../types'
  import { dropEndingZeros } from '../util'
  import { configStore } from '../stores/ConfigStore'

  export let isBuy = true

  $: isDonation = $configStore.intent === 'donate'
  $: ({
    destinationCurrency,
    intent,
    destinationAmount,
    sourceAmount,
    sourceCurrency,
  } = $transactionStore)
  $: precision = intent === TransactionIntents.BUY ? 8 : 2
</script>

<div class="total-container">
  {#if isDonation}
    Donate
  {:else}
    You Get ≈
  {/if}
  {#if isBuy}
    <strong>
      {dropEndingZeros(destinationAmount.toFixed(precision))}
      {destinationCurrency.ticker}
    </strong>
  {:else}
    <strong>
      {dropEndingZeros(destinationAmount.toFixed(2))}
      {destinationCurrency.ticker}
    </strong>
  {/if}
</div>

<style lang="scss">
  .total-container {
    // inherit surround color for success/disabled/etc...
    // color: var(--theme-text-color);
  }
</style>
