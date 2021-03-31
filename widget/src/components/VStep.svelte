<script lang="ts">
  export let success: boolean = false
  export let disabled: boolean = false
  export let onClick: () => {}
</script>

<li
  on:click={onClick}
  class:success
  class:disabled
  style={onClick ? 'cursor: pointer' : ''}
>
  <slot name="icon">
    <span class="default-icon" />
  </slot>
  <div class="step">
    <slot name="step" />
  </div>
</li>

<style lang="scss">
  @import '../styles/_vars.scss';
  li {
    position: relative;
    padding-left: 1.25rem;
    margin-left: 1rem;
    margin-top: 0.75rem;
    // icon surround
    :global(span.default-icon):before {
      position: absolute;
      content: '';
      background: var(--theme-text-color);
      border-radius: 50%;
      height: 8px;
      width: 8px;
      left: 0;
      top: 8px;
      z-index: 1;
    }
    > :global(span > svg) {
      position: absolute;
      left: -4px;
      z-index: 1;
      top: 3px;
    }
    div {
      display: inline-block;
    }
    // marker
    .default-icon:before {
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
    &:first-child:after {
      height: 500%;
      top: -3rem;
    }
    &:after {
      position: absolute;
      width: 2px;
      left: 4px;
      top: -1.25rem;
      opacity: 0.3;
      height: 120%;
      content: '';
      background-color: var(--theme-text-color-4);
      background-position: 0 0;
      background-size: 200% 200%;
      border-color: inherit;
      border-width: 0;
      outline: 0;
    }
    &.success {
      display: flex;
      align-items: center;
      :global(.total-container) {
        font-weight: bold;
      }
      &:after {
        background-color: var(--theme-success-color);
        background: linear-gradient(
          to bottom,
          var(--theme-text-color-4),
          var(--theme-success-color)
        );
        width: 2px;
      }
      :global(span:before) {
        display: none;
        border-color: var(--theme-success-color) !important;
        background: var(--theme-success-color);
      }
      & > :global(.icon) {
        margin-left: 0.4rem;
      }
      &:before {
        border: 4px solid var(--theme-success-color) !important;
      }
    }
    &.disabled {
      cursor: auto !important;
      :global(b) {
        font-weight: normal;
      }
    }
    .step {
      margin-left: 0.5rem;
      top: -1px;
      position: relative;
    }
  }
</style>
