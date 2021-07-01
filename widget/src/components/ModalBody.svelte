<script lang="ts">
  export let fullscreen: boolean = false
  export let padded: boolean = false
</script>

<div class="modal-body scroll-y" class:fullscreen class:padded>
  <slot />
</div>

<style lang="scss">
  .modal-body {
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    padding: 1rem 3rem 0 3rem;
    &.padded {
      padding: 1rem 3rem 0 3rem;
    }
    &.fullscreen {
      padding: 0;
    }
  }
</style>
