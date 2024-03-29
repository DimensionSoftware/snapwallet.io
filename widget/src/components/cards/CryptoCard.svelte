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
    const { destinationCurrency, sourceCurrency, sourceAmount } =
      $transactionStore
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
    dispatch('select', {
      destinationCurrency,
      sourceCurrency: crypto,
    })
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
    width: 99%;
    height: 3.5rem;
    display: flex;
    align-items: center;
    transition: transform 0.4s var(--theme-ease-out-expo);
    transform: scale(1);
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
      opacity: 0.3;
      background: linear-gradient(
        to right,
        transparent,
        var(--theme-color-lightened),
        var(--theme-color-lightened),
        var(--theme-color-lightened),
        var(--theme-color-lightened),
        var(--theme-color-lightened),
        transparent 99%
      );
      transform: scale(0);
      transition: transform 0.2s $easeOutExpo;
    }
    &:hover {
      border-top: 1px solid var(--theme-color-lightened);
      border-bottom: 1px solid var(--theme-color-lightened);
      transition: none;
      transform: scale(1.01);
      .crypto-icon {
        filter: grayscale(0);
      }
      &:before {
        transform: scale(1);
        transition: none;
      }
      .crypto-icon {
        transform: scale(1.05);
        transition: none;
      }
    }
  }

  .crypto-icon {
    position: relative;
    height: 33px;
    width: 33px;
    left: -2px;
    filter: grayscale(100%);
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
      line-height: 1.7rem;
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
