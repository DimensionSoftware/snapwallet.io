<script lang="ts">
  import { ParentMessenger } from '../util/parent_messenger'
  export let hideCloseButton = false
</script>
<div class="modal-footer">
  <slot />
  <span
    class:hidden={hideCloseButton}
    on:click={ParentMessenger.exit}
    class="modal-header-right-action"
  >
  Close
</span>
</div>

<style lang="scss">
  .modal-footer {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-end;
    padding: 0 initial;
    span {
      cursor: pointer;
      margin-top: .5rem;
      font-size: .9rem;
    }
  }
</style>
