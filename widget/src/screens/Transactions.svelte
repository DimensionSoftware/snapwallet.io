<script lang="ts">
  import { onMount } from 'svelte'
  import { fly } from 'svelte/transition'
  import { push } from 'svelte-spa-router'
  import FaIcon from 'svelte-awesome'
  import { faFileDownload } from '@fortawesome/free-solid-svg-icons'
  import { formatLocaleCurrency } from '../util'
  import { Routes } from '../constants'
  import { transactionsAsDataURI } from '../util/transactions'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import TransactionCard from '../components/cards/TransactionCard.svelte'
  import {
    transactionDetailsStore,
    transactionsStore,
  } from '../stores/TransactionsStore'

  $: transfers = $transactionsStore
  $: csvURI = ''
  $: loading = true
  $: csvFileName = getFileName()
  let csvElement: HTMLElement

  const getFileName = () => {
    return `snap_txn_history_${new Date().toISOString()}.csv`
  }

  onMount(async () => {
    // NOTE: let SideMenu prefetch do its thing :melder:
    if (!$transactionsStore.length) {
      await transactionsStore.fetchUserTransactions()
    }
    transactionDetailsStore.reset()
    loading = false
    csvURI = transactionsAsDataURI(transfers)
  })
</script>

<ModalContent>
  <ModalHeader>My Transactions</ModalHeader>
  <ModalBody fullscreen>
    {#if transfers?.length > 0}
      <a
        bind:this={csvElement}
        class="csv-link"
        href={csvURI}
        download={csvFileName}
        target="_blank"
      >
        <div style="display:flex;justify-content:center;align-items:center;">
          <div style="display:flex;justify-content:center;margin-left:0.5rem;">
            <FaIcon data={faFileDownload} />
          </div>
          Download
        </div>
      </a>
      <div class="line-items scroll-y">
        {#each transfers as transfer, i}
          <div
            on:click={() => {
              $transactionDetailsStore = { transaction: transfer }
              push(Routes.TRANSACTION_DETAILS)
            }}
            style="margin-bottom: 1rem;cursor:pointer;"
            in:fly={{ y: 25, duration: 350 + 50 * i }}
          >
            <TransactionCard transaction={transfer} />
          </div>
        {/each}
        <div class="line dashed" />
        <div class="total">
          <b>Total</b>
          <div>
            {formatLocaleCurrency(
              transfers[0].sourceCurrency,
              transfers.reduce((prev, cur) => {
                return prev + cur.sourceAmount
              }, 0),
            )}
          </div>
        </div>
      </div>
    {:else if loading}
      <h4 style="text-align: center;">Getting Your Transactions...</h4>
    {:else}
      <a on:mousedown={_ => push(Routes.ROOT)}>
        <h4 style="text-align: center;">Start Your First Transaction</h4>
      </a>
    {/if}
  </ModalBody>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';
  h4,
  a {
    color: var(--theme-text-color);
  }
  .csv-link {
    max-width: 50%;
    color: var(--theme-text-color);
    margin: 0 auto;
    display: flex;
    align-items: center;
    justify-content: center;
    text-decoration: none;
    :global(svg) {
      margin-right: 0.5rem;
    }
  }
  .checkout-item-box {
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
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

  .total {
    display: flex;
    justify-content: space-between;
    padding-bottom: 2rem;
    margin: 0.5rem 0 0 0.5rem;
  }
  .line {
    position: relative;
    height: 1px;
    max-height: 1px;
    width: 100%;
    margin: 0.5rem 0 0.5rem 0;
    &.dashed {
      &:after {
        content: '';
        position: absolute;
        background: linear-gradient(
          to right,
          transparent,
          var(--theme-text-color),
          var(--theme-text-color),
          transparent
        );
        opacity: 0.15;
        height: 1px;
        bottom: 0;
        left: 0;
        right: 0;
      }
    }
  }

  .line-items {
    line-height: 1.5rem;
    height: 100%;
    width: 100%;
    align-self: center;
    margin-top: 0.5rem;
    padding: 1.5rem 1.75rem 1rem 1.5rem;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    & > .line-item {
      display: flex;
      justify-content: space-between;
      align-items: flex-start;
      & > div:first-child {
        margin-right: 1rem;
        font-weight: 400;
      }
      &.muted {
        color: $textColor4;
        font-weight: 300;
      }
    }
    &:after {
      @include bottom-shadow;
    }
  }
</style>
