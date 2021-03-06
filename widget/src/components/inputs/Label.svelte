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
  label {
    cursor: pointer;
    position: relative;
    color: var(--theme-text-color-3);
    text-transform: uppercase;
    font-size: 0.7rem;
    font-weight: 600;
    margin-bottom: 1.25em;
    span {
      position: absolute;
      top: 12px;
      margin-left: 15px;
      z-index: 99;
    }
  }
  :global(label .input-container > input) {
    padding-top: 1.5em !important;
  }
</style>
