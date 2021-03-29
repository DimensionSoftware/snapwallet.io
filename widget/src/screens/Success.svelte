<script lang="ts">
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Button from '../components/Button.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { TransactionIntents } from '../types'
  import { faCheckCircle } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import { ParentMessenger } from '../util/parent_messenger'

  $: ({ intent, wyrePreview } = $transactionStore)
  $: ({ sourceCurrency, destCurrency: destinationCurrency } = wyrePreview)
  $: isBuy = intent === TransactionIntents.BUY
  $: cryptoTicker = isBuy ? destinationCurrency : sourceCurrency
</script>

<ModalContent>
  <ModalBody>
    <ModalHeader hideBackButton>Success</ModalHeader>
    <div class="icon-box">
      <FaIcon scale="4.5" data={faCheckCircle} />
    </div>
    <div class="text-center">
      <p>You're {cryptoTicker} checkout was successful!</p>
      <p>Please allow up to 5 business days for your purchase to complete.</p>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button on:click={ParentMessenger.exit}>Done</Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';

  .icon-box {
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    color: var(--theme-success-color);
    margin-bottom: 0.5rem;
  }
  .text-center {
    margin: 0 1rem;
  }
</style>
