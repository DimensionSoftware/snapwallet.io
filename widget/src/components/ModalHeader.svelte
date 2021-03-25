<script lang="ts">
  import { pop } from 'svelte-spa-router'
  import FaIcon from 'svelte-awesome'
  import { faChevronLeft } from '@fortawesome/free-solid-svg-icons'
  import { ParentMessenger } from '../util/parent_messenger'

  export let hideRightAction = false
  export let hideBackButton = false
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
    class:hidden={hideRightAction}
    on:click={ParentMessenger.exit}
    class="modal-header-right-action"
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
    & > .modal-header-right-action {
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
</style>
