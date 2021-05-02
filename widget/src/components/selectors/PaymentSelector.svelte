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
  export let onClick
  export let description
  export let disabled: boolean | null = null

  $: success =
    Boolean($transactionStore.selectedSourcePaymentMethod) ||
    $transactionStore.inMedium === TransactionMediums.DEBIT_CARD

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
  title="Select Your Payment Method"
  disabled={disabled !== null
    ? disabled
    : !$userStore.isProfilePending && !flags?.hasWyreAccount}
  {success}
  {onClick}
>
  <span slot="icon">
    <FaIcon
      data={!success
        ? faUniversity
        : faCheck}
    />
  </span>
  <b slot="step">
    <!-- Multiple PMs will be possible for buy and bank account is only option for sell atm -->
    {#if $transactionStore.selectedSourcePaymentMethod}
      {$transactionStore.selectedSourcePaymentMethod.name}
    {:else if $transactionStore.inMedium === TransactionMediums.DEBIT_CARD}
      Debit Card
    {:else if isBuy}
      Select Payment Method
    {:else}
      Select Bank Account
    {/if}
  </b>
  <div class:hidden={!description} class="description help" slot="info">
    {description}
  </div>
</VStep>

<style lang="scss">
  @import '../../styles/_vars.scss';
  @import '../../styles/text.scss';
  .description {
    min-height: 2.5rem;
    margin-left: 0.55rem;
    color: var(--theme-text-color) !important;
    opacity: 0.85;
  }
  .hidden {
    display: none;
  }
</style>
