<script lang="ts">
  import Card from './Card.svelte'
  import FaIcon from 'svelte-awesome'
  import { createEventDispatcher } from 'svelte'
  import Badge from '../Badge.svelte'
  const dispatch = createEventDispatcher()

  export let icon
  export let title: string = ''
  export let label: string
  export let paddingSmall = false
  export let blend = false
  export let badgeText = ''
  export let badgeType: 'error' | 'warning' | 'success' | 'info' | undefined =
    undefined
</script>

<Card on:click={() => dispatch('click')}>
  <div {title} class="icon-card-container" class:blend class:paddingSmall>
    <div class="content-container">
      <div class="header-container" style="">
        <div class="header-title">
          {#if icon}
            <FaIcon data={icon} />
          {/if}
          <div alt={label} class="label">{label}</div>
        </div>
        {#if badgeText}
          <Badge
            class="account-card-badge"
            error={badgeType === 'error'}
            success={['success', 'confirmed', 'completed'].includes(
              badgeType.toLocaleLowerCase(),
            )}
            warning={badgeType === 'warning'}
            info={badgeType === 'info'}
          >
            {badgeText}
          </Badge>
        {/if}
      </div>
      <slot />
    </div>
  </div>
</Card>

<style lang="scss">
  @import '../../styles/_vars.scss';

  .icon-card-container {
    position: relative;
    width: 100%;
    display: flex;
    justify-content: space-between;
    flex-direction: column;
    align-items: center;
    box-shadow: none;
    &:after {
      content: '';
      position: absolute;
      left: 30%;
      right: 30%;
      bottom: -3px;
      background-color: var(--theme-text-color);
      opacity: 0.2;
      height: 1px;
    }
    &:before {
      content: '';
      position: absolute;
      left: -2px;
      right: -2px;
      bottom: -2px;
      top: -2px;
      border-radius: 0.5rem;
      background-color: var(--theme-color-lightened);
      opacity: 0;
      transform: scale(0.975);
      transition: opacity 0.2s ease-in, transform 0.1s ease-out;
    }
    &:hover {
      transition: none;
      &:after {
        opacity: 0;
      }
      &:before {
        transform: scale(1);
        opacity: 0.3;
        transition: none;
      }
      .header-title {
        transition: none;
      }
    }
    &.paddingSmall {
      padding: 0.5rem;
    }
    & > .content-container {
      position: relative;
      height: 100%;
      width: 100%;
      display: flex;
      padding: 1rem;
      justify-content: flex-start;
      flex-direction: column;
      align-items: center;
    }
  }

  .header-container {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 100%;
    margin-bottom: 0.25rem;
  }

  .header-title {
    position: relative;
    display: flex;
    align-items: center;
    :global(svg) {
      margin-right: 0.5rem;
    }
    & > .label {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      font-weight: 500;
    }
  }
</style>
