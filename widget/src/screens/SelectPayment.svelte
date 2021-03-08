<script lang="ts">
  import { push } from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import FaIcon from 'svelte-awesome'
  import { faUniversity } from '@fortawesome/free-solid-svg-icons'
  import PaymentCard from '../components/cards/PaymentCard.svelte'
  import PlaidBanner from '../components/PlaidBanner.svelte'
  import { Routes } from '../constants'
  import { toaster } from '../stores/ToastStore'

  const onACHClicked = async () => {
    try {
      const { flags = {} } = await window.API().fluxViewerData()
      const { hasPlaidItems, hasWyreAccount } = flags
      let nextRoute = Routes.PLAID_LINK

      if (hasPlaidItems && hasWyreAccount) nextRoute = Routes.CHECKOUT_OVERVIEW
      else if (hasPlaidItems) nextRoute = Routes.PROFILE

      push(nextRoute)
    } catch (e) {
      toaster.pop({
        msg: 'Oops',
        error: true,
      })
    }
  }
</script>

<ModalContent>
  <ModalBody>
    <ModalHeader hideCloseButton>Payment Method</ModalHeader>
    <PaymentCard on:click={onACHClicked} label="Bank Account">
      <div class="icon-slot-container" slot="icon">
        <FaIcon data={faUniversity} />
      </div>
    </PaymentCard>
    <PlaidBanner />
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
