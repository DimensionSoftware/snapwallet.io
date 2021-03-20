<script lang="ts">
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { Logger, fileToBase64 } from '../util'
  import IconCard from '../components/cards/IconCard.svelte'
  import {
    faFileImage,
    faIdCard,
    faPassport,
  } from '@fortawesome/free-solid-svg-icons'
  import PopupSelector from '../components/inputs/PopupSelector.svelte'
  import Button from '../components/Button.svelte'
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import { transactionStore } from '../stores/TransactionStore'
  import type { UsGovernmentIdDocumentInputKind } from 'api-client';

  const allowedFileTypes = 'image/png,image/jpeg,image/jpg,application/pdf'

  let animation = 'left'
  let isFileTypeSelectorOpen = false
  let fileType = ''
  let fileEl: HTMLInputElement
  let selectedFileURI: string = ''
  let selectedFileName: string = ''
  let fileSizeError: string = ''

  const handleNextStep = async () => {
    const uploadResponse = await window.API.fluxUploadFile(fileEl.files[0])
    Logger.debug(uploadResponse)
    const profileResponse = await window.API.fluxSaveProfileData({
      usGovernmentIdDoc: {
        kind: fileType as UsGovernmentIdDocumentInputKind,
        fileIds: [uploadResponse.fileId],
      },
    })
    setTimeout(() => {
      const wyreApproved = profileResponse.wyre?.status === 'APPROVED'
      if (wyreApproved && $transactionStore.sourceAmount)
        push(Routes.CHECKOUT_OVERVIEW)
      else push(Routes.ROOT)
    }, 800)
  }

  const selectFileType = selectedType => () => {
    fileType = selectedType
    isFileTypeSelectorOpen = false
  }

  const openFileBrowser = () => document.getElementById('file-input').click()

  const getSelectorProps = ft => {
    if (ft === 'passport')
      return {
        icon: faPassport,
        label: 'Passport',
      }

    if (ft === 'drivers_license')
      return {
        icon: faIdCard,
        label: 'Drivers License',
      }

    return {
      icon: faFileImage,
      label: 'Document Type',
    }
  }

  $: iconCardProps = getSelectorProps(fileType)
</script>

<ModalContent {animation}>
  <ModalBody>
    <ModalHeader>Verify Identity</ModalHeader>
    <IconCard
      blend
      icon={iconCardProps.icon}
      on:click={() => (isFileTypeSelectorOpen = true)}
      label={iconCardProps.label}
    />
    <div on:click={openFileBrowser} class="dropzone">
      {#if selectedFileURI}
        <img
          class="selected-image"
          alt={selectedFileName}
          src={selectedFileURI}
        />
      {:else}
        <p class:underlined={!fileSizeError} class:error={fileSizeError}>
          {fileSizeError || 'Select a File'}
        </p>
      {/if}
    </div>
    <input
      accept={allowedFileTypes}
      id="file-input"
      hidden
      type="file"
      bind:this={fileEl}
      on:change={async e => {
        const file = e.target.files[0]
        if (file.size >= 7e6) {
          selectedFileName = ''
          selectedFileURI = ''
          fileSizeError = 'Please select a file smaller than 7mb'
          return
        }
        selectedFileName = file.name
        selectedFileURI = await fileToBase64(file)
      }}
    />
  </ModalBody>
  <Button disabled={!fileType} on:click={handleNextStep}>Upload</Button>
</ModalContent>

<PopupSelector
  on:close={() => {
    isFileTypeSelectorOpen = false
  }}
  visible={isFileTypeSelectorOpen}
  headerTitle="Select a Document Type"
>
  <div class="">
    <IconCard
      icon={faPassport}
      on:click={selectFileType('GI_US_PASSPORT')}
      label="Passport"
    />
    <IconCard
      icon={faIdCard}
      on:click={selectFileType('GI_US_DRIVING_LICENSE')}
      label="Drivers License"
    />
  </div>
</PopupSelector>

<style lang="scss">
  @import '../styles/_vars.scss';

  // TODO: make drop-able on desktop
  .dropzone {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 2rem;
    height: 60%;
    width: 100%;
    border: 1px dashed var(--theme-color);
    cursor: pointer;
    font-weight: 600;
  }

  .underlined {
    text-decoration: underline;
  }

  .error {
    color: var(--theme-error-color);
  }

  .selected-image {
    max-height: 95%;
    max-width: 95%;
  }
</style>
