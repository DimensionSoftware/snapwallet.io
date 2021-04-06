<script lang="ts">
  export let animation: string = 'left'
  export let fullscreen: boolean = false
</script>

<div class="modal-content {animation}" class:fullscreen>
  <slot />
</div>

<style lang="scss">
  @import '../styles/_vars.scss';
  @keyframes slideInLeft {
    0% {
      transform: translateX(100%);
    }
    100% {
      transform: translateX(0);
    }
  }
  @keyframes slideInRight {
    0% {
      transform: translateX(-100%);
    }
    100% {
      transform: translateX(0);
    }
  }
  .right {
    transform: translateX(-100%);
    animation: slideInRight 0.4s var(--theme-ease-out-expo) forwards;
  }
  .left {
    animation: slideInLeft 0.25s var(--theme-ease-out-back) forwards;
  }
  .modal-content {
    position: absolute;
    top: 1em;
    left: 1em;
    right: 1em;
    bottom: 8.5rem;
    display: flex;
    flex-direction: column;
    height: 100%;
    &.fullscreen {
      left: 0;
      right: 0;
      bottom: 0;
      top: 0;
      :global(.modal-header) {
        padding: 2rem 1rem 0;
      }
      :global(.modal-footer) {
        padding-top: 0.5rem;
        margin: 0 1rem;
      }
    }
  }
</style>
