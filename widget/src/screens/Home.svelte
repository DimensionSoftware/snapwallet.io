<script lang="ts">
  import { push } from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import IntentSelector from '../components/IntentSelector.svelte'
  import { toaster } from '../stores/ToastStore'
  import Button from '../components/Button.svelte'
  import { userStore } from '../stores/UserStore'
  import PopupSelector from '../components/inputs/PopupSelector.svelte'
  import CryptoCard from '../components/cards/CryptoCard.svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { priceStore } from '../stores/PriceStore'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import { onMount } from 'svelte'

  let selectorVisible = false
  const handleNextStep = () => {
    push('/checkout')
  }

  const cryptoCurrencies = [
    { name: 'Bitcoin', ticker: 'BTC' },
    { name: 'Ethereum', ticker: 'ETH' },
    { name: 'Tether', ticker: 'USDT' },
    { name: 'USDC', ticker: 'USDC' },
  ]

  $: selectedDirection = `${$transactionStore.sourceCurrency.ticker}_${$transactionStore.destinationCurrency.ticker}`
  $: selectedPriceMap = $priceStore.prices[selectedDirection]
  $: selectedSourcePrice =
    selectedPriceMap[$transactionStore.sourceCurrency.ticker]

  onMount(async () => {
    try {
      await priceStore.fetchPrices()
      priceStore.pollPrices()
    } catch (e) {
      toaster.pop({
        msg: 'Oops, there was a problem refreshing prices.',
        error: true,
      })
    }
  })
</script>

<ModalContent>
  <ModalBody>
    <IntentSelector />
    <div class="cryptocurrencies-container">
      <div
        style="display:flex;flex-direction:column;height:5rem;margin-bottom:1rem"
      >
        <Label>Currency</Label>
        <CryptoCard
          on:click={() => (selectorVisible = true)}
          crypto={$transactionStore.destinationCurrency}
        />
      </div>
      <div style="display:flex;flex-direction:column;height:5rem;">
        <div style="display:flex;justify-content:space-between;cursor:pointer;">
          <Label>Amount</Label>
          <Label>
            <b>{$transactionStore.sourceCurrency.ticker}</b>
            /
            <span class="muted"
              >{$transactionStore.destinationCurrency.ticker}</span
            >
          </Label>
        </div>
        <Input
          on:change={e => {
            transactionStore.setSourceAmount(Number(e.detail))
          }}
          defaultValue={$transactionStore.sourceAmount}
          type="number"
          placeholder="Amount"
        />
      </div>
      <div class="exchange-rate-container">
        ~ 1 {$transactionStore.destinationCurrency.ticker} @ {selectedSourcePrice}
        {$transactionStore.sourceCurrency.ticker}
      </div>
      <div class="total-container">
        ~ {($transactionStore.sourceAmount / selectedSourcePrice).toFixed(8)}
        {$transactionStore.destinationCurrency.ticker}
      </div>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button
      disabled={!$transactionStore.sourceAmount ||
        isNaN($transactionStore.sourceAmount)}
      on:click={handleNextStep}>Checkout</Button
    >
  </ModalFooter>
</ModalContent>

<PopupSelector
  on:close={() => (selectorVisible = false)}
  visible={selectorVisible}
  headerTitle="Select Cryptocurrency"
>
  <div class="cryptocurrencies-container">
    {#each cryptoCurrencies as cryptoCurrency (cryptoCurrency.ticker)}
      <div style="margin: 0.5rem 0">
        <CryptoCard
          on:click={() => (selectorVisible = false)}
          crypto={cryptoCurrency}
        />
      </div>
    {/each}
  </div>
</PopupSelector>

<style lang="scss">
  @import '../styles/_vars.scss';

  .cryptocurrencies-container {
    height: 100%;
    width: 100%;
    overflow: hidden;
    overflow-y: scroll;
    padding: 0 0.5rem;
    margin-top: 2rem;
  }

  .exchange-rate-container {
    display: flex;
    justify-content: flex-end;
    font-size: 0.9rem;
    color: var(--theme-text-color-muted);
  }

  .total-container {
    display: flex;
    justify-content: flex-end;
    color: var(--theme-text-color);
    font-weight: 500;
    font-size: 0.9rem;
  }

  .muted {
    color: var(--theme-text-color-muted);
  }
</style>
