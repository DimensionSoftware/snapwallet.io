<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
</script>

<button on:click={() => dispatch('click')}><slot /></button>

<style lang="scss">
  @import '../styles/_vars.scss';

  button {
    min-height: 50px;
    width: 100%;
    background-color: $greyBackground1;
    border-radius: 0.5rem;
    color: white;
    cursor: pointer;
    text-transform: capitalize;
    margin: 0;
  }
</style>
