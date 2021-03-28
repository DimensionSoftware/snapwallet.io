<script lang="ts">
  import Router, { push, location, pop } from 'svelte-spa-router'
  import wrap from 'svelte-spa-router/wrap'
  import FaIcon from 'svelte-awesome'
  import { faLock } from '@fortawesome/free-solid-svg-icons'
  import Toast from './components/Toast.svelte'
  import Home from './screens/Home.svelte'
  import SendOTP from './screens/SendOTP.svelte'
  import NotFound from './screens/NotFound.svelte'
  import Profile from './screens/Profile.svelte'
  import Transactions from './screens/Transactions.svelte'
  import Address from './screens/Address.svelte'
  import VerifyOTP from './screens/VerifyOTP.svelte'
  import Overview from './screens/Overview.svelte'
  import { onMount, setContext } from 'svelte'
  import PlaidWidget from './screens/PlaidWidget.svelte'
  import SelectPayment from './screens/SelectPayment.svelte'
  import { Routes, APIErrors } from './constants'
  import { authedRouteOptions, isJWTValid, Logger, onEscPressed } from './util'
  import { ParentMessenger } from './util/parent_messenger'
  import { userStore } from './stores/UserStore'
  import { toaster } from './stores/ToastStore'
  import FileUpload from './screens/FileUpload.svelte'
  import SendOtp from './screens/SendOTP.svelte'
  import VerifyOtp from './screens/VerifyOTP.svelte'
  import Success from './screens/Success.svelte'

  // Querystring provided props, see main.ts.
  export let appName: string
  export let intent: 'buy' | 'sell'
  export let apiKey: string
  export let theme: object
  export let focus: boolean

  // Handler for routing condition failure
  const routeConditionsFailed = (event: any): boolean => {
    Logger.debug(
      'route conditions failed',
      event.detail,
      $userStore.lastKnownRoute,
    )
    const isAccessingAuthRoutes = [Routes.SEND_OTP, Routes.VERIFY_OTP].includes(
      event.detail.location,
    )

    // Don't allow user to hit send or verify when authed.
    if (isAccessingAuthRoutes) {
      pop()
      return false
    }

    // Allow back from OTP screens
    if ($userStore.lastKnownRoute === $location) {
      userStore.updateLastKnownRoute(Routes.ROOT)
      push(Routes.ROOT)
      return
    }

    // Sets the last known route for redirect
    // upon successful auth/reauth.
    userStore.updateLastKnownRoute($location as Routes)
    push(Routes.SEND_OTP)
  }

  // close modal on escape and outside mouse click
  const onKeyDown = (e: Event) => {
      if (e.target !== document.body) onEscPressed(e, ParentMessenger.exit)
    },
    onMouseDown = (e: MouseEvent) => {
      if ((e.target as Element).id === 'modal') ParentMessenger.exit()
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
    [Routes.SUCCESS]: wrap({
      ...authedRouteOptions(Success),
    }),
    [Routes.PROFILE]: wrap({
      ...authedRouteOptions(Profile),
    }),
    [Routes.TRANSACTIONS]: wrap({
      ...authedRouteOptions(Transactions),
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
    [Routes.PROFILE_SEND_SMS]: wrap({
      ...authedRouteOptions(SendOtp),
      props: {
        phoneVerificationOnly: true,
      },
    }),
    [Routes.PROFILE_VERIFY_SMS]: wrap({
      ...authedRouteOptions(VerifyOtp),
      props: {
        phoneVerificationOnly: true,
      },
    }),
    '*': NotFound as any,
  }

  // Set theme context so theme can be used in JS also
  setContext('theme', {
    ...theme,
  })

  // Override theme css variables
  onMount(() => {
    if (focus) setTimeout(() => document.getElementById('amount')?.focus(), 350)
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

      // show toast
      const isWyreErr =
        reason instanceof String
          ? reason.match(/wyre.APIError.*Message:.(.+)?"/)
          : false
      toaster.pop({
        msg: isWyreErr
          ? isWyreErr[0]
          : reason?.body?.message || body?.message || msg,
        error: true,
      })
    }
  })
</script>

<svelte:window on:keydown={onKeyDown} on:mousedown={onMouseDown} />

<div id="modal">
  <div id="modal-body">
    <Router on:conditionsFailed={routeConditionsFailed} {routes} />
    <Toast />
  </div>
  <FaIcon class="lock" data={faLock} />
</div>

<svelte:head>
  <script
    src="https://js.pusher.com/7.0/pusher.min.js"
    on:load={window.tryInitializePusher}></script>
  <script
    async
    src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"></script>
  <script
    async
    src="https://maps.googleapis.com/maps/api/js?key=AIzaSyDr7FQk1bZV4Zght87YNUgCv5P4cg_1DIs&libraries=places"></script>
</svelte:head>

<style lang="scss">
  @import './styles/_vars.scss';
  @import './styles/animations.scss';

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

  :global(a) {
    color: var(--theme-color);
    position: relative;
    text-decoration: underline;
  }

  #modal {
    position: absolute;
    z-index: 1;
    left: 0;
    top: 0;
    right: 0;
    bottom: 0;
    width: 100%;
    height: 100%;
    overflow: hidden;
    background: var(--theme-modal-container-background-color);
    opacity: 0;
    animation: backgroundFadeIn 1s ease-out forwards;
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
    :global(.lock) {
      position: absolute;
      content: '';
      bottom: 20px;
      right: 20px;
      height: 61px;
      width: 61px;
      opacity: 0.4;
    }
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
    animation: scaleIn 0.25s var(--theme-ease-out-back);
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
