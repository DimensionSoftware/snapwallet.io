<script lang="ts">
  import { expoOut } from 'svelte/easing'

  import { blur } from 'svelte/transition'

  export let ticker: string = 'USD'
  export let isBuy: boolean = true
</script>

<div class="left" title={`Buy ${ticker}`}>
  {#if isBuy}
    <span>Buy</span>
    <small in:blur={{ opacity: 0.6, easing: expoOut, duration: 2000 }}
      >Sell</small
    >
  {:else}
    <small in:blur={{ opacity: 0.6, easing: expoOut, duration: 2000 }}
      >Buy</small
    >
    <span>Sell</span>
  {/if}
</div>

<style lang="scss">
  @import '../styles/_vars.scss';
  .left {
    text-align: left;
  }
  span,
  small {
    display: inline-block;
    text-align: center;
    width: 60px;
  }
  small {
    font-size: 1rem;
    margin: 0;
    opacity: 0.75;
  }
</style>
