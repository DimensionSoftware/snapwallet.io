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
  import { onEnterPressed } from '../util'
  import { Routes } from '../constants'
  import { Masks } from '../types'

  import { push } from 'svelte-spa-router'

  let animation = 'left'

  $: fullName = `${$userStore.firstName} ${$userStore.lastName}`.trim()
  let isSaving = false

  const handleNextStep = async () => {
    try {
      isSaving = true
      const {
          firstName,
          lastName,
          birthDate,
          socialSecurityNumber,
        } = $userStore,
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

      const [mm, dd, yyyy] = birthDate.split('-')

      await window.API.fluxSaveProfileData({
        ssn: socialSecurityNumber,
        dateOfBirth: `${yyyy}-${mm}-${dd}`,
        // TODO: capture fullname somewhere for full accuracy? or reprocessing later?
        legalName: `${firstName} ${lastName}`,
      })
      setTimeout(() => push(Routes.ADDRESS), 1000)
    } finally {
      setTimeout(() => (isSaving = false), 1000)
    }
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
        defaultValue={fullName}
        pattern={`[\\w\\s]+`}
        on:change={e => {
          const { firstName, lastName } = human.parseName(e.detail)
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
        placeholder="mm-dd-yyyy"
        pattern={`[\\d]{2}-[\\d]{2}-[\\d]{4}`}
        mask={Masks.US_DATE}
        defaultValue={$userStore.birthDate}
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
        pattern={`[\\d]{3}-[\\d]{2}-[\\d]{4}`}
        mask={Masks.SSN}
        defaultValue={$userStore.socialSecurityNumber}
        on:change={e => {
          userStore.setSocialSecurityNumber(e.detail)
        }}
      />
    </Label>
  </ModalBody>
  <ModalFooter>
    <Button disabled={isSaving} isLoading={isSaving} on:click={handleNextStep}
      >{isSaving ? 'Saving...' : 'Save'}</Button
    >
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
</style>
