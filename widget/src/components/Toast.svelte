<script lang="ts">
  import { toaster } from '../stores/ToastStore'
  import { fade, fly } from 'svelte/transition'
  import { backOut } from 'svelte/easing'
  import { capitalize } from '../util'

  $: success = Boolean($toaster?.success)
  $: warning = Boolean($toaster?.warning)
  $: error = Boolean($toaster?.error)
</script>

<div class="toast-wrapper">
  {#if $toaster}
    <div
      class="toast-item"
      class:success
      class:warning
      class:error
      in:fly={{
        delay: 0,
        duration: 300,
        x: 0,
        y: 50,
        opacity: 0.1,
        easing: backOut,
      }}
      out:fade={{ duration: 500 }}
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
    z-index: 2;
    // Don't overlay top of modal
    height: 0px;
  }

  .toast-item {
    position: relative;
    display: flex;
    align-items: center;
    padding: 0 1rem;
    width: 100%;
    min-height: 60px;
    max-height: 100px;
    color: white;
    font-weight: 500;
    z-index: 999;
    &:before {
      background: linear-gradient(to right, var(--theme-color), var(--theme-color-darkened));
      content: '';
      position: absolute;
      opacity: 0.98;
      top: 0;
      bottom: 0;
      left: 0;
      right: 0;
      z-index: -1;
    }
    &.error:before {
      background: linear-gradient(to right, var(--theme-color), var(--theme-color-darkened));
    }
    &.warning:before {
      background: linear-gradient(to right, var(--theme-color), var(--theme-color-darkened));
    }
    &.success:before {
      background-color: var(--theme-success-color);
    }
  }
</style>
