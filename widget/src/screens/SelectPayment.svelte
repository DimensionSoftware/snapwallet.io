<script lang="ts">
  import { push } from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import FaIcon from 'svelte-awesome'
  import { faUniversity } from '@fortawesome/free-solid-svg-icons'
  import IconCard from '../components/cards/IconCard.svelte'
  import PlaidBanner from '../components/PlaidBanner.svelte'
  import { Routes } from '../constants'

  const onACHClicked = async () => {
    const { flags = {} } = await window.API.fluxViewerData()
    const { hasWyrePaymentMethods, hasWyreAccount } = flags
    let nextRoute = Routes.PLAID_LINK

    if (hasWyrePaymentMethods && hasWyreAccount)
      nextRoute = Routes.CHECKOUT_OVERVIEW
    else if (hasWyrePaymentMethods) nextRoute = Routes.PROFILE

    push(nextRoute)
  }
</script>

<ModalContent>
  <ModalBody>
    <ModalHeader hideRightAction>Payment Method</ModalHeader>
    <IconCard on:click|once={onACHClicked} label="Bank Account">
      <div class="icon-slot-container" slot="icon">
        <FaIcon data={faUniversity} />
      </div>
    </IconCard>
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
