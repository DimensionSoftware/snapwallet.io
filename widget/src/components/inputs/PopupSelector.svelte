<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import { createEventDispatcher } from 'svelte'
  import { onKeysPressed } from '../../util'
  import { scale } from 'svelte/transition'
  import { expoOut } from 'svelte/easing'

  const dispatch = createEventDispatcher()

  export let headerTitle: string

  // lifecycle events
  onMount(() => {
    // XXX could use svelte's dispatcher
    window.dispatchEvent(new Event('blurry'))
  })
  onDestroy(() => {
    window.dispatchEvent(new Event('unblurry'))
  })

  function handleClose(e: Event) {
    if (onKeysPressed(e, ['Escape'])) {
      // close if esc pressed
      e.stopPropagation()
      dispatch('close')
    } else if (e instanceof MouseEvent) {
      // close when bg clicked
      if (e.target instanceof HTMLDivElement) {
        if (e.target.id.indexOf('modal-body') >= 0) {
          dispatch('close')
        }
      }
    }
  }
</script>

<svelte:window on:keydown={handleClose} on:click={handleClose} />

<div
  out:scale={{ opacity: 0, start: 1.15, duration: 175, easing: expoOut }}
  class="popup-selector"
>
  <div class="popup-selector-header">
    <div class="popup-title">{headerTitle}</div>
    <div on:click={() => dispatch('close')} class="close-icon">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        height="100%"
        width="100%"
        viewBox="0 0 24 24"
        ><path fill="none" d="M0 0h24v24H0z" /><path
          fill="currentColor"
          d="M19.03 6.03l-1.06-1.06L12 10.94 6.03 4.97 4.97 6.03 10.94 12l-5.97 5.97 1.06 1.06L12 13.06l5.97 5.97 1.06-1.06L13.06 12l5.97-5.97z"
        /></svg
      >
    </div>
  </div>
  <slot />
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';
  @import '../../styles/animations.scss';
  .popup-selector {
    display: flex;
    flex-direction: column;
    width: 100%;
    z-index: 1000;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    padding-top: 0.5rem;
    background: var(--theme-modal-popup-background);
    animation: slideUp 0.2s var(--theme-ease-out-expo) forwards;
    &:after {
      @include bottom-shadow;
    }
  }

  .popup-selector-header {
    display: flex;
    /* background: var(--theme-modal-background); */
    height: 50px;
    padding: 1rem;
    width: 100%;
    justify-content: space-between;
    align-items: center;
  }

  .popup-title {
    font-weight: 400;
    font-size: 1.25em;
    margin-left: 0.5rem;
  }

  .close-icon {
    width: 2rem;
    height: 2rem;
    border-radius: 100%;
    background: var(--theme-modal-popup-background);
    padding: 0.3rem;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
  }
</style>
