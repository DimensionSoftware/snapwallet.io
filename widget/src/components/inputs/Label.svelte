<script lang="ts">
  export let hidden: boolean
</script>

<div class="input-label">
  {#if !hidden}
    <slot />
  {/if}
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';

  .input-label {
    color: var(--theme-color);
    font-size: 0.7rem;
    font-weight: 600;
  }
</style>
