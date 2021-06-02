<div class="modal-footer">
  <slot />
</div>

<style lang="scss">
  .modal-footer {
    /* height: 10%; */
    flex-grow: 1;
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-end;
    // Match with body and header
    padding: 0 1.5rem 2rem 1.5rem;
    span {
      cursor: pointer;
      font-size: 0.9rem;
      margin: 0.5rem 0 -0.9rem 0;
    }
    .hidden {
      display: none;
      cursor: initial;
    }
  }
</style>
