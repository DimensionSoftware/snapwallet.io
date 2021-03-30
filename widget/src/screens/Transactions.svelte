<script lang="ts">
  import { onMount } from 'svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Button from '../components/Button.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { formatLocaleCurrency } from '../util'
  import type { WyreTransfer, WyreTransfers } from 'api-client'
  import { pop } from 'svelte-spa-router'
  import { exportTransactionsAsCSV } from '../util/transactions'
  import FaIcon from 'svelte-awesome'
  import { faFileDownload } from '@fortawesome/free-solid-svg-icons'

  $: transfers = []

  onMount(async () => {
    const res = await window.API.fluxWyreGetTransfers()
    transfers = res.transfers
  })

  const close = () => {
    pop()
  }
</script>

<ModalContent>
  <ModalBody>
    <ModalHeader>Transactions</ModalHeader>
    <div
      style="cursor:pointer;"
      on:click={() => exportTransactionsAsCSV(transfers)}
    >
      <FaIcon data={faFileDownload} />
    </div>
    <div class="line-items">
      <ol>
        {#each transfers as transfer, i}
          <li>
            <span>{transfer.createdAt}</span>
            <span>{transfer.sourceCurrency}</span>
            <span>{transfer.destCurrency}</span>
            <span>{transfer.destAmount}</span>
            <span>{transfer.status}</span>
          </li>
        {/each}
      </ol>
      <div class="line dashed" />
      <div class="line-item">
        <div><b>Total</b></div>
        <div>
          <!-- {formatLocaleCurrency(fiatTicker, total)} -->
        </div>
      </div>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button on:click={close}>Back</Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';

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
    width: 90%;
    align-self: center;
    margin-top: 1rem;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
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
