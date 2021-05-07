<script context="module" lang="ts">
  export const prerender = true
</script>

<script lang="ts">
  let Typewriter: any
  import { onMount } from 'svelte'
  import NFT from '$lib/features/NFT.svelte'
  import Donation from '$lib/features/Donation.svelte'
  import Buy from '$lib/features/Buy.svelte'
  import Footer from '$lib/Footer.svelte'
  const domain = 'https://snapwallet.io'

  let ifr: HTMLIFrameElement

  onMount(async () => {
    await import('flux-init')
    Typewriter = (await import('svelte-typewriter')).default

    const appName = 'Noir Checkout',
      // themeColor = '#E1143D',
      // themeColor = '#F1071C',
      themeColor = '#fffc00',
      SnapWallet = new (window as any).Snap({
        appName,
        intent: 'buy',
        wallets: [],
        focus: true,
        theme: {
          modalBackground: 'rgba(40,40,40,.9)',
          modalPopupBackground: 'rgba(50,50,50,.95)',
          color: 'rgba(0,0,0,.9)',
          badgeTextColor: '#333',
          colorLightened: 'rgba(5,5,5,.8)',
          shadowBottomColor: 'rgba(0,0,0,.25)',
          colorInverse: '#fff',
          buttonColor: themeColor,
          buttonTextColor: '#000',
          buttonGlowColor: themeColor,
          successColor: themeColor,
          textColor: '#fff',
          inputTextColor: '#333',
        },
      })

    // respond to widget events
    window.addEventListener(
      'message',
      ({ data: msg }) => {
        if (!msg) return
        try {
          const { event, data } = JSON.parse(msg)
          if (event === SnapWallet.events.RESIZE && data && ifr) {
            if (appName === data.appName) ifr.height = data.height
          }
        } catch (e) {
          console.warn('Unable to parse message', msg, e)
        }
      },
      false,
    )

    ifr.onload = () => {
      ifr.classList.add('loaded')
    }
    ifr.src = SnapWallet.generateURL()

    // front
    console.log(`
·▄▄▄▄  ▪  • ▌ ▄ ·. ▄▄▄ . ▐ ▄ .▄▄ · ▪         ▐ ▄
██▪ ██ ██ ·██ ▐███▪▀▄.▀·•█▌▐█▐█ ▀. ██ ▪     •█▌▐█
▐█· ▐█▌▐█·▐█ ▌▐▌▐█·▐▀▀▪▄▐█▐▐▌▄▀▀▀█▄▐█· ▄█▀▄ ▐█▐▐▌
██. ██ ▐█▌██ ██▌▐█▌▐█▄▄▌██▐█▌▐█▄▪▐█▐█▌▐█▌.▐▌██▐█▌
▀▀▀▀▀• ▀▀▀▀▀  █▪▀▀▀ ▀▀▀ ▀▀ █▪ ▀▀▀▀ ▀▀▀ ▀█▄▀▪▀▀ █▪
Hey, you-- join us!  https://dimensionsoftware.com
      `)
  })
</script>

<main>
  <div class="intro col">
    <h1>Snap Wallet</h1>
    <h2>
      Connect Crypto to Your
      {#if Typewriter}
        <Typewriter interval={50} delay={0} loop={1800}>
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
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 320"
  ><path
    fill="#ffffff"
    fill-opacity="1"
    d="M0,320L240,288L480,192L720,160L960,96L1200,192L1440,0L1440,320L1200,320L960,320L720,320L480,320L240,320L0,320Z"
  /></svg
>
<NFT />
<Donation />
<Buy />
<Footer />

<style lang="scss">
  @import '../../../widget/src/styles/animations.scss';
  $textColor: #333;
  $easeOutExpo: cubic-bezier(0.16, 1, 0.3, 1);
  $easeOutBack: cubic-bezier(0.34, 1.25, 0.64, 1);

  main {
    position: relative;
    z-index: 1;
    display: flex;
    max-width: 960px;
    height: 100vh;
    margin: 0 auto;
    transform: translateY(30%);
    .col {
      position: relative;
      max-width: 50%;
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
      top: -15%;
      right: -20%;
      opacity: 0;
    }
    :global(iframe.loaded) {
      opacity: 1;
      box-shadow: 0 0 25px rgba(0, 0, 0, 0.4);
      transition: opacity 0.3s $easeOutExpo, box-shadow 0.5s $easeOutExpo,
        height 0.3s $easeOutBack;
      will-change: box-shadow, height;
    }
  }

  video {
    position: absolute;
    z-index: 0;
    height: 100%;
    width: 100%;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    object-fit: fill;
    opacity: 0.05;
  }
  svg {
    position: absolute;
    bottom: 0;
    right: 0;
    left: 0;
    width: 100%;
  }
  section {
    position: relative;
    z-index: 1;
    display: flex;
    flex: 1;
    background: #fff;
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
  @media (max-width: 1000px) {
    :global(body) {
      overflow-y: scroll !important;
      main {
        transform: translateY(10%);
        flex-direction: column;
        > .col {
          max-width: 100%;
        }
        .intro {
          padding-left: 2rem;
          padding-right: 1rem;
          article {
            max-width: none;
            margin-right: 2rem;
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
      svg {
        position: fixed;
      }
    }
  }
</style>
