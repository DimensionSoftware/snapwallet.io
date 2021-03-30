<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import { scale } from 'svelte/transition'
  import { expoOut } from 'svelte/easing'
  import { focusFirstInput, onKeysPressed } from '../util'

  export let isVisible: boolean = false
  export let onClosed: Function

  // for countdown timer
  let _timer = null
  $: expiresAfter = 0

  const logout = () => {
      // close, yielding animations
      close()
      setTimeout(() => window.AUTH_MANAGER.logout(), 100)
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
      if ((e.target as HTMLElement).id === 'prelogout' && isVisible) close()
    },
    countdown = () => {
      const diff: number =
          +new Date(window.AUTH_MANAGER.getSessionExpiration()) - +new Date(),
        minutes = Math.round(diff / 1000 / 60)
      expiresAfter = minutes < 0 ? 0 : minutes
    }

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
  on:click={shouldClose}
  out:scale={{ opacity: 0, start: 1.15, duration: 175, easing: expoOut }}
>
  <article>
    <h3>Stay Signed In?</h3>
    <p>
      You will be securely signed out in <big>{expiresAfter}</big>
      {expiresAfter > 1 ? 'minutes' : 'minute'}.
    </p>
    <div class="flex">
      <button on:click={logout} class="logout">Logout</button>
      <button on:click={cont} class="continue">Continue</button>
    </div>
  </article>
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
    font-size: 1.2rem;
    z-index: 101;
    background: rgba(200, 200, 200, 0.7);
    article {
      position: relative;
      margin: 0;
      top: 20%;
      background: var(--theme-modal-background);
      border-radius: 3px;
      padding: 1.5rem;
      animation: slideUp 0.3s var(--theme-ease-out-back);
      h3 {
        margin: 0;
      }
      p {
        font-size: 0.9rem;
      }
      .flex {
        display: flex;
        justify-content: start;
        .continue {
          cursor: pointer;
          font-size: 1rem;
          margin-left: 0.5rem;
          color: var(--theme-text-color);
          background: transparent;
          border: none;
        }
        .logout {
          cursor: pointer;
          margin-right: 0.5rem;
          background: var(--theme-color);
          font-weight: 500;
          padding: 0.25rem 1rem;
          color: var(--theme-color-inverse);
          border: none;
        }
      }
    }
  }
</style>
