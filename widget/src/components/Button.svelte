<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
  export let disabled: boolean = false
</script>

<button {disabled} on:click={() => dispatch('click')}><slot /></button>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';

  button {
    position: relative;
    min-height: 50px;
    width: 100%;
    background: var(--theme-color);
    border-radius: 0.5rem;
    border: none;
    border-top: 1px solid lighten($themeColor, 10%);
    border-bottom: 1px solid var(--theme-color-darkened);
    color: white;
    cursor: pointer;
    text-transform: uppercase;
    box-shadow: 0 0 0 0 $themeColor, 0 6px 6px var(--theme-button-shadow-color);
    letter-spacing: 2px;
    margin: 0;
    font-weight: bold;
    transition: background 0.3s ease-in 0.1s, box-shadow 0.2s ease-in 0.123s;
    // gloss fix
    &:before {
      position: absolute;
      content: '';
      height: 50%;
      top: -1px;
      right: -1px;
      bottom: -1px;
      left: -1px;
      opacity: 0.2;
      background: linear-gradient(
        to bottom,
        rgba(#fff, 0.2) 0%,
        rgba(#fff, 0.15) 100%
      );
      white-space: nowrap;
      border-radius: 0.5em 0.5em 6em 6em/0.1em 0.1em 1em 1em;
      transition: opacity 0.3s ease-in 0.1s;
    }
    &:hover {
      box-shadow: 0 0 0 1px $themeColor,
        0 6px 6px var(--theme-button-shadow-color);
      transition: none;
      &:before {
        opacity: 0.3;
      }
    }
    &:active,
    &:focus {
      background: var(--theme-color-lightened);
      box-shadow: 0 0 0 1px $themeColor,
        0 6px 6px var(--theme-button-shadow-color);
      text-shadow: 0 1px 0 --var(--theme-text-color);
      transition: none;
      animation: infocus 0.35s;
      animation-timing-function: $easeOutBack;
      &:before {
        opacity: 0.08;
        transition: none;
      }
    }
    &:disabled {
      background: var(--theme-color);
      cursor: not-allowed;
      text-shadow: none;
      opacity: 0.83;
      box-shadow: none;
    }
  }
</style>
