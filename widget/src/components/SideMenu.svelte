<script lang="ts">
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import { focusFirstInput, onKeysPressed } from '../util'
  import { userStore } from '../stores/UserStore'

  export let isExpanded: boolean = false

  function logout() {
    close()
    // yield to ui animations
    setTimeout(() => window.AUTH_MANAGER.logout(), 100)
  }
  function login() {
    push(Routes.SEND_OTP)
    close()
  }
  function go(route) {
    push(route)
    close()
  }

  function close() {
    isExpanded = false
    focusFirstInput()
  }
  function handleClose(e) {
    // close if esc pressed
    if (onKeysPressed(e, ['Escape'])) close()
  }

  $: isLoggedIn = $userStore.isLoggedIn
</script>

<svelte:window on:keydown={handleClose} />

<div
  class="container"
  class:active={isExpanded}
  on:click={_ => {
    if (isExpanded) {
      close()
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
      stroke-width={isExpanded ? 5 : isLoggedIn ? 4 : 3}
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
    <a on:click={_ => go(Routes.ROOT)}>Buy Crypto Assets</a>
    <a class="hr" on:click={_ => go(Routes.TRANSACTIONS)}>My Transactions</a>
    <a on:click={_ => go(Routes.PROFILE)}>My Profile</a>
    {#if isLoggedIn}
      <a class="hr" on:click={logout}>Logout</a>
    {:else}
      <a class="hr" on:click={login}>Login</a>
    {/if}
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
    top: 0.1rem;
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
    background: var(--theme-modal-background);
    top: 0;
    left: -0.5rem;
    right: -0.5rem;
    bottom: 0;
    width: 125%;
    height: 125%;
    padding: 25% 2rem 0;
    transform: translateX(105%);
    transition: transform 0.35s var(--theme-ease-in-expo);
    z-index: 100;
    nav a {
      position: relative;
      display: block;
      margin: 1rem 0;
      font-size: 1.25rem;
      transform: translateX(50px);
      transition: transform 0s ease-out 0.6s;
      &.hr {
        margin-top: 2rem;
        position: relative;
      }
    }
    &.active {
      transition: transform 0.2s var(--theme-ease-out-expo);
      transform: translateX(0);
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
    hr {
      opacity: 0.1;
      height: 1px;
      margin: 1.5rem 0;
    }
  }
</style>
