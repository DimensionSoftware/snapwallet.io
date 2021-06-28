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
  import {
    Logger,
    onEnterPressed,
    focus,
    resizeWidget,
    focusFirstInput,
  } from '../util'
  import { Routes } from '../constants'
  import { Masks } from '../types'
  import { unMaskValue } from '../masks'
  import { configStore } from '../stores/ConfigStore'
  import CountrySelector from '../components/selectors/CountrySelector.svelte'
  import PhoneInput from '../components/inputs/PhoneInput.svelte'

  let countrySelectorVisible = false
  $: isUSPhoneNumber = $userStore.phoneNumberCountry.code.toUpperCase() === 'US'

  onMount(() => {
    resizeWidget(400, $configStore.appName)
  })

  export let phoneVerificationOnly: boolean = false

  let isMakingRequest = false
  let isUsingPhoneNumber = false

  const timeout = 700

  const handleNextStep = async () => {
    if (!phoneVerificationOnly && !isUsingPhoneNumber) {
      let emailIsValid = vld8.isEmail($userStore.emailAddress)
      if (!emailIsValid) {
        focus(document.querySelector('input[type="email"]'))
        throw new Error('Enter a valid email address.')
      }
    } else {
      const rawPhone =
        $userStore.phoneNumberCountry.dial_code +
        unMaskValue($userStore.phoneNumber, Masks.PHONE)
      let isPhoneValid = vld8.isMobilePhone(rawPhone)
      if (!isPhoneValid) {
        focus(document.querySelector('input[type="tel"]'))
        throw new Error('Enter a valid phone number.')
      }
    }

    try {
      isMakingRequest = true
      await window.API.fluxOneTimePasscode({
        emailOrPhone:
          phoneVerificationOnly || isUsingPhoneNumber
            ? `${
                $userStore.phoneNumberCountry.dial_code
              }${$userStore.phoneNumber.replace(/[^0-9]/g, '')}`
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
      Verify Identity
    {/if}
  </ModalHeader>
  <ModalBody>
    {#if configStore.environment === 'sandbox'}
      <small title="Testing in Sandbox Mode">Test Mode</small>
    {/if}
    {#if !phoneVerificationOnly && (!$userStore.flags?.hasEmail || !isUsingPhoneNumber)}
      <div class="email" in:fade={{ duration: 300 }}>
        <h3>We'll email a 6-digit code to secure you.</h3>
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
        <h3>We'll send a 6-digit code to secure you.</h3>
        <Label label="Your Phone Number">
          <PhoneInput
            on:select={() => (countrySelectorVisible = true)}
            inputmode="phone"
            autocapitalize="none"
            autocomplete="on"
            autofocus={!countrySelectorVisible}
            required
            type="tel"
            mask={isUSPhoneNumber ? Masks.PHONE : undefined}
            placeholder={$userStore.virtual.phone
              ? $userStore.virtual.phone
              : isUSPhoneNumber
              ? '222 333-4444'
              : '222333444'}
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
    <Button
      glow={$userStore.emailAddress.indexOf('@') > -1}
      isLoading={isMakingRequest}
      on:mousedown={handleNextStep}
    >
      {#if isMakingRequest}
        Sending...
      {:else}
        Get Code
      {/if}
    </Button>
  </ModalFooter>
</ModalContent>

{#if countrySelectorVisible}
  <CountrySelector
    visible
    selectedCountryCode={$userStore.phoneNumberCountry?.code ||
      $userStore.address.country ||
      $userStore.geo.country}
    on:close={() => {
      countrySelectorVisible = false
      focusFirstInput()
    }}
    on:select={e => {
      const { country } = e?.detail
      country && userStore.setPhoneNumberCountry(country)
      countrySelectorVisible = false
      focusFirstInput()
    }}
  />
{/if}

<style lang="scss">
  @import '../styles/_vars.scss';
  .email,
  .phone {
    margin: 7% 0.5rem 0 0.5rem;
  }
  .link {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 0.8rem;
  }
  h3 {
    color: var(--theme-text-color);
    opacity: 0.8;
    font-size: 0.85rem;
    font-weight: 400;
    margin: 0 0 1rem 0;
  }
</style>
