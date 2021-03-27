<script lang="ts">
  import { isEmbedded } from '../util'
  export let animation: string = 'left'
</script>

<div class="modal-content {animation}" style={`bottom: ${isEmbedded ? '5.5em' : '4.5em'}`}>
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
    animation: slideInRight 0.2s var(--theme-ease-out-expo) forwards;
  }
  .left {
    animation: slideInLeft 0.25s var(--theme-ease-out-back) forwards;
  }
  .modal-content {
    position: absolute;
    top: 1em;
    left: 1em;
    right: 1em;
  }
</style>
