<script lang="ts">
  import { fade } from 'svelte/transition'
  import { onMount } from 'svelte'
  import { push } from 'svelte-spa-router'
  import vld8 from 'validator'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { userStore } from '../stores/UserStore'
  import { toaster } from '../stores/ToastStore'
  import { Logger, onEnterPressed, focus, resizeWidget } from '../util'
  import { Routes } from '../constants'
  import { Masks } from '../types'
  import { unMaskValue } from '../masks'
  import { configStore } from '../stores/ConfigStore'

  onMount(() => {
    resizeWidget(425, $configStore.appName)
  })

  export let phoneVerificationOnly: boolean = false

  let isMakingRequest = false
  let isUsingPhoneNumber = false

  const timeout = 700

  const handleNextStep = async () => {
    try {
      if (!phoneVerificationOnly && !isUsingPhoneNumber) {
        let emailIsValid = vld8.isEmail($userStore.emailAddress)
        if (!emailIsValid)
          return focus(document.querySelector('input[type="email"]'))
      } else {
        const rawPhone = unMaskValue($userStore.phoneNumber, Masks.PHONE)
        let isPhoneValid = vld8.isMobilePhone(rawPhone)
        if (!isPhoneValid)
          return focus(document.querySelector('input[type="tel"]'))
      }

      isMakingRequest = true
      await window.API.fluxOneTimePasscode({
        emailOrPhone:
          phoneVerificationOnly || isUsingPhoneNumber
            ? `+${unMaskValue($userStore.phoneNumber, Masks.PHONE)}`
            : $userStore.emailAddress,
      })

      let nextRoute = phoneVerificationOnly
        ? Routes.PROFILE_VERIFY_SMS
        : Routes.VERIFY_OTP
      setTimeout(() => push(nextRoute), timeout)
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

<ModalContent>
  <ModalHeader>
    {#if phoneVerificationOnly}
      Phone Verification
    {:else}
      Login or Sign Up
    {/if}
  </ModalHeader>
  <ModalBody>
    {#if !phoneVerificationOnly && (!$userStore.flags?.hasEmail || !isUsingPhoneNumber)}
      <div class="email" in:fade={{ duration: 300 }}>
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
        {#if $userStore.flags?.hasEmail}
          <div class="link">
            <a
              on:mousedown={() => {
                isUsingPhoneNumber = true
                // clear so verify doesn't use this value
                userStore.setEmailAddress('')
              }}>Use my phone</a
            >
          </div>
        {/if}
      </div>
    {:else}
      <div class="phone" in:fade={{ duration: 300 }}>
        <Label label="Your Phone Number">
          <Input
            inputmode="phone"
            autocapitalize="none"
            autocomplete="on"
            autofocus
            required
            type="tel"
            mask={Masks.PHONE}
            placeholder="1 (222) 333-4444"
            defaultValue={$userStore.phoneNumber}
            on:change={e => {
              userStore.setPhoneNumber(e.detail)
            }}
          />
        </Label>
        {#if $userStore.flags?.hasPhone && !$userStore.flags?.hasEmail}
          <div class="link">
            <a
              on:mousedown={() => {
                isUsingPhoneNumber = false
                // clear so verify doesn't use this value
                userStore.setPhoneNumber('')
              }}>Use my email</a
            >
          </div>
        {/if}
      </div>
    {/if}
  </ModalBody>
  <ModalFooter>
    <Button isLoading={isMakingRequest} on:mousedown={handleNextStep}>
      {#if isMakingRequest}
        Sending Code...
      {:else}
        Get Code
      {/if}
    </Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  .email,
  .phone {
    margin: 10% 0.5rem 0 0.5rem;
  }
  .link {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 0.8rem;
  }
</style>
