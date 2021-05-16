<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
  export let id
  export let disabled: boolean = false
  export let isLoading: boolean = false
  export let title: string = ''
  export let glow: boolean = false
</script>

<button
  {id}
  disabled={disabled || isLoading}
  class:glow
  class:isLoading
  on:mousedown={() => dispatch('mousedown')}
  {title}
>
  <div class="lds-circle"><div /></div>
  <slot /></button
>

<style lang="scss">
  @import '../../../widget/src/styles/animations.scss';
  button {
    position: relative;
    overflow: hidden;
    display: block;
    background: #fffc00;
    padding: 0.75rem 1.75rem;
    color: #000;
    border: 1px solid rgba(0, 0, 0, 0.2);
    border-radius: 2rem;
    cursor: pointer;
    white-space: nowrap;
    letter-spacing: 2px;
    margin: 0;
    font-weight: bolder;
    font-size: 1.25rem;
    letter-spacing: 0;
    transition: border 0.3s ease-in 0.05s;
    z-index: 10;
    &.glow {
      overflow: hidden;
      &:hover {
        animation: inherit;
      }
      box-shadow: 0 0 0 0 rgba(var(--theme-button-glow-color), 0.5);
      animation: glow 1.5s linear;
      animation-iteration-count: infinite;
    }
    &:hover {
      position: relative;
      background: var(--theme-button-color);
      border-color: var(--theme-button-color);
      transform: scale(1.0025);
      box-shadow: 0 2px 2px 1px rgba(0, 0, 0, 0.1);
      transition: none;
    }
    &:active,
    &:focus {
      &.glow {
        animation: infocus 0.35s !important;
        animation-timing-function: var(--theme-ease-out-back);
      }
      transition: none;
      border-color: var(--theme-button-color);
      box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.1);
      background: var(--theme-button-color);
      animation: infocus 0.35s;
      animation-timing-function: var(--theme-ease-out-back);
    }
    &:disabled {
      animation: infocus 0.75s;
      background: var(--theme-button-color);
      cursor: not-allowed;
      text-shadow: none;
      opacity: 0.83;
    }
  }
</style>
