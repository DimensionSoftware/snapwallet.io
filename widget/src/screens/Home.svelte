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
  import ModalHeader from '../components/ModalHeader.svelte'
  import VStep from '../components/VStep.svelte'
  import VerticalStepper from '../components/VerticalStepper.svelte'

  let cryptoSelectorVisible = false
  let paymentSelectorVisible = false
  let isLoadingPrices = !Boolean($transactionStore.sourceAmount)

  $: ({
    sourceCurrency,
    destinationCurrency,
    sourceAmount,
    intent,
  } = $transactionStore)

  $: ({ flags } = $userStore)

  $: selectedDirection = `${$transactionStore.sourceCurrency.ticker}_${$transactionStore.destinationCurrency.ticker}`
  $: isBuy = intent === TransactionIntents.BUY

  $: selectedPriceMap = $priceStore.prices[selectedDirection]
  $: selectedDestinationPrice =
    selectedPriceMap[$transactionStore.destinationCurrency.ticker]
  $: exchangeRate = isBuy
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
    if (window.AUTH_MANAGER.viewerIsLoggedIn()) {
      const { flags = {} } = await window.API.fluxViewerData()
      userStore.setFlags(flags)
      const { hasWyrePaymentMethods, hasWyreAccount } = flags

      if (hasWyrePaymentMethods && hasWyreAccount)
        nextRoute = Routes.CHECKOUT_OVERVIEW
      else if (hasWyrePaymentMethods) nextRoute = Routes.PROFILE
    }

    nextRoute = Routes.CHECKOUT_OVERVIEW
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
    <ModalHeader hideBackButton>
      <IntentSelector />
    </ModalHeader>
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
            pattern={`[\\d,\\.]+`}
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
          <ExchangeRate {fakePrice} {isLoadingPrices} {exchangeRate} />
        </Label>
      </div>
      <ul class="vertical-stepper">
        {#if flags?.hasWyreAccount}
          <VStep success>
            <slot name="icon">
              <FaIcon data={faCheck} />
            </slot>
            <b slot="step">Verify Identity</b>
          </VStep>
        {:else}
          <VStep onClick={() => push(Routes.PROFILE)}>
            <span slot="icon">
              <FaIcon data={faIdCard} />
            </span>
            <b slot="step"> Verify Identity </b>
          </VStep>
        {/if}
        <VStep
          disabled={!flags?.hasWyreAccount}
          onClick={() =>
            flags?.hasWyreAccount && (paymentSelectorVisible = true)}
        >
          <span slot="icon">
            <FaIcon data={faUniversity} />
          </span>
          <b slot="step">
            <!-- Multiple PMs will be possible for buy and bank account is only option for sell atm -->
            {#if $transactionStore.selectedSourcePaymentMethod}
              {$transactionStore.selectedSourcePaymentMethod.name}
            {:else if isBuy}
              Select Payment Method
            {:else}
              Select Bank Account
            {/if}
          </b>
        </VStep>
        <VStep>
          <b slot="step">
            <TotalContainer />
          </b>
        </VStep>
      </ul>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button on:click={handleNextStep}>Checkout</Button>
  </ModalFooter>
</ModalContent>

<!-- Cryptocurrency Selector (remount for onMount trigger) -->
{#if cryptoSelectorVisible}
  <CryptoSelector visible on:close={() => (cryptoSelectorVisible = false)} />
{/if}

<!-- Payment Method Selector (remount for onMount trigger) -->
{#if paymentSelectorVisible}
  <AccountSelector visible on:close={() => (paymentSelectorVisible = false)} />
{/if}

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/text.scss';

  .dst-container {
    margin-top: 2rem;
    margin-left: 0.5rem;
    margin-right: 0.5rem;
    display: flex;
    flex-direction: column;
    height: 5rem;
  }

  .vertical-stepper {
    list-style: none;
    padding: 0;
  }
</style>
