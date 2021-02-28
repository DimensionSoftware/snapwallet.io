<script lang="ts">
  import { pop } from 'svelte-spa-router'
  import FaIcon from 'svelte-awesome'
  import { faArrowLeft, faTimes } from '@fortawesome/free-solid-svg-icons'

  export let hideCloseButton = false
  export let hideBackButton = false
</script>

<div class="modal-header">
  <div
    on:click={pop}
    class:hidden={hideBackButton}
    class="modal-header-back-button"
  >
    <FaIcon data={faArrowLeft} />
  </div>
  <div class="modal-header-title">
    <slot />
  </div>
  <div class:hidden={hideCloseButton} class="modal-header-close-button">
    <FaIcon data={faTimes} />
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
      justify-content: flex-start;
      cursor: pointer;
    }
  }

  .hidden {
    visibility: hidden;
    cursor: initial;
  }
</style>
