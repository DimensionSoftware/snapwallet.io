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
  import { onEnterPressed } from '../util'
  import type {
    FluxApi,
    OneTimePasscodeRequest,
    OneTimePasscodeVerifyResponse,
  } from 'api-client'
  import { linear } from 'svelte/easing'

  let animation = 'left'

  const handleNextStep = () => {
    // validate
    let emailIsValid = vld8.isEmail($userStore.emailAddress)
    if (!emailIsValid) {
      ;(document.querySelector('input[type="email"]') as any).focus()
      return
    }

    ;((window as any).API() as FluxApi)
      .fluxOneTimePasscode({
        emailOrPhone: $userStore.emailAddress,
      })
      .then((resp: any) => {
        // TODO: instead of profile should go to verify otp screen with keypad numeric only enabled (6 digits)
        push('#/profile')
      })
      .catch((resp: any) => {
        // InvalidArgument code 3 (same as http 400)
        if (resp.body.code === 3) {
          // FIXME: bubble up to user in a nice way
          return alert(resp.body.message.match(/desc = (.+)/)[1])
        }
        // unhandled error default
        throw resp
      })
  }

  const onKeyDown = (e: Event) => {
    onEnterPressed(e, handleNextStep)
  }
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent {animation}>
  <ModalBody>
    <ModalHeader hideCloseButton>Welcome</ModalHeader>
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
    <Button on:click={handleNextStep}>Login <small>or</small> SignUp</Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
</style>
