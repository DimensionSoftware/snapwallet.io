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
          <FaIcon data={icon} />
          <div alt={label} class="label">{label}</div>
        </div>
        {#if badgeText}
          <Badge
            class="account-card-badge"
            error={badgeType === 'error'}
            success={badgeType === 'success'}
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
    border: 1px solid var(--theme-color);
    box-shadow: none;
    border-radius: 1.25rem;
    &:before {
      content: '';
      position: absolute;
      border-radius: 1.5rem;
      left: -2px;
      right: -2px;
      bottom: -2px;
      top: -2px;
      background-color: var(--theme-modal-background-color);
      border: 2px solid transparent;
      opacity: 0;
      transition: opacity 0.2s ease-in, border 0s ease-out 0.524s;
    }
    &:hover {
      transition: none;
      &:before {
        left: -3px;
        right: -3px;
        bottom: -3px;
        top: -3px;
        transform: scale(1);
        opacity: 1;
        border: 2px solid var(--theme-color);
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
    & > .label {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      font-weight: 500;
      margin-left: 0.5rem;
    }
  }
</style>
