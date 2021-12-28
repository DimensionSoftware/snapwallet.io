<script lang="ts">
  import { expoOut } from 'svelte/easing'

  import { blur } from 'svelte/transition'

  export let ticker: string = 'USD'
  export let isBuy: boolean = true
</script>

<div class="left" title={`Buy ${ticker}`}>
  {#if isBuy}
    Buy
    <small in:blur={{ opacity: 0.5, easing: expoOut, duration: 2000 }}
      >Sell</small
    >
  {:else}
    Sell
    <small in:blur={{ opacity: 0.5, easing: expoOut, duration: 2000 }}
      >Buy</small
    >
  {/if}
</div>

<style lang="scss">
  @import '../styles/_vars.scss';
  .left {
    text-align: left;
  }
  small {
    margin: 0 0 0 0.7rem;
    opacity: 0.75;
  }
</style>
