<script lang="ts">
  import type { WyreTransfer } from 'api-client'
  import {
    formatHumanDate,
    formatDate,
    capitalize,
    formatLocaleCurrency,
  } from '../../util'
  import Badge from '../Badge.svelte'
  export let transaction: WyreTransfer

  const status = transaction.status.toLowerCase(),
    badgeText = capitalize(status),
    badgeType = status,
    label = `${transaction.sourceCurrency} • ${transaction.destCurrency}`
</script>

<div class="container">
  <div class="content-container">
    <h4>{label}</h4>
    <h5 title={formatHumanDate(transaction.createdAt)}>
      {formatDate(transaction.createdAt)}
    </h5>
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
  <Badge
    error={badgeType === 'error'}
    success={['success', 'confirmed', 'completed'].includes(
      status.toLocaleLowerCase(),
    )}
    warning={badgeType === 'pending'}
  >
    {badgeText}
  </Badge>
{/if}

<style lang="scss">
  @import '../../styles/_vars.scss';

  .container {
    flex: 1;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.25rem 0 0 0;
    margin-left: 0.2rem;
    & > .content-container {
      height: 100%;
      width: 100%;
      h4,
      h5 {
        margin: 0;
      }
      h4 {
        line-height: 1.5rem;
      }
      h5 {
        font-weight: 500;
      }
      &.right {
        text-align: right;
      }
    }
  }
</style>
