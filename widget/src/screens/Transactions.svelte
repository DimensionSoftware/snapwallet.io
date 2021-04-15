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
  import { transactionsStore } from '../stores/TransactionsStore'

  $: transfers = $transactionsStore
  $: csvURI = ''
  $: loading = true
  $: csvFileName = getFileName()
  let csvElement: HTMLElement

  const getFileName = () => {
    return `snap_txn_history_${new Date().toISOString()}.csv`
  }

  onMount(async () => {
    await transactionsStore.fetchUserTransactions()
    loading = false
    csvURI = transactionsAsDataURI(transfers)
  })
</script>

<ModalContent fullscreen>
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
      <div class="line-items">
        {#each transfers as transfer, i}
          <div
            style="margin-bottom: 1rem;"
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
    height: 1px;
    max-height: 1px;
    width: 100%;
    border-bottom: 0.5px solid $textColor4;
    margin: 0.5rem 0 0.5rem 0;
    &.dashed {
      border-bottom: 0.7px dashed $textColor4;
    }
  }

  .line-items {
    height: 100%;
    width: 100%;
    align-self: center;
    margin-top: 0.5rem;
    padding: 0.5rem 1.5rem 0 1.5rem;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    overflow: hidden;
    overflow-y: scroll;
    scrollbar-width: thin;
    & > .line-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
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
