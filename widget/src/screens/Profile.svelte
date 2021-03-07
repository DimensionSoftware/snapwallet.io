<script lang="ts">
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

  import { push } from 'svelte-spa-router'

  let animation = 'left'

  const defaultName = `${$userStore.firstName} ${$userStore.lastName}`

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

    window
      .API()
      .fluxSaveProfileData({
        ssn: socialSecurityNumber,
        dateOfBirth: birthDate,
        // TODO: capture fullname somewhere for full accuracy? or reprocessing later?
        legalName: `${firstName} ${lastName}`,
      })
      .then(() => {
        push('/overview')
      })
      .catch(e => {
        const err = e as { body: { code: number; message: string } }
        Logger.error(err)

        toaster.pop({
          msg: err.body.message,
          error: true,
        })
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
    <Label label="Full Name">
      <Input
        inputmode="text"
        autocapitalize="true"
        autocomplete="on"
        required
        autofocus
        type="text"
        placeholder="Your Full Name"
        defaultValue={defaultName.trim()}
        on:change={e => {
          userStore.setFullName(e.detail)
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
        placeholder="mm/dd/yyyy"
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
        defaultValue={$userStore.socialSecurityNumber}
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
