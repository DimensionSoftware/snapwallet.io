<script lang="ts">
  // HACK: this lib. does not offer a good
  // way to import icons dynamically.
  import * as ICONS from 'svelte-cryptoicon'
  let tickerIcons = {}
  Object.entries(ICONS).forEach(([k, v]) => {
    tickerIcons[k.toUpperCase()] = v
  })

  export let name
  export let ticker
</script>

<div class="crypto-card">
  <svelte:component this={tickerIcons[ticker]} />
  <div class="crypto-name">{name}</div>
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
