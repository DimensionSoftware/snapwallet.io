<script lang="ts">
  import { configStore } from '../stores/ConfigStore'
</script>

<div class="modal-footer">
  <slot />
  {#if configStore.environment === 'sandbox'}
    <small title="Testing in Sandbox Mode">Test Mode</small>
  {/if}
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
    padding: 0 3rem 3rem 3rem;
    span {
      cursor: pointer;
      font-size: 0.9rem;
      margin: 0.5rem 0 -0.9rem 0;
    }

    small {
      position: absolute;
      bottom: 0.25rem;
      color: var(--theme-button-color);
    }
    .hidden {
      display: none;
      cursor: initial;
    }
  }
</style>
