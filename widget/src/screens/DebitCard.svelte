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
  import { userStore } from '../stores/UserStore'
  import { Masks } from '../types'
  import PhoneInput from '../components/inputs/PhoneInput.svelte'
  import CountrySelector from '../components/selectors/CountrySelector.svelte'
  import { onMount } from 'svelte'
  import { focusFirstInput, onEnterPressed } from '../util'
  import { debitCardValidationRules, validateForm } from '../util/validation'
  import { formatExpiration } from '../util/transactions'
  import { transactionStore } from '../stores/TransactionStore'
  import TimeTicker from '../components/TimeTicker.svelte'

  let countrySelectorVisible = false
  $: isUSPhoneNumber =
    $debitCardStore.phoneNumberCountry.code.toUpperCase() === 'US'

  const handleNextStep = async () => {
    const phoneNumber = `${$debitCardStore.phoneNumberCountry?.dial_code}${$debitCardStore.phoneNumber}`.replace(
      /(-|\s)/g,
      '',
    )
    const { isValid, error } = validateForm(debitCardValidationRules, {
      phoneNumber,
      firstName: $debitCardStore.firstName,
      lastName: $debitCardStore.lastName,
      cardNumber: $debitCardStore.number,
      cardExpiration: $debitCardStore.expirationDate,
      cardVerificationCode: $debitCardStore.verificationCode,
    })

    if (!isValid) throw new Error(error)

    push(Routes.DEBIT_CARD_ADDRESS)
  }

  onMount(() => {
    focusFirstInput()
  })

  const onKeyDown = (e: Event) => {
    onEnterPressed(e, handleNextStep)
  }
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent>
  <ModalHeader>Card Information</ModalHeader>
  <ModalBody>
    <TimeTicker
      time={formatExpiration($transactionStore.transactionExpirationSeconds)}
    />
    <Label label="Name on Card">
      <Input
        autocomplete="cc-name"
        defaultValue={`${$debitCardStore.firstName} ${$debitCardStore.lastName}`.trim()}
        placeholder="John Smith"
        on:change={e => {
          let [firstName = '', ...lastName] = (e.detail || '').split(' ')
          lastName = lastName.join(' ').trim()
          debitCardStore.update({ firstName: firstName.trim(), lastName })
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
        autocomplete="cc-number"
        mask={Masks.DEBIT_CARD}
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
          autocomplete="cc-exp"
          maxlength={5}
          mask={Masks.DEBIT_CARD_EXPIRATION_DATE}
          placeholder="MM/YY"
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
          autocomplete="cc-csc"
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
      if (country) {
        userStore.setPhoneNumberCountry(country)
        debitCardStore.update({ phoneNumberCountry: country })
      }
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
