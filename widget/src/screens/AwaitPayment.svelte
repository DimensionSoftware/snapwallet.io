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

  const dstCurrency = $transactionStore.destinationCurrency?.toString()

  onMount(() => {
    // TODO build-in
  })
</script>

<ModalContent>
  <ModalHeader>Awaiting Payment</ModalHeader>
  <ModalBody>
    <Surround>
      <h3>Send</h3>
      <div class="row">
        <div class="crypto-icon">
          <svelte:component
            this={CryptoIcons[
              $transactionStore.destinationCurrency?.toString().toUpperCase() ??
                'BTC'
            ]}
          />
        </div>
        <h4>
          {dstCurrency}
          {dstCurrency}
        </h4>
        <Clipboard value="foo" />
      </div>
      <FaIcon data={faArrowDown} />
      <span>qrcode</span>
      <strong>Or, Copy & Paste</strong>
    </Surround>
  </ModalBody>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';
  .row {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
  }
</style>
