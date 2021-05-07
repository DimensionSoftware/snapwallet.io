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
    faGlobe,
  } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import { TransactionIntents, TransactionMediums } from '../types'
  import ExchangeRate from '../components/ExchangeRate.svelte'
  import PaymentSelector from '../components/selectors/PaymentSelector.svelte'
  import AccountSelector from '../components/selectors/AccountSelector.svelte'
  import CryptoSelector from '../components/selectors/CryptoSelector.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import VStep from '../components/VStep.svelte'
  import { paymentMethodStore } from '../stores/PaymentMethodStore'
  import { toaster } from '../stores/ToastStore'
  import { configStore } from '../stores/ConfigStore'
  import CountrySelector from '../components/selectors/CountrySelector.svelte'
  import { debitCardStore } from '../stores/DebitCardStore'
  import { countries, WYRE_SUPPORTED_COUNTRIES } from '../util/country'
  import { getMissingFieldMessages } from '../util/profiles'

  let cryptoSelectorVisible = false
  let paymentSelectorVisible = false
  let countrySelectorVisible = false
  let isLoadingPrices = !Boolean($transactionStore.sourceAmount)
  let glow = false
  let isLoggedIn = window.AUTH_MANAGER.viewerIsLoggedIn()

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

  $: country = countries[$debitCardStore.address.country]
  $: missingInfo = getMissingFieldMessages($userStore.profileItems)

  let verificationNextStep
  $: {
    // NOTE: these should remain in this order

    if (!missingInfo.personal.isComplete) {
      verificationNextStep = Routes.PROFILE
    } else if (!missingInfo.address.isComplete) {
      verificationNextStep = Routes.ADDRESS
    } else if (!missingInfo.contact.isComplete) {
      verificationNextStep = Routes.PROFILE_SEND_SMS
    } else if (!missingInfo.document.isComplete) {
      verificationNextStep = Routes.FILE_UPLOAD
    } else {
      verificationNextStep = Routes.PROFILE_STATUS
    }
  }

  const animateRandomPrice = () => {
    window.requestAnimationFrame(_ts => {
      if (isLoadingPrices) {
        fakePrice = fakePrice + 213.02
        animateRandomPrice()
      }
    })
  }

  const processDebitTransaction = async (isLoggedIn: boolean) => {
    if (!isLoggedIn) push(Routes.SEND_OTP)
    try {
      isCreatingTxnPreview = true
      const dest = // TODO: move srn prefix to server
        $transactionStore.destinationCurrency.ticker.toLowerCase() !== 'btc'
          ? '0xf636B6aA45C554139763Ad926407C02719bc22f7'
          : 'n1F9wb29WVFxEZZVDE7idJjpts7qdS8cWU'
      const {
        reservationId,
        quote,
      } = await window.API.fluxWyreCreateDebitCardQuote({
        dest,
        sourceCurrency: $transactionStore.sourceCurrency.ticker,
        lockFields: ['sourceAmount'],
        amountIncludesFees: false,
        country: $debitCardStore.address.country,
        sourceAmount: $transactionStore.sourceAmount,

        destCurrency: $transactionStore.destinationCurrency?.ticker,
      })

      debitCardStore.update({ reservationId, dest })
      transactionStore.setWyrePreview(quote)
      return push(Routes.CHECKOUT_OVERVIEW)
    } finally {
      isCreatingTxnPreview = false
    }
  }

  const handleNextStep = async () => {
    const { sourceAmount, selectedSourcePaymentMethod } = $transactionStore

    isLoggedIn = window.AUTH_MANAGER.viewerIsLoggedIn()

    // guards
    if (!sourceAmount || !isValidNumber(sourceAmount)) {
      focus(document.querySelector('input'))
      throw new Error(`Input an amount to ${isBuy ? 'Buy' : 'Sell'} in USD.`)
    }

    if (sourceAmount < 0.01) {
      focus(document.querySelector('input'))
      throw new Error('The minimum trade amount is $0.01')
    }

    if ($transactionStore.inMedium === TransactionMediums.DEBIT_CARD) {
      return await processDebitTransaction(isLoggedIn)
    }

    getNextPath()

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
          // dest: $configStore.wallets.find(
          //   w =>
          //     w.asset ===
          //     $transactionStore.destinationCurrency.ticker.toLowerCase(),
          // )?.address,
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
  const getNextPath = () => {
    if (window.AUTH_MANAGER.viewerIsLoggedIn()) {
      const { hasWyrePaymentMethods, hasWyreAccount } = flags
      if (hasWyrePaymentMethods && hasWyreAccount)
        return Routes.CHECKOUT_OVERVIEW
      else if (!hasWyrePaymentMethods) return Routes.PLAID_LINK
      else return verificationNextStep
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
      if (!$transactionStore.sourceAmount) {
        transactionStore.setSourceAmount(
          $configStore.sourceAmount,
          selectedDestinationPrice,
        )
      }
    } finally {
      setTimeout(() => (isLoadingPrices = false), 250)
    }
  }

  onMount(() => {
    resizeWidget(525, $configStore.appName)
    getInitialPrices()
    getNextPath()
    const interval = priceStore.pollPrices()
    if (window.AUTH_MANAGER.viewerIsLoggedIn()) {
      // Profile should be updated when user comes back here from any other route
      userStore.fetchUserProfile()
      userStore.fetchFlags()
    }
    // TODO: @khoerling if ($configStore.intent === 'donate') {transactionStore.update({inMedium: TransactionMediums.DEBIT_CARD})}
    return () => clearInterval(interval)
  })

  $: hasCountryIcon = WYRE_SUPPORTED_COUNTRIES.includes(
    country?.code?.toUpperCase(),
  )
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
              glow = !!val
              transactionStore.setSourceAmount(val, selectedDestinationPrice)
            }}
            defaultValue={sourceAmount
              ? sourceAmount
              : $configStore.sourceAmount}
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
        <PaymentSelector
          {isBuy}
          onClick={() => (paymentSelectorVisible = true)}
        />
        {#if $transactionStore.inMedium === TransactionMediums.ACH}
          {#if $userStore.isProfilePending}
            <VStep disabled>
              <span class="glow" slot="icon">
                <FaIcon data={faExclamationCircle} />
              </span>
              <b slot="step">Reviewing Identity</b>
              <div class="description help" slot="info">
                We're reviewing your identity. This should only take a few
                minutes.
                {#if !$paymentMethodStore.wyrePaymentMethods?.length}
                  Please add a payment method.
                {/if}
              </div>
            </VStep>
          {:else if flags?.hasWyreAccount}
            <VStep success>
              <span slot="icon">
                <FaIcon data={faCheck} />
              </span>
              <b slot="step">Verify Identity</b>
            </VStep>
          {:else}
            <VStep disabled onClick={() => push(verificationNextStep)}>
              <span
                class:glow={$transactionStore.selectedSourcePaymentMethod}
                slot="icon"
              >
                <FaIcon data={faIdCard} />
              </span>
              <b slot="step"> Verify Identity </b>
            </VStep>
          {/if}
        {:else if $transactionStore.inMedium === TransactionMediums.DEBIT_CARD}
          <VStep
            disabled
            custom={!!hasCountryIcon}
            success={!!country}
            title="Select Payment Country"
            onClick={() => {
              countrySelectorVisible = true
            }}
          >
            <span class:glow={!$debitCardStore.address.country} slot="icon">
              <FaIcon data={country ? faCheck : faGlobe} />
            </span>
            <b slot="step">
              {#if country}
                {`${country.name}`}
                &nbsp;<small>( change )</small>
              {:else}
                Select Payment Country
              {/if}
            </b>
          </VStep>
        {/if}
      </ul>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button
      {glow}
      isLoading={isCreatingTxnPreview}
      on:mousedown={handleNextStep}
    >
      <div style="display:flex;justify-content:center;align-items:center;">
        <span style="margin-right:0.75rem;">
          {isCreatingTxnPreview
            ? 'Previewing'
            : isLoggedIn
            ? 'Preview'
            : 'Continue'}
        </span>
        {#if isLoggedIn}
          <FaIcon data={faLock} />
        {/if}
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

{#if countrySelectorVisible}
  <CountrySelector
    whiteList={WYRE_SUPPORTED_COUNTRIES}
    on:close={() => (countrySelectorVisible = false)}
    on:select={e => {
      const { country } = e?.detail
      if (country) {
        userStore.setPhoneNumberCountry(country)
        debitCardStore.updateAddress({ country: country.code })
      }
      countrySelectorVisible = false
    }}
  />
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
    :global(.flag > svg) {
      position: absolute;
      left: -12px;
      z-index: 1;
    }
  }

  .description {
    margin-left: 0.55rem;
    color: var(--theme-text-color);
  }
</style>
