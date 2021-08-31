<script lang="ts">
  import { onMount } from 'svelte'
  import { faClipboard, faArrowDown } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import { CryptoIcons, formatLocaleCurrency, dropEndingZeros } from '../util'
  import { TransactionIntents, TransactionMediums } from '../types'
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import { ParentMessenger } from '../util/parent_messenger'
  import { configStore } from '../stores/ConfigStore'
  import { transactionStore } from '../stores/TransactionStore'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Surround from '../components/cards/Surround.svelte'
  import Button from '../components/Button.svelte'
  import Clipboard from '../components/Clipboard.svelte'

  const dstCurrency = $transactionStore.destinationCurrency?.toString(),
    dstAmount = $transactionStore.destinationAmount,
    dstAddress = '0xDEADBEEF'

  onMount(() => {
    // TODO build-in
  })
</script>

<ModalContent>
  <ModalHeader>Awaiting Payment</ModalHeader>
  <ModalBody>
    <Surround>
      <h2>Send</h2>
      <div class="row">
        <div class="crypto-icon">
          <svelte:component
            this={CryptoIcons[dstCurrency.toUpperCase() ?? 'BTC']}
          />
        </div>
        <h4>
          {dstAmount}
        </h4>
        <Clipboard value={dstAmount} />
      </div>
      <FaIcon class="down-arrow" data={faArrowDown} />
      <span>qrcode</span>
      <strong>Or, Copy & Paste</strong>
      <div class="row">
        <p>{dstAddress}</p>
        <Clipboard value={dstAddress} />
      </div>
    </Surround>
  </ModalBody>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';
  :global(.surround) {
    display: flex;
    text-align: center;
  }
  :global(.down-arrow) {
    align-self: center;
  }
  .row {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
  }
</style>
