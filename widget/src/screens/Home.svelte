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
    getPrimaryPaymentMethodID,
    isValidNumber,
    onEnterPressed,
    focus,
  } from '../util'
  import TotalContainer from '../components/TotalContainer.svelte'
  import { Routes } from '../constants'
  import {
    faCheck,
    faIdCard,
    faLock,
    faUniversity,
    faSpinner,
    faExclamationCircle,
  } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import { TransactionIntents } from '../types'
  import ExchangeRate from '../components/ExchangeRate.svelte'
  import AccountSelector from '../components/selectors/AccountSelector.svelte'
  import CryptoSelector from '../components/selectors/CryptoSelector.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import VStep from '../components/VStep.svelte'
  import { paymentMethodStore } from '../stores/PaymentMethodStore'
  import { toaster } from '../stores/ToastStore'

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
    const { sourceAmount, selectedSourcePaymentMethod } = $transactionStore,
      isLoggedIn = window.AUTH_MANAGER.viewerIsLoggedIn()
    if (
      selectedSourcePaymentMethod &&
      selectedSourcePaymentMethod?.status !== 'ACTIVE'
    ) {
      return toaster.pop({
        msg: 'Please select an active payment method.',
        error: true,
      })
    }

    // guards
    if (!sourceAmount || !isValidNumber(sourceAmount)) {
      focus(document.querySelector('input'))
      throw new Error('Input an Amount in USD')
    }
    if (isLoggedIn && !$transactionStore.selectedSourcePaymentMethod) {
      paymentSelectorVisible = true
      return
    }
    // if they're not logged in, forward them instead to login
    if (!isLoggedIn) return push(Routes.SEND_OTP)

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
    const isLoggedIn = window.AUTH_MANAGER.viewerIsLoggedIn()
    if (isLoggedIn) {
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

  onMount(() => {
    getInitialPrices()
    getNextPath()
    // cleanup
    return () => clearInterval(priceStore.pollPrices())
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
        <VStep
          title="Select Your Payment Method"
          disabled={!$userStore.isProfilePending && !flags?.hasWyreAccount}
          success={$transactionStore.selectedSourcePaymentMethod}
          onClick={() =>
            ($userStore.isProfilePending || flags?.hasWyreAccount) &&
            (paymentSelectorVisible = true)}
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
        </VStep>
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
  }
</style>
