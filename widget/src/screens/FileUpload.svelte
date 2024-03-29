<script lang="ts">
  import { fly, blur } from 'svelte/transition'
  import { push } from 'svelte-spa-router'
  import {
    faCheck,
    faExclamationCircle,
    faFileImage,
    faHome,
    faIdCard,
    faPassport,
    faUniversity,
  } from '@fortawesome/free-solid-svg-icons'
  import type { UsGovernmentIdDocumentInputKind } from 'api-client'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import { Logger, fileToBase64 } from '../util'
  import IconCard from '../components/cards/IconCard.svelte'
  import PopupSelector from '../components/inputs/PopupSelector.svelte'
  import Button from '../components/Button.svelte'
  import { Routes } from '../constants'
  import { transactionStore } from '../stores/TransactionStore'
  import { configStore } from '../stores/ConfigStore'
  import { userStore } from '../stores/UserStore'
  import { FileUploadTypes } from '../types'
  import { getMissingFieldMessages } from '../util/profiles'
  import FaIcon from 'svelte-awesome'

  export let isUpdatingFiles: boolean = false

  const allowedFileTypes = 'image/png,image/jpeg,image/jpg,application/pdf'

  let isFileTypeSelectorOpen = false
  let isUploadingFile = false
  let fileType = FileUploadTypes.US_DRIVER_LICENSE
  let fileEl: HTMLInputElement
  let selectedFileURI: string = ''
  let selectedFileName: string = ''
  let fileSizeError: string = ''
  let browseTitle

  $: fileIds = []
  $: minimumFiles = fileType === FileUploadTypes.US_DRIVER_LICENSE ? 2 : 1
  $: {
    if (minimumFiles <= 1) {
      browseTitle = 'Select Document'
    } else if (fileIds.length < 1) {
      browseTitle = 'Select Document (Front)'
    } else {
      browseTitle = 'Select Document (Back)'
    }
  }

  $: missingInfo = getMissingFieldMessages($userStore.profileItems)

  const handleNextStep = async () => {
    try {
      isUploadingFile = true
      const uploadResponse = await window.API.fluxUploadFile(fileEl.files[0])
      fileIds = [...fileIds, uploadResponse.fileId]
      if (fileIds.length < minimumFiles) {
        setTimeout(() => {
          selectedFileURI = ''
          selectedFileName = ''
          // Clear the element val so chrome doesn't
          // silently fail to show the preview [ch266]
          fileEl.value = ''
        }, 800)
      }
      if (fileIds.length >= minimumFiles) {
        const isGovId = [
          FileUploadTypes.US_DRIVER_LICENSE,
          FileUploadTypes.US_PASSPORT,
        ].includes(fileType)
        const isACHForm = FileUploadTypes.ACH_AUTHORIZATION_FORM === fileType
        const isProofOfAddress = FileUploadTypes.PROOF_OF_ADDRESS === fileType
        const profileResponse = await window.API.fluxSaveProfileData({
          ...(isProofOfAddress && { proofOfAddressDoc: { fileIds } }),
          ...(isACHForm && { achAuthorizationFormDoc: { fileIds } }),
          ...(isGovId && {
            usGovernmentIdDoc: {
              kind: fileType as UsGovernmentIdDocumentInputKind,
              fileIds,
            },
          }),
        })

        Logger.debug('Profile response', profileResponse)
        await userStore.fetchUserProfile()
        const wyreApproved = profileResponse.wyre?.status === 'APPROVED'
        if (isUpdatingFiles) push(Routes.PROFILE_STATUS)
        else if (wyreApproved && $transactionStore.sourceAmount)
          push(Routes.CHECKOUT_OVERVIEW)
        else push(Routes.ROOT)
      }
    } finally {
      setTimeout(() => (isUploadingFile = false), 800)
    }
  }

  const selectFileType = selectedType => () => {
    fileType = selectedType
    isFileTypeSelectorOpen = false
  }

  const openFileBrowser = () => fileEl.click()

  const SELECTOR_OPTIONS = {
    [FileUploadTypes.US_PASSPORT]: {
      icon: faPassport,
      label: 'Passport',
    },
    [FileUploadTypes.US_DRIVER_LICENSE]: {
      icon: faIdCard,
      label: 'Drivers License',
    },
  }

  // Only allow these file types after initial KYC has been done
  if (isUpdatingFiles) {
    SELECTOR_OPTIONS[FileUploadTypes.ACH_AUTHORIZATION_FORM] = {
      icon: faUniversity,
      label: 'Bank Authorization Form',
    }

    SELECTOR_OPTIONS[FileUploadTypes.PROOF_OF_ADDRESS] = {
      icon: faHome,
      label: 'Proof of Address',
    }
  }

  const getSelectorProps = fileType => {
    return (
      SELECTOR_OPTIONS[fileType] || {
        icon: faFileImage,
        label: 'Document Type',
      }
    )
  }

  $: iconCardProps = getSelectorProps(fileType)
</script>

<ModalContent>
  <ModalHeader>
    {#if missingInfo.document.isComplete}
      Identity Received
    {:else}
      Verify Identity
    {/if}
  </ModalHeader>
  <ModalBody padded>
    {#if missingInfo.document.submitted.size}
      <div style="display:flex;align-items:center;">
        <span style="margin-right:0.5rem;">
          <FaIcon
            scale={0.75}
            data={!missingInfo.document.isValid ? faExclamationCircle : faCheck}
          />
        </span>
        <h5 in:blur={{ duration: 300 }}>
          {missingInfo.document.submitted.size}
          {missingInfo.document.submitted.size > 1 ? 'Documents' : 'Document'} Uploaded
        </h5>
      </div>
    {:else if $configStore.environment === 'sandbox'}
      <h3 class="test">Upload any image in Test Mode</h3>
    {:else}
      <h5 in:blur={{ duration: 300 }}>Upload one of:</h5>
    {/if}
    <div style="margin-top:1rem;margin-bottom:0.75rem;">
      <IconCard
        title="Select a Document Type"
        blend
        icon={iconCardProps.icon}
        on:click={() => {
          // Can't change this once one doc side is uploaded
          if (fileIds.length && minimumFiles > 1) {
            return
          }
          isFileTypeSelectorOpen = true
        }}
        label={iconCardProps.label}
      />
    </div>
    <div
      on:mousedown={openFileBrowser}
      title="Click to Upload a Document"
      class="dropzone"
    >
      {#if selectedFileURI}
        <img
          class="selected-image"
          alt={selectedFileName}
          src={selectedFileURI}
        />
      {:else}
        <p class:underlined={!fileSizeError} class:error={fileSizeError}>
          {fileSizeError || browseTitle}
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
  <ModalFooter>
    <Button
      glow={selectedFileName !== ''}
      isLoading={isUploadingFile}
      disabled={!fileType}
      on:mousedown={handleNextStep}
      >{isUploadingFile
        ? // uploading file
          'Uploading'
        : minimumFiles <= 1
        ? // only need 1 document, so--
          'Upload'
        : // otherwise; guide customer Front/Back
        fileIds.length < 1
        ? 'Upload Front'
        : 'Upload Back'}</Button
    >
  </ModalFooter>
</ModalContent>

{#if isFileTypeSelectorOpen}
  <PopupSelector
    on:close={() => {
      isFileTypeSelectorOpen = false
    }}
    headerTitle="Select a Document Type"
  >
    <div class="scroll-y selector-container">
      {#each Object.entries(SELECTOR_OPTIONS) as [optionFileType, options], i}
        <div
          in:fly={{ y: 25, duration: 250 + 50 * (i + 1) }}
          class="card-vertical-margin"
        >
          <IconCard
            icon={options.icon}
            on:click={selectFileType(optionFileType)}
            label={options.label}
          />
        </div>
      {/each}
      <div class="spacer" />
    </div>
  </PopupSelector>
{/if}

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/selectors.scss';

  // TODO: make drop-able on desktop
  .dropzone {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 60%;
    width: auto;
    border: 1px dashed var(--theme-text-color);
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
