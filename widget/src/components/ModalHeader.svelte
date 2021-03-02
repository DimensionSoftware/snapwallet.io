<script lang="ts">
  import { pop } from 'svelte-spa-router'
  import FaIcon from 'svelte-awesome'
  import { faChevronLeft } from '@fortawesome/free-solid-svg-icons'

  export let hideCloseButton = false
  export let hideBackButton = false

  const handleExit = () => {
    const event = JSON.stringify({
      event: '__FLUX_EXIT',
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
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"
      ><path fill="none" d="M0 0h24v24H0z" /><path
        fill="currentColor"
        d="M19.03 6.03l-1.06-1.06L12 10.94 6.03 4.97 4.97 6.03 10.94 12l-5.97 5.97 1.06 1.06L12 13.06l5.97 5.97 1.06-1.06L13.06 12l5.97-5.97z"
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
    margin-top: 1rem;
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
