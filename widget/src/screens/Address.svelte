<script lang="ts">
  import { push } from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { userStore } from '../stores/UserStore'
  import { Logger, onEnterPressed } from '../util'
  import { Routes } from '../constants'
  import type { Address } from 'api-client'
  import { transactionStore } from '../stores/TransactionStore'

  let animation = 'left'
  let isSubmittingProfile = false
  let autocomplete: google.maps.places.Autocomplete

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
      const nextRoute = $transactionStore.sourceAmount
        ? Routes.CHECKOUT_OVERVIEW
        : Routes.ROOT
      setTimeout(() => push(nextRoute), 800)
    } finally {
      setTimeout(() => {
        isSubmittingProfile = false
        userStore.clearAddress()
      }, 800)
    }
  }

  const onKeyDown = (e: Event) => {
    onEnterPressed(e, handleNextStep)
  }
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent {animation}>
  <ModalBody>
    <ModalHeader>Where do you live?</ModalHeader>
    <Label label="Street 1">
      <Input
        id="autocomplete"
        defaultValue={$userStore.address.street1}
        placeholder="Street 1"
      />
    </Label>
    <Label label="Street 2">
      <Input
        placeholder="Street 2"
        defaultValue={$userStore.address.street2}
        on:change={e => (address.street2 = e.detail)}
      />
    </Label>
    <Label label="City">
      <Input placeholder="City" defaultValue={$userStore.address.city} />
    </Label>
    <div class="inline-inputs">
      <Label label="Country" style="margin-right:1rem;">
        <Input
          placeholder="Country"
          defaultValue={$userStore.address.country}
        />
      </Label>
      <Label class="state" label="State">
        <Input placeholder="State" defaultValue={$userStore.address.state} />
      </Label>
    </div>
    <Label label="Postal Code">
      <Input
        placeholder="Postal Code"
        defaultValue={$userStore.address.postalCode}
      />
    </Label>
  </ModalBody>
  <ModalFooter>
    <Button on:click={handleNextStep}
      >{isSubmittingProfile ? 'Saving...' : 'Save'}</Button
    >
  </ModalFooter>
</ModalContent>

<svelte:head>
  <script
    src="https://maps.googleapis.com/maps/api/js?key=AIzaSyDr7FQk1bZV4Zght87YNUgCv5P4cg_1DIs&libraries=places"
    on:load={initAutoComplete}></script>
</svelte:head>

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
</style>
