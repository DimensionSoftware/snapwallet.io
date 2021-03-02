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
  import { Logger, onEnterPressed, setFluxSession } from '../util'
  import { toaster } from '../stores/ToastStore'

  let animation = 'left'
  let code = ''
  let isMakingRequest = false
  let isSendingEmail = false

  const resendEmail = async () => {
    Logger.debug('Resending email')
    isSendingEmail = true
    return await window.API().fluxOneTimePasscode({
      emailOrPhone: $userStore.emailAddress,
    })
  }

  const verifyOTP = async () => {
    Logger.debug('Verifying using OTP code:', code)
    return await window.API().fluxOneTimePasscodeVerify({
      code,
      emailOrPhone: $userStore.emailAddress,
    })
  }

  const handleResend = async () => {
    isMakingRequest = true

    try {
      await resendEmail()
      Logger.debug('Email sent')
      setTimeout(() => {
        code = ''
        toaster.pop({
          msg: 'Success! Please check your email inbox.',
          success: true,
        })
      }, 600)
    } catch (e) {
      Logger.error(e)
      // TODO: move error messages to the server
      let msg = 'An unknown error occurred. Please try again later.'
      const code = e.body?.code

      if ([3].includes(code)) {
        msg = 'The email provided was invalid. Please re-enter and try again.'
      }

      toaster.pop({
        msg,
        error: true,
      })
      setTimeout(() => {
        push('#/checkout')
      }, 1700)
    } finally {
      setTimeout(() => {
        isSendingEmail = false
        isMakingRequest = false
      }, 1000)
    }
  }

  const handleNextStep = async () => {
    isMakingRequest = true

    try {
      const { jwt } = await verifyOTP()
      Logger.debug('Logged in')
      setFluxSession(jwt)
      setTimeout(() => push('#/profile'), 700)
    } catch (e) {
      const err = e as { body: { code: number; message: string } }
      Logger.error(err)

      toaster.pop({
        msg: err.body.message,
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
    <ModalHeader>Verify Your Email</ModalHeader>
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
    <div class="resend">
      Didn't get an email? <a on:click={handleResend}>Resend Code</a>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button disabled={isMakingRequest} on:click={handleNextStep}>
      {#if isMakingRequest && !isSendingEmail}
        Verifying Code...
      {:else if isMakingRequest && isSendingEmail}
        Resending Email...
      {:else}
        Verify Email
      {/if}
    </Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  .resend {
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 0.8rem;
    margin-top: -1rem;
    & > a {
      margin-left: 0.5rem;
    }
  }
</style>
