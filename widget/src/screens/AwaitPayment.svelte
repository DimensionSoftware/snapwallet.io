<script lang="ts">
  import { onMount } from 'svelte'
  import { faClipboard } from '@fortawesome/free-solid-svg-icons'
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
  import Surround from './cards/Surround.svelte'
  import Button from '../components/Button.svelte'

  onMount(() => {
    // TODO build-in
  })
</script>

<ModalContent>
  <ModalHeader>Awaiting Payment</ModalHeader>
  <ModalBody>
    <Surround>
      <h3>Send</h3>
      <h4>Send</h4>
      <strong>Or, Copy & Paste</strong>
    </Surround>
  </ModalBody>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';
</style>
