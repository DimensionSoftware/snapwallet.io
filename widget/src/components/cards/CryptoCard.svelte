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
  import FaIcon from 'svelte-awesome'
  import { faChevronRight } from '@fortawesome/free-solid-svg-icons'

  export let crypto

  const dispatch = createEventDispatcher()
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
  <div class="crypto-arrow"><FaIcon data={faChevronRight} /></div>
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';

  .crypto-card {
    margin: 0.5rem 0;
    padding: 0;
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
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    margin-left: 1rem;
  }

  .crypto-arrow {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    margin-left: auto;
    color: $textColor1;
  }
</style>
