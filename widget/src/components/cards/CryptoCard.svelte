<script lang="ts">
  import { transactionStore } from '../../stores/TransactionStore'
  import { createEventDispatcher } from 'svelte'
  import FaIcon from 'svelte-awesome'
  import {
    faChevronDown,
    faChevronRight,
  } from '@fortawesome/free-solid-svg-icons'
  import { TransactionIntents } from '../../types'
  import { CryptoIcons } from '../../util'
  import { ParentMessenger } from '../../util/parent_messenger'

  export let crypto
  export let isDown

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
    ParentMessenger.currencySelected(crypto)
    dispatch('mousedown')
  }}
  class="crypto-card"
>
  <div class="crypto-icon">
    <svelte:component this={CryptoIcons[crypto.ticker.toUpperCase()]} />
  </div>
  <div class="crypto-name">
    <span>{crypto.ticker}</span>
    <small style="text-transform:capitalize;">{crypto.name}</small>
  </div>
  <div class="crypto-arrow">
    <FaIcon data={isDown ? faChevronDown : faChevronRight} />
  </div>
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
      opacity: 0.5;
      background: linear-gradient(
        to right,
        transparent,
        var(--theme-color-lightened),
        transparent
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
          background: var(--theme-color);
          animation: currency 0.3s var(--theme-ease-out-back),
            background 0s ease-out 0.3s;
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
    transition: 0.15s var(--theme-ease-out-back) 0.05s;
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
      background: linear-gradient(
        var(--theme-color),
        var(--theme-color-lightened)
      );
    }
  }

  .crypto-name {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: row;
    height: 100%;
    margin-left: 0.4rem;
    font-size: 2em;
    color: var(--theme-text-color);
    span {
      font-size: 1.25rem;
    }
    small {
      font-weight: 300;
      margin-left: 0.25rem;
      text-transform: capitalize;
      white-space: nowrap;
      text-overflow: ellipsis;
      overflow: hidden;
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
