<script lang="ts">
  import { pop } from 'svelte-spa-router'
  import FaIcon from 'svelte-awesome'
  import { faChevronLeft } from '@fortawesome/free-solid-svg-icons'

  export let hideCloseButton = false
  export let hideBackButton = false

  const handleExit = () => {
    const event = JSON.stringify({
      event: '__SNAP_EXIT',
    })
    if (window.parent) {
      window.parent.postMessage(event, '*')
    }
    if ((window as any).ReactNativeWebView) {
      ;(window as any).ReactNativeWebView?.postMessage(event)
    }
  }
</script>

<div class="modal-header">
  <div
    on:click={pop}
    class:hidden={hideBackButton}
    class="modal-header-back-button"
  >
    <FaIcon data={faChevronLeft} />
  </div>
  <div class="modal-header-title">
    <slot />
  </div>
  <div
    class:hidden={hideCloseButton}
    on:click={handleExit}
    class="modal-header-close-button"
  >
    <svg
      class="text-main"
      fill="none"
      height="24"
      stroke="currentColor"
      stroke-linecap="round"
      stroke-linejoin="round"
      stroke-width="2"
      viewBox="0 0 24 24"
      width="24"
      xmlns="http://www.w3.org/2000/svg"
      ><line x1="18" y1="6" x2="6" y2="18" /><line
        x1="6"
        y1="6"
        x2="18"
        y2="18"
      /></svg
    >
  </div>
</div>

<style lang="scss">
  @import '../styles/_vars.scss';

  @mixin flex-align-center {
    display: flex;
    align-items: center;
    flex: 1;
    height: 100%;
  }

  .modal-header {
    margin-bottom: 1rem;
    height: 2.75rem;
    width: 100%;
    display: flex;
    justify-content: space-evenly;
    align-items: center;
    & > .modal-header-title {
      @include flex-align-center();
      flex: 3;
      justify-content: center;
      font-weight: bold;
      font-size: 1.2rem;
    }
    & > .modal-header-close-button {
      @include flex-align-center();
      margin-right: 0.2em;
      justify-content: flex-end;
      cursor: pointer;
    }
    & > .modal-header-back-button {
      @include flex-align-center();
      margin-left: 0.2em;
      justify-content: flex-start;
      cursor: pointer;
    }
  }

  .hidden {
    visibility: hidden;
    cursor: initial;
  }
</style>
