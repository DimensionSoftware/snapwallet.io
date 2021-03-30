<div class="modal-body">
  <slot />
</div>

<style lang="scss">
  .modal-body {
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    flex: 1;
    &:nth-child(1) {
      // modal body content
      padding: 0 0.2rem;
    }
  }
</style>
