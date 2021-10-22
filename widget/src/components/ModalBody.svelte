<script lang="ts">
  export let fullscreen: boolean = false
  export let padded: boolean = false
  export let scrollable: boolean = true
</script>

<div
  class="modal-body"
  class:scroll-y={scrollable}
  class:fullscreen
  class:padded
>
  <slot />
</div>

<style lang="scss">
  .modal-body {
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    padding: 1rem 2rem 0 2rem;
    &.padded {
      padding: 1rem 3rem 0 3rem;
    }
    &.fullscreen {
      padding: 0;
    }
  }
</style>
