<script lang="ts">
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import IntentSelector from '../components/IntentSelector.svelte'
  import { toaster } from '../stores/ToastStore'
  import Button from '../components/Button.svelte'
  import { userStore } from '../stores/UserStore'

  const handleClick = () => {
    toaster.pop({ msg: 'Success', success: true })
  }
</script>

<ModalContent>
  <ModalBody>
    <IntentSelector />
  </ModalBody>
  <ModalFooter>
    <Button on:click={handleClick}>{$userStore.intent}</Button></ModalFooter
  >
</ModalContent>
