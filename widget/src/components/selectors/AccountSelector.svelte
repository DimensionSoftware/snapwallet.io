<script lang="ts">
  import { onMount } from 'svelte'
  import { fly } from 'svelte/transition'
  import { faCreditCard, faUniversity } from '@fortawesome/free-solid-svg-icons'
  import { createEventDispatcher } from 'svelte'
  import { push } from 'svelte-spa-router'
  import { Routes } from '../../constants'
  import { transactionStore } from '../../stores/TransactionStore'
  import { userStore } from '../../stores/UserStore'
  import { TransactionIntents, TransactionMediums } from '../../types'
  import IconCard from '../cards/IconCard.svelte'
  import PopupSelector from '../inputs/PopupSelector.svelte'
  import { paymentMethodStore } from '../../stores/PaymentMethodStore'
  import { cachePrimaryPaymentMethodID } from '../../util'
  import { findNextKYCRoute } from '../../util/profiles'
  const dispatch = createEventDispatcher()

  export let visible = false

  $: ({ intent } = $transactionStore)
  $: ({ flags } = $userStore)
  $: isSell = intent === TransactionIntents.SELL
  $: allPaymentMethods = [
    ...$paymentMethodStore.wyrePaymentMethods.sort(pm =>
      pm.status.toLowerCase() === 'active' ? -1 : 1,
    ),
  ]

  let isLoadingPaymentMethods = true
  let copy
  let verificationNextStep = findNextKYCRoute($userStore.profileItems)

  $: {
    if (isSell) {
      copy = {
        headerTitle: 'Bank Accounts',
        sectionOneTitle: 'Add a Bank Account',
        sectionTwoTitle: 'Select a Bank Account',
        unavailable: 'No bank accounts available',
      }
    } else {
      copy = {
        headerTitle: 'Payment Methods',
        sectionOneTitle: 'Add a Payment Method',
        sectionTwoTitle: 'Select a Payment Method',
        unavailable: 'No payment methods available',
      }
    }

    verificationNextStep = findNextKYCRoute($userStore.profileItems)
  }

  onMount(() => {
    if (flags && !flags.hasWyreAccount) {
      isLoadingPaymentMethods = false
      return
    }
    try {
      // Load latest payment methods on open
      paymentMethodStore.fetchWyrePaymentMethods()
    } finally {
      setTimeout(() => (isLoadingPaymentMethods = false), 1000)
    }
  })
</script>

<PopupSelector
  {visible}
  headerTitle={copy.headerTitle}
  on:close={() => dispatch('close')}
>
  <div class="scroll-y selector-container">
    <h5 style="margin-bottom:1rem;">{copy.sectionOneTitle}</h5>
    <div
      class="card-vertical-margin"
      in:fly={{ y: 25, duration: 200 + 50 * 1 }}
    >
      <IconCard
        icon={faUniversity}
        on:click={() => {
          // Route user to next KYC step when they don't have an active Wyre acct
          const route =
            !flags?.hasWyreAccount &&
            verificationNextStep !== Routes.PROFILE_STATUS
              ? verificationNextStep
              : Routes.PLAID_LINK
          push(route)
        }}
        label="Bank Account"
        title="Connect Your Bank Account"
      />
    </div>
    <h5 style="margin-top:2rem;margin-bottom:1rem;">{copy.sectionTwoTitle}</h5>
    <div
      class="card-vertical-margin"
      in:fly={{ y: 25, duration: 200 + 50 * 2 }}
    >
      <IconCard
        label="Debit Card"
        icon={faCreditCard}
        badgeText="Active"
        badgeType="success"
        on:click={() => {
          transactionStore.update({ inMedium: TransactionMediums.DEBIT_CARD })
          dispatch('close')
        }}
      />
    </div>
    {#if !allPaymentMethods.length && isLoadingPaymentMethods}
      <p class="help">Retrieving Payment Methods...</p>
    {:else}
      {#each allPaymentMethods as pm, i (i + pm.id)}
        <div
          class="card-vertical-margin"
          in:fly={{ y: 25, duration: 300 + 50 * i }}
        >
          <IconCard
            label={pm.name}
            icon={faUniversity}
            badgeText={pm.status === 'ACTIVE' ? 'Active' : 'Pending'}
            badgeType={pm.status === 'ACTIVE' ? 'success' : 'warning'}
            on:click={() => {
              transactionStore.update({
                selectedSourcePaymentMethod: pm,
                inMedium: TransactionMediums.ACH,
              })
              cachePrimaryPaymentMethodID(pm.id)
              dispatch('close')
            }}
          />
        </div>
      {/each}
      <div class="spacer" />
    {/if}
  </div>
</PopupSelector>

<style lang="scss">
  @import '../../styles/selectors.scss';
  h5 {
    margin-bottom: 0.5rem !important;
  }
</style>
