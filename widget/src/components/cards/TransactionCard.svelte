<script lang="ts">
  import type { WyreTransfer } from 'api-client'
  import { formatDate, capitalize, formatLocaleCurrency } from '../../util'
  export let transaction: WyreTransfer

  const status = transaction.status.toLowerCase(),
    badgeText = capitalize(status),
    badgeType = status,
    label = `${transaction.sourceCurrency} • ${transaction.destCurrency}`
</script>

<div class="container">
  <div class="content-container">
    <h4>{label}</h4>
    <h5>{formatDate(transaction.createdAt)}</h5>
  </div>
  <div class="content-container right">
    <h4>{transaction.destAmount}</h4>
    <h5>
      {formatLocaleCurrency(
        transaction.sourceCurrency,
        transaction.sourceAmount,
      )}
    </h5>
  </div>
</div>
{#if badgeText}
  <div
    class="badge"
    class:error={badgeType === 'error'}
    class:success={badgeType === 'success'}
    class:warning={badgeType === 'pending'}
  >
    {badgeText}
  </div>
{/if}

<style lang="scss">
  @import '../../styles/_vars.scss';

  .container {
    flex: 1;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.25rem 0;
    margin-left: 0.2rem;
    & > .content-container {
      height: 100%;
      width: 100%;
      h4,
      h5 {
        margin: 0;
      }
      h4 {
        line-height: 1.15rem;
      }
      h5 {
        font-weight: 500;
      }
      &.right {
        text-align: right;
      }
    }
  }

  .badge {
    border-radius: 0.5rem;
    padding: 0 0.5rem;
    font-size: 0.75rem;
    float: left;
    &.success {
      color: var(--theme-text-color);
      border: 1px solid var(--theme-success-color);
      background-color: var(--theme-success-color);
    }

    &.warning {
      color: var(--theme-text-color);
      border: 1px solid var(--theme-warning-color);
      background-color: var(--theme-warning-color);
    }

    &.error {
      color: var(--theme-text-color);
      border: 1px solid var(--theme-error-color);
      background-color: var(--theme-error-color);
    }
  }
</style>
