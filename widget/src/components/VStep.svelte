<script lang="ts">
  export let success: boolean = false
  export let disabled: boolean = false
  export let onClick: () => {}
</script>

<li on:click={onClick} class:success class:disabled style={onClick ? 'cursor: pointer' : ''}>
  <slot name="icon">
    <span class="icon" />
  </slot>
  <slot name="step" />
</li>

<style lang="scss">
  @import '../styles/_vars.scss';
  li {
    position: relative;
    padding-left: 1.25rem;
    margin-left: 1rem;
    margin-top: 0.5rem;
    // marker
    &:before {
      content: '';
      border: 4px solid $textColor4;
      border-radius: 100%;
      position: absolute;
      height: 0;
      width: 0;
      left: 0;
      right: 0;
      bottom: 0;
      top: 8px;
      z-index: 1;
    }
    // line
    &:after {
      position: absolute;
      width: 1px;
      left: 4px;
      top: -1.25rem;
      opacity: 0.3;
      height: 110%;
      content: '';
      background-color: $textColor4;
      background-position: 0 0;
      background-size: 200% 200%;
      border-color: inherit;
      border-width: 0;
      outline: 0;
    }
    &.success {
      display: flex;
      align-items: center;
      color: var(--theme-text-color-4);
      & > .icon {
        margin-left: 0.4rem;
        color: var(--theme-success-color);
      }
      &:before {
        border: 4px solid var(--theme-success-color) !important;
      }
    }
    &.disabled {
      cursor: inherit;
      color: var(--theme-text-color-4);
      cursor: auto !important;
    }
    b {
      cursor: pointer;
      margin-left: 0.5rem;
      text-decoration: underline;
    }
  }
</style>
