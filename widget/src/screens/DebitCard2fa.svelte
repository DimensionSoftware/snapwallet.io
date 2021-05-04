<script lang="ts">
  import { push } from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { Routes } from '../constants'
  import { debitCardStore } from '../stores/DebitCardStore'
  import { onMount } from 'svelte'

  let pollTimer
  let cardCode = ''
  let smsCode = ''
  $: smsCodeRequired = false
  $: cardCodeRequired = false
  let submittingAuth = false
  // Wyre may never require auth codes
  // but we have no way of knowing if they were sent
  // so we wait for the timeout and then proceed
  let verificationWaitTimeout = 10_000

  const handleNextStep = async () => {
    try {
      submittingAuth = true
      await window.API.fluxWyreSubmitDebitCardAuthorizations({
        reservationId: $debitCardStore.reservationId,
        orderId: $debitCardStore.orderId,
        sms2faCode: smsCode,
        card2faCode: cardCode,
      })
      push(Routes.SUCCESS)
    } finally {
      submittingAuth = false
    }
  }

  const fetchAuthorizations = async () => {
    const {
      card2faNeeded,
      smsNeeded,
    } = await window.API.fluxWyreGetDebitCardAuthorizations(
      $debitCardStore.orderId,
    )
    smsCodeRequired = smsNeeded
    cardCodeRequired = card2faNeeded
  }

  const pollAuthorizations = () => {
    const t = setInterval(() => {
      // Only one of these may be required
      if (!smsCodeRequired || !cardCodeRequired) {
        fetchAuthorizations()
      } else {
        clearInterval(t)
      }
    }, 4000)
    return t
  }

  const authorizationDoneWaitingTimer = () => {
    const t = setTimeout(() => {
      // Only one of these may be required
      if (!smsCodeRequired || !cardCodeRequired) {
        return push(Routes.SUCCESS)
      }
      clearTimeout(t)
    }, verificationWaitTimeout)
    return t
  }

  onMount(() => {
    const authzTimer = authorizationDoneWaitingTimer()
    pollTimer = pollAuthorizations()
    return () => {
      clearInterval(pollTimer)
      clearTimeout(authzTimer)
    }
  })
</script>

<ModalContent>
  <ModalHeader>Card Authorization</ModalHeader>
  <ModalBody>
    {#if smsCodeRequired}
      <Label label="SMS Code">
        <Input
          id="autocomplete"
          defaultValue={smsCode}
          placeholder="123456"
          on:change={e => (smsCode = e?.detail)}
        />
      </Label>
    {/if}
    {#if cardCodeRequired}
      <Label label="Card Code">
        <Input
          id="autocomplete"
          defaultValue={cardCode}
          placeholder="123456"
          on:change={e => (cardCode = e?.detail)}
        />
      </Label>
    {/if}
    <!-- TODO: add an animation or something-->
    {#if !smsCodeRequired && !cardCodeRequired}
      Authorizing card...
      <br />
      This can take up to 1 minute.
    {/if}
  </ModalBody>
  <ModalFooter>
    <Button
      disabled={!cardCodeRequired && !smsCodeRequired}
      isLoading={submittingAuth}
      on:mousedown={handleNextStep}
      >{submittingAuth ? 'Buying...' : 'Buy Now'}</Button
    >
  </ModalFooter>
</ModalContent>

<style lang="scss">
</style>
