<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import { fly } from 'svelte/transition'
  import CryptoCard from '../cards/CryptoCard.svelte'
  import Label from '../inputs/Label.svelte'
  import PopupSelector from '../inputs/PopupSelector.svelte'
  const dispatch = createEventDispatcher()

  const cryptoCurrencies = [
    { name: 'Aave', ticker: 'AAVE', color: '#EFB914' },
    { name: 'Basic Attention Token', ticker: 'BAT', color: '#FF5000' },
    { name: 'Binance USD', ticker: 'BUSD', color: '#EFB914' },
    { name: 'Bitcoin', ticker: 'BTC', popular: true, color: '#F7931A' },
    { name: 'Curve', ticker: 'CRV', color: '#EFB914' },
    { name: 'Compound', ticker: 'COMP', color: '#00D395' },
    { name: 'DAI', ticker: 'DAI', color: '#F4B731' },
    { name: 'Ethereum', ticker: 'ETH', popular: true, color: '#627EEA' },
    { name: 'Gemini Dollar', ticker: 'GUSD', color: '#00DCFA' },
    { name: 'Link', ticker: 'LINK', color: '#2A5ADA' },
    { name: 'MakerDAO', ticker: 'MKR', color: '#1AAB9B' },
    { name: 'Paxos Standard', ticker: 'PAX', color: '#EDE708' },
    { name: 'Stably Dollar', ticker: 'USDS', color: '#EFB914' },
    { name: 'Synthetix', ticker: 'SNX', color: '#EFB914' },
    { name: 'Tether', ticker: 'USDT', popular: true, color: '#26A17B' },
    { name: 'UMA', ticker: 'UMA', color: '#FF4A4A' },
    { name: 'USDC', ticker: 'USDC', color: '#2775C9' },
    { name: 'Uniswap', ticker: 'UNI', popular: true, color: '#FF007A' },
    { name: 'Wrapped Bitcoin', ticker: 'WBTC', color: '#323544' },
    { name: 'Yearn.Finance', ticker: 'YFI', color: '#006AE3' },
  ]

  export let visible = false
</script>

{#if visible}
  <PopupSelector
    on:close={() => dispatch('close')}
    headerTitle="Select Currency"
  >
    <div class="scroll selector-container">
      <h5>Popular</h5>
      {#each cryptoCurrencies.filter(c => c.popular) as cryptoCurrency, i (cryptoCurrency.ticker)}
        <div
          in:fly={{ y: 25, duration: 250 + 50 * (i + 1) }}
          style="margin: 0.5rem 0"
        >
          <Label fx={false}>
            <CryptoCard
              on:mousedown={() => dispatch('close')}
              crypto={cryptoCurrency}
            />
          </Label>
        </div>
      {/each}
      <h5 style="margin-top: 1.25rem">All</h5>
      {#each cryptoCurrencies.filter(c => !c.popular) as cryptoCurrency, i (cryptoCurrency.ticker)}
        <div
          in:fly={{ y: 25, duration: 250 + 50 * (i + 4) }}
          style="margin: 0.5rem 0"
        >
          <Label>
            <CryptoCard
              on:mousedown={() => dispatch('close')}
              crypto={cryptoCurrency}
            />
          </Label>
        </div>
      {/each}
      <div class="spacer" />
    </div>
  </PopupSelector>
{/if}

<style lang="scss">
  @import '../../styles/selectors.scss';
</style>
