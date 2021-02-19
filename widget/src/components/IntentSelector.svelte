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

  .buy-toggle {
    height: 100%;
    width: 100%;
    display: flex;
    color: white;
    background-color: $greyBackground1;
    border-top-left-radius: 0.5rem;
    border-bottom-left-radius: 0.5rem;
    align-items: center;
    justify-content: center;
    opacity: 1;
    &.active {
      opacity: 0.4;
    }
  }

  .sell-toggle {
    height: 100%;
    width: 100%;
    display: flex;
    color: white;
    background-color: $greyBackground1;
    border-top-right-radius: 0.5rem;
    border-bottom-right-radius: 0.5rem;
    align-items: center;
    justify-content: center;
    opacity: 1;
    &.active {
      opacity: 0.4;
    }
  }
</style>
