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
  let fetchingAuth = true
  let cardCode = ''
  let smsCode = ''
  let smsCodeRequired = false
  let cardCodeRequired = false

  const handleNextStep = async () => {
    await window.API.fluxWyreSubmitDebitCardAuthorizations({
      reservationId: $debitCardStore.reservationId,
      orderId: $debitCardStore.orderId,
      sms2faCode: smsCode,
      card2faCode: cardCode,
    })
    push(Routes.SUCCESS)
  }

  const fetchAuthorizations = async () => {
    try {
      const {
        card2faNeeded,
        smsNeeded,
      } = await window.API.fluxWyreGetDebitCardAuthorizations(
        $debitCardStore.orderId,
      )
      smsCodeRequired = smsNeeded
      cardCodeRequired = card2faNeeded
      fetchingAuth = false
    } finally {
      clearInterval(pollTimer)
    }
  }

  const pollAuthorizations = () => {
    return setInterval(fetchAuthorizations, 4000)
  }

  onMount(() => {
    if (fetchingAuth) {
      pollTimer = pollAuthorizations()
      return () => clearInterval(pollTimer)
    }
  })
</script>

<ModalContent>
  <ModalHeader>Card Authorization</ModalHeader>
  <ModalBody>
    {#if fetchingAuth}
      Retrieving required authorizations...
    {:else if smsCodeRequired}
      <Label label="SMS Code">
        <Input
          id="autocomplete"
          defaultValue={smsCode}
          placeholder="123456"
          on:change={e => (smsCode = e?.detail)}
        />
      </Label>
    {:else if cardCodeRequired}
      <Label label="Card Code">
        <Input
          id="autocomplete"
          defaultValue={cardCode}
          placeholder="123456"
          on:change={e => (cardCode = e?.detail)}
        />
      </Label>
    {/if}
  </ModalBody>
  <ModalFooter>
    <Button on:mousedown={handleNextStep}>Continue</Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
</style>
