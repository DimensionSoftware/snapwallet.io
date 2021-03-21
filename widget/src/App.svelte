<script lang="ts">
  import Router, { push, location, pop } from 'svelte-spa-router'
  import wrap from 'svelte-spa-router/wrap'
  import Toast from './components/Toast.svelte'
  import Home from './screens/Home.svelte'
  import SendOTP from './screens/SendOTP.svelte'
  import NotFound from './screens/NotFound.svelte'
  import Profile from './screens/Profile.svelte'
  import Address from './screens/Address.svelte'
  import VerifyOTP from './screens/VerifyOTP.svelte'
  import Overview from './screens/Overview.svelte'
  import { onMount, setContext } from 'svelte'
  import PlaidWidget from './screens/PlaidWidget.svelte'
  import SelectPayment from './screens/SelectPayment.svelte'
  import { Routes, APIErrors } from './constants'
  import { authedRouteOptions, isJWTValid, Logger, parseJwt } from './util'
  import { userStore } from './stores/UserStore'
  import { toaster } from './stores/ToastStore'
  import Address2 from './screens/Address2.svelte'
  import FileUpload from './screens/FileUpload.svelte'

  // Querystring provided props, see main.ts.
  export let appName: string
  export let intent: 'buy' | 'sell'
  export let apiKey: string
  export let theme: object

  // Handler for routing condition failure
  const routeConditionsFailed = (event: any): boolean => {
    Logger.debug('route conditions failed', event.detail)
    const isAccessingAuthRoutes = [Routes.SEND_OTP, Routes.VERIFY_OTP].includes(
      event.detail.location,
    )
    if (isAccessingAuthRoutes) {
      pop()
      return false
    }
    // Sets the last known route for redirect
    // upon successful auth/reauth.
    userStore.updateLastKnownRoute($location as Routes)
    push(Routes.SEND_OTP)
  }

  const routes = {
    [Routes.ROOT]: wrap({
      component: Home as any,
      props: { appName, intent, apiKey },
    }),
    [Routes.SELECT_PAYMENT]: wrap({
      component: SelectPayment as any,
    }),
    [Routes.SEND_OTP]: wrap({
      component: SendOTP as any,
      props: { appName, intent, apiKey },
      conditions: [() => !isJWTValid()],
    }),
    [Routes.VERIFY_OTP]: wrap({
      component: VerifyOTP as any,
      conditions: [() => !isJWTValid()],
    }),
    // Authenticated
    [Routes.PROFILE]: wrap({
      ...authedRouteOptions(Profile),
    }),
    [Routes.PLAID_LINK]: wrap({
      ...authedRouteOptions(PlaidWidget),
    }),
    [Routes.CHECKOUT_OVERVIEW]: wrap({
      ...authedRouteOptions(Overview),
    }),
    [Routes.ADDRESS]: wrap({
      ...authedRouteOptions(Address),
    }),
    [Routes.FILE_UPLOAD]: wrap({
      ...authedRouteOptions(FileUpload),
    }),
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

    // Centralized error handler
    window.onunhandledrejection = e => {
      Logger.error(e)
      const msg = 'Oops, an unexpected error occurred. Please try again later.'
      const { reason, body } = e

      if (
        reason?.body?.code === APIErrors.UNAUTHORIZED &&
        ($location as Routes) !== Routes.VERIFY_OTP
      ) {
        // expired session, so-- automagically logout
        // - handle re-routing elsewhere
        return window.AUTH_MANAGER.logout()
      }

      toaster.pop({
        msg: reason?.body?.message || body?.message || msg,
        error: true,
      })
    }
  })

</script>

<div id="modal">
  <div id="modal-body">
    <Router on:conditionsFailed={routeConditionsFailed} {routes} />
    <Toast />
  </div>
</div>

<svelte:head>
  <script
    src="https://js.pusher.com/7.0/pusher.min.js"
    on:load={window.tryInitializePusher}></script>
  <script
    src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"></script>
</svelte:head>

<style lang="scss">
  @import './styles/_vars.scss';

  :global(*) {
    box-sizing: border-box;
  }

  :global(a, a:hover, a:active) {
    cursor: pointer;
    text-decoration: none;
  }

  :root {
    // theme
    --theme-color: #{$themeColor};
    --theme-color-inverse: #{$themeColorInverse};
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
    width: 360px;
    height: 608px;
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
