<script lang="ts">
  import { onMount } from 'svelte'
  import { fly } from 'svelte/transition'
  import {
    faClock,
    faCreditCard,
    faPlusCircle,
    faUniversity,
  } from '@fortawesome/free-solid-svg-icons'
  import { createEventDispatcher } from 'svelte'
  import { push } from 'svelte-spa-router'
  import { Routes } from '../../constants'
  import { transactionStore } from '../../stores/TransactionStore'
  import { userStore } from '../../stores/UserStore'
  import { TransactionIntents, TransactionMediums } from '../../types'
  import PaymentMethodCard from '../cards/PaymentMethodCard.svelte'
  import PopupSelector from '../inputs/PopupSelector.svelte'
  import { paymentMethodStore } from '../../stores/PaymentMethodStore'
  import { cachePrimaryPaymentMethodID } from '../../util'
  import { findNextKYCRoute } from '../../util/profiles'
  import FaIcon from 'svelte-awesome'
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
  let showAvailablePms = false

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
        sectionTwoTitle: 'Available',
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
    {#if !showAvailablePms}
      <div
        class="card-vertical-margin"
        in:fly={{ y: 25, duration: 200 + 50 * 1 }}
      >
        <PaymentMethodCard
          badgeText="Lowest Fee"
          badgeType="info"
          on:click={() => {
            if (!window.AUTH_MANAGER.viewerIsLoggedIn()) {
              push(Routes.SEND_OTP)
              return
            }
            if (allPaymentMethods.length) {
              showAvailablePms = true
              return
            }

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
        >
          <div
            class="flex"
            style="display:flex;flex-direction:column;opacity:0.85;font-size:0.75rem;justify-content:center;width:100%;height:100%;"
          >
            <div style="opacity:0.85">Identity verification required</div>
            <div
              style="display:flex;align-items:center;justify-content:space-between;"
            >
              <strong>U.S. Fee</strong>
              <div style="display:flex;align-items:center;">
                <span style="margin-right:0.35rem"
                  >0.75% or mid-market rate</span
                >
                <FaIcon scale="0.75" data={faUniversity} />
              </div>
            </div>
            <div
              style="display:flex;align-items:center;justify-content:space-between;"
            >
              <strong>Delivery time</strong>
              <div style="display:flex;align-items:center;">
                <span style="margin-right:0.35rem">5 business days</span>
                <FaIcon scale="0.75" data={faClock} />
              </div>
            </div>
          </div>
        </PaymentMethodCard>
      </div>
      <div
        class="card-vertical-margin"
        in:fly={{ y: 25, duration: 200 + 50 * 2 }}
      >
        <PaymentMethodCard
          label="Debit Card"
          badgeText="Instant Buy"
          badgeType="success"
          on:click={() => {
            transactionStore.update({ inMedium: TransactionMediums.DEBIT_CARD })
            dispatch('close')
          }}
        >
          <div
            class="flex"
            style="display:flex; flex-direction:column;opacity:0.85;font-size:0.75rem;justify-content:center;width:100%;height:100%;"
          >
            <div style="margin-top: .2rem;">
              <img height="14" src="/widget/card_mastercard.png" />
              <img height="16" src="/widget/card_visa.png" />
              <img height="16" src="/widget/card_discover.png" />
            </div>
            <div
              style="display:flex;align-items:center;justify-content:space-between;"
            >
              <strong>U.S. Fee</strong>
              <div style="display:flex;align-items:center">
                <span style="margin-right:0.35rem">2.9% + 30¢ or $5</span>
                <FaIcon scale="0.75" data={faCreditCard} />
              </div>
            </div>
            <div
              style="display:flex;align-items:center;justify-content:space-between;"
            >
              <strong>Delivery time</strong>
              <div style="display:flex;align-items:center;">
                <span style="margin-right:0.35rem">1-2 hours</span>
                <FaIcon scale="0.75" data={faClock} />
              </div>
            </div>
          </div>
        </PaymentMethodCard>
      </div>
    {/if}
    {#if showAvailablePms}
      {#if !allPaymentMethods.length && isLoadingPaymentMethods}
        <p class="help">Retrieving Payment Methods...</p>
      {:else}
        <div
          style="display:flex;justify-content:space-between;align-items:center;margin-bottom:1.75rem;"
          in:fly={{ y: 25, duration: 300 + 50 }}
        >
          <div
            style="display:flex;justify-content:flex-end;text-decoration:underline;cursor:pointer;font-size:0.75rem;opacity:0.85;align-items:center"
            on:mousedown={() => (showAvailablePms = false)}
          >
            Back
          </div>
          <div
            style="display:flex;justify-content:flex-start;cursor:pointer;align-items:center;"
            on:mousedown={() => {
              // Route user to next KYC step when they don't have an active Wyre acct
              const route =
                !flags?.hasWyreAccount &&
                verificationNextStep !== Routes.PROFILE_STATUS
                  ? verificationNextStep
                  : Routes.PLAID_LINK
              push(route)
            }}
          >
            <FaIcon data={faPlusCircle} />
            <span style="margin-left:0.5rem">Add Account</span>
          </div>
        </div>
        {#each allPaymentMethods as pm, i (i + pm.id)}
          <div
            class="card-vertical-margin"
            in:fly={{ y: 25, duration: 300 + 50 * i }}
          >
            <PaymentMethodCard
              label={pm.name}
              icon={faUniversity}
              badgeText={pm.status === 'ACTIVE' ? undefined : 'Pending'}
              badgeType={pm.status === 'ACTIVE' ? undefined : 'warning'}
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
      {/if}
      <div class="spacer" />
    {/if}
  </div>
</PopupSelector>

<style lang="scss">
  @import '../../styles/selectors.scss';
  h5 {
    margin-bottom: 0.5rem !important;
  }
  .flex > div {
    margin-bottom: 0.25rem;
  }
  .flex > div:last-child {
    margin-bottom: 0.5rem;
  }
</style>
