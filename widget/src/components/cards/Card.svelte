<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
</script>

<div on:click={() => dispatch('click')} class="card-container">
  <slot />
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';

  .card-container {
    width: 100%;
    border-radius: 0.5rem;
    cursor: pointer;
    display: flex;
  }
</style>
