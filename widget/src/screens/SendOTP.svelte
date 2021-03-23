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
  import { Routes } from '../constants'
  import { Masks } from '../types'
  import { unMaskValue } from '../masks'

  let animation = 'left'
  let isMakingRequest = false
  let isUsingPhoneNumber = false

  const timeout = 700

  const handleNextStep = async () => {
    try {
      if (!isUsingPhoneNumber) {
        let emailIsValid = vld8.isEmail($userStore.emailAddress)
        if (!emailIsValid)
          return (document.querySelector('input[type="email"]') as any).focus()
      } else {
        const rawPhone = unMaskValue($userStore.phoneNumber, Masks.PHONE)
        let isPhoneValid = vld8.isMobilePhone(rawPhone)
        if (!isPhoneValid)
          return (document.querySelector('input[type="tel"]') as any).focus()
      }

      isMakingRequest = true
      await window.API.fluxOneTimePasscode({
        emailOrPhone: isUsingPhoneNumber
          ? `+${unMaskValue($userStore.phoneNumber, Masks.PHONE)}`
          : $userStore.emailAddress,
      })

      setTimeout(() => push(Routes.VERIFY_OTP), timeout)
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
    <ModalHeader hideBackButton>Login or Sign Up</ModalHeader>
    {#if !isUsingPhoneNumber}
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
      <div class="link">
        <a
          on:click={() => {
            isUsingPhoneNumber = true
            userStore.setEmailAddress('')
          }}>Use my phone</a
        >
      </div>
    {:else}
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
      <div class="link">
        <a
          on:click={() => {
            isUsingPhoneNumber = false
            userStore.setPhoneNumber('')
          }}>Use my email</a
        >
      </div>
    {/if}
  </ModalBody>
  <ModalFooter>
    <Button isLoading={isMakingRequest} on:click={handleNextStep}>
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
  .link {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    font-size: 0.75rem;
  }
</style>
