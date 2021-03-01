<script lang="ts">
  import vld8 from 'validator'
  import { push } from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { userStore } from '../stores/UserStore'
  import { onEnterPressed } from '../util'

  let animation = 'left'

  const handleNextStep = () => {
    // validate
    if (!vld8.isEmail($userStore.emailAddress))
      return document.querySelector('input[type="email"]')?.focus()
    // next
    push('#/profile')
  }

  const onKeyDown = (e: Event) => {
    onEnterPressed(e, handleNextStep)
  }
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent {animation}>
  <ModalBody>
    <ModalHeader hideCloseButton>
      Welcome
    </ModalHeader>
    <Label label="Your Email">
      <Input
        inputmode="email"
        autocapitalize="none"
        autocomplete="on"
        autofocus
        required
        type="email"
        placeholder="your@email.address"
        defaultValue={$userStore.emailAddress}
        on:change={e => userStore.setEmailAddress(e.detail)}
      />
    </Label>
  </ModalBody>
  <ModalFooter>
    <Button on:click={handleNextStep}>Login <small>or</small> SignUp</Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
</style>
