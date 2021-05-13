<script lang="ts">
  import { toaster } from '../stores/ToastStore'
  import { fade, fly } from 'svelte/transition'
  import { backOut } from 'svelte/easing'
  import { capitalize } from '../util'

  $: success = Boolean($toaster?.success)
  $: warning = Boolean($toaster?.warning)
  $: error = Boolean($toaster?.error)
</script>

<div
  class="toast-wrapper"
  title="Click to Dismiss"
  on:mousedown={toaster.dismiss}
>
  {#if $toaster}
    <div
      class="toast-item"
      class:success
      class:warning
      class:error
      in:fly={{
        delay: 0,
        duration: 250,
        x: 0,
        y: 25,
        opacity: 0.1,
        easing: backOut,
      }}
      out:fade={{ duration: 250 }}
    >
      {capitalize($toaster?.msg)}
    </div>
  {/if}
</div>

<style lang="scss">
  @import '../styles/_vars.scss';

  .toast-wrapper {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    width: 100%;
    display: flex;
    justify-content: center;
    z-index: 9999;
    // Don't overlay top of modal
    height: 0px;
    padding: 0 0.75rem;
    cursor: pointer;
  }

  .toast-item {
    position: relative;
    display: flex;
    align-items: center;
    padding: 0 1rem;
    width: 100%;
    min-height: 100px;
    max-height: 250px;
    color: var(--theme-text-color);
    font-weight: 500;
    z-index: 9999;
    &:before {
      background: var(--theme-color);
      content: '';
      position: absolute;
      opacity: 0.05;
      top: 0;
      bottom: 0;
      left: 0;
      right: 0;
      z-index: -1;
    }
    &.error:before {
      background: var(--theme-error-color);
      background: var(--theme-modal-background);
    }
    &.warning:before {
      background: var(--theme-warning-color);
      background: var(--theme-modal-background);
    }
    &.success:before {
      background-color: var(--theme-success-color);
      background: var(--theme-modal-background);
    }
  }
</style>
