<script lang="ts">
  import { getContext, onMount } from 'svelte'
  import { push } from 'svelte-spa-router'
  import { fade } from 'svelte/transition'
  import { Masks } from '../types'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { userStore } from '../stores/UserStore'
  import { Logger, onEnterPressed, focus, resizeWidget } from '../util'
  import { toaster } from '../stores/ToastStore'
  import { Routes } from '../constants'
  import type { OneTimePasscodeVerifyResponse } from 'api-client'
  import { paymentMethodStore } from '../stores/PaymentMethodStore'
  import { configStore } from '../stores/ConfigStore'

  export let phoneVerificationOnly: boolean = false

  const inputs = [1, 2, 3, 4, 5, 6]
  let isMakingRequest = false
  let isSendingCode = false
  let codes = Array(6)
  let cur = 0
  let code = ''

  onMount(() => {
    resizeWidget(425, configStore.appName)
  })

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
    if (!(code?.length > 5) || !emailOrPhone) {
      // reset
      codes = Array(6)
      code = ''
      cur = 0
      focus(document.getElementById('code-0'))

      toaster.pop({
        msg: 'Check for your code and try again!',
        error: true,
      })

      return
    }

    if (phoneVerificationOnly)
      return window.API.fluxChangeViewerPhone({
        code,
        phone: $userStore.phoneNumber,
      })

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
      }, 250)
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
        let nextRoute = phoneVerificationOnly
          ? Routes.PROFILE_SEND_SMS
          : Routes.SEND_OTP
        push(nextRoute)
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
      let nextRoute = phoneVerificationOnly
        ? Routes.FILE_UPLOAD
        : $userStore.lastKnownRoute

      if (!phoneVerificationOnly) {
        window.AUTH_MANAGER.login(resp.tokens)
        userStore.fetchFlags()
        userStore.fetchUserProfile()
        paymentMethodStore.fetchWyrePaymentMethods()
        userStore.setIsLoggedIn(true)
        Logger.debug('Logged in')
        window.tryInitializePusher()
      }

      setTimeout(() => push(nextRoute), 700)
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

<ModalContent>
  <ModalHeader>Enter Your Code</ModalHeader>
  <ModalBody>
    <div class="code" in:fade={{ duration: 300 }}>
      <Label label="Your Code">
        {#each inputs as input, i}
          <Input
            id={`code-${i}`}
            inputmode="numeric"
            autocapitalize="none"
            autocomplete="one-time-code"
            autofocus={i === 0}
            mask={Masks.CODE}
            size={1}
            required
            type="text"
            placeholder={i + 1}
            maxlength="1"
            autoselect
            on:change={e => {
              const num = e.detail
              if (num.match(/\d/)) {
                // set or replace current code
                cur = i
                codes[cur] = num
                code = codes.reduce((acc, cur) => acc + cur ?? 0, '')
                if (code.length === 6) {
                  handleNextStep()
                } else {
                  // next
                  if (num) cur = cur >= 5 ? 0 : cur + 1
                  document.getElementById(`code-${cur}`)?.focus()
                }
              }
            }}
          />
        {/each}
      </Label>
      <div class="resend" title="Check SPAM">
        Didn't get a code?
        <!-- svelte-ignore a11y-missing-attribute -->
        <a on:click={handleResend}>Resend Code</a>
      </div>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button isLoading={isMakingRequest} on:mousedown={handleNextStep}>
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
  .code {
    :global(label) {
      display: flex;
      flex-direction: row;
      :global(> span.input-label) {
        top: -2rem !important;
        margin-left: 0 !important;
      }
      :global(.input-container) {
        :global(#code-0, #code-1, #code-2, #code-3, #code-4, #code-5) {
          text-align: center !important;
          padding: 25px 0 25px 0 !important;
          border-radius: 0;
          border-right: 1px solid rgba(0, 0, 0, 0.1);
          text-indent: 0;
        }
        :global(#code-0) {
          border-top-left-radius: 0.5rem;
          border-bottom-left-radius: 0.5rem;
        }
        :global(#code-5) {
          border-top-right-radius: 0.5rem;
          border-bottom-right-radius: 0.5rem;
          border-right: 0 solid transparent;
        }
      }
    }
  }

  .code {
    margin: 10% 0.5rem 0 0.5rem;
  }
  .resend {
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 0.95rem;
    margin-top: 0.35rem;
    & > a {
      color: var(--theme-text-color);
      margin-left: 0.25em;
    }
  }
</style>
