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
    background-color: $themeColor;
    border-radius: 0.5rem;
    border: none;
    border-bottom: 2px solid darken($themeColor, 10%);
    color: white;
    cursor: pointer;
    text-transform: capitalize;
    margin: 0;
    font-weight: bold;
    &:active,
    &:focus {
      background-color: lighten($themeColor, 5%);
    }
    &:disabled {
      opacity: 0.9;
      cursor: initial;
    }
  }
</style>
