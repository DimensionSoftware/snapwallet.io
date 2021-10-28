<script lang="ts">
  export let glow: boolean = false
</script>

<div class="surround" class:glow on:click={() => {}}>
  <slot />
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';
  @import '../../styles/animations.scss';

  .surround {
    position: relative;
    overflow: hidden;
    padding: 0;
    margin-top: 1rem;
    width: 100%;
    display: flex;
    flex-direction: column;
    padding: 0 0.5rem 0.5rem 0.5rem;
    border: 3px solid var(--theme-button-color);
    border-radius: 1rem;
    &.glow {
      box-shadow: 0 0 0 0 rgba(var(--theme-button-glow-color), 0.35);
      animation: glowBasic 1.5s linear;
      animation-iteration-count: infinite;
    }
  }
</style>
