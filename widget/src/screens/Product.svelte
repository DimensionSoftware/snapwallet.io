<script lang="ts">
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Button from '../components/Button.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import type { ProductType } from '../types'
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import {
    faCheck,
    faIdCard,
    faLock,
    faUniversity,
  } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  import { onMount } from 'svelte'
  import { priceStore } from '../stores/PriceStore'
  import { formatLocaleCurrency } from '../util'
  import AccountSelector from '../components/selectors/AccountSelector.svelte'
  import VStep from '../components/VStep.svelte'
  import { userStore } from '../stores/UserStore'
  import { transactionStore } from '../stores/TransactionStore'

  export let product: ProductType

  let isPreviewing = false
  let isPaymentSelectorVisible = false

  $: ({ flags } = $userStore)
  $: nextRoute = Routes.PROFILE
  $: priceMap = $priceStore.prices[`USD_${product.destinationTicker}`]
  $: exchangeRate = priceMap[product.destinationTicker]

  const handleNextStep = async () => {
    const { sourceAmount, selectedSourcePaymentMethod } = $transactionStore,
      isLoggedIn = window.AUTH_MANAGER.viewerIsLoggedIn()
    // if they're not logged in, forward them instead to login
    if (!isLoggedIn) return push(Routes.SEND_OTP)

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
      paymentSelectorVisible = true
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
    getViewer()
    getPrices()
    const interval = priceStore.pollPrices()
    getNextPath()
    return clearInterval(interval)
  })
</script>

<ModalContent>
  <ModalHeader isProductCheckout hideBackButton>Buying</ModalHeader>
  <ModalBody>
    <div class="container">
      <h3 class="nft-title">{product.title}</h3>
      <small class="nft-title">by {product.author}</small>
      {#if product.videoURL}
        <video loop playsinline autoplay muted class="nft-video">
          <source src={product.videoURL} />
        </video>
      {/if}
      {#if product.imageURL && !product.videoURL}
        <img alt={product.title} class="nft-image" src={product.imageURL} />
      {/if}
      <ul class="vertical-stepper">
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
          success={$transactionStore.selectedSourcePaymentMethod}
          onClick={() =>
            flags?.hasWyreAccount && (isPaymentSelectorVisible = true)}
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
            {:else}
              Select Payment Method
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
    <Button on:mousedown={handleNextStep} isLoading={isPreviewing}>
      <div class="btn-content">
        <span class="btn-text">Preview</span>
        <FaIcon data={faLock} />
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
      margin: 0;
      & + small {
        margin-bottom: 1rem;
        display: block;
      }
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
    margin-top: 1.75rem;
    list-style: none;
    padding: 0 0.5rem;
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
