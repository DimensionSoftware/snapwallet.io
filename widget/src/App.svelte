<script lang="ts">
  import Router from 'svelte-spa-router'
  import wrap from 'svelte-spa-router/wrap'
  import Toast from './components/Toast.svelte'
  import Home from './screens/Home.svelte'
  import NotFound from './screens/NotFound.svelte'

  // Querystring provided props, see main.ts.
  export let appName: string
  export let intent: 'buy' | 'sell'
  export let apiKey: string

  const routes = {
    '/': wrap({ component: Home as any, props: { appName, intent, apiKey } }),
    '*': NotFound as any,
  }
</script>

<div class="modal">
  <div id="modal-body">
    <Router {routes} />
    <Toast />
  </div>
</div>

<style lang="scss">
  @import './styles/_vars.scss';

  :global(body, html) {
    color: $textColor1;
    margin: 0;
    padding: 0;
    height: 100%;
    width: 100%;
    position: relative;
    box-sizing: border-box;
  }

  .modal {
    position: absolute;
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
    display: flex;
    flex-direction: column;
    // Used by toast
    position: relative;
  }

  @media screen and (max-width: 450px) {
    #modal-body {
      border-radius: 0;
      height: 100%;
      width: 100%;
    }
  }
</style>
