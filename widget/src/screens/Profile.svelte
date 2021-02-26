<script lang="ts">
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'

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
    />
    <Input
      label="Last Name"
      inputmode="text"
      autocapitalize="true"
      autocomplete="on"
      type="text"
      placeholder="Last Name"
    />
    <Input
      label="Birthdate"
      inputmode="text"
      autocapitalize="true"
      autocomplete="on"
      type="text"
      placeholder="Birthdate"
    />
    <Input
      label="Social Security Number"
      inputmode="text"
      autocapitalize="true"
      autocomplete="on"
      type="text"
      placeholder="Social Security Number"
    />
  </ModalBody>
  <ModalFooter>
    <Button>Next</Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
</style>
