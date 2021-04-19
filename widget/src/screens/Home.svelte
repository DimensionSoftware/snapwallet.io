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
  import {
    focusFirstInput,
    isValidNumber,
    onEnterPressed,
    focus,
    resizeWidget,
  } from '../util'
  import TotalContainer from '../components/TotalContainer.svelte'
  import { Routes } from '../constants'
  import {
    faCheck,
    faIdCard,
    faLock,
    faExclamationCircle,
  } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import { TransactionIntents } from '../types'
  import ExchangeRate from '../components/ExchangeRate.svelte'
  import PaymentSelector from '../components/selectors/PaymentSelector.svelte'
  import AccountSelector from '../components/selectors/AccountSelector.svelte'
  import CryptoSelector from '../components/selectors/CryptoSelector.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import VStep from '../components/VStep.svelte'
  import { paymentMethodStore } from '../stores/PaymentMethodStore'
  import { toaster } from '../stores/ToastStore'
  import { configStore } from '../stores/ConfigStore'

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
    getNextPath()
    const { sourceAmount, selectedSourcePaymentMethod } = $transactionStore,
      isLoggedIn = window.AUTH_MANAGER.viewerIsLoggedIn()
    if (
      selectedSourcePaymentMethod &&
      selectedSourcePaymentMethod?.status !== 'ACTIVE'
    ) {
      paymentSelectorVisible = true
      return toaster.pop({
        msg: 'Please select an active payment method.',
        error: true,
      })
    }

    // guards
    if (!sourceAmount || !isValidNumber(sourceAmount)) {
      focus(document.querySelector('input'))
      throw new Error('Input an amount in USD')
    }

    if (sourceAmount < 0.01) {
      focus(document.querySelector('input'))
      throw new Error('The minimum trade amount is $0.01')
    }

    // Only do this when the user has a Wyre account
    if (
      isLoggedIn &&
      (flags?.hasWyreAccount || $userStore.isProfilePending) &&
      !$transactionStore.selectedSourcePaymentMethod
    ) {
      paymentSelectorVisible = true
      return
    }
    // if they're not logged in, forward them instead to login
    if (!isLoggedIn) return push(Routes.SEND_OTP)

    const nextRoute = getNextPath()
    if (nextRoute === Routes.CHECKOUT_OVERVIEW) {
      try {
        isCreatingTxnPreview = true
        const preview = await window.API.fluxWyreCreateTransfer({
          source: $transactionStore.selectedSourcePaymentMethod?.id,
          sourceAmount: $transactionStore.sourceAmount,
          // TODO: get this from app config wallets
          dest: $configStore.wallets.find(
            w =>
              w.asset ===
              $transactionStore.destinationCurrency.ticker.toLowerCase(),
          )?.address,
          // $transactionStore.destinationCurrency.ticker.toLowerCase() !== 'btc'
          //   ? '0xf636B6aA45C554139763Ad926407C02719bc22f7'
          //   : 'n1F9wb29WVFxEZZVDE7idJjpts7qdS8cWU',
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
  const getNextPath = () => {
    if (window.AUTH_MANAGER.viewerIsLoggedIn()) {
      let nextRoute = Routes.PROFILE
      const { hasWyrePaymentMethods, hasWyreAccount } = flags

      if (hasWyrePaymentMethods && hasWyreAccount)
        nextRoute = Routes.CHECKOUT_OVERVIEW
      else if (!hasWyrePaymentMethods) nextRoute = Routes.PLAID_LINK
      else if ($userStore.isProfileComplete) nextRoute = Routes.ADDRESS
      else nextRoute = Routes.PROFILE_STATUS
      return nextRoute
    }
    return Routes.SEND_OTP
  }

  const onKeyDown = (e: Event) => {
    onEnterPressed(e, handleNextStep)
  }

  const getInitialPrices = async () => {
    try {
      setTimeout(animateRandomPrice, 275)
      await priceStore.fetchPrices()
      // Auto input source amount after awaiting prices
      transactionStore.setSourceAmount(
        $configStore.sourceAmount,
        selectedDestinationPrice,
      )
    } finally {
      setTimeout(() => (isLoadingPrices = false), 250)
    }
  }

  onMount(() => {
    resizeWidget(525, $configStore.appName)
    getInitialPrices()
    getNextPath()
    const interval = priceStore.pollPrices()
    return () => clearInterval(interval)
  })
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent animation="right">
  <ModalHeader hideBackButton
    >{isBuy ? 'Buy' : 'Sell'} {destinationCurrency.ticker}</ModalHeader
  >
  <ModalBody>
    <div class="cryptocurrencies-container">
      <div class="dst-container">
        <Label fx={false}>
          <CryptoCard
            on:mousedown={() => (cryptoSelectorVisible = true)}
            crypto={isBuy ? destinationCurrency : sourceCurrency}
          />
        </Label>
      </div>
      <div
        style="display:flex;flex-direction:column;height:5rem;margin-top: -1rem;"
      >
        <Label label="Amount">
          <span class="dst-currency">$</span>
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
        <VStep success={!!$transactionStore.sourceAmount}>
          <span
            class:default-icon={!$transactionStore.sourceAmount}
            slot="icon"
          >
            {#if $transactionStore.sourceAmount}
              <FaIcon data={faCheck} />
            {/if}
          </span>
          <b slot="step">
            <TotalContainer />
          </b>
        </VStep>
        {#if flags?.hasWyreAccount}
          <VStep success>
            <span slot="icon">
              <FaIcon data={faCheck} />
            </span>
            <b slot="step">Verify Identity</b>
          </VStep>
        {:else if $userStore.isProfilePending}
          <VStep>
            <span slot="icon">
              <FaIcon data={faExclamationCircle} />
            </span>
            <b slot="step">Reviewing Identity</b>
            <div class="description help" slot="info">
              We're reviewing your identity.
              {#if !$paymentMethodStore.wyrePaymentMethods?.length}
                Please add a payment method below.
              {/if}
            </div>
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
        <PaymentSelector
          {isBuy}
          onClick={() => (paymentSelectorVisible = true)}
        />
      </ul>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button isLoading={isCreatingTxnPreview} on:mousedown={handleNextStep}>
      <div style="display:flex;justify-content:center;align-items:center;">
        <span style="margin-right:0.75rem;">
          {isCreatingTxnPreview ? 'Previewing' : 'Preview'}
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

  .cryptocurrencies-container {
    padding: 0 0.5rem;
  }
  .dst-container {
    margin-top: 2rem;
    margin-left: 0.5rem;
    margin-right: 0.5rem;
    display: flex;
    flex-direction: column;
    height: 5rem;
  }
  .dst-currency {
    position: absolute;
    left: 1rem;
    bottom: 1.35rem;
    z-index: 5;
    font-size: 0.8rem;
  }
  :global(#amount) {
    text-indent: 1.55rem;
  }

  .vertical-stepper {
    margin-top: 2rem;
    list-style: none;
    padding: 0;
  }

  .description {
    margin-left: 0.55rem;
    color: var(--theme-text-color);
  }
</style>
