<script lang="ts">
  import { push } from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import IntentSelector from '../components/IntentSelector.svelte'
  import Button from '../components/Button.svelte'
  import CryptoCard from '../components/cards/CryptoCard.svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { userStore } from '../stores/UserStore'
  import { priceStore } from '../stores/PriceStore'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import { onMount } from 'svelte'
  import { isValidNumber, onEnterPressed } from '../util'
  import TotalContainer from '../components/TotalContainer.svelte'
  import { Routes } from '../constants'
  import {
    faCheck,
    faIdCard,
    faUniversity,
  } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import { TransactionIntents } from '../types'
  import ExchangeRate from '../components/ExchangeRate.svelte'
  import AccountSelector from '../components/selectors/AccountSelector.svelte'
  import CryptoSelector from '../components/selectors/CryptoSelector.svelte'

  let cryptoSelectorVisible = false
  let paymentSelectorVisible = false

  let isEnteringSourceAmount = true
  let isLoadingPrices = !Boolean($transactionStore.sourceAmount)

  $: ({
    sourceCurrency,
    destinationCurrency,
    sourceAmount,
    intent,
  } = $transactionStore)

  $: ({ flags } = $userStore)

  $: selectedDirection = `${sourceCurrency.ticker}_${destinationCurrency.ticker}`
  $: isBuy = intent === TransactionIntents.BUY

  $: selectedPriceMap = $priceStore.prices[selectedDirection]
  $: selectedDestinationPrice = selectedPriceMap[destinationCurrency.ticker]
  $: exchangeRate = isEnteringSourceAmount
    ? 1 / selectedDestinationPrice
    : selectedDestinationPrice

  $: fakePrice = 10_000
  $: nextRoute = Routes.PLAID_LINK

  const animateRandomPrice = () => {
    window.requestAnimationFrame(_ts => {
      if (isLoadingPrices) {
        fakePrice = fakePrice + 213.02
        animateRandomPrice()
      }
    })
  }

  const handleNextStep = async () => {
    const { sourceAmount } = $transactionStore
    if (!sourceAmount || !isValidNumber(sourceAmount))
      // focus input
      return document.querySelector('input')?.focus()

    push(nextRoute)
  }

  // Find the next path based on user data
  const getNextPath = async () => {
    // TODO: move this request somewhere sane
    const { flags = {} } = await window.API.fluxViewerData()
    userStore.setFlags(flags)
    const { hasWyrePaymentMethods, hasWyreAccount } = flags

    if (hasWyrePaymentMethods && hasWyreAccount)
      nextRoute = Routes.CHECKOUT_OVERVIEW
    else if (hasWyrePaymentMethods) nextRoute = Routes.PROFILE
  }

  const onKeyDown = (e: Event) => {
    onEnterPressed(e, handleNextStep)
  }

  const getInitialPrices = async () => {
    try {
      animateRandomPrice()
      await priceStore.fetchPrices()
    } finally {
      setTimeout(() => (isLoadingPrices = false), 1200)
    }
  }

  onMount(() => {
    getInitialPrices()
    getNextPath()
    const priceInterval = priceStore.pollPrices()
    return () => clearInterval(priceInterval)
  })
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent>
  <ModalBody>
    <IntentSelector
      on:change={() => (isEnteringSourceAmount = !isEnteringSourceAmount)}
    />
    <div class="cryptocurrencies-container">
      <div class="dst-container">
        <Label>
          <CryptoCard
            on:click={() => (cryptoSelectorVisible = true)}
            crypto={isBuy ? destinationCurrency : sourceCurrency}
          />
        </Label>
      </div>
      <div style="display:flex;flex-direction:column;height:5rem;">
        <Label label="Amount">
          <Input
            pattern={`[\\d,\\.]*`}
            on:change={e => {
              const val = Number(e.detail)
              transactionStore.setSourceAmount(val, selectedDestinationPrice)
            }}
            defaultValue={sourceAmount}
            required
            type="number"
            inputmode="number"
            placeholder="0"
          />
        </Label>
      </div>
      <ul class="vertical-stepper">
        {#if flags?.hasWyreAccount}
          <li class="success">
            Verify Identity
            <span class="icon">
              <FaIcon data={faCheck} />
            </span>
          </li>
        {:else}
          <li
            style="cursor:pointer;display:flex;align-items:center;"
            on:click={() => (paymentSelectorVisible = true)}
          >
            <FaIcon data={faIdCard} />
            <b style="margin-left:0.5rem;text-decoration:underline">
              Verify Identity
            </b>
          </li>
        {/if}
        <li
          class:disabled={!flags?.hasWyreAccount}
          style="cursor:pointer;display:flex;align-items:center;"
          on:click={() =>
            flags?.hasWyreAccount && (paymentSelectorVisible = true)}
        >
          <FaIcon data={faUniversity} />
          <b style="margin-left:0.5rem;text-decoration:underline">
            <!-- Multiple PMs will be possible for buy and bank account is only option for sell atm -->
            {isBuy ? 'Select Payment Method' : 'Select Bank Account'}
          </b>
        </li>
        <li>
          <ExchangeRate {fakePrice} {isLoadingPrices} {exchangeRate} />
        </li>
        <li>
          <TotalContainer />
        </li>
      </ul>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button
      disabled={!sourceAmount ||
        !flags?.hasWyreAccount ||
        !flags?.hasWyrePaymentMethods}
      on:click={handleNextStep}>Checkout</Button
    >
  </ModalFooter>
</ModalContent>

<!-- Cryptocurrency Selector -->
<CryptoSelector
  visible={cryptoSelectorVisible}
  on:close={() => (cryptoSelectorVisible = false)}
/>

<!-- Payment Method Selector -->
<AccountSelector
  on:close={() => (paymentSelectorVisible = false)}
  visible={paymentSelectorVisible}
/>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/text.scss';

  .dst-container {
    margin-top: 3rem;
    display: flex;
    flex-direction: column;
    height: 5rem;
  }

  .vertical-stepper {
    list-style: none;
    padding: 0;
    li {
      position: relative;
      padding-left: 1.25rem;
      margin-left: 1rem;
      margin-top: 0.5rem;
      // marker
      &:before {
        content: '';
        border: 4px solid $textColor4;
        border-radius: 100%;
        position: absolute;
        height: 0;
        width: 0;
        left: 0;
        right: 0;
        bottom: 0;
        top: 8px;
        z-index: 1;
      }
      // line
      &:after {
        position: absolute;
        width: 1px;
        left: 4px;
        top: -1.25rem;
        opacity: 0.3;
        height: 110%;
        content: '';
        background-color: $textColor4;
        background-position: 0 0;
        background-size: 200% 200%;
        border-color: inherit;
        border-width: 0;
        outline: 0;
      }
      &.success {
        display: flex;
        align-items: center;
        color: var(--theme-text-color-4);
        & > .icon {
          margin-left: 0.4rem;
          color: var(--theme-success-color);
        }
        &:before {
          border: 4px solid var(--theme-success-color) !important;
        }
      }
      &.disabled {
        color: var(--theme-text-color-4);
        cursor: auto !important;
      }
    }
  }
</style>
