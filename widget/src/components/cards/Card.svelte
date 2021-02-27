<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
</script>

<div class="card-container">
  <slot />
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';

  .card-container {
    border: 1px solid var(--theme-text-color-4);
    min-height: 3rem;
    width: 100%;
    border-radius: 0.5rem;
    padding: 0.2rem;
  }
</style>
