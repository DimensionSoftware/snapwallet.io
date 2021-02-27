<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
  export let disabled: boolean
</script>

<button {disabled} on:click={() => dispatch('click')}><slot /></button>

<style lang="scss">
  @import '../styles/_vars.scss';

  button {
    min-height: 50px;
    width: 100%;
    background-color: var(--theme-color);
    border-radius: 0.5rem;
    border: none;
    border-bottom: 2px solid var(--theme-color-darkened);
    color: white;
    cursor: pointer;
    text-transform: uppercase;
    box-shadow: 0 6px 6px var(--theme-button-shadow-color);
    letter-spacing: 2px;
    margin: 0;
    font-weight: bold;
    &:active,
    &:focus {
      background-color: var(--theme-color-lightened);
    }
    &:disabled {
      opacity: 0.9;
      cursor: initial;
    }
  }
</style>
