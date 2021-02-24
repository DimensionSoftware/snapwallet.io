<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import Icon from 'svelte-awesome'
  import { faTimesCircle } from '@fortawesome/free-regular-svg-icons'

  const dispatch = createEventDispatcher()

  export let visible: boolean
  export let headerTitle: string

  function handleClose(e: Event) {
    if (e instanceof KeyboardEvent) {
      // close on esc
      if (e.key === 'Escape') dispatch('close')
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

<div class="popup-selector" class:visible>
  <div class="popup-selector-header">
    <div class="placeholder-box" />
    <div class="popup-title">{headerTitle}</div>
    <div on:click={() => dispatch('close')} class="close-icon">
      <Icon data={faTimesCircle} />
    </div>
  </div>
  <slot />
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';

  .popup-selector {
    padding: 1rem;
    display: flex;
    flex-direction: column;
    height: 95%;
    width: 100%;
    border-radius: 1rem;
    z-index: 1000;
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: white;
    box-shadow: 0px -7px 36px -10px $commonShadowColor;
    transform: translateY(100%);
    transition: transform 0.08s $easeInExpo;
    &.visible {
      transform: translateY(0);
      transition: transform 0.25s $easeOutExpo;
    }
  }

  .popup-selector-header {
    display: flex;
    height: 50px;
    width: 100%;
    justify-content: space-between;
    align-items: center;
  }

  .placeholder-box {
    width: 40px;
    height: 40px;
  }

  .popup-title {
    font-weight: bold;
  }

  .close-icon {
    min-width: 40px;
    min-height: 40px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
  }
</style>
