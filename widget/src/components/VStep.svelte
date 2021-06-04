<script lang="ts">
  export let success: boolean = false
  export let disabled: boolean = false
  export let custom: boolean = false
  export let onClick: () => void
  export let title: string
</script>

<li
  on:click={onClick}
  {title}
  class:custom
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
  <slot name="info" />
</li>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';
  li {
    position: relative;
    z-index: 2;
    padding-left: 1.25rem;
    margin-left: 1rem;
    margin-top: 0.55rem;
    &.custom {
      :global(span):before {
        display: none;
      }
    }
    // icon surround
    :global(span.default-icon:before) {
      position: absolute;
      content: '';
      background: var(--theme-text-color);
      border-radius: 50%;
      height: 6px;
      width: 6px;
      left: 1px;
      top: 9px;
      opacity: 0.7;
      z-index: 1;
    }
    > :global(span > svg) {
      position: absolute;
      left: -4px;
      z-index: 4;
      top: 4px;
    }
    div {
      display: inline-block;
    }
    // marker dot
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
      opacity: 1;
      top: 8px;
      z-index: 1;
    }
    // line
    &:first-child:after {
      background: linear-gradient(transparent, var(--theme-text-color));
      top: -40px;
      height: calc(100% + 32px);
    }
    &:last-child:after {
      background: linear-gradient(var(--theme-text-color), transparent);
    }
    &:after {
      position: absolute;
      width: 2px;
      left: 3px;
      top: -20px;
      opacity: 0.3;
      height: calc(100% + 9px);
      content: '';
      background-color: var(--theme-text-color);
      background-position: 0 0;
      background-size: 200% 200%;
      border-color: inherit;
      border-width: 0;
      outline: 0;
      z-index: 0;
    }
    :global(span:before) {
      background: var(--theme-modal-background);
      position: absolute;
      content: '';
      border-radius: 50%;
      opacity: 0.25;
      height: 26px;
      width: 26px;
      left: -9px;
      top: -1px;
      z-index: 1;
    }
    :global(span.glow:before) {
      box-shadow: 0 0 0 0 rgba(var(--theme-button-glow-color), 0.5);
      animation: glow 1.5s linear;
      animation-iteration-count: infinite;
    }
    &.success {
      animation: scaleIn 0.25s ease-out;
      z-index: 1;
      :global(.total-container) {
        font-weight: bold;
      }
      // connecting line
      &:first-child:after {
        background: linear-gradient(transparent, var(--theme-success-color));
      }
      &:last-child:after {
        background: var(--theme-success-color);
        background: linear-gradient(var(--theme-success-color), transparent);
      }
      &:after {
        background-color: var(--theme-success-color);
        background: var(--theme-success-color);
        width: 2px;
        z-index: 0;
      }
      // hide the dot
      :global(span:before) {
        z-index: 3;
        opacity: 1;
        // background: var(--theme-success-color);
        background: var(--theme-modal-background);
      }
      & > :global(.icon) {
        margin-left: 0.4rem;
      }
      &:before {
        border: 4px solid var(--theme-success-color) !important;
      }

      & > :global(span > svg) {
        color: var(--theme-success-color);
      }
    }
    &.disabled {
      :global(b) {
        font-weight: normal;
      }
    }
    .step {
      white-space: nowrap;
      margin-left: 0.5rem;
      top: -1px;
      position: relative;
    }
  }
</style>
