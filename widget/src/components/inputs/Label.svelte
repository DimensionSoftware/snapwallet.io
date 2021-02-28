<script lang="ts">
  export let hidden: boolean = false
  export let label: string = ''
</script>

{#if !hidden}
  <label>
    <span class="input-label">{label}</span>
    <slot />
  </label>
{/if}

<style lang="scss">
  @import '../../styles/_vars.scss';
  .input-label {
    cursor: pointer;
    color: var(--theme-text-color-3);
    text-transform: uppercase;
    font-size: 0.7rem;
    font-weight: 600;
  }
</style>
