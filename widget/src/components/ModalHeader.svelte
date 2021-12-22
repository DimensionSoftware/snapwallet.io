<script lang="ts">
  import { pop } from 'svelte-spa-router'
  import FaIcon from 'svelte-awesome'
  import { faChevronLeft } from '@fortawesome/free-solid-svg-icons'
  import SideMenu from './SideMenu.svelte'

  export let onBack: () => any = undefined
  export let hideRightAction = false
  export let hideBackButton = false
  export let onClick: () => any = undefined
</script>

<div
  class="modal-header"
  style={`cursor: ${onClick ? 'pointer' : 'auto'}`}
  on:click={onClick}
>
  <div
    on:mousedown={onBack ? onBack : pop}
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
    flex: 1;
  }

  .modal-header {
    margin-bottom: 0.5rem;
    // height: 2.75rem;
    line-height: 1.25rem;
    width: 100%;
    display: flex;
    color: var(--theme-text-color);
    // Match this with body and footer
    padding: 2rem 1.5rem 0 1.5rem;
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
      margin-right: 0;
      justify-content: flex-end;
      @include backface;
    }
    & > .modal-header-back-button {
      @include flex-align-center();
      margin-left: 0.2em;
      justify-content: flex-start;
      cursor: pointer;
      @include backface;
    }
    .hidden {
      visibility: hidden;
    }
  }
</style>
