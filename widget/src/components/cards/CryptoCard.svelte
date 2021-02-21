<script lang="ts">
  // HACK: this lib. does not offer a good
  // way to import icons dynamically.
  import * as ICONS from 'svelte-cryptoicon'
  let tickerIcons = {}
  Object.entries(ICONS).forEach(([k, v]) => {
    tickerIcons[k.toUpperCase()] = v
  })

  import { transactionStore } from '../../stores/TransactionStore'
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
  export let crypto
</script>

<div
  on:click={() => {
    transactionStore.setDestinationCurrency(crypto)
    dispatch('click')
  }}
  class="crypto-card"
>
  <svelte:component this={tickerIcons[crypto.ticker]} />
  <div class="crypto-name">{crypto.name}</div>
</div>

<style lang="scss">
  .crypto-card {
    margin: 0.5rem 0;
    padding: 0 1rem;
    width: 100%;
    height: 3rem;
    display: flex;
    align-items: center;
    cursor: pointer;
    transition: all 0.1s ease-in-out;
    &:hover {
      transform: scale(1.01);
    }
  }

  .crypto-name {
    margin-left: 1rem;
  }
</style>
