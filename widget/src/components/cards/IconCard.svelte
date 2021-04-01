<script lang="ts">
  import Card from './Card.svelte'
  import FaIcon from 'svelte-awesome'
  import { faChevronRight } from '@fortawesome/free-solid-svg-icons'
  import { createEventDispatcher } from 'svelte'
  import Badge from '../Badge.svelte'
  const dispatch = createEventDispatcher()

  export let icon
  export let label: string
  export let paddingSmall = false
  export let blend = false
  export let badgeText = ''
  export let badgeType: 'error' | 'warning' | 'success' | undefined = undefined
</script>

<Card on:click={() => dispatch('click')}>
  <div class="icon-card-container" class:blend class:paddingSmall>
    <div class="content-container">
      <FaIcon data={icon} />
      <div class="label">{label}</div>
    </div>
    {#if badgeText}
      <Badge
        class="icon-card-badge"
        error={badgeType === 'error'}
        success={badgeType === 'success'}
        warning={badgeType === 'warning'}
      >
        {badgeText}
      </Badge>
    {/if}
    <div class="chevron-right">
      <FaIcon height="1rem" width="1rem" data={faChevronRight} />
    </div>
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
    /* &.disabled {
      color: var(--theme-text-color-4);
      cursor: auto !important;
      & > .content-container {
        & > .label {
          font-weight: 300 !important;
        }
      }
    } */
  }

  .chevron-right {
    display: flex;
    justify-content: flex-end;
    align-items: center;
  }

  :global(.icon-card-badge) {
    color: red;
    margin-right: 1rem;
  }
</style>
