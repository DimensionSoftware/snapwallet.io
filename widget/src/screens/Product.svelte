<script lang="ts">
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Button from '../components/Button.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import {
    faCheck,
    faGlobe,
    faIdCard,
    faLock,
    faUniversity,
  } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import { onMount } from 'svelte'
  import { priceStore } from '../stores/PriceStore'
  import { formatLocaleCurrency } from '../util'
  import AccountSelector from '../components/selectors/AccountSelector.svelte'
  import CountrySelector from '../components/selectors/CountrySelector.svelte'
  import VStep from '../components/VStep.svelte'
  import { userStore } from '../stores/UserStore'
  import { transactionStore } from '../stores/TransactionStore'
  import { configStore } from '../stores/ConfigStore'
  import { toaster } from '../stores/ToastStore'
  import { TransactionMediums } from '../types'
  import { countries, WYRE_SUPPORTED_COUNTRIES } from '../util/country'
  import { debitCardStore } from '../stores/DebitCardStore'

  $: ({ product } = $configStore)
  $: isDebitCard = $transactionStore.inMedium === TransactionMediums.DEBIT_CARD

  let isPreviewing = false
  let isPaymentSelectorVisible = false
  let countrySelectorVisible = false

  $: ({ flags } = $userStore)
  $: nextRoute = Routes.PROFILE
  $: priceMap = $priceStore.prices[`USD_${product.destinationTicker}`] || {}
  $: exchangeRate = priceMap[product.destinationTicker] || 0

  let selectedCountryCode
  $: {
    if (isDebitCard) {
      selectedCountryCode =
        $debitCardStore.address.country || $userStore.geo.country
    } else {
      selectedCountryCode = $userStore.address.country || $userStore.geo.country
    }
  }

  const handleNextStep = async () => {
    const { selectedSourcePaymentMethod } = $transactionStore,
      isLoggedIn = window.AUTH_MANAGER.viewerIsLoggedIn()
    // if they're not logged in, forward them instead to login
    if (!isLoggedIn) return push(Routes.SEND_OTP)

    if (isDebitCard) {
      try {
        isPreviewing = true
        const dest = // TODO: move srn prefix to server
          product.destinationTicker.toLowerCase() !== 'btc'
            ? '0xf636B6aA45C554139763Ad926407C02719bc22f7'
            : 'n1F9wb29WVFxEZZVDE7idJjpts7qdS8cWU'
        const { reservationId, quote } =
          await window.API.fluxWyreCreateDebitCardQuote({
            dest,
            sourceCurrency: $transactionStore.sourceCurrency.ticker,
            lockFields: ['destAmount'],
            amountIncludesFees: false,
            country: $debitCardStore.address.country,
            destAmount: product.destinationAmount,
            destCurrency: product.destinationTicker,
          })

        debitCardStore.update({ reservationId, dest })
        transactionStore.setWyrePreview(quote)
        return push(Routes.CHECKOUT_OVERVIEW)
      } finally {
        isPreviewing = false
      }
    }

    if (
      selectedSourcePaymentMethod &&
      selectedSourcePaymentMethod?.status !== 'ACTIVE'
    ) {
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
      isPaymentSelectorVisible = true
      return
    }

    if (nextRoute === Routes.CHECKOUT_OVERVIEW) {
      try {
        isPreviewing = true
        const preview = await window.API.fluxWyreCreateTransfer({
          source: $transactionStore.selectedSourcePaymentMethod?.id,
          destAmount: product.destinationAmount,
          dest: product.destinationAddress,
          destCurrency: product.destinationTicker,
        })

        transactionStore.setWyrePreview(preview)
      } finally {
        isPreviewing = false
      }
    }

    push(nextRoute)
  }

  // Find the next path based on user data
  const getNextPath = async () => {
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
      else if (flags?.hasWyreAccount && !hasWyrePaymentMethods)
        nextRoute = Routes.PLAID_LINK
      return
    }

    nextRoute = Routes.SEND_OTP
  }

  // TODO: prefetch or something
  const getPrices = async () => {
    await priceStore.fetchPrices()
  }

  const getViewer = async () => {
    const isLoggedIn = window.AUTH_MANAGER.viewerIsLoggedIn()
    if (isLoggedIn) {
      await userStore.fetchFlags()
    }
  }

  onMount(() => {
    // select debit by default
    transactionStore.update({ inMedium: TransactionMediums.DEBIT_CARD })
    getViewer()
    getPrices()
    const interval = priceStore.pollPrices()
    getNextPath()
    return clearInterval(interval)
  })
</script>

<ModalContent>
  <ModalHeader hideBackButton>{product.title}</ModalHeader>
  <ModalBody>
    <div class="container">
      {#if product.author}
        <h4 class="nft-title">by {product.author}</h4>
      {/if}
      {#if product.videoURL}
        <video loop playsinline autoplay muted class="nft-video">
          <source src={product.videoURL} />
        </video>
      {/if}
      {#if product.imageURL && !product.videoURL}
        <img alt={product.title} class="nft-image" src={product.imageURL} />
      {/if}
      <ul class="vertical-stepper">
        {#if !isDebitCard && $transactionStore.selectedSourcePaymentMethod}
          {#if flags?.hasWyreAccount}
            <VStep success>
              <span slot="icon">
                <FaIcon data={faCheck} />
              </span>
              <b slot="step">Verify Identity</b>
            </VStep>
          {:else}
            <VStep
              onClick={() =>
                push(
                  $userStore.isProfileComplete
                    ? Routes.ADDRESS
                    : Routes.PROFILE,
                )}
            >
              <span slot="icon">
                <FaIcon data={faIdCard} />
              </span>
              <b slot="step"> Verify Identity </b>
            </VStep>
          {/if}
        {/if}
        <VStep
          title="Click to Change Payment Method"
          success={isDebitCard ||
            Boolean($transactionStore.selectedSourcePaymentMethod)}
          onClick={() => (isPaymentSelectorVisible = true)}
        >
          <span slot="icon">
            <FaIcon
              data={!isDebitCard &&
              !$transactionStore.selectedSourcePaymentMethod
                ? faUniversity
                : faCheck}
            />
          </span>
          <b slot="step">
            <!-- Multiple PMs will be possible for buy and bank account is only option for sell atm -->
            {#if $transactionStore.selectedSourcePaymentMethod}
              {$transactionStore.selectedSourcePaymentMethod.name}
            {:else if isDebitCard}
              Pay with Debit Card
            {:else}
              Select Payment Method
            {/if}
          </b>
        </VStep>
        {#if isDebitCard}
          <VStep
            disabled
            onClick={() => {
              countrySelectorVisible = true
            }}
            success={Boolean($debitCardStore.address.country)}
          >
            <span slot="icon">
              <FaIcon
                data={$debitCardStore.address.country ? faCheck : faGlobe}
              />
            </span>
            <b slot="step"
              >{countries[$debitCardStore.address.country]?.name ||
                'Select Location'}</b
            >
          </VStep>
        {/if}
        <VStep success={!!$transactionStore.sourceAmount}>
          <span
            class:default-icon={!$transactionStore.sourceAmount}
            slot="icon"
          >
            {#if $transactionStore.sourceAmount}
              <FaIcon data={faCheck} />
            {/if}
          </span>
          <div slot="step">
            Subtotal ≈ {formatLocaleCurrency(
              'USD',
              product.destinationAmount / exchangeRate,
            )}
          </div>
        </VStep>
      </ul>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button glow on:mousedown={handleNextStep} isLoading={isPreviewing}>
      <div class="btn-content">
        <span class="btn-text">Buy</span>
      </div>
    </Button>
  </ModalFooter>
</ModalContent>

<!-- Payment Method Selector (remount for onMount trigger) -->
{#if isPaymentSelectorVisible}
  <AccountSelector
    visible
    on:close={() => (isPaymentSelectorVisible = false)}
  />
{/if}

{#if countrySelectorVisible}
  <CountrySelector
    {selectedCountryCode}
    whiteList={WYRE_SUPPORTED_COUNTRIES}
    on:close={() => (countrySelectorVisible = false)}
    on:select={e => {
      const { country } = e?.detail
      country && debitCardStore.updateAddress({ country: country.code })
      countrySelectorVisible = false
    }}
  />
{/if}

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';

  .container {
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    .nft-title {
      margin: -1rem 0 1rem;
      font-weight: 300;
    }
    .nft-video {
      position: relative;
      z-index: 1;
      height: 50%;
    }
    .nft-image {
      position: relative;
      z-index: 1;
      height: 50%;
    }
  }

  .vertical-stepper {
    margin-top: 2rem;
    list-style: none;
    padding: 0 0.5rem;
    :global(li) {
      margin-top: 0.25rem !important;
    }
  }

  .btn-content {
    display: flex;
    justify-content: center;
    align-items: center;
    .btn-text {
      margin-right: 0.75rem;
    }
  }
</style>
