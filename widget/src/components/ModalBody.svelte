<script lang="ts">
  export let fullscreen: boolean = false
</script>

<div class="modal-body" class:fullscreen>
  <slot />
</div>

<style lang="scss">
  .modal-body {
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    overflow: hidden !important;
    overflow-y: scroll;
    scrollbar-width: thin;
    padding: 1rem 1rem 0 1rem;
    &.fullscreen {
      padding: 0;
    }
  }
</style>
