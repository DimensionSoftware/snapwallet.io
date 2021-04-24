<script lang="ts">
  import { onMount } from 'svelte'
  import { fly } from 'svelte/transition'
  import { faUniversity } from '@fortawesome/free-solid-svg-icons'
  import { createEventDispatcher } from 'svelte'
  import { push } from 'svelte-spa-router'
  import { Routes } from '../../constants'
  import { transactionStore } from '../../stores/TransactionStore'
  import { userStore } from '../../stores/UserStore'
  import { TransactionIntents } from '../../types'
  import IconCard from '../cards/IconCard.svelte'
  import PopupSelector from '../inputs/PopupSelector.svelte'
  import { paymentMethodStore } from '../../stores/PaymentMethodStore'
  import { cachePrimaryPaymentMethodID } from '../../util'
  const dispatch = createEventDispatcher()

  export let visible = false

  $: ({ intent } = $transactionStore)
  $: ({ flags } = $userStore)
  $: isSell = intent === TransactionIntents.SELL
  $: allPaymentMethods = [...$paymentMethodStore.wyrePaymentMethods]

  let isLoadingPaymentMethods = true
  let copy
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
    <div class="card-vertical-margin">
      <IconCard
        icon={faUniversity}
        on:click={() => push(Routes.PLAID_LINK)}
        label="Bank Account"
        title="Connect Your Bank Account"
      />
    </div>
    <h5 style="margin-top:2rem;margin-bottom:1rem;">{copy.sectionTwoTitle}</h5>
    {#if !allPaymentMethods.length && isLoadingPaymentMethods}
      <p class="help">Retrieving Payment Methods...</p>
    {:else if !allPaymentMethods.length && !isLoadingPaymentMethods}
      <p class="help">{copy.unavailable}</p>
    {:else}
      {#each allPaymentMethods as pm, i (pm.id)}
        <div class="card-vertical-margin" in:fly={{ y: 25, duration: 50 * i }}>
          <IconCard
            label={pm.name}
            icon={faUniversity}
            badgeText={['OPEN', 'PENDING'].includes(pm.status)
              ? 'Pending'
              : 'Active'}
            badgeType={['OPEN', 'PENDING'].includes(pm.status)
              ? 'warning'
              : 'success'}
            on:click={() => {
              transactionStore.setSelectedSourcePaymentMethod(pm)
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
