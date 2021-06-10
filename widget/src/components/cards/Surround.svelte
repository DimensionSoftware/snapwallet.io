<script lang="ts">
</script>

<div class="surround" on:click={() => {}}>
  <slot />
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';
  @import '../../styles/animations.scss';

  .surround {
    position: relative;
    padding: 0;
    margin-top: 1rem;
    width: 100%;
    display: flex;
    flex-direction: column;
    padding: 0 1rem 0.5rem;
    border: 2px solid var(--theme-button-color);
    border-radius: 0.75rem;
  }
</style>
