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
  let submittingAuth = false

  $: smsCodeRequired = false
  $: cardCodeRequired = false
  $: isOneCodeRequired = Boolean(smsCodeRequired) || Boolean(cardCodeRequired)

  const handleNextStep = async () => {
    try {
      submittingAuth = true
      await window.API.fluxWyreSubmitDebitCardAuthorizations({
        reservationId: $debitCardStore.reservationId,
        orderId: $debitCardStore.orderId,
        sms2faCode: smsCode,
        card2faCode: cardCode,
      })
      // Make sure screens do not read this data
      // for other types of future transfers
      debitCardStore.clear()
      push(Routes.SUCCESS)
    } finally {
      submittingAuth = false
    }
  }

  /**
   * Fetch Wyre debit card authorization codes
   * Both SMS and Card (micro deposit) codes may be required.
   */
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

  /**
   * Fetch authorizations regularly
   * until either the overall authz timeout
   * is met or codes are required.
   */
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

  onMount(() => {
    pollTimer = pollAuthorizations()
    return () => clearInterval(pollTimer)
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
    {#if !isOneCodeRequired}
      Authorizing card...
    {/if}
  </ModalBody>
  <ModalFooter>
    <Button
      disabled={!isOneCodeRequired}
      isLoading={submittingAuth}
      on:mousedown={handleNextStep}
      >{submittingAuth ? 'Buying...' : 'Buy Now'}</Button
    >
  </ModalFooter>
</ModalContent>

<style lang="scss">
</style>
