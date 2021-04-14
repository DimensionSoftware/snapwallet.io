<script lang="ts">
  import { push } from 'svelte-spa-router'
  import FaIcon from 'svelte-awesome'
  import {
    faLock,
    faExchangeAlt,
    faList,
    faSignOutAlt,
    faSignInAlt,
    faUserCircle,
  } from '@fortawesome/free-solid-svg-icons'
  import { Routes } from '../constants'
  import {
    cachePrimaryPaymentMethodID,
    focusFirstInput,
    onKeysPressed,
  } from '../util'
  import { userStore } from '../stores/UserStore'
  import { transactionsStore } from '../stores/TransactionsStore'
  import { transactionStore } from '../stores/TransactionStore'
  import { paymentMethodStore } from '../stores/PaymentMethodStore'
  import { ParentMessenger } from '../util/parent_messenger'

  export let isExpanded: boolean = false
  export let isProductCheckout: boolean = false
  let slow: boolean = false

  function logout() {
    close()
    // yield to ui animations
    setTimeout(() => {
      window.AUTH_MANAGER.logout()
      cachePrimaryPaymentMethodID('')
      transactionStore.reset()
      userStore.reset()
      paymentMethodStore.reset()
      push(Routes.ROOT)
    }, 100)
  }
  function login() {
    push(Routes.SEND_OTP)
    close()
  }
  function go(route) {
    push(route)
    close()
  }

  function close(isSlow = false) {
    slow = isSlow
    isExpanded = false
    focusFirstInput()
  }
  function handleClose(e) {
    // close if esc pressed
    if (onKeysPressed(e, ['Escape'])) close(true)
  }

  $: isLoggedIn = $userStore.isLoggedIn
  $: {
    // open/close
    setTimeout(
      () => window.dispatchEvent(new Event(isExpanded ? 'blurry' : 'unblurry')),
      isExpanded ? 0 : slow ? 300 : 150,
    )

    if (isExpanded && window.AUTH_MANAGER.viewerIsLoggedIn()) {
      // pre-cache transactions
      transactionsStore.fetchUserTransactions()
    }
  }
</script>

<svelte:window on:keydown={handleClose} />

<div
  class="container"
  class:active={isExpanded}
  on:mousedown={_ => {
    if (isExpanded) {
      close(true)
    } else {
      isExpanded = !isExpanded
    }
  }}
>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="50"
    height="50"
    viewBox="0 0 200 200"
  >
    <g
      stroke-width={isExpanded ? 5 : isLoggedIn ? 8 : 4}
      stroke-linecap="round"
    >
      <path
        d="M72 82.286h28.75"
        fill="#009100"
        fill-rule="evenodd"
        stroke="#333"
      />
      <path
        d="M100.75 103.714l72.482-.143c.043 39.398-32.284 71.434-72.16 71.434-39.878 0-72.204-32.036-72.204-71.554"
        fill="none"
        stroke="#333"
      />
      <path
        d="M72 125.143h28.75"
        fill="#009100"
        fill-rule="evenodd"
        stroke="#333"
      />
      <path
        d="M100.75 103.714l-71.908-.143c.026-39.638 32.352-71.674 72.23-71.674 39.876 0 72.203 32.036 72.203 71.554"
        fill="none"
        stroke="#333"
      />
      <path
        d="M100.75 82.286h28.75"
        fill="#009100"
        fill-rule="evenodd"
        stroke="#333"
      />
      <path
        d="M100.75 125.143h28.75"
        fill="#009100"
        fill-rule="evenodd"
        stroke="#333"
      />
    </g>
  </svg>
</div>
<aside class:active={isExpanded}>
  <nav>
    <div>
      <FaIcon data={faExchangeAlt} />
      <a on:mousedown={_ => go(Routes.ROOT)}
        >{isProductCheckout ? 'View Cart' : 'Buy Crypto Assets'}</a
      >
    </div>
    <div />
    <div>
      <FaIcon data={faList} />
      <a class="hr" on:mousedown={_ => go(Routes.TRANSACTIONS)}
        >My Transactions</a
      >
    </div>
    <div>
      <FaIcon data={faUserCircle} />
      <a on:mousedown={_ => go(Routes.PROFILE_STATUS)}>My Profile</a>
    </div>
    <div />
    {#if isLoggedIn}
      <div>
        <FaIcon data={faSignOutAlt} />
        <a class="hr" on:mousedown={logout}>Logout</a>
      </div>
    {:else}
      <div>
        <FaIcon data={faSignInAlt} />
        <a class="hr" on:mousedown={login}>Login</a>
      </div>
    {/if}
    <div>
      <FaIcon data={faLock} />
      <a
        on:mousedown={() => {
          if (isExpanded) {
            close(true)
          }
          ParentMessenger.exit()
        }}>Exit</a
      >
    </div>
  </nav>
</aside>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';
  .container {
    font-size: 1.2rem;
    z-index: 101;
    position: relative;
    right: -0.9rem;
    top: 0.15rem;
    cursor: pointer;
  }
  // not-hamburger hamburger toggle + fx
  svg {
    transition: transform 300ms cubic-bezier(0.4, 0, 0.2, 1);
  }
  .active svg {
    transform: rotate(90deg);
  }
  path {
    transition: transform 300ms cubic-bezier(0.4, 0, 0.2, 1),
      stroke-dasharray 300ms cubic-bezier(0.4, 0, 0.2, 1),
      stroke-dashoffset 300ms cubic-bezier(0.4, 0, 0.2, 1);
    stroke: var(--theme-text-color);
  }
  path:nth-child(1) {
    transform-origin: 36% 40%;
  }
  path:nth-child(2) {
    stroke-dasharray: 29 299;
  }
  path:nth-child(3) {
    transform-origin: 35% 63%;
  }
  path:nth-child(4) {
    stroke-dasharray: 10 299;
  }
  path:nth-child(5) {
    transform-origin: 61% 52%;
  }
  path:nth-child(6) {
    transform-origin: 62% 52%;
  }
  .active path:nth-child(1) {
    transform: translateX(9px) translateY(1px) rotate(45deg);
  }
  .active path:nth-child(2) {
    stroke-dasharray: 225 299;
    stroke-dashoffset: -72px;
  }
  .active path:nth-child(3) {
    transform: translateX(9px) translateY(1px) rotate(-45deg);
  }
  .active path:nth-child(4) {
    stroke-dasharray: 225 299;
    stroke-dashoffset: -72px;
  }
  .active path:nth-child(5) {
    transform: translateX(9px) translateY(1px) rotate(-45deg);
  }
  .active path:nth-child(6) {
    transform: translateX(9px) translateY(1px) rotate(45deg);
  }
  // menu
  aside {
    position: absolute;
    background: var(--theme-modal-popup-background);
    top: -1rem;
    left: -1rem;
    right: -0.5rem;
    bottom: -0.5rem;
    width: 125%;
    height: 150%;
    padding: 25% 1rem 0 3rem;
    transform: translateX(105%);
    transition: transform 0.35s var(--theme-ease-in-expo);
    z-index: 100;
    nav > div {
      display: flex;
      align-items: center;
      margin: 1rem 0 0 0;
      :global(svg) {
        opacity: 0;
        transition: opacity 0s ease-in .8s;
      }
    }
    nav a {
      position: relative;
      display: block;
      margin: -.1rem 0 0 1rem !important;
      color: var(--theme-text-color);
      font-size: 1.35rem;
      transform: translateX(50px);
      transition: transform 0s ease-out 0.5s;
      &.hr {
        margin-top: 2.5rem;
        position: relative;
      }
    }
    &.active {
      transition: transform 0.2s var(--theme-ease-out-expo);
      transform: translateX(0);
      nav > div :global(svg) {
        opacity: 1;
        &:nth-child(0) {
          transition: opacity 0.4s ease-out;
        }
        &:nth-child(1) {
          transition: opacity 0.5s ease-out;
        }
        &:nth-child(2) {
          transition: opacity 0.6s ease-out;
        }
        &:nth-child(3) {
          transition: opacity 0.7s ease-out;
        }
        &:nth-child(4) {
          transition: opacity .8s ease-out;
        }
        &:nth-child(5) {
          transition: opacity .9s ease-out;
        }
        &:nth-child(6) {
          transition: opacity 1.0s ease-out;
        }
        &:nth-child(7) {
          transition: opacity 1.1s ease-out;
        }
      }
      nav a {
        transform: translateX(0);
        &:nth-child(0) {
          transition: transform 0.3s var(--theme-ease-out-expo);
        }
        &:nth-child(1) {
          transition: transform 0.4s var(--theme-ease-out-expo);
        }
        &:nth-child(2) {
          transition: transform 0.5s var(--theme-ease-out-expo);
        }
        &:nth-child(3) {
          transition: transform 0.6s var(--theme-ease-out-expo);
        }
        &:nth-child(4) {
          transition: transform 0.7s var(--theme-ease-out-expo);
        }
        &:nth-child(5) {
          transition: transform 0.8s var(--theme-ease-out-expo);
        }
        &:nth-child(6) {
          transition: transform 0.9s var(--theme-ease-out-expo);
        }
        &:nth-child(7) {
          transition: transform 1s var(--theme-ease-out-expo);
        }
      }
    }
  }
</style>
