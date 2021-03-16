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
import type { Address } from 'api-client';
  
    let animation = 'left'
    let isMakingRequest = false
  
    const timeout = 700

    let fileEl: HTMLInputElement
    const handleNextStep = async () => {
      const resp = await window.API.fluxUploadFile(fileEl.files[0])
      console.log(resp)
      await window.API.fluxSaveProfileData({
        proofOfAddressDoc: {
          fileIds: [resp.fileId],
        },
      })

    }
  </script>
  
  <ModalContent {animation}>
    <ModalBody>
      <ModalHeader hideBackButton>you have an address; prove it!</ModalHeader>
      <form>
        <label label="Upload Proof of Address">
          <input name="file" type="file" bind:this={fileEl} />
        </label>
      </form>
    </ModalBody>
    <button on:click={handleNextStep}>Plz give me now</button>
  </ModalContent>
  
  <style lang="scss">
    @import '../styles/_vars.scss';
    label {
        size: smaller;
    }
  </style>
  