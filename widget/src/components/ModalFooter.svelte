<script lang="ts">
  import { ParentMessenger } from '../util/parent_messenger'
  export let hideClose = false
</script>

<div class="modal-footer">
  <slot />
  <!-- <span
    class:hidden={hideClose}
    on:click={ParentMessenger.exit}
    class="modal-header-right-action"
  >
    Close
  </span> -->
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
    padding: 0 1rem 1rem 1rem;
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
  :global(#modal.hide-close) .modal-footer > span {
    display: none;
  }
</style>
