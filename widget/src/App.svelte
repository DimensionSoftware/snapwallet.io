<script lang="ts">
  import Router from 'svelte-spa-router'
  import wrap from 'svelte-spa-router/wrap'
  import Toast from './components/Toast.svelte'
  import Home from './screens/Home.svelte'
  import Checkout from './screens/Checkout.svelte'
  import NotFound from './screens/NotFound.svelte'
  import Profile from './screens/Profile.svelte'
  import VerifyOTP from './screens/VerifyOTP.svelte'
  import { onMount, setContext } from 'svelte'
  import PlaidWidget from './screens/PlaidWidget.svelte'

  // Querystring provided props, see main.ts.
  export let appName: string
  export let intent: 'buy' | 'sell'
  export let apiKey: string
  export let theme: object

  const routes = {
    '/': wrap({ component: Home as any, props: { appName, intent, apiKey } }),
    '/checkout': wrap({
      component: Checkout as any,
      props: { appName, intent, apiKey },
    }),
    '/profile': Profile,
    '/verify-otp': VerifyOTP,
    '/link-bank': PlaidWidget,
    '*': NotFound as any,
  }

  // Set theme context so theme can be used in JS also
  setContext('theme', {
    ...theme,
  })

  // Override theme css variables
  onMount(() => {
    Object.entries(theme).forEach(([k, v]) => {
      k = k.replace(/[A-Z]/g, (k, i) =>
        i === 0 ? k.toLowerCase() : `-${k.toLowerCase()}`,
      )
      document.documentElement.style.setProperty(`--theme-${k}`, v, 'important')
    })
  })
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

  :root {
    // theme
    --theme-color: #{$themeColor};
    --theme-color-darkened: #{$themeColorDarkened};
    --theme-color-lightened: #{$themeColorLightened};
    --theme-font: #{$themeFont};
    --heading-font: #{$headingFont};

    // colors
    --theme-text-color: #{$textColor};
    --theme-text-color-2: #{$textColor2};
    --theme-text-color-3: #{$textColor3};
    --theme-text-color-4: #{$textColor4};
    --theme-text-color-muted: #{$textColorMuted};
    --theme-modal-background: #{$modalBackground};
    --theme-modal-container-background-color: #{$modalContainerBackgroundColor};
    --theme-shadow-color: #{$shadowColor};
    --theme-success-color: #{$success};
    --theme-error-color: #{$error};
    --theme-warning-color: #{$warning};
    --theme-info-color: #{$info};
    --theme-button-shadow-color: #{$buttonShadowColor};

    // easing
    --theme-ease-bounce: #{$bounce};
    --theme-ease-in-back: #{$easeInBack};
    --theme-ease-out-back: #{$easeOutBack};
    --theme-ease-in-expo: #{$easeInExpo};
    --theme-ease-out-expo: #{$easeOutExpo};
  }

  #modal {
    position: absolute;
    z-index: 1;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: hidden;
    background: var(--theme-modal-container-background-color);
    display: flex;
    justify-content: center;
    align-items: center;
    color: var(--theme-text-color);
    font-family: var(--theme-font);
    font-size: 1rem;
    line-height: 1.5rem;
    text-rendering: optimizeLegibility;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  #modal-body {
    width: 320px;
    height: 540px;
    background: var(--theme-modal-background);
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
