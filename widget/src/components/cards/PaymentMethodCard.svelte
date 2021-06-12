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
  export let badgeType:
    | 'error'
    | 'warning'
    | 'success'
    | 'info'
    | undefined = undefined
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
    height: 8rem;
    width: 100%;
    display: flex;
    justify-content: space-between;
    flex-direction: column;
    align-items: center;
    border: 1px solid var(--theme-color-lightened);
    border-radius: 0.5rem;
    &.paddingSmall {
      padding: 0.5rem;
    }
    & > .content-container {
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
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 100%;
  }

  .header-title {
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
