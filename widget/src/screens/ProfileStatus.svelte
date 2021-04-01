<script lang="ts">
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import {
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
</script>

<ModalContent>
  <ModalBody>
    <ModalHeader onBack={() => push(Routes.ROOT)}>Profile</ModalHeader>
    <div style="padding:0 0.5rem;font-weight:bold;">
      <FaIcon scale="3" data={faUserCircle} />
    </div>
    <ul class="vertical-stepper">
      <VStep onClick={() => push(Routes.PROFILE_UPDATE)}>
        <span slot="icon">
          <FaIcon data={faIdCard} />
        </span>
        <b slot="step">Personal</b>
        <div class="description help" slot="info">
          Personal identity information used for verification purposes.
        </div>
      </VStep>
      <VStep onClick={() => push(Routes.ADDRESS_UPDATE)}>
        <span slot="icon">
          <FaIcon data={faHome} />
        </span>
        <b slot="step"> Address </b>
        <div class="description help" slot="info">
          Up to date residential address information used for identity
          verification.
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
