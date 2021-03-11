<script lang="ts">
  //@ts-ignore
  import human from 'humanparser'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Label from '../components/inputs/Label.svelte'
  import Input from '../components/inputs/Input.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { userStore } from '../stores/UserStore'
  import { Logger, onEnterPressed } from '../util'
  import { toaster } from '../stores/ToastStore'
  import { Routes } from '../constants'

  import { push } from 'svelte-spa-router'

  let animation = 'left'

  $: fullName = `${$userStore.firstName} ${$userStore.lastName}`.trim()
  const defaultName = `${$userStore.firstName} ${$userStore.lastName}`.trim()

  const handleNextStep = () => {
    const { firstName, lastName, birthDate, socialSecurityNumber } = $userStore,
      focus = (ndx: number) => {
        const thingie = document.querySelectorAll('input[type="text"]')[
          ndx
        ] as any
        thingie.focus()
      }

    // validate inputs
    if (!firstName || !lastName?.length) return focus(0)
    if (!birthDate) return focus(1)
    if (!socialSecurityNumber) return focus(2)

    window.API.fluxSaveProfileData({
      ssn: socialSecurityNumber,
      dateOfBirth: birthDate,
      // TODO: capture fullname somewhere for full accuracy? or reprocessing later?
      legalName: `${firstName} ${lastName}`,
    }).then(() => {
      push(Routes.CHECKOUT_OVERVIEW)
    })
  }

  const onKeyDown = (e: Event) => {
    onEnterPressed(e, handleNextStep)
  }
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent {animation}>
  <ModalBody>
    <ModalHeader>Tell Us About You</ModalHeader>
    <Label label={fullName || 'Full Name'}>
      <Input
        inputmode="text"
        autocapitalize="true"
        autocomplete="on"
        required
        autofocus
        type="text"
        placeholder="Your Full Name"
        defaultValue={defaultName}
        pattern={`[\\w]+\\s`}
        on:change={e => {
          const {firstName, lastName} = human.parseName(e.detail)
          userStore.setFirstName(firstName ?? '')
          userStore.setLastName(lastName ?? '')
        }}
      />
    </Label>
    <Label label="Birthdate">
      <Input
        inputmode="text"
        autocapitalize="true"
        autocomplete="bday"
        required
        type="text"
        placeholder="yyyy-mm-dd"
        pattern={`[\\d]{4}-[\\d]{2}-[\\d]{2}`}
        maskChar="[\d-]"
        defaultValue={$userStore.birthDate}
        value={$userStore.birthDate}
        on:change={e => {
          userStore.setBirthDate(e.detail)
        }}
      />
    </Label>
    <Label label="Social Security Number">
      <Input
        inputmode="text"
        autocapitalize="true"
        autocomplete="on"
        required
        type="text"
        placeholder="xxx-xx-xxxx"
        maskChar="[\d-]"
        pattern={`[\\d]{3}-[\\d]{2}-[\\d]{4}`}
        defaultValue={$userStore.socialSecurityNumber}
        value={$userStore.socialSecurityNumber}
        on:change={e => {
          userStore.setSocialSecurityNumber(e.detail)
        }}
      />
    </Label>
  </ModalBody>
  <ModalFooter>
    <Button on:click={handleNextStep}>Save</Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
</style>
