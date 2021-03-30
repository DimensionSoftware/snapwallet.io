<script lang="ts">
  import { push } from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import CryptoCard from '../components/cards/CryptoCard.svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { userStore } from '../stores/UserStore'
  import { priceStore } from '../stores/PriceStore'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import { onMount } from 'svelte'
  import { focusFirstInput, isValidNumber, onEnterPressed } from '../util'
  import TotalContainer from '../components/TotalContainer.svelte'
  import { Routes } from '../constants'
  import {
    faCheck,
    faIdCard,
    faLock,
    faUniversity,
  } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import { TransactionIntents } from '../types'
  import ExchangeRate from '../components/ExchangeRate.svelte'
  import AccountSelector from '../components/selectors/AccountSelector.svelte'
  import CryptoSelector from '../components/selectors/CryptoSelector.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import VStep from '../components/VStep.svelte'

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

  $: fakePrice = 1_000
  $: nextRoute = Routes.PROFILE
  $: isCreatingTxnPreview = false

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

    // guards
    if (!sourceAmount || !isValidNumber(sourceAmount))
      return document.querySelector('input')?.focus()
    if (!$transactionStore.selectedSourcePaymentMethod)
      throw new Error('Select a Payment Method:')

    if (nextRoute === Routes.CHECKOUT_OVERVIEW) {
      try {
        isCreatingTxnPreview = true
        const preview = await window.API.fluxWyreCreateTransfer({
          source: $transactionStore.selectedSourcePaymentMethod?.id,
          sourceAmount: $transactionStore.sourceAmount,
          // TODO: get this from app config wallets
          dest:
            $transactionStore.destinationCurrency.ticker.toLowerCase() !== 'btc'
              ? '0xf636B6aA45C554139763Ad926407C02719bc22f7'
              : 'n1F9wb29WVFxEZZVDE7idJjpts7qdS8cWU',
          destCurrency: $transactionStore.destinationCurrency?.ticker,
        })

        transactionStore.setWyrePreview(preview)
      } finally {
        isCreatingTxnPreview = false
      }
    }

    push(nextRoute)
  }

  // Find the next path based on user data
  const getNextPath = async () => {
    // TODO: move this request somewhere sane
    if (window.AUTH_MANAGER.viewerIsLoggedIn()) {
      const { flags = {}, user = {} } = await window.API.fluxViewerData()
      userStore.setFlags({
        ...flags,
        hasEmail: Boolean(user.email),
        hasPhone: Boolean(user.phone),
      })
      const { hasWyrePaymentMethods, hasWyreAccount } = flags

      if (hasWyrePaymentMethods && hasWyreAccount)
        nextRoute = Routes.CHECKOUT_OVERVIEW
      else if ($userStore.isProfileComplete) nextRoute = Routes.ADDRESS
      else if (!hasWyrePaymentMethods) nextRoute = Routes.PLAID_LINK
      return
    }

    nextRoute = Routes.SEND_OTP
  }

  const onKeyDown = (e: Event) => {
    onEnterPressed(e, handleNextStep)
  }

  const getInitialPrices = async () => {
    try {
      setTimeout(animateRandomPrice, 275)
      await priceStore.fetchPrices()
    } finally {
      setTimeout(() => (isLoadingPrices = false), 250)
    }
  }

  onMount(() => {
    getInitialPrices()
    getNextPath()
    // cleanup
    return () => clearInterval(priceStore.pollPrices())
  })
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent animation="right">
  <ModalBody>
    <ModalHeader hideBackButton>
      Buy {destinationCurrency.name}
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
            id="amount"
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
            <span name="icon">
              <FaIcon data={faCheck} />
            </span>
            <b slot="step">Verify Identity</b>
          </VStep>
        {:else}
          <VStep
            onClick={() =>
              push(
                $userStore.isProfileComplete ? Routes.ADDRESS : Routes.PROFILE,
              )}
          >
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
    <Button isLoading={isCreatingTxnPreview} on:click={handleNextStep}>
      <div style="display:flex;justify-content:center;align-items:center;">
        <span style="margin-right:0.75rem;">
          {#if isCreatingTxnPreview}
            Checking Out
          {:else}
            Checkout
          {/if}
        </span>
        <FaIcon data={faLock} />
      </div>
    </Button>
  </ModalFooter>
</ModalContent>

<!-- Cryptocurrency Selector (remount for onMount trigger) -->
{#if cryptoSelectorVisible}
  <CryptoSelector
    visible
    on:close={() => {
      cryptoSelectorVisible = false
      focusFirstInput(150)
    }}
  />
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
    margin-top: 2.5rem;
    list-style: none;
    padding: 0;
  }
</style>
