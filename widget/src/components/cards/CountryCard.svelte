<script lang="ts">
  import Card from './Card.svelte'
  import FaIcon from 'svelte-awesome'
  import { faChevronRight } from '@fortawesome/free-solid-svg-icons'
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()

  export let paddingSmall = false
  export let blend = false
</script>

<Card on:click={() => dispatch('click')}>
  <div class="country-card-container" class:blend class:paddingSmall>
    <div class="content-container">
      <slot />
    </div>
    <div class="chevron-right">
      <FaIcon data={faChevronRight} />
    </div>
  </div>
</Card>

<style lang="scss">
  @import '../../styles/_vars.scss';

  .country-card-container {
    flex: 1;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.75rem 0;
    &.paddingSmall {
      padding: 0.5rem;
    }
    & > .content-container {
      height: 100%;
      width: 100%;
      display: flex;
      justify-content: flex-start;
      align-items: center;
    }
  }

  .chevron-right {
    display: flex;
    justify-content: flex-end;
    align-items: center;
  }
</style>
