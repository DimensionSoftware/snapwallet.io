<script lang="ts">
  export let animation: string = 'left'
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
    display: flex;
    flex-direction: column;
    height: 100%;
    overflow: hidden;
  }

  @media (max-width: 480px) {
    .left {
      transform: none;
      animation: none;
    }
  }
</style>
