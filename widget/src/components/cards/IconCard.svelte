<script lang="ts">
  import Card from './Card.svelte'
  import FaIcon from 'svelte-awesome'
  import { faChevronRight } from '@fortawesome/free-solid-svg-icons'
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()

  export let icon
  export let label: string
  export let paddingSmall = false
  export let blend = false
</script>

<Card on:click={() => dispatch('click')}>
  <div class="icon-card-container" class:blend class:paddingSmall>
    <div class="content-container">
      <FaIcon data={icon} />
      <div class="label">{label}</div>
    </div>
    <FaIcon data={faChevronRight} />
  </div>
</Card>

<style lang="scss">
  @import '../../styles/_vars.scss';

  .icon-card-container {
    flex: 1;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.25rem 0;
    &.paddingSmall {
      padding: 0.5rem;
    }
    &.blend {
      background-color: var(--theme-modal-background);
    }
    & > .content-container {
      height: 100%;
      width: 100%;
      display: flex;
      justify-content: flex-start;
      align-items: center;
      & > .label {
        font-weight: 500;
        margin-left: 0.5rem;
      }
    }
  }
</style>
