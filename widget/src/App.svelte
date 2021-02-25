<script lang="ts">
  import Router from 'svelte-spa-router'
  import wrap from 'svelte-spa-router/wrap'
  import Toast from './components/Toast.svelte'
  import Home from './screens/Home.svelte'
  import Checkout from './screens/Checkout.svelte'
  import NotFound from './screens/NotFound.svelte'

  // Querystring provided props, see main.ts.
  export let appName: string
  export let intent: 'buy' | 'sell'
  export let apiKey: string

  const routes = {
    '/': wrap({ component: Home as any, props: { appName, intent, apiKey } }),
    '/checkout': wrap({ component: Checkout as any, props: { appName, intent, apiKey } }),
    '*': NotFound as any,
  }
</script>

<div id="modal">
  <div id="modal-body">
    <Router {routes} />
    <Toast />
  </div>
</div>

<style lang="scss">
  @import './styles/_vars.scss';

  :global(*) {
    box-sizing: border-box;
  }

  #modal {
    position: absolute;
    z-index: 1;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: hidden;
    background: $modalContainerBackgroundColor;
    display: flex;
    justify-content: center;
    align-items: center;
    color: $textColor;
    font-family: $themeFont;
    font-size: 1rem;
    line-height: 1.5rem;
    text-rendering: optimizeLegibility;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  #modal-body {
    width: 320px;
    height: 540px;
    background-color: $modalBackgroundColor;
    padding: 1rem;
    border-radius: 1rem;
    overflow: hidden;
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
