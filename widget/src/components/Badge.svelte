<script lang="ts">
  export let error: boolean = false
  export let success: boolean = false
  export let warning: boolean = false
</script>

<div
  class={`badge ${$$props.class || ''}`.trim()}
  class:success
  class:warning
  class:error
>
  <slot />
</div>

<style lang="scss">
  .badge {
    border-radius: 0.5rem;
    padding: 0 0.5rem;
    font-size: 0.75rem;
    float: left;
    &.success {
      color: var(--theme-text-color);
      /* border: 1px solid var(--theme-success-color); */
      background-color: var(--theme-success-color);
    }

    &.warning {
      color: var(--theme-text-color);
      /* border: 1px solid var(--theme-warning-color); */
      background-color: var(--theme-warning-color);
    }

    &.error {
      color: var(--theme-text-color);
      /* border: 1px solid var(--theme-error-color); */
      background-color: var(--theme-error-color);
    }
  }
</style>
