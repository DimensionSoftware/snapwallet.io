<script lang="ts">
  import { faUniversity } from '@fortawesome/free-solid-svg-icons'
  import { createEventDispatcher } from 'svelte'
  import { push } from 'svelte-spa-router'
  import { Routes } from '../../constants'
  import { transactionStore } from '../../stores/TransactionStore'
  import { userStore } from '../../stores/UserStore'
  import { TransactionIntents } from '../../types'
  import IconCard from '../cards/IconCard.svelte'
  import PopupSelector from '../inputs/PopupSelector.svelte'
  const dispatch = createEventDispatcher()

  export let visible = false

  $: ({ intent } = $transactionStore)
  $: ({ flags } = $userStore)
  $: isSell = intent === TransactionIntents.SELL

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
</script>

<PopupSelector
  {visible}
  headerTitle={copy.headerTitle}
  on:close={() => dispatch('close')}
>
  <div class="scroll selector-container">
    <h5>{copy.sectionOneTitle}</h5>
    <IconCard
      icon={faUniversity}
      on:click={() => push(Routes.PLAID_LINK)}
      label="Bank Account"
    />
    <h5 style="margin-top:2rem">{copy.sectionTwoTitle}</h5>
    {#if !flags?.hasWyrePaymentMethods}
      <p class="help">{copy.unavailable}</p>
    {/if}
  </div>
</PopupSelector>

<style lang="scss">
  @import '../../styles/selectors.scss';
</style>
