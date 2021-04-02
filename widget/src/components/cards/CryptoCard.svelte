<script lang="ts">
  import { transactionStore } from '../../stores/TransactionStore'
  import { createEventDispatcher } from 'svelte'
  import FaIcon from 'svelte-awesome'
  import { faChevronRight } from '@fortawesome/free-solid-svg-icons'
  import { TransactionIntents } from '../../types'
  import { CryptoIcons } from '../../util'

  export let crypto

  const dispatch = createEventDispatcher()
</script>

<div
  on:click={() => {
    const { destinationCurrency, sourceCurrency } = $transactionStore
    if ($transactionStore.intent === TransactionIntents.BUY) {
      transactionStore.setCurrencies({
        destinationCurrency: crypto,
        sourceCurrency,
      })
    } else {
      transactionStore.setCurrencies({
        destinationCurrency,
        sourceCurrency: crypto,
      })
    }
    dispatch('click')
  }}
  class="crypto-card"
>
  <div class="crypto-icon">
    <svelte:component this={CryptoIcons[crypto.ticker.toUpperCase()]} />
  </div>
  <div class="crypto-name">
    <span>{crypto.ticker}</span> <small>{crypto.name}</small>
  </div>
  <div class="crypto-arrow"><FaIcon data={faChevronRight} /></div>
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';
  @import '../../styles/animations.scss';

  .crypto-card {
    position: relative;
    padding: 0;
    width: 100%;
    height: 3rem;
    display: flex;
    align-items: center;
    cursor: pointer;
    &:before {
      // background fx
      content: '';
      position: absolute;
      z-index: 0;
      top: 0;
      right: -2px;
      left: -2px;
      bottom: 0;
      border-radius: 0.5rem;
      background: linear-gradient(
        to right,
        rgba($themeColor, 0.02),
        rgba($themeColor, 0.3),
        rgba($themeColor, 0.02)
      );
      transform: scale(0);
      transition: transform 0.2s $easeOutExpo;
    }
    &:hover {
      &:before {
        transform: scale(1);
        transition: none;
      }
      .crypto-icon {
        transform: scale(1.05);
        transition: none;
        &:before {
          background: $themeColor;
          animation: currency 0.3s $easeOutBack, background 0s ease-out 0.3s;
          top: -1px;
          right: -1px;
          left: -1px;
          bottom: 0;
        }
      }
    }
  }

  .crypto-icon {
    position: relative;
    height: 33px;
    left: -4px;
    transition: 0.15s $easeOutBack 0.05s;
    :global(svg) {
      position: relative;
      z-index: 2;
    }
    &:before {
      content: '';
      position: absolute;
      z-index: 0;
      top: 0;
      right: 0;
      left: 0;
      bottom: 0;
      border-radius: 100%;
      background: linear-gradient($themeColor, lighten($themeColor, 25%));
    }
  }

  .crypto-name {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: row;
    height: 100%;
    margin-left: .4rem;
    font-size: 1.5em;
    color: var(--theme-text-color);
    span {
      font-size: 1.25rem;
    }
    small {
      font-weight: 500;
      margin-left: 0.75rem;
      text-transform: lowercase;
    }
  }
  .crypto-arrow {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    margin-left: auto;
    color: var(--theme-text-color);
  }
</style>
