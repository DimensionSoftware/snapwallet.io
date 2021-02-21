<script lang="ts">
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
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'

  let selectorVisible = false
  const handleNextStep = () => {
    toaster.pop({ msg: 'Success', success: true })
  }

  const cryptoCurrencies = [
    { name: 'Bitcoin', ticker: 'BTC' },
    { name: 'Ethereum', ticker: 'ETH' },
    { name: 'Tether', ticker: 'USDT' },
    { name: 'USDC', ticker: 'USDC' },
  ]
</script>

<ModalContent>
  <ModalBody>
    <IntentSelector />
    <div class="cryptocurrencies-container">
      <div
        style="display: flex;flex-direction:column;height:5rem;margin-bottom:1rem"
      >
        <Label>Currency</Label>
        <CryptoCard
          on:click={() => (selectorVisible = true)}
          crypto={$transactionStore.destinationCurrency}
        />
      </div>
      <div style="display:flex;flex-direction:column;height:5rem;">
        <div style="display:flex;justify-content:space-between;">
          <Label>Amount</Label>
          <Label
            >{$transactionStore.sourceCurrency.ticker} / {$transactionStore
              .destinationCurrency.ticker}</Label
          >
        </div>
        <Input forceLabel={true} type="number" placeholder="Amount" />
      </div>
      <div class="exchange-rate-container">
        ~ 1 {$transactionStore.destinationCurrency.ticker} @ 55,000 {$transactionStore
          .sourceCurrency.ticker}
      </div>
    </div>

    <PopupSelector
      on:close={() => (selectorVisible = false)}
      visible={selectorVisible}
      headerTitle="Select Cryptocurrency"
    >
      <div class="cryptocurrencies-container">
        {#each cryptoCurrencies as cryptoCurrency (cryptoCurrency.ticker)}
          <CryptoCard
            on:click={() => (selectorVisible = false)}
            crypto={cryptoCurrency}
          />
        {/each}
      </div>
    </PopupSelector>
  </ModalBody>
  <ModalFooter>
    <Button on:click={handleNextStep}>Checkout</Button>
  </ModalFooter>
</ModalContent>

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
    color: lighten($textColor3, 20%);
  }
</style>
