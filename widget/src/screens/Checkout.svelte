<script lang="ts">
  import {pop} from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'

</script>

<ModalContent>
  <svg on:click={pop} viewBox="0 0 20 11" fill="none" class="back-arrow"><path fill-rule="evenodd" clip-rule="evenodd" d="M20 5.267L14.343 0l-1.414 1.217 3.243 3.01H0v1.958h16.172L12.929 9.53 14.343 11 20 5.267z" fill="black"></path></svg>
  <ModalBody>
    <h3>Enter Your Email</h3>
    <Input inputmode="email" autocapitalize="none" autocomplete="on" type="email" placeholder="your@email.address" />
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

  .back-arrow {
    height: 10px;
    transform: rotate(180deg);
    cursor: pointer;
    position: absolute;
    left: 0;
    top: 10px;
  }
</style>
