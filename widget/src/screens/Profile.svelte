<script lang="ts">
  //@ts-ignore
  import human from 'humanparser'
  import { blur } from 'svelte/transition'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Label from '../components/inputs/Label.svelte'
  import Input from '../components/inputs/Input.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { configStore } from '../stores/ConfigStore'
  import { userStore } from '../stores/UserStore'
  import { onEnterPressed, focus as focusElement } from '../util'
  import { Routes } from '../constants'
  import { Masks } from '../types'

  import { push } from 'svelte-spa-router'

  export let isUpdateScreen: boolean = false

  $: fullName = `${$userStore.firstName} ${$userStore.lastName}`.trim()
  let isSaving = false

  const handleNextStep = async () => {
    try {
      isSaving = true
      const { firstName, lastName, birthDate, socialSecurityNumber } =
          $userStore,
        focus = (ndx: number) => {
          const thingie = document.querySelectorAll('input[type="text"]')[
            ndx
          ] as any
          focusElement(thingie)
        }

      const minFieldFilled = [
        firstName,
        lastName,
        birthDate,
        socialSecurityNumber,
      ].join('')

      if (!minFieldFilled) {
        isSaving = false
        focus(0)
        throw new Error('Enter and update your details below.')
      }

      const parsedBirthDate =
          Date.now() - Number(new Date(birthDate.replace(/-/g, '/'))),
        isEighteen = !isNaN(parsedBirthDate) && parsedBirthDate >= 5.676e11

      if (!isEighteen) throw new Error('You must be 18 years of age or older.')

      const [mm, dd, yyyy] = birthDate.split('-')

      await window.API.fluxSaveProfileData({
        ...(socialSecurityNumber && { ssn: socialSecurityNumber }),
        ...(birthDate && { dateOfBirth: `${yyyy}-${mm}-${dd}` }),
        // TODO: capture fullname somewhere for full accuracy? or reprocessing later?
        ...(firstName && lastName && { legalName: `${firstName} ${lastName}` }),
      })
      setTimeout(() => {
        userStore.clearProfile()
        let nextRoute = isUpdateScreen ? Routes.PROFILE_STATUS : Routes.ADDRESS
        push(nextRoute)
      }, 1000)
    } finally {
      setTimeout(() => (isSaving = false), 1000)
    }
  }

  const onKeyDown = (e: Event) => {
    onEnterPressed(e, handleNextStep)
  }

  const fillTestInfo = e => {
    e.preventDefault()
    userStore.setBirthDate('01-01-1980')
    userStore.setFirstName('John')
    userStore.setLastName('Smith')
    userStore.setSocialSecurityNumber('123-12-1234')
  }
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent>
  <ModalHeader>Your Identity</ModalHeader>
  <ModalBody padded>
    {#if $userStore.isProfileComplete}
      <h5 in:blur={{ duration: 300 }}>Identity received and may be updated:</h5>
    {:else if $configStore.environment === 'sandbox'}
      <h3 class="test">
        <a on:click={fillTestInfo} href="">Fill With Test Info</a>
      </h3>
    {:else}
      <h5>&nbsp;</h5>
    {/if}
    <Label label="Full Name">
      <Input
        inputmode="text"
        autocapitalize="true"
        autocomplete="on"
        required
        autofocus
        type="text"
        placeholder={$userStore.virtual?.fullName || 'Your Full Name'}
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
        placeholder={$userStore.virtual?.birthDate || 'mm-dd-yyyy'}
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
        placeholder={$userStore.virtual?.socialSecurityNumber || 'xxx-xx-xxxx'}
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
    <Button
      disabled={isSaving}
      isLoading={isSaving}
      on:mousedown={handleNextStep}>{isSaving ? 'Saving' : 'Save'}</Button
    >
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  h5 {
    margin-top: 0;
    margin-left: 0.25rem;
  }
  .test {
    text-align: left;
  }
</style>
