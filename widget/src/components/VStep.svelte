<script lang="ts">
  export let success: boolean = false
  export let active: boolean = false
  export let disabled: boolean = false
  export let custom: boolean = false
  export let onClick: () => void = () => {}
  export let title: string = ''
  export let line: boolean = false
</script>

<li
  on:click={onClick}
  {title}
  class:custom
  class:success
  class:active
  class:disabled
  class:line
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
    margin-top: 0.255rem;
    line-height: 1.5rem;
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
      height: 3px;
      width: 3px;
      left: 2px;
      top: 10px;
      opacity: 0.7;
      z-index: 1;
    }

    // line
    &.line:first-child:after {
      background: linear-gradient(transparent, var(--theme-text-color));
      top: -32px;
      height: calc(100% + 37px);
    }
    &.line:last-child:after {
      background: linear-gradient(var(--theme-text-color), transparent);
    }
    &.line:after {
      position: absolute;
      width: 1px;
      left: 3px;
      top: 0;
      opacity: 0.2;
      height: calc(100% + 4px);
      content: '';
      background-color: var(--theme-text-color);
      background-position: 0 0;
      background-size: 200% 200%;
      border-color: inherit;
      border-width: 0;
      outline: 0;
      z-index: -1;
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
      border-radius: 100%;
      position: absolute;
      height: 3px;
      width: 3px;
      left: 2px;
      right: 0;
      bottom: 0;
      opacity: 0.7;
      top: 10px;
      z-index: 1;
    }
    :global(span:before) {
      background: transparent;
      position: absolute;
      content: '';
      border-radius: 50%;
      opacity: 0.75;
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
      :global(span.default-icon:before) {
        background: var(--theme-color);
      }
      :global(.step),
      :global(.total-container),
      :global(.step > span),
      :global(b) {
        opacity: 1 !important;
        font-weight: bold;
        color: var(--theme-color);
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
        margin-left: 0.1rem;
      }
      &:before {
        border: 2px solid var(--theme-color) !important;
      }

      & > :global(span > svg) {
        color: var(--theme-color);
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
      :global(> span),
      :global(b) {
        opacity: 0.75;
      }
    }
  }
</style>
