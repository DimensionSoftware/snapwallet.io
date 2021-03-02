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
  import { onEnterPressed } from '../util'
  import type {
    OneTimePasscodeVerifyResponse,
    ResponseContext,
  } from 'api-client'
  import { handle_promise } from 'svelte/internal'

  function initializePlaid() {
    const handler = window.Plaid.create({
      token: 'GENERATED_LINK_TOKEN',
      environment: 'sandbox',
      onSuccess: (public_token, metadata) => {},
      onLoad: () => {},
      onExit: (err, metadata) => {},
      onEvent: (eventName, metadata) => {},
      receivedRedirectUri: null,
    })
    /* handler.destroy() <-- cleanup function for plaid */
  }

  /*
  let animation = 'left'
  let code = ''

  const handleNextStep = () => {
    const c = code
    console.log('Verifying using OTP code:', c)
    window
      .API()
      .fluxOneTimePasscodeVerify({
        emailOrPhone: $userStore.emailAddress,
        code: c,
      })
      .then((resp: OneTimePasscodeVerifyResponse) => {
        // login (update jwt)
        window.API(resp.jwt)

        // TODO: use returned user data to update store if necessary
        console.log('LOGGED IN:', resp.user)
        push('#/profile')
      })
      .catch((resp: { body: { code: number; message: string } }) => {
        // InvalidArgument code 3 (same as http 400), or Unauthenticated code 16 (same ass http 401)
        // FIXME: bubble up to user in a nice way
        if (resp.body.code === 3) {
          return alert(resp.body.message.match(/desc = (.+)/)[1])
        } else if (resp.body.code == 16) {
          return alert(resp.body.message)
        }

        // unhandled error default
        throw resp
      })
  }
  */
</script>

<svelte:head>
  <script
    src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"
    on:load={initializePlaid}></script>
</svelte:head>

<style lang="scss">
  @import '../styles/_vars.scss';
</style>
