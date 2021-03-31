<script lang="ts">
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import {
    faHome,
    faIdCard,
    faMailBulk,
    faUserCircle,
  } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import VStep from '../components/VStep.svelte'
</script>

<ModalContent>
  <ModalBody>
    <ModalHeader>Profile</ModalHeader>
    <div style="padding-left:1rem;font-weight:bold;">
      <FaIcon scale="3" data={faUserCircle} />
    </div>
    <ul class="vertical-stepper">
      <VStep>
        <span slot="icon">
          <FaIcon data={faIdCard} />
        </span>
        <b slot="step">Personal</b>
        <div class="description help" slot="info">
          Personal information used for verification purposes.
        </div>
      </VStep>
      <VStep>
        <span slot="icon">
          <FaIcon data={faHome} />
        </span>
        <b slot="step"> Address </b>
        <div class="description help" slot="info">
          Up to date residential address information used for verification.
        </div>
      </VStep>
      <VStep>
        <span slot="icon">
          <FaIcon data={faMailBulk} />
        </span>
        <b slot="step"> Contact </b>
        <div class="description help" slot="info">
          Contact information used for verification and security.
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
    margin-top: 2.5rem;
    margin-left: 1rem;
    list-style: none;
    padding: 0;
  }

  .description {
    min-height: 60px;
    cursor: pointer;
  }
</style>
