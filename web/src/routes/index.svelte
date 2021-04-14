<script context="module" lang="ts">
  export const prerender = true
</script>

<script lang="ts">
  let Typewriter: any
  import { onMount } from 'svelte'
  const domain = 'https://snapwallet.io'

  let ifr: HTMLIFrameElement

  onMount(async () => {
    await import('flux-init')
    Typewriter = (await import('svelte-typewriter')).default

    // respond to widget events
    window.addEventListener(
      'message',
      ({ data }) => {
        if (data) {
          const payload = JSON.parse(data)
          // resize
          if (payload.event === '__SNAP_RESIZE')
            ifr?.height = payload.data?.size
        }
      },
      false,
    )

    const SnapWallet = new (window as any).Snap({
      appName: 'Snap Wallet',
      intent: 'buy',
      wallets: [],
      focus: true,
      // theme: {
      //   modalBackground: 'rgba(40,40,40,.9)',
      //   modalPopupBackground: 'rgba(50,50,50,.95)',
      //   color: 'rgba(0,0,0,.9)',
      //   badgeTextColor: '#333',
      //   colorLightened: 'rgba(5,5,5,.8)',
      //   shadowBottomColor: 'rgba(0,0,0,.25)',
      //   colorInverse: '#fff',
      //   buttonColor: '#fffc00',
      //   buttonTextColor: '#000',
      //   successColor: '#fffc00',
      //   textColor: '#fff',
      //   inputTextColor: '#333',
      // },
      theme: {
        modalBackground: 'rgb(28, 28, 40)',
        modalPopupBackground: 'rgb(40, 41, 61, .95)',
        color: 'rgba(0,0,0,.9)',
        badgeTextColor: '#333',
        colorLightened: 'rgba(5,5,5,.8)',
        shadowBottomColor: 'rgba(0,0,0,.25)',
        colorInverse: '#fff',
        buttonColor: 'rgb(172, 93, 217)',
        successColor: 'rgb(0, 211, 149)',
        warningColor: 'rgb(253, 221, 72)',
        textColor: '#e4f0fb',
        inputTextColor: '#333',
      },
    })

    ifr.onload = () => {
      ifr.classList.add('loaded')
    }
    ifr.src = SnapWallet.generateURL()
  })
</script>

<main>
  <div class="intro col">
    <h1>Snap Wallet</h1>
    <h2>
      Connect Crypto to Your
      {#if Typewriter}
        <Typewriter interval={50} loop={1800}>
          <span>Idea</span>
          <span>NFT</span>
          <span>App</span>
          <span>Company</span>
          <span>Site</span>
          <span>Donations</span>
        </Typewriter>
      {:else}
        Idea
      {/if}
    </h2>
    <article>
      The "Add Money" button for Crypto Currency, a fully configurable, gorgeous
      wallet that delights customers.
    </article>
    <div class="buttons col">
      <a class="button" href={`${domain}/docs/guide`} target="_blank"
        >Get Started</a
      >
      <!-- <a href={`https://snapwallet.io/docs/api`} target="_blank"
        >API Documentation</a
      > -->
    </div>
    <!-- Features, Advantages, Benefits -->
    <!--ul>
      <li>Fast</li>
      <li>Secure</li>
      <li>Embeddable</li>
      <li>Customs pay in USD; you receive your preferred currency</li>
    </ul-->
  </div>
  <div class="col wallet" style="margin: 0 auto;">
    <iframe
      title="Snap Wallet"
      frameborder="0"
      height="608px"
      width="360px"
      bind:this={ifr}
    />
  </div>
</main>

<style lang="scss">
  @import '../../../widget/src/styles/animations.scss';
  $textColor: #333;
  $easeOutExpo: cubic-bezier(0.16, 1, 0.3, 1);
  $easeOutBack: cubic-bezier(0.34, 1.25, 0.64, 1);
  main {
    display: flex;
    max-width: 960px;
    margin: 15% auto 0;
    .col {
      position: relative;
      max-width: 50%;
      .buttons {
        margin: 3rem 0 0;
        max-width: 100%;
        .button {
          display: inline-block;
          font-weight: 600;
          border-radius: 3px;
          background: rgba($textColor, 0.9);
          color: #fff;
          padding: 1rem 3rem;
          margin-right: 1.5rem;
          transition: background 0.3s ease-out;
          &:hover {
            animation: scaleIn 0.5s ease-out forwards;
            background: $textColor;
            transition: none;
          }
        }
      }
    }
    h1,
    h2 {
      white-space: nowrap;
      font-size: 2rem;
      line-height: 1.1;
      margin: 0.5rem 0;
    }
    h1 {
      font-weight: 500;
      margin-bottom: 0.25rem;
    }
    h2 {
      margin: 0 0 1.75rem;
      font-weight: 500;
      font-size: 1.25rem;
      :global(div) {
        display: inline-block;
      }
    }
    a {
      color: $textColor;
      text-decoration: none;
      white-space: nowrap;
      margin-bottom: 1rem;
    }
    article {
      font-size: 1.1rem;
      margin: 0;
      line-height: 1.35;
    }
    :global(iframe) {
      position: relative;
      border-radius: 20px;
      top: -20%;
      right: -20%;
    }
    :global(iframe.loaded) {
      box-shadow: 0 0 25px rgba(0, 0, 0, 0.4);
      transition: box-shadow 0.8s $easeOutExpo 0.1s, height 0.3s $easeOutBack;
      will-change: box-shadow, height;
    }
  }
  @media (min-width: 480px) {
    h1,
    h2 {
      max-width: none;
    }
    article {
      max-width: none;
    }
  }

  // responsive
  @media (max-width: 375px) {
    :global(body) {
      overflow-y: scroll !important;
      main {
        padding: 0;
      }
    }
  }
  @media (max-width: 1024px) {
    :global(body) {
      overflow-y: scroll !important;
      main {
        flex-direction: column;
        > .col {
          max-width: 100%;
        }
        .intro {
          padding-left: 2rem;
          padding-right: 1rem;
          article {
            max-width: none;
          }
        }
        .wallet {
          max-width: inherit;
          padding-top: 5rem;
          padding-bottom: 5rem;
          iframe {
            top: inherit;
            right: inherit;
          }
        }
      }
    }
  }
</style>
