<script lang="ts">
  import { ParentMessenger } from '../util/parent_messenger'
  import { isEmbedded } from '../util'
  export let hideCloseButton = false
</script>

<div class="modal-footer">
  <slot />
  <span
    class:hidden={!isEmbedded || hideCloseButton}
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
    height: 85px;
    span {
      cursor: pointer;
      margin-top: 0.5rem;
      font-size: 0.9rem;
    }
    .hidden {
      visibility: hidden;
      cursor: initial;
    }
  }
</style>
