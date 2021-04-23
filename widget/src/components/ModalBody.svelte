<script lang="ts">
  export let fullscreen: boolean = false
</script>

<div class="modal-body scroll-y" class:fullscreen>
  <slot />
</div>

<style lang="scss">
  .modal-body {
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    padding: 1rem 1.25rem 0 1.25rem;
    &.fullscreen {
      padding: 0;
    }
  }
</style>
