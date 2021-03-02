<script lang="ts">
  import { quintOut } from 'svelte/easing'
  import { fly } from 'svelte/transition'
  export let animation: string = ''
</script>

<div class="modal-content {animation}">
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
      transform: translateX(0);
    }
    100% {
      transform: translateX(100%);
    }
  }
  .right {
    animation: slideInRight 0.23s var(--theme-ease-out-expo);
  }
  .left {
    animation: slideInLeft 0.25s var(--theme-ease-out-back);
  }
  .modal-content {
    position: absolute;
    top: 1em;
    bottom: 4.5em;
    left: 1em;
    right: 1em;
  }
</style>
