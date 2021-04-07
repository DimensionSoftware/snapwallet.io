<script lang="ts">
  import { faCheck, faUniversity } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import VStep from '../../components/VStep.svelte'
  import { paymentMethodStore } from '../../stores/PaymentMethodStore'
  import { transactionStore } from '../../stores/TransactionStore'
  import { userStore } from '../../stores/UserStore'

  export let isBuy: boolean = true
  export let onClick
  export let description

  $: ({ flags } = $userStore)
</script>

<VStep
  title="Select Your Payment Method"
  disabled={!$userStore.isProfilePending && !flags?.hasWyreAccount}
  success={$transactionStore.selectedSourcePaymentMethod}
  {onClick}
>
  <span slot="icon">
    <FaIcon
      data={!$transactionStore.selectedSourcePaymentMethod
        ? faUniversity
        : faCheck}
    />
  </span>
  <b slot="step">
    <!-- Multiple PMs will be possible for buy and bank account is only option for sell atm -->
    {#if $transactionStore.selectedSourcePaymentMethod}
      {$transactionStore.selectedSourcePaymentMethod.name}
    {:else if isBuy && !$paymentMethodStore.wyrePaymentMethods?.length}
      Add Payment Method
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
    min-height: 60px;
    margin-left: 0.55rem;
    color: var(--theme-text-color) !important;
    opacity: 0.85;
  }
  .hidden {
    display: none;
  }
</style>
