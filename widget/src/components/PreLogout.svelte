<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import { scale } from 'svelte/transition'
  import { expoOut } from 'svelte/easing'
  import { focusFirstInput, Logger, onKeysPressed } from '../util'
  import { userStore } from '../stores/UserStore'
  import { transactionStore } from '../stores/TransactionStore'
  import { paymentMethodStore } from '../stores/PaymentMethodStore'
  import Surround from './cards/Surround.svelte'
  import Button from './Button.svelte'

  export let isVisible: boolean = false
  export let onClosed: Function

  // for countdown timer
  let _timer = null
  $: expiresAfter = 0
  $: expiresAfterResolution = 'minutes'

  const logout = () => {
      // close, yielding animations
      close()
      setTimeout(() => {
        Logger.debug('Logout called from PreLogout')
        window.AUTH_MANAGER.logout()
        transactionStore.reset()
        userStore.reset()
        paymentMethodStore.reset()
      }, 100)
    },
    cont = () => {
      // refresh token and close
      window.AUTH_MANAGER.getAccessToken()
      close()
    },
    close = () => {
      isVisible = false
      focusFirstInput()
      onClosed()
    },
    keydown = e => {
      // dismiss if esc pressed
      if (onKeysPressed(e, ['Escape'])) close()
    },
    shouldClose = (e: MouseEvent) => {
      // dismiss if clicked outside content
      if ((e.target as HTMLElement)?.id === 'prelogout' && isVisible) close()
    },
    countdown = () => {
      const diff: number =
          +new Date(window.AUTH_MANAGER.getSessionExpiration()) - +new Date(),
        seconds = Math.round(diff / 1000),
        minutes = Math.round(seconds / 60)
      // set ui
      ;(expiresAfter = seconds < 0 ? 0 : seconds > 59 ? minutes : seconds),
        (expiresAfterResolution =
          seconds > 59
            ? minutes > 1
              ? 'minutes'
              : 'minute'
            : seconds > 1
            ? 'seconds'
            : 'second')
    }

  // component lifecycle events
  onMount(() => {
    countdown()
    _timer = setInterval(countdown, 1000)
  })
  onDestroy(() => {
    clearInterval(_timer)
  })
</script>

<svelte:window on:keydown={keydown} />

<div
  id="prelogout"
  class="container"
  class:visible={isVisible}
  on:mousedown={shouldClose}
  out:scale={{ opacity: 0, start: 1.15, duration: 175, easing: expoOut }}
>
  <Surround>
    <h3>Stay Signed In?</h3>
    <p>
      You will be securely signed out in <big>{expiresAfter}</big>
      {expiresAfterResolution}
    </p>
  </Surround>
  <div class="flex">
    <Button on:mousedown={cont} class="continue">Continue</Button>
    <button on:mousedown={logout} class="logout">Logout</button>
  </div>
</div>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';
  .container {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    padding: 5rem 3rem;
    font-size: 1.2rem;
    z-index: 101;
    background: var(--theme-modal-popup-background);
    :global(.surround) {
      padding: 0.5rem 2rem;
    }
    h3 {
      margin-bottom: 0.25rem;
    }
    p {
      font-size: 0.9rem;
      margin: 0.24rem 0 1.75rem;
    }
    .flex {
      margin-top: 1rem;
      display: flex;
      justify-content: start;
      flex-direction: column;
      .logout {
        cursor: pointer;
        font-size: 1rem;
        width: 100%;
        margin-top: 0.5rem;
        color: var(--theme-text-color);
        background: transparent;
        border: none;
      }
      .continue {
        cursor: pointer;
        margin-right: 0.5rem;
        background: var(--theme-color);
        font-weight: 500;
        padding: 0.4rem 1rem;
        color: var(--theme-color-inverse);
        border: none;
      }
    }
  }
</style>
