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
  import { transactionStore } from '../stores/TransactionStore'
  import { onEnterPressed } from '../util'
  import { Routes } from '../constants'

  let animation = 'left'
  let isSubmittingProfile = false

  const handleNextStep = async () => {
    try {
      isSubmittingProfile = true
      await window.API.fluxSaveProfileData({
        address: $userStore.address,
      })
      const nextRoute = $transactionStore.sourceAmount
        ? Routes.CHECKOUT_OVERVIEW
        : Routes.ROOT
      push(nextRoute)
    } finally {
      setTimeout(() => {
        isSubmittingProfile = false
        userStore.clearAddress()
      }, 800)
    }
  }

  const onKeyDown = (e: Event) => {
    onEnterPressed(e, handleNextStep)
  }
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent {animation}>
  <ModalBody>
    <ModalHeader>Where do you live?</ModalHeader>
    <Label label="Country">
      <Input placeholder="Country" defaultValue={$userStore.address.country} />
    </Label>
    <Label label="State">
      <Input placeholder="State" defaultValue={$userStore.address.state} />
    </Label>
    <Label label="Postal Code">
      <Input
        placeholder="Postal Code"
        defaultValue={$userStore.address.postalCode}
      />
    </Label>
  </ModalBody>
  <ModalFooter>
    <Button on:click={handleNextStep}
      >{isSubmittingProfile ? 'Saving...' : 'Save'}</Button
    >
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  label {
    size: smaller;
  }
</style>
