<script lang="ts">
  import { onMount } from 'svelte'
  import { fly } from 'svelte/transition'
  import { pop, push } from 'svelte-spa-router'
  import FaIcon from 'svelte-awesome'
  import { faFileDownload } from '@fortawesome/free-solid-svg-icons'
  import { formatLocaleCurrency } from '../util'
  import { Routes } from '../constants'
  import { transactionsAsDataURI } from '../util/transactions'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Button from '../components/Button.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import TransactionCard from '../components/cards/TransactionCard.svelte'

  $: transfers = []
  $: csvURI = ''
  $: loading = true

  onMount(async () => {
    // TODO move transfers into store?
    try {
      const res = await window.API.fluxWyreGetTransfers()
      transfers = res.transfers
    } finally {
      setTimeout(() => (loading = false), 1000)
    }
    csvURI = transactionsAsDataURI(transfers)
  })
</script>

<ModalContent fillHeight>
  <ModalBody>
    <ModalHeader>Transactions</ModalHeader>
    {#if transfers?.length > 0}
      <a
        class="csv-link"
        href={csvURI}
        title="Download a CSV"
        download={`snap_txn_history_${new Date().toISOString()}.csv`}
        target="_blank"
      >
        <FaIcon data={faFileDownload} />
        <span>Download Transactions</span>
      </a>
      <div class="line-items">
        {#each transfers as transfer, i}
          <div
            style="margin-bottom: 1rem;"
            in:fly={{ y: 25, duration: 50 * i }}
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
      <a on:click={_ => push(Routes.ROOT)}>
        <h4 style="text-align: center;">Start Your First Transaction</h4>
      </a>
    {/if}
  </ModalBody>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';
  .csv-link {
    display: flex;
    align-items: center;
    margin: 0 0 0 0.3rem;
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
    width: 100%;
    align-self: center;
    margin-top: 1rem;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    overflow: hidden;
    overflow-y: scroll;
    & > .line-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
      &.muted {
        color: $textColor4;
        font-weight: 300;
      }
    }
  }
</style>
