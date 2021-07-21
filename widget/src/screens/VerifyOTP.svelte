<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
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
  import {
    Logger,
    onEnterPressed,
    focus,
    resizeWidget,
    focusFirstInput,
  } from '../util'
  import { toaster } from '../stores/ToastStore'
  import { Routes } from '../constants'
  import type { OneTimePasscodeVerifyResponse } from 'api-client'
  import { paymentMethodStore } from '../stores/PaymentMethodStore'
  import { configStore } from '../stores/ConfigStore'

  export let phoneVerificationOnly: boolean = false

  const inputs = [1, 2, 3, 4, 5, 6]
  let isMakingRequest = false
  let isSendingCode = false
  let code = ''

  $: codes = Array(6).fill('')
  $: cur = 0

  onMount(() => {
    resizeWidget(400, $configStore.appName)
    window.addEventListener('paste', handlePaste)

    return () => {
      // cleanup
      window.removeEventListener('paste', handlePaste)
    }
  })

  const handlePaste = e => {
    e.preventDefault()
    const numString = e.clipboardData.getData('Text').slice(0, 6)
    numString.split('').forEach((n, idx) => {
      document.getElementById(`code-${idx}`).value = n
      codes[idx] = n
    })
    code = codes.join('')
    cur = codes.length > 0 ? codes.length - 1 : 0
    focus(document.getElementById(`code-${cur}`))
    if (codes.length >= 6) {
      // submit
      handleNextStep()
    }
  }

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
      toaster.pop({
        msg: 'Please provide a valid email address and try again.',
        error: true,
      })

      setTimeout(() => {
        codes = Array(6).fill('')
        code = codes.join('')
        push(Routes.SEND_OTP)
      }, 800)

      return
    }

    if (!(code?.length > 5)) {
      toaster.pop({
        msg: 'Please check your code and try again.',
        error: true,
      })
      return focus(document.querySelector('input[type="text"]:first-child'))
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
      for (let cur = 0; cur < 6; cur++) {
        document.getElementById(`code-${cur}`).value = codes[cur] = ''
      }
      code = ''
      setTimeout(() => {
        code = ''
        toaster.pop({
          msg: 'Success! Please check your email inbox.',
          success: true,
        })
        focusFirstInput()
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
      <Label on:click={focusFirstInput} label="Your Code" />
      <div class="row">
        {#each inputs as input, i}
          <Input
            id={`code-${i}`}
            inputmode="numeric"
            autocapitalize="none"
            autocomplete="off"
            autofocus={i === 0}
            mask={Masks.CODE}
            size={1}
            required
            type="text"
            maxlength="1"
            autoselect
            value={codes[i]}
            on:focus={() => {
              cur = i
            }}
            on:keydown={e => {
              if (isSendingCode || isMakingRequest) {
                return e.preventDefault()
              }
              if (e.keyCode === 8) {
                e.preventDefault()
                document.getElementById(`code-${cur}`).value = codes[i] = ''
                code = codes.join('')
                // backspace over input
                cur = cur <= 0 ? 0 : cur - 1
                const el = document.getElementById(`code-${cur}`)
                el?.focus()
              } else if ([38, 40].includes(e.keyCode)) {
                // up/down arrows
                cur = i
                const el = document.getElementById(`code-${cur}`)
                let v = parseInt(el.value) || 0
                if (e.keyCode == 38) v = v >= 9 ? 9 : v + 1
                if (e.keyCode == 40) v = v <= 0 ? 0 : v - 1
                el.value = v
              } else if (e.keyCode === 37) {
                // Left arrow
                cur = cur > 0 ? cur - 1 : 0
                const el = document.getElementById(`code-${cur}`)
                el?.focus()
              } else if (e.keyCode === 39) {
                // Right arrow
                cur = cur < 5 ? cur + 1 : 5
                const el = document.getElementById(`code-${cur}`)
                el?.focus()
              }
            }}
            on:change={e => {
              const num = e.detail
              codes[i] = num
              code = codes.join('')

              cur = !num ? i : i >= 5 ? 5 : i + 1
              document.getElementById(`code-${cur}`)?.focus()

              if (code.length === 6) {
                handleNextStep()
              }
            }}
          />
        {/each}
      </div>
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
    :global(.input-container:after) {
      content: '';
      position: absolute;
      right: -0.05rem;
      top: 45%;
      bottom: 20px;
      z-index: 10;
      width: 1px;
      height: 20px;
      opacity: 0.25;
      background: var(--theme-text-color-3);
    }
    :global(.input-container:last-child:after) {
      display: none;
    }
    :global(label) {
      display: flex;
      flex-direction: row;
      margin-bottom: 0 !important;
    }
    :global(.code label > span.input-label) {
      top: -1.75rem !important;
      margin-left: 0 !important;
    }
    :global(#code-0, #code-1, #code-2, #code-3, #code-4, #code-5) {
      text-align: center !important;
      padding: 25px 0 25px 0 !important;
      border-radius: 0;
      text-indent: 0;
      &:focus,
      &:hover {
        z-index: 9;
      }
    }
    :global(#code-0) {
      border-top-left-radius: 0.5rem;
      border-bottom-left-radius: 0.5rem;
    }
    :global(#code-5) {
      border-top-right-radius: 0.5rem;
      border-bottom-right-radius: 0.5rem;
    }
  }

  .row {
    display: flex;
    flex-direction: row;
  }
  .code {
    margin: 10% 0.5rem 0 0.5rem;
  }
  .resend {
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 0.95rem;
    margin-top: 1.25rem;
    & > a {
      color: var(--theme-text-color);
      margin-left: 0.25em;
    }
  }
</style>
