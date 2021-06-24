<script lang="ts">
  export let error: boolean = false
  export let success: boolean = false
  export let warning: boolean = false
  export let info: boolean = false
</script>

<div
  class={`badge ${$$props.class || ''}`.trim()}
  class:success
  class:warning
  class:error
  class:info
>
  <slot />
</div>

<style lang="scss">
  .badge {
    position: relative;
    border-radius: 0.5rem;
    padding: 0 0.5rem;
    font-size: 0.75rem;
    float: left;
    color: var(--theme-badge-text-color);
    overflow: hidden;
    &:before {
      position: absolute;
      content: '';
      top: 0;
      bottom: 0;
      left: 0;
      right: 0;
      opacity: 0.15;
    }
    &.success {
      color: var(--theme-success-color);
      &:before {
        background-color: var(--theme-success-color);
      }
    }

    &.warning {
      color: var(--theme-warning-color);
      &:before {
        background-color: var(--theme-warning-color);
      }
    }

    &.error {
      color: var(--theme-error-color);
      &:before {
        background-color: var(--theme-error-color);
      }
    }

    &.info {
      color: var(--theme-info-color);
      &:before {
        background-color: var(--theme-info-color);
      }
    }
  }
</style>
