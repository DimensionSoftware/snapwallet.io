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
      <div
        class="badge"
        class:error={badgeType === 'error'}
        class:success={badgeType === 'success'}
        class:warning={badgeType === 'warning'}
      >
        {badgeText}
      </div>
    {/if}
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

  .badge {
    border-radius: 0.5rem;
    padding: 0 0.5rem;
    margin-right: 1rem;
    font-size: 0.75rem;

    &.success {
      color: var(--theme-text-color);
      border: 1px solid var(--theme-success-color);
      background-color: lighten($success, 35%);
    }

    &.warning {
      color: var(--theme-text-color);
      border: 1px solid var(--theme-warning-color);
      background-color: lighten($warning, 35%);
    }

    &.error {
      color: var(--theme-text-color);
      border: 1px solid var(--theme-error-color);
      background-color: lighten($error, 35%);
    }
  }
</style>
