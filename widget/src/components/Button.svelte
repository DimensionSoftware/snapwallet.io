<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
  export let id
  export let disabled: boolean = false
  export let isLoading: boolean = false
  export let title: string = ''
  export let glow: boolean = false

  $: if (isLoading) glow = true
</script>

<button
  {id}
  disabled={disabled || isLoading}
  class:glow
  class:isLoading
  on:click={() => dispatch('mousedown')}
  {title}
>
  <div class="lds-circle"><div /></div>
  <slot /></button
>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';

  button {
    position: relative;
    min-height: 50px;
    width: 100%;
    background: var(--theme-button-color);
    border-radius: 0.5rem;
    border: none;
    border-top: none;
    color: var(--theme-button-text-color);
    cursor: pointer;
    text-transform: uppercase;
    box-shadow: 0 0 0 0 var(--theme-color), 0 6px 6px var(--theme-shadow-color);
    letter-spacing: 2px;
    margin: 0;
    font-weight: bold;
    transition: none;
    z-index: 10;
    overflow: hidden;
    // gloss fix
    &:before {
      position: absolute;
      content: '';
      height: 50%;
      top: -1px;
      right: -1px;
      bottom: -1px;
      left: -1px;
      opacity: 0.3;
      background: linear-gradient(
        to bottom,
        rgba(#fff, 0.3) 0%,
        rgba(#fff, 0.1) 100%
      );
      white-space: nowrap;
      border-radius: 0.5em 0.5em 6em 6em/0.1em 0.1em 1em 1em;
      border-top-left-radius: 0.5rem;
      border-top-right-radius: 0.5rem;
      transform: scale(1);
      transition: transform 0.1s ease-out, opacity 0.1s ease-in;
    }
    &:hover {
      box-shadow: 0 0 0 1px var(--theme-button-color),
        0 6px 6px var(--theme-shadow-color);
      transition: none;
      &:before {
        transform: scale(1.1);
        transition: none;
        opacity: 0.4;
      }
    }
    &:active,
    &:focus {
      &.glow {
        animation: infocus 0.35s !important;
        animation-timing-function: var(--theme-ease-out-back);
      }
      box-shadow: 0 0 0 1px var(--theme-button-color),
        0 4px 4px var(--theme-shadow-color);
      text-shadow: 0 1px 0 --var(--theme-button-text-color);
      transition: none;
      animation: infocus 0.35s;
      animation-timing-function: var(--theme-ease-out-back);
      &:before {
        transform: scale(1.1);
        opacity: 0.3;
        transition: none;
      }
    }
    &:disabled {
      animation: infocus 0.75s;
      background: var(--theme-button-color);
      cursor: not-allowed;
      text-shadow: none;
      opacity: 0.83;
      box-shadow: none;
      &:before {
        transform: scale(1);
        opacity: 0;
      }
    }
    &.glow {
      &:hover {
        animation: inherit;
      }
      box-shadow: 0 0 0 0 rgba(var(--theme-button-glow-color), 0.55);
      animation: glow 1.5s linear;
      animation-iteration-count: infinite;
    }
    &.isLoading {
      .lds-circle {
        visibility: visible;
      }
    }
    .lds-circle {
      position: absolute;
      left: 0.5rem;
      top: 0;
      display: inline-block;
      transform: translateZ(1px);
      visibility: hidden;
      transition: opacity 0.2s ease-out;
      > div {
        display: inline-block;
        width: 32px;
        height: 32px;
        margin: 8px;
        border-radius: 50%;
        background: var(--theme-button-text-color);
        animation: lds-circle 2.4s cubic-bezier(0, 0.2, 0.8, 1) infinite;
      }
    }
    @keyframes lds-circle {
      0%,
      100% {
        animation-timing-function: cubic-bezier(0.5, 0, 1, 0.5);
      }
      0% {
        transform: rotateY(0deg);
      }
      50% {
        transform: rotateY(1800deg);
        animation-timing-function: cubic-bezier(0, 0.5, 0.5, 1);
      }
      100% {
        transform: rotateY(3600deg);
      }
    }
  }
</style>
