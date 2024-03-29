<script lang="ts">
  import Router, { push, location, pop } from 'svelte-spa-router'
  import wrap from 'svelte-spa-router/wrap'
  import FaIcon from 'svelte-awesome'
  import { faLock } from '@fortawesome/free-solid-svg-icons'
  import Toast from './components/Toast.svelte'
  import PreLogout from './components/PreLogout.svelte'
  import Home from './screens/Home.svelte'
  import SendOTP from './screens/SendOTP.svelte'
  import NotFound from './screens/NotFound.svelte'
  import Profile from './screens/Profile.svelte'
  import Transactions from './screens/Transactions.svelte'
  import Address from './screens/Address.svelte'
  import VerifyOTP from './screens/VerifyOTP.svelte'
  import Overview from './screens/Overview.svelte'
  import AwaitPayment from './screens/AwaitPayment.svelte'
  import CartCheckout from './screens/CartCheckout.svelte'
  import { onMount, setContext } from 'svelte'
  import PlaidWidget from './screens/PlaidWidget.svelte'
  import SelectPayment from './screens/SelectPayment.svelte'
  import {
    Routes,
    APIErrors,
    ParentMessages,
    SUPPORTED_CRYPTOCURRENCY_ASSETS,
  } from './constants'
  import {
    authedRouteOptions,
    capitalize,
    isJWTValid,
    Logger,
    onEscPressed,
  } from './util'
  import { ParentMessenger } from './util/parent_messenger'
  import { userStore } from './stores/UserStore'
  import { toaster } from './stores/ToastStore'
  import FileUpload from './screens/FileUpload.svelte'
  import SendOtp from './screens/SendOTP.svelte'
  import VerifyOtp from './screens/VerifyOTP.svelte'
  import Success from './screens/Success.svelte'
  import CartSuccess from './screens/CartSuccess.svelte'
  import { transactionStore } from './stores/TransactionStore'
  import { paymentMethodStore } from './stores/PaymentMethodStore'
  import ProfileStatus from './screens/ProfileStatus.svelte'
  import Product from './screens/Product.svelte'
  import { configStore } from './stores/ConfigStore'
  import DebitCard from './screens/DebitCard.svelte'
  import DebitCardAddress from './screens/DebitCardAddress.svelte'
  import DebitCard2fa from './screens/DebitCard2fa.svelte'
  import { debitCardStore } from './stores/DebitCardStore'
  import TransactionDetails from './screens/TransactionDetails.svelte'
  import { transactionDetailsStore } from './stores/TransactionsStore'

  $: isPreLogout = false
  $: isBlurred = false
  $: isHeaderBlurred = false

  $: isFramed = window != window.top
  $: isOpen = false // was modal opened via openWeb()?

  // auth bits
  window.addEventListener('logout', _ => {
    Logger.debug('viewer has logged out')
    isPreLogout = false
    userStore.setIsLoggedIn(false)
    push(Routes.ROOT)
    toaster.pop({
      msg: 'You have been securely logged out.',
      error: true,
    })
  })
  window.addEventListener('prelogout', _ => {
    Logger.debug('viewer is prelogout')
    isPreLogout = true
  })

  // blurry fx
  window.addEventListener('blurryHeader', _ => {
    isHeaderBlurred = true
  })
  window.addEventListener('unblurryHeader', _ => {
    isHeaderBlurred = false
  })
  window.addEventListener('blurry', _ => {
    isBlurred = true
  })
  window.addEventListener('unblurry', _ => {
    isBlurred = false
  })

  // screen height events
  const HEIGHT = '608px', // default screen height
    WIDTH = '360px' // " width
  let height: string = HEIGHT,
    width: string = WIDTH,
    lastLocation: string = null
  window.addEventListener(ParentMessages.RESIZE, (event: Event) => {
    // respond to custom screen heights
    height = event.detail?.height || HEIGHT
    width = event.detail?.width || WIDTH
    ParentMessenger.resize({ height, width }, $configStore.appName)
  })
  $: {
    if (lastLocation !== $location) {
      // reset screen height & width at every change
      height = HEIGHT
      width = WIDTH
      ParentMessenger.resize({ height, width }, $configStore.appName) // iframe
      lastLocation = $location
    }
  }

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
      if ((e.target as Element)?.id === 'modal') ParentMessenger.exit()
    }

  const routes = {
    [Routes.ROOT]: wrap({
      component: ($configStore.product?.destinationTicker
        ? Product
        : $configStore.intent === 'cart'
        ? CartCheckout
        : Home) as any,
    }),
    [Routes.CART_CHECKOUT]: wrap({
      component: CartCheckout as any,
    }),
    [Routes.AWAIT_PAYMENT]: wrap({
      component: AwaitPayment as any,
    }),
    [Routes.SELECT_PAYMENT]: wrap({
      component: SelectPayment as any,
    }),
    [Routes.SEND_OTP]: wrap({
      component: SendOTP as any,
      conditions: [() => !isJWTValid()],
    }),
    [Routes.VERIFY_OTP]: wrap({
      component: VerifyOTP as any,
      conditions: [() => !isJWTValid()],
    }),
    [Routes.CART_SUCCESS]: wrap({
      component: CartSuccess as any,
    }),
    // Authenticated
    [Routes.SUCCESS]: wrap({
      ...authedRouteOptions(Success),
      conditions: [isJWTValid, () => Boolean($transactionStore.wyrePreview)],
    }),
    [Routes.PROFILE]: wrap({
      ...authedRouteOptions(Profile),
    }),
    [Routes.PROFILE_UPDATE]: wrap({
      ...authedRouteOptions(Profile),
      props: {
        isUpdateScreen: true,
      },
    }),
    [Routes.PROFILE_STATUS]: wrap({
      ...authedRouteOptions(ProfileStatus),
    }),
    [Routes.TRANSACTIONS]: wrap({
      ...authedRouteOptions(Transactions),
    }),
    [Routes.TRANSACTION_DETAILS]: wrap({
      ...authedRouteOptions(TransactionDetails),
      conditions: [
        isJWTValid,
        () => Boolean($transactionDetailsStore.transaction),
      ],
    }),
    [Routes.PLAID_LINK]: wrap({
      ...authedRouteOptions(PlaidWidget),
    }),
    [Routes.CHECKOUT_OVERVIEW]: wrap({
      ...authedRouteOptions(Overview),
      conditions: [isJWTValid, () => Boolean($transactionStore.wyrePreview)],
    }),
    [Routes.ADDRESS]: wrap({
      ...authedRouteOptions(Address),
    }),
    [Routes.ADDRESS_UPDATE]: wrap({
      ...authedRouteOptions(Address),
      props: {
        isUpdateScreen: true,
      },
    }),
    [Routes.FILE_UPLOAD]: wrap({
      ...authedRouteOptions(FileUpload),
    }),
    [Routes.FILE_UPLOAD_UPDATE]: wrap({
      ...authedRouteOptions(FileUpload),
      props: {
        isUpdatingFiles: true,
      },
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
    [Routes.DEBIT_CARD]: wrap({
      ...authedRouteOptions(DebitCard),
      conditions: [
        isJWTValid,
        () => {
          return Boolean(
            $debitCardStore.address.country &&
              $transactionStore.wyrePreview &&
              $transactionStore.sourceAmount,
          )
        },
      ],
    }),
    [Routes.DEBIT_CARD_ADDRESS]: wrap({
      ...authedRouteOptions(DebitCardAddress),
      conditions: [
        isJWTValid,
        () => {
          return Boolean(
            $debitCardStore.address.country &&
              $transactionStore.wyrePreview &&
              $transactionStore.sourceAmount &&
              $debitCardStore.reservationId,
          )
        },
      ],
    }),
    [Routes.DEBIT_CARD_2FA]: wrap({
      ...authedRouteOptions(DebitCard2fa),
      conditions: [
        isJWTValid,
        () => {
          return Boolean(
            $debitCardStore.address.country &&
              $debitCardStore.address.street1 &&
              $transactionStore.wyrePreview &&
              $transactionStore.sourceAmount &&
              $debitCardStore.reservationId,
          )
        },
      ],
    }),
    '*': NotFound as any,
  }

  // Set theme context so theme can be used in JS also
  setContext('theme', {
    ...$configStore.theme,
  })

  userStore.subscribe(state => {
    // Set this once fetchGeo runs
    if (!$debitCardStore.address?.country) {
      debitCardStore.updateAddress({ country: state.geo?.country || '' })
    }
  })

  onMount(() => {
    const defaultDestinationAsset = SUPPORTED_CRYPTOCURRENCY_ASSETS.find(
      sca => {
        return (
          sca.ticker.toLowerCase() ===
          $configStore.defaultDestinationAsset?.toLowerCase()
        )
      },
    )

    if (defaultDestinationAsset) {
      transactionStore.setDestinationCurrency(defaultDestinationAsset)
    }

    // pre-fetch user
    userStore.fetchGeo()
    if (window.AUTH_MANAGER.viewerIsLoggedIn()) {
      userStore.fetchUserProfile()
      paymentMethodStore.fetchWyrePaymentMethods()
      // set user flags up, non-blocking
      userStore.fetchFlags()
    }
    // Override theme css variables
    Object.entries($configStore.theme).forEach(([k, v]) => {
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
        Logger.debug('Logout called from onunhandledrejection because of:', e)
        // expired session, so-- automagically logout
        window.AUTH_MANAGER.logout()
        transactionStore.reset()
        userStore.reset()
        paymentMethodStore.reset()
        push(Routes.ROOT)
      }

      // show toast
      toaster.pop({
        msg: capitalize(
          reason?.body?.message || body?.message || reason?.message || msg,
        ),
        error: true,
      })
    }

    // init ui triggers
    isOpen = $configStore.isOpen
    isFramed = window != window.top
  })
</script>

<svelte:window on:keydown={onKeyDown} on:mousedown={onMouseDown} />

<div id="modal" class:isFramed class:isOpen>
  <div
    id="modal-body"
    style={`height: ${height}; width: ${width}`}
    class:blur={isPreLogout || isBlurred}
    class:blur-header={isHeaderBlurred}
  >
    <Router on:conditionsFailed={routeConditionsFailed} {routes} />
    <Toast />
    {#if isPreLogout}
      <PreLogout
        onClosed={() => (isPreLogout = false)}
        isVisible={isPreLogout}
      />
    {/if}
  </div>
  <div class="lock">
    <FaIcon scale="3" data={faLock} />
  </div>
</div>

<svelte:head>
  <script
    src="https://js.pusher.com/7.0/pusher.min.js"
    on:load={window.tryInitializePusher}></script>
  <script
    defer
    src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"></script>
  <script
    defer
    src="https://maps.googleapis.com/maps/api/js?key=AIzaSyCrIHFoF52WJV5iH4-IyPd1CbBRwXXOymw&libraries=places"></script>
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
    --theme-input-text-color: #{inputTextColor};
    --theme-badge-text-color: #{$badgeTextColor};
    --theme-text-color: #{$textColor};
    --theme-text-color-2: #{$textColor2};
    --theme-text-color-3: #{$textColor3};
    --theme-text-color-4: #{$textColor4};
    --theme-text-color-muted: #{$textColorMuted};
    --theme-text-color-no-background: #{$textColorNoBackground};
    --theme-button-color: #{$buttonColor};
    // XXX glow color format is "r,g,b", eg. "255,50,15"
    --theme-button-glow-color: #{toRGB($buttonGlowColor)};
    --theme-button-text-color: #{$buttonTextColor};
    --theme-modal-background: #{$modalBackground};
    --theme-modal-background-color: #{$modalBackgroundColor};
    --theme-modal-popup-background: #{$modalPopupBackground};
    --theme-modal-container-background-color: #{$modalContainerBackgroundColor};
    --theme-shadow-color: #{$shadowColor};
    --theme-shadow-bottom-color: #{$shadowBottomColor};
    --theme-success-color: #{$success};
    --theme-error-color: #{$error};
    --theme-warning-color: #{$warning};
    --theme-info-color: #{$info};

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
    &:hover {
      text-decoration: underline;
    }
  }
  :global(.spacer) {
    margin-top: 1.5rem;
  }
  :global(.scroll-y) {
    overflow: hidden;
    overflow-y: scroll;
    overscroll-behavior: contain;
    scrollbar-width: none;
    -ms-overflow-style: none;
  }
  :global(.scroll-y::-webkit-scrollbar) {
    display: none;
    background: transparent;
    height: 0;
    width: 0;
  }
  :global(iframe.modal #modal:before) {
    // modal overlay bg
    opacity: 0 !important;
    animation: backgroundFadeIn 1s ease-out forwards !important;
    // background: var(--theme-modal-container-background-color) !important;
    background: transparent;
    background-color: transparent;
  }
  #modal:before {
    content: '';
    position: absolute;
    left: 0;
    right: 0;
    bottom: 0;
    top: 0;
    opacity: 0;
    animation: backgroundFadeIn 1s ease-out forwards;
    background-color: transparent;
    // background: var(--thme-modal-container-background-color);
    background: transparent;
  }
  #modal.isFramed :global(.lock) {
    display: none;
  }
  #modal.isOpen.isFramed:before,
  :global(iframe.sw-open #modal:before) {
    background: var(--theme-modal-container-background-color) !important;
  }
  #modal,
  :global(#plaid-link-iframe-1) {
    position: absolute;
    z-index: 1;
    left: 0;
    top: 0;
    right: 0;
    bottom: 0;
    width: 100%;
    height: 100%;
    overflow: hidden !important;
    opacity: 1;
    display: flex;
    justify-content: center;
    align-items: center;
    color: var(--theme-text-color);
    font-family: var(--theme-font);
    font-size: 1rem;
    line-height: 1.1rem;
    opacity: 0;
    animation: backgroundFadeIn 0.2s ease-out forwards;
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
    transition: height 0.3s var(--theme-ease-out-back),
      width 0.4s var(--theme-ease-out-back) 0.301s;
    will-change: height, width, transform;
    background: var(--theme-modal-background);
    border-radius: 1rem;
    overflow: hidden;
    transform: translateZ(0);
    display: flex;
    flex-direction: column;
    // Used by toast
    position: relative;
    :global(.popup-selector-header),
    :global(.modal-content .modal-body),
    :global(.modal-content .modal-header-title),
    :global(.modal-content .modal-header-back-button),
    :global(.modal-content .modal-footer) {
      will-change: filter;
      transition: filter 0.5s var(--theme-ease-out-expo);
      backface-visibility: hidden;
      perspective: 1000;
      transform: translate3d(0, 0, 0);
      transform: translateZ(0);
    }
    &.blur-header {
      :global(.popup-selector-header),
      :global(.modal-content .modal-header-title),
      :global(.modal-content .modal-header-right-action svg),
      :global(.modal-content .modal-header-back-button) {
        filter: blur(20px) contrast(150%);
        transition: none;
      }
    }
    &.blur {
      :global(.modal-content .modal-body),
      :global(.modal-content .modal-header-title),
      :global(.modal-content .modal-header-back-button),
      :global(.modal-content .modal-footer) {
        filter: blur(50px) contrast(175%);
        transition: none;
      }
    }
    :global(h3) {
      color: var(--theme-text-color);
      opacity: 0.8;
      font-size: 0.85rem;
      font-weight: 400;
      margin: 0 0 1rem 0;
      text-align: center;
    }
    :global(h3.test),
    :global(h3.test a) {
      color: var(--theme-color);
      font-weight: bold;
      opacity: 1;
      text-align: center;
    }
  }

  @media screen and (max-width: 550px) {
    #modal-body {
      border-radius: 0;
      height: 100% !important;
      width: 100%;
    }
    .lock {
      display: none;
    }
  }
</style>
