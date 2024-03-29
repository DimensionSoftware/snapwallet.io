<script lang="ts">
  import { push } from 'svelte-spa-router'
  import { blur } from 'svelte/transition'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { configStore } from '../stores/ConfigStore'
  import { userStore } from '../stores/UserStore'
  import { focusFirstInput, Logger, onEnterPressed } from '../util'
  import { Routes, UserProfileFieldTypes } from '../constants'
  import type { Address } from 'api-client'
  import { transactionStore } from '../stores/TransactionStore'
  import { onMount } from 'svelte'
  import { getMissingFieldMessages } from '../util/profiles'

  export let isUpdateScreen: boolean = false

  let isSubmittingProfile = false
  let autocomplete: google.maps.places.Autocomplete

  $: missingInfo = getMissingFieldMessages($userStore.profileItems)

  const componentForm = {
    street_number: 'short_name',
    route: 'long_name',
    locality: 'long_name',
    administrative_area_level_1: 'short_name',
    country: 'short_name',
    postal_code: 'short_name',
  }

  const googleTypeToFluxType = {
    country: 'country',
    locality: 'city',
    postal_code: 'postalCode',
    street_number: 'streetNumber',
    route: 'street1',
    administrative_area_level_1: 'state',
  }

  const address: Address = {
    street1: '',
    street2: '',
    city: '',
    state: '',
    postalCode: '',
    country: '',
  }

  const initAutoComplete = () => {
    // Create the autocomplete object, restricting the search predictions to
    // geographical location types.
    autocomplete = new google.maps.places.Autocomplete(
      document.getElementById('autocomplete') as HTMLInputElement,
      { types: ['geocode'] },
    )

    // Avoid paying for data that you don't need by restricting the set of
    // place fields that are returned to just the address components.
    autocomplete.setFields(['address_component'])

    // When the user selects an address from the drop-down, populate the
    // address fields in the form.
    autocomplete.addListener('place_changed', fillInAddress)

    return () => (autocomplete = undefined)
  }

  onMount(() => {
    focusFirstInput()
    const waitForGoogle = () => {
      if (window.google?.maps) {
        initAutoComplete()
      } else {
        setTimeout(waitForGoogle, 100)
      }
    }
    // wait for address api to load (since we're 'defer')
    waitForGoogle()
  })

  function fillInAddress() {
    // Get the place details from the autocomplete object.
    const place = autocomplete.getPlace()
    Logger.debug('Place', place)

    // // Get each component of the address from the place details,
    // // and then fill-in the corresponding field on the form.
    let streetNumberSet = false
    let routeSet = false

    for (const component of place.address_components) {
      const addressType = component.types[0]

      if (componentForm[addressType]) {
        const val = component[componentForm[addressType]]
        const fluxType = googleTypeToFluxType[addressType]
        const isStreetNumber = fluxType === googleTypeToFluxType.street_number
        const isStreetRoute = fluxType === googleTypeToFluxType.route

        if (isStreetNumber) {
          address.street1 = !routeSet ? val : `${val} ${address.street1}`
          streetNumberSet = true
          continue
        }

        if (isStreetRoute) {
          address.street1 = !streetNumberSet
            ? address.street1
            : `${address.street1} ${val}`
          routeSet = true
          continue
        }

        address[fluxType] = val
      }
    }
    userStore.setFullAddress(address)
  }

  const handleNextStep = async () => {
    try {
      isSubmittingProfile = true
      await window.API.fluxSaveProfileData({
        address: $userStore.address,
      })
      setTimeout(() => {
        if (isUpdateScreen) push(Routes.PROFILE_STATUS)
        else if (!$userStore.flags?.hasPhone) push(Routes.PROFILE_SEND_SMS)
        else if (!$userStore.flags.hasWyreAccount) push(Routes.FILE_UPLOAD)
        else if ($transactionStore.sourceAmount) push(Routes.CHECKOUT_OVERVIEW)
        else push(Routes.ROOT)
      }, 800)
    } finally {
      setTimeout(() => {
        isSubmittingProfile = false
        userStore.clearAddress()
      }, 800)
    }
  }

  const onKeyDown = (e: Event) => {
    // Stop "Save" from occurring when enter
    // is clicked during google autocomplete
    const addressVal = Object.values($userStore.address).join('')
    if (addressVal.length) onEnterPressed(e, handleNextStep)
  }

  const fillTestInfo = e => {
    e.preventDefault()
    userStore.setFullAddress({
      street1: '1 Crypto',
      street2: '',
      city: 'Metasphere',
      state: 'CA',
      postalCode: '90123',
      country: 'US',
    })
  }
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent>
  <ModalHeader>Your Address</ModalHeader>
  <ModalBody padded>
    {#if missingInfo.address.isComplete}
      <h5 in:blur={{ duration: 300 }}>Address received and may be updated:</h5>
    {:else if $configStore.environment === 'sandbox'}
      <h3 class="test">
        <a on:click={fillTestInfo} href="">Fill With Test Info</a>
      </h3>
    {:else}
      <h5>&nbsp;</h5>
    {/if}
    <Label label="Street 1">
      <Input
        id="autocomplete"
        defaultValue={$userStore.address.street1}
        placeholder={$userStore.virtual?.address?.street1 || 'Street 1'}
      />
    </Label>
    <div class="inline-inputs">
      <Label label="Street 2" style="max-width: 40%; margin-right: 1rem;">
        <Input
          placeholder={$userStore.virtual?.address?.street2 || 'Street 2'}
          defaultValue={$userStore.address.street2}
          on:change={e => (address.street2 = e.detail)}
        />
      </Label>
      <Label class="postal" label="Postal Code">
        <Input
          placeholder={$userStore.virtual?.address?.postalCode || 'Postal Code'}
          defaultValue={$userStore.address.postalCode}
        />
      </Label>
    </div>
    <Label label="City">
      <Input
        placeholder={$userStore.virtual?.address?.city || 'City'}
        defaultValue={$userStore.address.city}
      />
    </Label>
    <div class="inline-inputs">
      <Label label="Country" style="margin-right: 1rem;">
        <Input
          placeholder={$userStore.virtual?.address?.country || 'Country'}
          defaultValue={$userStore.address.country}
        />
      </Label>
      <Label class="state" label="State">
        <Input
          placeholder={$userStore.virtual?.address?.state || 'State'}
          defaultValue={$userStore.address.state}
        />
      </Label>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button
      id="address_save"
      isLoading={isSubmittingProfile}
      on:mousedown={handleNextStep}
      >{isSubmittingProfile ? 'Saving' : 'Save'}</Button
    >
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  label {
    size: smaller;
  }
  .inline-inputs {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  h5 {
    margin-top: 0;
  }
  .test {
    text-align: left;
  }
</style>
