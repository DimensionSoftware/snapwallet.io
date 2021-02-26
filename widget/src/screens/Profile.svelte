<script lang="ts">
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { userStore } from '../stores/UserStore'

  let animation = 'left'
</script>

<ModalContent {animation}>
  <ModalBody>
    <ModalHeader>Personal Details</ModalHeader>
    <Input
      label="First Name"
      inputmode="text"
      autocapitalize="true"
      autocomplete="on"
      type="text"
      placeholder="First Name"
      on:change={e => {
        userStore.setFirstName(e.detail)
      }}
    />
    <Input
      label="Last Name"
      inputmode="text"
      autocapitalize="true"
      autocomplete="on"
      type="text"
      placeholder="Last Name"
      on:change={e => {
        userStore.setLastName(e.detail)
      }}
    />
    <Input
      label="Birthdate"
      inputmode="text"
      autocapitalize="true"
      autocomplete="on"
      type="text"
      placeholder="Birthdate"
      on:change={e => {
        userStore.setBirthDate(e.detail)
      }}
    />
    <Input
      label="Social Security Number"
      inputmode="text"
      autocapitalize="true"
      autocomplete="on"
      type="text"
      placeholder="Social Security Number"
      on:change={e => {
        userStore.setSocialSecurityNumber(e.detail)
      }}
    />
  </ModalBody>
  <ModalFooter>
    <Button>Next</Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
</style>
