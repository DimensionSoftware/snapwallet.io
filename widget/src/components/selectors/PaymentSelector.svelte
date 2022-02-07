<script lang="ts">
  import { faCheck, faUniversity } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import { getPrimaryPaymentMethodID } from '../../util'
  import { paymentMethodStore } from '../../stores/PaymentMethodStore'
  import { transactionStore } from '../../stores/TransactionStore'
  import { userStore } from '../../stores/UserStore'
  import VStep from '../../components/VStep.svelte'
  import { TransactionMediums } from '../../types'

  export let isBuy: boolean = true
  export let active: boolean = false
  export let onClick
  export let description: string = ''
  export let disabled: boolean | null = null

  $: isDebitCard = $transactionStore.inMedium === TransactionMediums.DEBIT_CARD
  $: success =
    Boolean($transactionStore.selectedSourcePaymentMethod) || isDebitCard

  $: ({ flags } = $userStore)

  // Select last used pm when request completes.
  paymentMethodStore.subscribe(({ wyrePaymentMethods }) => {
    if ($transactionStore.selectedSourcePaymentMethod) return
    const primaryPaymentMethodID = getPrimaryPaymentMethodID()
    if (!primaryPaymentMethodID) return
    const primaryPaymentMethod = wyrePaymentMethods.find(
      pm => pm.id === primaryPaymentMethodID,
    )
    if (primaryPaymentMethod) {
      transactionStore.setSelectedSourcePaymentMethod(primaryPaymentMethod)
    }
  })
</script>

<VStep
  title={isBuy ? 'Select Your Payment Method' : 'Select Your Destination Bank'}
  disabled={disabled !== null
    ? disabled
    : !$userStore.isProfilePending && !flags?.hasWyreAccount}
  {active}
  {success}
  {onClick}
>
  <span slot="icon" class="payment-icon">
    {#if success}
      <FaIcon data={!success ? faUniversity : faCheck} />
    {:else}
      <span class="default-icon" />
    {/if}
  </span>
  <span slot="step">
    <!-- Multiple PMs will be possible for buy and bank account is only option for sell atm -->
    {#if !isDebitCard && $transactionStore.selectedSourcePaymentMethod}
      {$transactionStore.selectedSourcePaymentMethod.name}
    {:else if isDebitCard}
      Debit Card
    {:else if isBuy}
      Select Payment Method
    {:else}
      Select Bank Account
    {/if}
  </span>
  <div class:hidden={!description} class="description help" slot="info">
    {description}
  </div>
</VStep>

<style lang="scss">
  @import '../../styles/_vars.scss';
  @import '../../styles/text.scss';
  .description {
    min-height: 2.5rem;
    margin-left: 1rem;
    opacity: 0.85;
  }
  .payment-icon:before {
    height: 2px;
    width: 2px;
    left: 3px;
    top: 11px;
  }
  .hidden {
    display: none;
  }
</style>
