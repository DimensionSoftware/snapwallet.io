<script lang="ts">
  import { userStore } from '../stores/UserStore'
  $: isSell = $userStore.intent === 'sell'
</script>

<div class="intent-selector-container">
  <div class="intent-selector">
    <div
      on:click={() => userStore.setIntent('buy')}
      class:active={!isSell}
      class="buy-toggle"
    >
      Buy
    </div>
    <div
      on:click={() => userStore.setIntent('sell')}
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
    height: 40px;
    align-items: center;
    justify-content: center;
  }

  .intent-selector {
    width: 50%;
    height: 100%;
    display: flex;
    cursor: pointer;
  }

  .sell-toggle, .buy-toggle {
    position: relative;
    height: 100%;
    width: 100%;
    display: flex;
    color: $textColor;
    overflow: hidden;
    align-items: center;
    justify-content: center;
    background-color: rgba($themeColor, .1);
    border-bottom: 1px solid rgba($themeColor, .1);
    font-weight: 500;
    border-top-left-radius: 0.5rem;
    border-bottom-left-radius: 0.5rem;
    opacity: 1;
    transition: border-bottom .2s ease-out .1s;
    &:hover {
      border-bottom: 1px solid rgba($themeColor, .5);
      transition: none;
    }
    &.active {
      cursor: default;
      color: white;
      opacity: 0.9;
      background-color: rgba($themeColor, .8);
      border-bottom: 1px solid darken($themeColor, 30%);
      &:before {
        // background fx
        transform: scale(1) translateY(1px);
        transition: transform .3s $easeOutExpo;
      }
    }
    &.sell-toggle {
        border-radius: 0;
        border-top-right-radius: 0.5rem;
        border-bottom-right-radius: 0.5rem;
      &:before {
        // background fx
        border-radius: 0;
        border-top-right-radius: 0.5rem;
        border-bottom-right-radius: 0.5rem;
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
      background-color: $themeColor;
      transform: scale(0) translateY(90%);
      transition: transform .1s $easeInExpo;
      border-top-left-radius: 0.5rem;
      border-bottom-left-radius: 0.5rem;
    }
  }
</style>
