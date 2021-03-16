<script lang="ts">
    import vld8 from 'validator'
    import { push } from 'svelte-spa-router'
    import ModalBody from '../components/ModalBody.svelte'
    import ModalContent from '../components/ModalContent.svelte'
    import ModalFooter from '../components/ModalFooter.svelte'
    import Button from '../components/Button.svelte'
    import Input from '../components/inputs/Input.svelte'
    import Label from '../components/inputs/Label.svelte'
    import ModalHeader from '../components/ModalHeader.svelte'
    import { userStore } from '../stores/UserStore'
    import { toaster } from '../stores/ToastStore'
    import { Logger, onEnterPressed } from '../util'
    import { APIErrors, Routes } from '../constants'
import type { Address, UsGovernmentIdDocumentInputKind } from 'api-client';
  
    let animation = 'left'
    let isMakingRequest = false
  
    const timeout = 700

    let fileEl: HTMLInputElement
    let file2El: HTMLInputElement
    let usIDKind: string
    const handleNextStep = async () => {
      const resp = await window.API.fluxUploadFile(fileEl.files[0])
      const resp2 = await window.API.fluxUploadFile(file2El.files[0])
      console.log(resp)
      await window.API.fluxSaveProfileData({
        usGovernmentIdDoc: {
          kind: usIDKind as UsGovernmentIdDocumentInputKind,
          fileIds: [resp.fileId, resp2.fileId],
        },
      })

    }
  </script>
  
  <ModalContent {animation}>
    <ModalBody>
      <ModalHeader hideBackButton>CITIZEN?: gonna need to see some ID</ModalHeader>
      <label label="Upload US Government ID (Front)">
        <strong>Upload US Government ID (Front)</strong><br />
        <input type="file" bind:this={fileEl} />
      </label>
      <label label="Upload US Government ID (Back)">
        <strong>Upload US Government ID (Back)</strong><br />
        <input type="file" bind:this={file2El} />
      </label>
      <label label="US Government ID Kind">
        <strong>US Government ID Kind</strong><br />
        <select bind:value={usIDKind}>
          <option value="GI_US_DRIVING_LICENSE">US Drivers' license (front/back)</option>
          <option value="GI_US_PASSPORT_CARD">US Passport Card (front/back)</option>
          <option value="GI_US_GOVERNMENT_ID">US Government ID (front/back)</option>
          <option value="GI_US_PASSPORT">US Passport (picture page)</option>
        </select>
      </label>
    </ModalBody>
    <button on:click={handleNextStep}>Plz give me now</button>
  </ModalContent>
  
  <style lang="scss">
    @import '../styles/_vars.scss';
    label {
        size: smaller;
    }
  </style>
  