<script lang="ts">
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import {
    faExclamationCircle,
    faFolder,
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
  import { groupRemediations, reducePersonalInfoFields } from '../util/profiles'

  let isPersonalInfoError = false
  let personalInfoMessage =
    'Personal identity information used for verification purposes.'

  let isAddressError = false
  let addressMessage =
    'Up to date residential address information used for identity verification.'

  $: remediationGroups = groupRemediations($userStore.profileRemediations)

  $: {
    isPersonalInfoError = remediationGroups.personal.length > 0
    if (isPersonalInfoError)
      personalInfoMessage = reducePersonalInfoFields(remediationGroups.personal)

    isAddressError = remediationGroups.address.length > 0
    if (isAddressError)
      addressMessage =
        'An address update is required. Please provide your current residential address.'
  }

  const getLatestProfile = async () => {
    await userStore.fetchUserProfile()
    remediationGroups = groupRemediations($userStore.profileRemediations)
  }

  onMount(() => {
    getLatestProfile()
  })
</script>

<ModalContent>
  <ModalBody>
    <ModalHeader onBack={() => push(Routes.ROOT)}>Profile</ModalHeader>
    <div style="padding:0 0.5rem;font-weight:bold;">
      <FaIcon scale="3" data={faUserCircle} />
    </div>
    <ul class="vertical-stepper">
      <VStep onClick={() => push(Routes.PROFILE_UPDATE)}>
        <span class:error={isPersonalInfoError} slot="icon">
          <FaIcon data={isPersonalInfoError ? faExclamationCircle : faIdCard} />
        </span>
        <b slot="step">Personal</b>
        <div class="description help" slot="info">
          {personalInfoMessage}
        </div>
      </VStep>
      <VStep onClick={() => push(Routes.ADDRESS_UPDATE)}>
        <span class:error={isAddressError} slot="icon">
          <FaIcon data={isPersonalInfoError ? faExclamationCircle : faHome} />
        </span>
        <b slot="step"> Address </b>
        <div class="description help" slot="info">
          {addressMessage}
        </div>
      </VStep>
      <VStep onClick={() => push(Routes.PROFILE_SEND_SMS)}>
        <span slot="icon">
          <FaIcon data={faMailBulk} />
        </span>
        <b slot="step"> Contact </b>
        <div class="description help" slot="info">
          Contact information used for verification, communication and security.
        </div>
      </VStep>
      <VStep onClick={() => push(Routes.FILE_UPLOAD)}>
        <span slot="icon">
          <FaIcon data={faFolder} />
        </span>
        <b slot="step"> Documents </b>
        <div class="description help" slot="info">
          Documents used for verifying your identity or residence.
        </div>
      </VStep>
    </ul>
  </ModalBody>
  <ModalFooter />
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/text.scss';

  .vertical-stepper {
    margin-top: 0.5rem;
    list-style: none;
    padding: 0 0.5rem;
  }

  .description {
    min-height: 60px;
    margin-left: 0.55rem;
  }
</style>
