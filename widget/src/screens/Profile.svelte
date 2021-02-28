<script lang="ts">
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Label from '../components/inputs/Label.svelte'
  import Input from '../components/inputs/Input.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { userStore } from '../stores/UserStore'

  let animation = 'left'
</script>

<ModalContent {animation}>
  <ModalBody>
    <ModalHeader>Personal Details</ModalHeader>
    <Label label="Name">
      <Input
        inputmode="text"
        autocapitalize="true"
        autocomplete="on"
        autofocus
        type="text"
        placeholder="First Name"
        defaultValue={$userStore.firstName}
        on:change={e => {
          userStore.setFirstName(e.detail)
        }}
      />
    </Label>
    <Label label="Family Name">
    <Input
      inputmode="text"
      autocapitalize="true"
      autocomplete="on"
      type="text"
      placeholder="Last Name"
      defaultValue={$userStore.lastName}
      on:change={e => {
        userStore.setLastName(e.detail)
      }}
    />
    </Label>
    <Label label="Birthdate">
    <Input
      inputmode="text"
      autocapitalize="true"
      autocomplete="bday"
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
    <Button>Continue</Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
</style>
