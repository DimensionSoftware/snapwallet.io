<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import { SUPPORTED_CRYPTOCURRENCY_ASSETS } from '../../constants'
  import { fly } from 'svelte/transition'
  import CryptoCard from '../cards/CryptoCard.svelte'
  import Label from '../inputs/Label.svelte'
  import PopupSelector from '../inputs/PopupSelector.svelte'
  const dispatch = createEventDispatcher()

  export let visible = false
</script>

{#if visible}
  <PopupSelector
    on:close={() => dispatch('close')}
    headerTitle="Select Currency"
  >
    <div class="scroll-y selector-container">
      <h5>Popular</h5>
      {#each SUPPORTED_CRYPTOCURRENCY_ASSETS.filter(c => c.popular) as cryptoCurrency, i (cryptoCurrency.ticker)}
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
      {#each SUPPORTED_CRYPTOCURRENCY_ASSETS.filter(c => !c.popular) as cryptoCurrency, i (cryptoCurrency.ticker)}
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
