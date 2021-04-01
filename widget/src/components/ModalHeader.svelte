<script lang="ts">
  import { pop } from 'svelte-spa-router'
  import FaIcon from 'svelte-awesome'
  import { faChevronLeft } from '@fortawesome/free-solid-svg-icons'
  import SideMenu from './SideMenu.svelte'

  export let onBack: () => any = undefined
  export let hideRightAction = false
  export let hideBackButton = false
</script>

<div class="modal-header">
  <div
    on:click={onBack ? onBack : pop}
    class:hidden={hideBackButton}
    class="modal-header-back-button"
  >
    <FaIcon data={faChevronLeft} />
  </div>
  <div class="modal-header-title">
    <slot />
  </div>
  <div class:hidden={hideRightAction} class="modal-header-right-action">
    <slot name="right">
      <SideMenu />
    </slot>
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
      text-align: center;
    }
    & > .modal-header-right-action {
      @include flex-align-center();
      margin-right: 0.2em;
      justify-content: flex-end;
    }
    & > .modal-header-back-button {
      @include flex-align-center();
      margin-left: 0.2em;
      justify-content: flex-start;
      cursor: pointer;
    }
    .hidden {
      visibility: hidden;
    }
  }
</style>
