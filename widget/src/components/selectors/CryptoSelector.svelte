<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import { fly } from 'svelte/transition'
  import CryptoCard from '../cards/CryptoCard.svelte'
  import Label from '../inputs/Label.svelte'
  import PopupSelector from '../inputs/PopupSelector.svelte'
  const dispatch = createEventDispatcher()

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
          <Label>
            <CryptoCard
              on:click={() => dispatch('close')}
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
              on:click={() => dispatch('close')}
              crypto={cryptoCurrency}
            />
          </Label>
        </div>
      {/each}
    </div>
  </PopupSelector>
{/if}

<style lang="scss">
  @import '../../styles/selectors.scss';
</style>
