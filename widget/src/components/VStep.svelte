<script lang="ts">
  export let success: boolean = false
  export let active: boolean = false
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
  class:active
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
      height: 4px;
      width: 4px;
      left: 2px;
      top: 10px;
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
    :global(span:before) {
      background: transparent;
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
    &.active {
      > :global(span:before),
      > :global(.default-icon:before) {
        box-shadow: 0 0 0 0 rgba(var(--theme-button-glow-color), 0.5);
        animation: glow 1.5s linear;
        animation-iteration-count: infinite;
      }
      :global(b) {
        font-weight: bold;
      }
    }
    &.success {
      animation: scaleIn 0.25s ease-out;
      z-index: 1;
      // hide the dot
      :global(span:before) {
        z-index: 3;
        opacity: 1;
        background: transparent;
      }
      :global(&.active span.default-icon:before) {
        box-shadow: 0 0 0 0 rgba(var(--theme-button-glow-color), 0.5);
        animation: glow 1.5s linear;
        animation-iteration-count: infinite;
        margin-left: 0.4rem;
      }
      &:before {
        border: 4px solid var(--theme-success-color) !important;
      }

      & > :global(span > svg) {
        color: var(--theme-success-color);
      }
    }
    &.disabled,
    &.success {
      :global(b) {
        font-weight: normal;
      }
    }
    .step {
      white-space: nowrap;
      margin-left: 1rem;
      top: -1px;
      position: relative;
    }
  }
</style>
