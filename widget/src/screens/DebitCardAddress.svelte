<script lang="ts">
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { focusFirstInput, Logger, onEnterPressed } from '../util'
  import { debitCardStore } from '../stores/DebitCardStore'
  import { onMount } from 'svelte'
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import { userStore } from '../stores/UserStore'
  import {
    debitCardAddressValidationRules,
    validateForm,
  } from '../util/validation'

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

  const address = {
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
    const waitForGoogle = () => {
      if (window.google?.maps) {
        initAutoComplete()
      } else {
        setTimeout(waitForGoogle, 100)
      }
    }
    // wait for address api to load (since we're 'defer')
    waitForGoogle()
    focusFirstInput()
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
    debitCardStore.update({ address })
  }

  const handleNextStep = () => {
    const { isValid, error } = validateForm(debitCardAddressValidationRules, {
      ...$debitCardStore.address,
    })
    if (!isValid) throw new Error(error)
    push(Routes.DEBIT_CARD_2FA)
  }

  const onKeyDown = (e: Event) => {
    // Stop "Save" from occurring when enter
    // is clicked during google autocomplete

    const { isValid } = validateForm(debitCardAddressValidationRules, {
      ...$debitCardStore.address,
    })
    if (isValid) onEnterPressed(e, handleNextStep)
  }
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent>
  <ModalHeader>Card Address</ModalHeader>
  <ModalBody>
    <Label label="Street 1">
      <Input
        id="autocomplete"
        defaultValue={$debitCardStore.address.street1}
        placeholder="Street 1"
      />
    </Label>
    <div class="inline-inputs">
      <Label label="Street 2" style="max-width: 40%; margin-right: 1rem;">
        <Input
          placeholder="Street 2"
          defaultValue={$debitCardStore.address.street2}
          on:change={e => (address.street2 = e.detail)}
        />
      </Label>
      <Label class="postal" label="Postal Code">
        <Input
          placeholder="Postal Code"
          defaultValue={$debitCardStore.address.postalCode}
        />
      </Label>
    </div>
    <Label label="City">
      <Input placeholder="City" defaultValue={$debitCardStore.address.city} />
    </Label>
    <div class="inline-inputs">
      <Label label="Country" style="margin-right: 1rem;">
        <Input
          placeholder="Country"
          defaultValue={$debitCardStore.address.country ||
            $userStore.geo?.country?.toUpperCase()}
        />
      </Label>
      <Label class="state" label="State">
        <Input
          placeholder="State"
          defaultValue={$debitCardStore.address.state}
        />
      </Label>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button on:mousedown={handleNextStep}>Confirm</Button>
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
</style>
