<script lang="ts">
  import { push } from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { Routes } from '../constants'
  import { debitCardStore } from '../stores/DebitCardStore'
  import { Masks } from '../types'
  import PhoneInput from '../components/inputs/PhoneInput.svelte'
  import CountrySelector from '../components/selectors/CountrySelector.svelte'
  import { onMount } from 'svelte'
  import { focusFirstInput } from '../util'

  let countrySelectorVisible = false
  $: isUSPhoneNumber =
    $debitCardStore.phoneNumberCountry.code.toUpperCase() === 'US'

  const handleNextStep = async () => {
    push(Routes.DEBIT_CARD_ADDRESS)
  }

  onMount(() => {
    focusFirstInput()
  })
</script>

<ModalContent>
  <ModalHeader>Card Information</ModalHeader>
  <ModalBody>
    <Label label="Name on Card">
      <Input
        id="autocomplete"
        defaultValue={$debitCardStore.firstName}
        placeholder="John Smith"
        on:change={e => {
          const [firstName = '', lastName = ''] = (e.detail || '').split(' ')
          debitCardStore.update({ firstName, lastName })
        }}
      />
    </Label>
    <Label label="Phone Number">
      <PhoneInput
        on:select={() => (countrySelectorVisible = true)}
        inputmode="phone"
        autocapitalize="none"
        autocomplete="on"
        required
        type="tel"
        mask={isUSPhoneNumber ? Masks.PHONE : undefined}
        placeholder={isUSPhoneNumber ? '222 333-4444' : '222333444'}
        defaultValue={$debitCardStore.phoneNumber}
        on:change={e => {
          debitCardStore.update({ phoneNumber: e.detail })
        }}
      />
    </Label>
    <Label label="Card Number">
      <Input
        mask={Masks.DEBIT_CARD}
        id="autocomplete"
        defaultValue={$debitCardStore.number}
        placeholder="4444 3333 2222 1111"
        on:change={e => {
          e.detail &&
            debitCardStore.update({
              number: e.detail,
            })
        }}
      />
    </Label>
    <div class="inline-inputs">
      <Label label="Expiration" style="margin-right:1rem;">
        <Input
          maxlength={7}
          mask={Masks.DEBIT_CARD_EXPIRATION_DATE}
          placeholder={`10/${String(new Date().getFullYear() + 1)}`}
          defaultValue={$debitCardStore.expirationDate}
          on:change={({ detail = '' }) => {
            debitCardStore.update({
              expirationDate: detail,
            })
          }}
        />
      </Label>
      <Label label="CVC">
        <Input
          maxlength={4}
          placeholder="0000"
          defaultValue={$debitCardStore.verificationCode}
          on:change={e => {
            debitCardStore.update({ verificationCode: e.detail })
          }}
        />
      </Label>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button on:mousedown={handleNextStep}>Continue</Button>
  </ModalFooter>
</ModalContent>

{#if countrySelectorVisible}
  <CountrySelector
    visible
    on:close={() => {
      countrySelectorVisible = false
    }}
    on:select={e => {
      const { country } = e?.detail
      country && debitCardStore.update({ phoneNumberCountry: country })
      countrySelectorVisible = false
    }}
  />
{/if}

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
