<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { TransactionIntents } from '../types'
  const dispatch = createEventDispatcher()

  $: isSell = $transactionStore.intent === TransactionIntents.SELL
  const onClick = () => {
    transactionStore.toggleIntent()
    dispatch('change')
  }
</script>

<div class="intent-selector-container">
  <div class="intent-selector">
    <div
      on:click={() => isSell && onClick()}
      class:active={!isSell}
      class="buy-toggle"
    >
      Buy
    </div>
    <div
      on:click={() => !isSell && onClick()}
      class:active={isSell}
      class="sell-toggle"
    >
      Sell
    </div>
  </div>
</div>

<style lang="scss">
  @import '../styles/_vars.scss';
  .intent-selector-container {
    display: flex;
    width: 100%;
    height: 55px;
    align-items: center;
    justify-content: center;
  }

  .intent-selector {
    width: 50%;
    height: 100%;
    display: flex;
    cursor: pointer;
    border-radius: 1.5rem;
    padding: 5px;
    background: white;
    box-shadow: 0 0 2px -1px var(--theme-text-color);
  }

  .sell-toggle,
  .buy-toggle {
    position: relative;
    height: 100%;
    width: 100%;
    display: flex;
    color: var(--theme-text-color);
    overflow: hidden;
    align-items: center;
    justify-content: center;
    background-color: rgba($themeColor, 0.1);
    border-bottom: 1px solid rgba($themeColor, 0.1);
    font-weight: 500;
    border-top-left-radius: 1.5rem;
    border-bottom-left-radius: 1.5rem;
    opacity: 1;
    transition: border-bottom 0.2s ease-out 0.1s;
    &:hover {
      border-bottom: 1px solid rgba($themeColor, 0.5);
      transition: none;
    }
    &.active {
      cursor: default;
      color: white;
      opacity: 0.9;
      background-color: rgba($themeColor, 0.8);
      border-bottom: 1px solid darken($themeColor, 30%);
      &:before {
        // background fx
        transform: scale(1) translateY(1px);
        transition: transform 0.3s var(--theme-ease-out-expo);
      }
    }
    &.sell-toggle {
      border-radius: 0;
      border-top-right-radius: 1.5rem;
      border-bottom-right-radius: 1.5rem;
      &:before {
        // background fx
        border-radius: 0;
        border-top-right-radius: 1.5rem;
        border-bottom-right-radius: 1.5rem;
      }
    }
    &:before {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      z-index: -1;
      background-color: var(--theme-color);
      transform: scale(0) translateY(90%);
      transition: transform 0.1s var(--theme-ease-in-expo);
      border-top-left-radius: 1.5rem;
      border-bottom-left-radius: 1.5rem;
    }
  }
</style>
