<script lang="ts">
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import PaymentSelector from '../components/selectors/PaymentSelector.svelte'
  import AccountSelector from '../components/selectors/AccountSelector.svelte'
  import {
    faExclamationCircle,
    faFolder,
    faCheck,
    faHome,
    faIdCard,
    faMailBulk,
    faUserCircle,
  } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import VStep from '../components/VStep.svelte'
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import { onMount } from 'svelte'
  import { userStore } from '../stores/UserStore'
  import {
    groupRemediations,
    reduceDocumentFields,
    reducePersonalInfoFields,
    reduceAddressFields,
    reduceContactFields,
    getMissingFieldMessages,
  } from '../util/profiles'
  import Address from './Address.svelte'

  let paymentSelectorVisible = false
  let step

  $: remediationGroups = groupRemediations($userStore.profileRemediations)

  $: isPersonalInfoError = remediationGroups.personal.length > 0
  $: personalInfoMessage = reducePersonalInfoFields(remediationGroups.personal)

  $: isAddressError = remediationGroups.address.length > 0
  $: addressMessage = reduceAddressFields(remediationGroups.address)

  $: isContactError = remediationGroups.contact.length > 0
  $: contactMessage = reduceContactFields(remediationGroups.contact)

  $: isDocumentError = remediationGroups.document.length > 0
  $: documentMessage = reduceDocumentFields(remediationGroups.document)

  $: missingInfo = getMissingFieldMessages($userStore.profileItems)
  $: step =
    !missingInfo.personal.isComplete || isPersonalInfoError
      ? 'personal'
      : !missingInfo.contact.isComplete || isContactError
      ? 'contact'
      : !missingInfo.address.isComplete || isAddressError
      ? 'address'
      : !missingInfo.document.isComplete || isDocumentError
      ? 'document'
      : 'payment'

  const getLatestProfile = async () => {
    await userStore.fetchUserProfile()
    remediationGroups = groupRemediations($userStore.profileRemediations)
  }

  const pollProfile = () => setInterval(getLatestProfile, 30_000)

  onMount(() => {
    getLatestProfile()
    const interval = pollProfile()
    return () => clearInterval(interval)
  })
</script>

<ModalContent>
  <ModalHeader onBack={() => push(Routes.ROOT)}>Profile</ModalHeader>
  <ModalBody>
    <div style="padding:0 0 0 0.3rem;font-weight:bold;">
      <FaIcon scale="3" data={faUserCircle} />
    </div>
    <ul class="vertical-stepper">
      <VStep
        title="Edit Your Profile"
        onClick={() => push(Routes.PROFILE_UPDATE)}
        success={missingInfo.personal.isComplete && !isPersonalInfoError}
        disabled={step !== 'personal' && !missingInfo.personal.isComplete}
        line
      >
        <span
          class:info={!missingInfo.personal.isValid}
          class:error={isPersonalInfoError}
          class:glow={step === 'personal'}
          slot="icon"
        >
          {#if missingInfo.personal.isComplete && !isPersonalInfoError}
            <FaIcon data={faCheck} />
          {:else if (missingInfo.personal.isComplete && !missingInfo.personal.isValid) || isPersonalInfoError}
            <FaIcon data={faExclamationCircle} />
          {:else}
            <FaIcon data={faIdCard} />
          {/if}
        </span>
        <b slot="step">Identity</b>
        <div class="description help" slot="info">
          {missingInfo.personal.message || personalInfoMessage}
        </div>
      </VStep>
      <!-- <PaymentSelector
        disabled={step !== 'payment'}
        onClick={() => (paymentSelectorVisible = true)}
        description="Payment to buy and sell"
      /> -->
      <VStep
        title="Edit Your Contact"
        onClick={() => push(Routes.PROFILE_SEND_SMS)}
        success={missingInfo.contact.isComplete && !isContactError}
        disabled={step !== 'contact' && !missingInfo.contact.isComplete}
        line
      >
        <span
          class:info={missingInfo.contact.isComplete &&
            !missingInfo.contact.isValid}
          class:error={isContactError}
          class:glow={step === 'contact'}
          slot="icon"
        >
          {#if missingInfo.contact.isComplete && !isContactError}
            <FaIcon data={faCheck} />
          {:else if (missingInfo.contact.isComplete && !missingInfo.contact.isValid) || isContactError}
            <FaIcon data={faExclamationCircle} />
          {:else}
            <FaIcon data={faMailBulk} />
          {/if}
        </span>
        <b slot="step"> Contact </b>
        <div class="description help" slot="info">
          {missingInfo.contact.message || contactMessage}
        </div>
      </VStep>
      <VStep
        title="Edit Your Address"
        onClick={() => push(Routes.ADDRESS_UPDATE)}
        success={missingInfo.address.isComplete && !isAddressError}
        disabled={step !== 'address' && !missingInfo.address.isComplete}
        line
      >
        <span
          class:info={!missingInfo.address.isValid}
          class:error={isAddressError}
          class:glow={step === 'address'}
          slot="icon"
        >
          {#if missingInfo.address.isComplete && !isAddressError}
            <FaIcon data={faCheck} />
          {:else if (missingInfo.address.isComplete && !missingInfo.address.isValid) || isAddressError}
            <FaIcon data={faExclamationCircle} />
          {:else}
            <FaIcon data={faHome} />
          {/if}
        </span>
        <b slot="step"> Address </b>
        <div class="description help" slot="info">
          {missingInfo.address.message || addressMessage}
        </div>
      </VStep>
      <VStep
        title="Edit Your Documents"
        onClick={() => push(Routes.FILE_UPLOAD_UPDATE)}
        success={missingInfo.document.isComplete && !isDocumentError}
        disabled={step !== 'document' && !missingInfo.document.isComplete}
        line
      >
        <span
          class:info={!missingInfo.document.isValid}
          class:error={isDocumentError}
          class:glow={step === 'document'}
          slot="icon"
        >
          {#if missingInfo.document.isComplete && !isDocumentError}
            <FaIcon data={faCheck} />
          {:else if (missingInfo.document.isComplete && !missingInfo.document.isValid) || isDocumentError}
            <FaIcon data={faExclamationCircle} />
          {:else}
            <FaIcon data={faFolder} />
          {/if}
        </span>
        <b slot="step"> Documents </b>
        <div class="description help" slot="info">
          {missingInfo.document.message || documentMessage}
        </div>
      </VStep>
    </ul>
  </ModalBody>
  <ModalFooter />
</ModalContent>

{#if paymentSelectorVisible}
  <AccountSelector visible on:close={() => (paymentSelectorVisible = false)} />
{/if}

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/text.scss';

  .vertical-stepper {
    margin-top: 2rem;
    list-style: none;
    padding: 0 0.5rem;
  }
  .description {
    min-height: 2.5rem;
    margin-left: 1rem;
    opacity: 0.65;
  }
  :global(span.info) {
    &:before {
      opacity: 0.9 !important;
      background: var(--theme-info-color) !important;
    }
    color: var(--text-color) !important;
  }
</style>
