<script lang="ts">
  import Toast from './components/Toast.svelte'
  import { toaster } from './stores/ToastStore'

  // Querystring provided props, see main.ts.
  export let name: string
  export let intent: 'buy' | 'sell'
  export let apiKey: string
</script>

<div class="modal">
  <div id="modal-body">
    <p>Product name: {name}</p>
    <p>Intent: {intent}</p>
    <p>API Key: {apiKey}</p>

    <button
      on:click={() => {
        toaster.pop({ msg: 'Success', success: true })
      }}>Toaster</button
    >
    <Toast />
  </div>
</div>

<style lang="scss">
  @import './styles/_vars.scss';

  :global(body, html, main, div) {
    color: $textColor1;
    margin: 0;
    padding: 0;
    height: 100%;
    width: 100%;
    position: relative;
  }

  .modal {
    position: fixed;
    z-index: 1;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: hidden;
    background-color: $modalContainerBackgroundColor;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  #modal-body {
    width: 320px;
    height: 540px;
    background-color: $modalBackgroundColor;
    padding: 1rem;
    border-radius: 1rem;
    overflow: hidden;
    overflow-y: scroll;
  }
</style>
