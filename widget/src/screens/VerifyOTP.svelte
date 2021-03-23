<script lang="ts">
  import { push } from 'svelte-spa-router'
  import { fly } from 'svelte/transition'
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
import type { OneTimePasscodeVerifyResponse } from 'api-client';

  let animation = 'left'
  let code = ''
  let isMakingRequest = false
  let isSendingCode = false

  const resendCode = async () => {
    Logger.debug('Resending email')
    isSendingCode = true
    return await window.API.fluxOneTimePasscode({
      emailOrPhone: $userStore.phoneNumber || $userStore.emailAddress,
    })
  }

  const verifyOTP = async (): Promise<OneTimePasscodeVerifyResponse> => {
    Logger.debug('Verifying using OTP code:', code)
    const emailOrPhone = $userStore.phoneNumber || $userStore.emailAddress
    if (!emailOrPhone) {
      document.getElementById('code').focus()

      toaster.pop({
        msg: 'Check for your code and try again!',
        error: true,
      })

      return
    }

    return window.API.fluxOneTimePasscodeVerify({
      code,
      emailOrPhone,
    })
  }

  const handleResend = async () => {
    isMakingRequest = true

    try {
      await resendCode()
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
        isSendingCode = false
        isMakingRequest = false
      }, 1000)
    }
  }

  const handleNextStep = async () => {
    isMakingRequest = true

    try {
      const resp = await verifyOTP()
      window.AUTH_MANAGER.login(resp.tokens)
      Logger.debug('Logged in')
      window.tryInitializePusher()
      setTimeout(() => push($userStore.lastKnownRoute), 700)
    } catch (e) {
      if (e.body?.code) {
        toaster.pop({ msg: e.body?.message, error: true })
        setTimeout(() => {
          push(Routes.SEND_OTP)
        }, 800)
      }
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
    <div in:fly={{ y: 25, duration: 300 }}>
      <Label label="Your Code">
        <Input
          id="code"
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
      <div class="resend" title="Check SPAM">

        Didn't get a code?
        <!-- svelte-ignore a11y-missing-attribute -->
        <a on:click={handleResend}>Resend Code</a>
      </div>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button disabled={isMakingRequest} on:click={handleNextStep}>
      {#if isMakingRequest && !isSendingCode}
        Confirming...
      {:else if isMakingRequest && isSendingCode}
        Resending...
      {:else}
        Confirm
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
    margin-top: 0.35rem;
    & > a {
      margin-left: 0.25em;
    }
  }
</style>
