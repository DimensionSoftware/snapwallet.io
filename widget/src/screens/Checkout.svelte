<script lang="ts">
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'

  let animation = 'left'
</script>

<ModalContent {animation}>
  <ModalBody>
    <ModalHeader />
    <h3>Enter Your Email</h3>
    <Input
      inputmode="email"
      autocapitalize="none"
      autocomplete="on"
      type="email"
      placeholder="your@email.address"
    />
  </ModalBody>
  <ModalFooter>
    <Button>Continue</Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';

  h3 {
    margin: 25% 0 0;
  }
</style>
