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
  import { toaster } from '../stores/ToastStore'
  import { Routes } from '../constants'

  let animation = 'left'
  let code = ''
  let isMakingRequest = false
  let isSendingEmail = false

  const resendEmail = async () => {
    Logger.debug('Resending email')
    isSendingEmail = true
    return await window.API.fluxOneTimePasscode({
      emailOrPhone: $userStore.emailAddress,
    })
  }

  const verifyOTP = async () => {
    Logger.debug('Verifying using OTP code:', code)
    return await window.API.fluxOneTimePasscodeVerify({
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
        push(Routes.SEND_OTP)
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
      const { tokens } = await verifyOTP()
      window.AUTH_MANAGER.login(tokens)
      Logger.debug('Logged in')
      setTimeout(() => push($userStore.lastKnownRoute), 700)
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
    <ModalHeader>Enter Your Code</ModalHeader>
    <Label label="Your Email Code">
      <Input
        inputmode="numeric"
        autocapitalize="none"
        autocomplete="one-time-code"
        autofocus
        required
        type="number"
        placeholder="123456"
        defaultValue={code}
        on:change={e => {
          code = e.detail
          if (code.length >= 6) {
            handleNextStep()
          }
        }}
      />
    </Label>
    <div class="resend" title="Check your SPAM folder!">
      Didn't get an email? <a on:click={handleResend}>Resend Code</a>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button disabled={isMakingRequest} on:click={handleNextStep}>
      {#if isMakingRequest && !isSendingEmail}
        Confirming Code...
      {:else if isMakingRequest && isSendingEmail}
        Resending Email...
      {:else}
        Confirm Code
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
      margin-left: 0.75em;
    }
  }
</style>
