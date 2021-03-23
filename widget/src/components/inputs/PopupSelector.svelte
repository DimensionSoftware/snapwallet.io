<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import { onKeysPressed } from '../../util'
  import { fly } from 'svelte/transition'

  const dispatch = createEventDispatcher()

  export let headerTitle: string

  function handleClose(e: Event) {
    if (onKeysPressed(e, ['Escape'])) {
      // close if esc pressed
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

<div out:fly={{ y: 15, duration: 125 }} class="popup-selector">
  <div class="popup-selector-header">
    <div class="popup-title">{headerTitle}</div>
    <div on:click={() => dispatch('close')} class="close-icon">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"
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
    padding: 1rem;
    display: flex;
    flex-direction: column;
    width: 100%;
    border-radius: 1rem;
    z-index: 1000;
    position: absolute;
    top: 2%;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: white;
    box-shadow: 0px -7px 25px 10px var(--theme-shadow-color);
    animation: slideUp 0.2s $easeOutExpo forwards;
  }

  .popup-selector-header {
    display: flex;
    height: 50px;
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
    width: 25px;
    height: 25px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
  }
</style>
