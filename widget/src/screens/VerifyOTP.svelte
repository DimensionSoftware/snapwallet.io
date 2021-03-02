<script lang="ts">
  import { push } from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { userStore } from '../stores/UserStore'
  import { Logger, onEnterPressed } from '../util'
  import type { OneTimePasscodeVerifyResponse } from 'api-client'
  import { toaster } from '../stores/ToastStore'

  let animation = 'left'
  let code = ''
  let isMakingRequest = false

  const handleNextStep = async () => {
    isMakingRequest = true
    const c = code
    Logger.debug('Verifying using OTP code:', c)

    try {
      const response: OneTimePasscodeVerifyResponse = await window
        .API()
        .fluxOneTimePasscodeVerify({
          emailOrPhone: $userStore.emailAddress,
          code: c,
        })
      Logger.debug('LOGGED IN:', response.user)
      push('#/profile')
    } catch (e) {
      Logger.error(e)
      // TODO: move error messages to the server
      let msg = 'An unknown error occurred. Please try again later.'
      const code = e.body?.code

      if ([16, 3].includes(code)) {
        msg = 'The email code provided was not valid. Please try again.'
      }

      toaster.pop({
        msg,
        error: true,
      })
    } finally {
      setTimeout(() => (isMakingRequest = false), 700)
    }
  }

  const onKeyDown = (e: Event) => {
    onEnterPressed(e, handleNextStep)
  }
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent {animation}>
  <ModalBody>
    <ModalHeader hideCloseButton>Verify Your Email</ModalHeader>
    <Label label="Your Email Code">
      <Input
        inputmode="numeric"
        autocapitalize="none"
        autocomplete="one-time-code"
        autofocus
        required
        type="number"
        placeholder="123456"
        on:change={e => {
          code = e.detail
          if (code.length >= 6) {
            handleNextStep()
          }
        }}
      />
    </Label>
  </ModalBody>
  <ModalFooter>
    <Button disabled={isMakingRequest} on:click={handleNextStep}
      >Verify Email</Button
    >
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
</style>
