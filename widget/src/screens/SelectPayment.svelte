<script lang="ts">
  import { push } from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import FaIcon from 'svelte-awesome'
  import { faUniversity } from '@fortawesome/free-solid-svg-icons'
  import PaymentCard from '../components/cards/PaymentCard.svelte'
</script>

<ModalContent>
  <ModalBody>
    <ModalHeader>Payment Method</ModalHeader>
    <PaymentCard on:click={() => push('#/checkout')} label="Bank Account">
      <div class="icon-slot-container" slot="icon">
        <FaIcon data={faUniversity} />
      </div>
    </PaymentCard>
  </ModalBody>
  <ModalFooter />
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  .icon-slot-container {
    display: flex;
    align-items: center;
    justify-content: center;
  }
</style>
