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

  $: priceMap = $priceStore.prices[`USD_${product.destinationTicker}`]
  $: exchangeRate = priceMap[product.destinationTicker]

  const fetchPreview = async () => {
    try {
      isPreviewing = true
      const txn = await window.API.fluxWyreCreateTransfer({
        dest: product?.destinationAddress,
        destCurrency: product?.destinationTicker,
        source: $transactionStore.selectedSourcePaymentMethod?.id,
        // TODO: make this optional at the API level
        sourceAmount: 0,
      })
    } finally {
      isPreviewing = false
    }
  }

  // TODO: prefetch or something
  const getPrices = async () => {
    await priceStore.fetchPrices()
  }

  onMount(() => {
    getPrices()
    return clearInterval(priceStore.pollPrices())
  })
</script>

<ModalContent>
  <ModalBody>
    <ModalHeader hideBackButton>Buying</ModalHeader>
    <div class="container">
      <b class="nft-title">{product.title}</b>
      {#if product.videoURL}
        <video loop autoplay muted class="nft-video">
          <source src={product.videoURL} type="video/mp4" />
        </video>
      {/if}
      {#if product.imageURL && !product.videoURL}
        <img
          alt={product.title}
          class="nft-image"
          style=""
          src={product.imageURL}
        />
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
    <Button isLoading={isPreviewing}>
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
      margin-bottom: 1rem;
    }
    .nft-video {
      height: 50%;
    }
    .nft-image {
      height: 50%;
    }
  }

  .vertical-stepper {
    margin-top: 0.5rem;
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
