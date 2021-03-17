<script lang="ts">
  import { push } from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import IntentSelector from '../components/IntentSelector.svelte'
  import Button from '../components/Button.svelte'
  import PopupSelector from '../components/inputs/PopupSelector.svelte'
  import CryptoCard from '../components/cards/CryptoCard.svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { userStore } from '../stores/UserStore'
  import { priceStore } from '../stores/PriceStore'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import { onMount } from 'svelte'
  import { isValidNumber, onEnterPressed, formatLocaleCurrency } from '../util'
  import TotalContainer from '../components/TotalContainer.svelte'
  import { Routes } from '../constants'
  import IconCard from '../components/cards/IconCard.svelte'
  import { faUniversity } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'

  let selectorVisible = false
  let paymentSelectorVisible = false

  const cryptoCurrencies = [
    { name: 'Bitcoin', ticker: 'BTC', popular: true },
    { name: 'Ethereum', ticker: 'ETH', popular: true },
    { name: 'USDC', ticker: 'USDC' },
    { name: 'Tether', ticker: 'USDT', popular: true },
    { name: 'DAI', ticker: 'DAI' },
    { name: 'MakerDAO', ticker: 'MKR' },
    { name: 'Gemini Dollar', ticker: 'GUSD' },
    { name: 'Paxos Standard', ticker: 'PAX' },
    { name: 'Link', ticker: 'LINK' },
  ]

  $: selectedDirection = `${$transactionStore.sourceCurrency.ticker}_${$transactionStore.destinationCurrency.ticker}`
  $: selectedPriceMap = $priceStore.prices[selectedDirection]
  $: selectedSourcePrice =
    selectedPriceMap[$transactionStore.sourceCurrency.ticker]
  $: selectedDestinationPrice =
    selectedPriceMap[$transactionStore.destinationCurrency.ticker]
  $: exchangeRate = 1 / selectedDestinationPrice
  $: destinationRate = $transactionStore.sourceAmount * selectedDestinationPrice
  $: sourceRate = $transactionStore.destinationAmount / selectedDestinationPrice

  let isEnteringSourceAmount = true
  let isLoadingPrices = !Boolean($transactionStore.sourceAmount)

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
    if (!sourceAmount || isNaN(sourceAmount) || !isValidNumber(destinationRate))
      // focus input
      return document.querySelector('input')?.focus()

    push(nextRoute)
  }

  // Find the next path based on user data
  const getNextPath = async () => {
    // TODO: move this request somewhere sane
    const { flags = {} } = await window.API.fluxViewerData()
    userStore.setFlags(flags)
    const { hasPlaidItems, hasWyreAccount } = flags

    if (hasPlaidItems && hasWyreAccount) nextRoute = Routes.CHECKOUT_OVERVIEW
    else if (hasPlaidItems) nextRoute = Routes.PROFILE
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

  const srcTicker = $transactionStore.sourceCurrency.ticker
</script>

<svelte:window on:keydown={onKeyDown} />

<ModalContent>
  <ModalBody>
    <IntentSelector />
    <div class="cryptocurrencies-container">
      <div class="dst-container">
        <Label>
          <CryptoCard
            on:click={() => (selectorVisible = true)}
            crypto={$transactionStore.destinationCurrency}
          />
        </Label>
      </div>
      <div style="display:flex;flex-direction:column;height:5rem;">
        <div class="dstCurrency">
          <Label>
            <span
              class:bold-pointer={isEnteringSourceAmount}
              class:muted={!isEnteringSourceAmount}
              on:click={() => {
                const sourceAmount = Number(sourceRate.toFixed(2))
                transactionStore.setSourceAmount(sourceAmount)
                isEnteringSourceAmount = true
              }}
            >
              {srcTicker}
            </span>
            /
            <span
              class:bold-pointer={!isEnteringSourceAmount}
              class:muted={isEnteringSourceAmount}
              on:click={() => {
                transactionStore.setDestinationAmount(destinationRate)
                isEnteringSourceAmount = false
              }}
            >
              {$transactionStore.destinationCurrency.ticker}
            </span>
          </Label>
        </div>
        <Label label="Amount">
          <Input
            pattern={`[\\d,\\.]*`}
            maskChar="[\d,\.]"
            on:change={e => {
              const val = Number(e.detail)
              isEnteringSourceAmount
                ? transactionStore.setSourceAmount(val)
                : transactionStore.setDestinationAmount(val)
            }}
            defaultValue={isEnteringSourceAmount
              ? $transactionStore.sourceAmount
              : $transactionStore.destinationAmount}
            required
            type="number"
            inputmode="number"
            placeholder="0"
          />
        </Label>
      </div>
      <ul class="vertical-stepper">
        <li class="exchange-rate-container">
          1 {$transactionStore.destinationCurrency.ticker} ≈
          {isLoadingPrices
            ? formatLocaleCurrency(
                $transactionStore.sourceCurrency.ticker,
                fakePrice,
              )
            : formatLocaleCurrency(
                $transactionStore.sourceCurrency.ticker,
                exchangeRate,
              )}
          {srcTicker}
        </li>
        <li>
          <TotalContainer
            rate={isEnteringSourceAmount ? destinationRate : sourceRate}
            isDestination={isEnteringSourceAmount}
          />
        </li>
        <li
          style="cursor:pointer;display:flex;align-items:center;"
          on:click={() => (paymentSelectorVisible = true)}
        >
          <FaIcon data={faUniversity} />
          <b style="margin-left:0.5rem;text-decoration:underline"
            >Select Payment</b
          >
        </li>
      </ul>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button on:click={handleNextStep}>Checkout</Button>
  </ModalFooter>
</ModalContent>

<PopupSelector
  on:close={() => (selectorVisible = false)}
  visible={selectorVisible}
  headerTitle="Select Currency"
>
  <div class="scroll cryptocurrencies-container">
    <h5>Popular</h5>
    {#each cryptoCurrencies.filter(c => c.popular) as cryptoCurrency (cryptoCurrency.ticker)}
      <div style="margin: 0.5rem 0">
        <Label>
          <CryptoCard
            on:click={() => (selectorVisible = false)}
            crypto={cryptoCurrency}
          />
        </Label>
      </div>
    {/each}
    <h5 style="margin-top: 1.25rem">All</h5>
    {#each cryptoCurrencies.filter(c => !c.popular) as cryptoCurrency (cryptoCurrency.ticker)}
      <div style="margin: 0.5rem 0">
        <Label>
          <CryptoCard
            on:click={() => (selectorVisible = false)}
            crypto={cryptoCurrency}
          />
        </Label>
      </div>
    {/each}
  </div>
</PopupSelector>

<!-- TODO: handle pm selection here -->
<PopupSelector
  on:close={() => {
    paymentSelectorVisible = false
  }}
  visible={paymentSelectorVisible}
  headerTitle="Payment Methods"
>
  <div class="scroll cryptocurrencies-container">
    <h5>Add a Payment Method</h5>
    <IconCard
      icon={faUniversity}
      on:click={() => push(Routes.PLAID_LINK)}
      label="Bank Account"
    />
    <h5 style="margin-top: 1.25rem">Select a Payment Method</h5>
    {#if !$userStore.flags?.hasPlaidItems}
      <p class="help">No payment methods available</p>
    {/if}
  </div>
</PopupSelector>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/text.scss';

  h5 {
    margin: 0;
  }
  .scroll {
    overflow-y: scroll;
  }
  .cryptocurrencies-container {
    position: relative;
    height: 100%;
    width: 100%;
    padding: 0 0.5rem;
    margin-top: 1rem;
  }

  .exchange-rate-container {
    height: 1.5rem;
    position: relative;
    z-index: 2;
    font-size: 0.9rem;
    color: var(--theme-text-color);
  }

  .muted {
    cursor: pointer;
    color: var(--theme-text-color-muted);
  }

  .bold-pointer {
    font-weight: bold;
  }

  span {
    width: 35px;
  }

  .dst-container {
    margin-top: 1rem;
    display: flex;
    flex-direction: column;
    height: 5rem;
    margin-bottom: 1rem;
  }
  .dstCurrency {
    position: absolute;
    right: 0.5rem;
    display: flex;
    justify-content: space-between;
    cursor: pointer;
  }

  .vertical-stepper {
    list-style: none;
    padding: 0;
    li {
      position: relative;
      padding-left: 1.25rem;
      margin-left: 1rem;
      margin-top: 0.25rem;
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
    }
  }
</style>
