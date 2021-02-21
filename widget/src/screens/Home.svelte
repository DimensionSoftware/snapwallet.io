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
    <CryptoCard
      on:click={() => (selectorVisible = true)}
      crypto={$transactionStore.destinationCurrency}
    />
    <PopupSelector
      on:close={() => (selectorVisible = false)}
      visible={selectorVisible}
      headerTitle="Select Cryptocurrency"
    >
      {#each cryptoCurrencies as cryptoCurrency (cryptoCurrency.ticker)}
        <CryptoCard
          on:click={() => (selectorVisible = false)}
          crypto={cryptoCurrency}
        />
      {/each}
    </PopupSelector>
  </ModalBody>
  <ModalFooter>
    <Button on:click={handleNextStep}>
      {$userStore.intent}
    </Button>
  </ModalFooter>
</ModalContent>
