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
  import { toaster } from '../stores/ToastStore'
  import { Logger, onEnterPressed } from '../util'

  let animation = 'left'
  let isMakingRequest = false

  const timeout = 700

  const handleNextStep = async () => {
    isMakingRequest = true

    try {
      let emailIsValid = vld8.isEmail($userStore.emailAddress)
      if (!emailIsValid)
        (document.querySelector('input[type="email"]') as any).focus()

      await window.API().fluxOneTimePasscode({
        emailOrPhone: $userStore.emailAddress,
      })

      setTimeout(() => push('#/verify-otp'), timeout)
    } catch (e) {
      const err = e as { body: { code: number; message: string } }
      Logger.error(err)

      toaster.pop({
        msg: err.body.message,
        error: true,
      })
    } finally {
      setTimeout(() => (isMakingRequest = false), timeout)
    }
  }

  const onKeyDown = (e: Event) => {
    onEnterPressed(e, handleNextStep)
  }
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent {animation}>
  <ModalBody>
    <ModalHeader>Get Email Code</ModalHeader>
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
    <Button disabled={isMakingRequest} on:click={handleNextStep}>
      {#if isMakingRequest}
        Getting Code...
      {:else}
        Get Code
      {/if}
    </Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
</style>
