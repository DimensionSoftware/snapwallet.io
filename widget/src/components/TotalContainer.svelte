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
  $: dst = dropEndingZeros(destinationAmount.toFixed(precision))
</script>

<div class="total-container">
  {#if isDonation}
    Donate
  {:else}
    You Get ≈
  {/if}
  {#if isBuy}
    <strong>
      {dst === 'NaN' ? '0' : dst}
      {destinationCurrency.ticker}
    </strong>
  {:else}
    <strong>
      {dst === 'NaN' ? '0' : dst}
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
