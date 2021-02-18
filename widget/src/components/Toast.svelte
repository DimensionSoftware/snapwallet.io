<script lang="ts">
  import { toaster } from '../stores/ToastStore'
  import { fade, fly } from 'svelte/transition'
  import { backOut } from 'svelte/easing'

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
      {$toaster?.msg}
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
    display: flex;
    align-items: center;
    padding: 0 1rem;
    width: 100%;
    min-height: 50px;
    max-height: 100px;
    background-color: $info;
    color: white;
    &.error {
      background-color: $error;
    }
    &.warning {
      background-color: $warning;
    }
    &.success {
      background-color: $success;
    }
  }
</style>
