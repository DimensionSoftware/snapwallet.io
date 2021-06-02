<script context="module" lang="ts">
  export const prerender = true
</script>

<script lang="ts">
  let Typewriter: any
  import { onMount } from 'svelte'
  import { fly } from 'svelte/transition'
  import Overview from '$lib/features/Overview.svelte'
  import NFT from '$lib/features/NFT.svelte'
  import Donation from '$lib/features/Donation.svelte'
  import Buy from '$lib/features/Buy.svelte'
  import Footer from '$lib/Footer.svelte'
  import LiquidContent from '$lib/LiquidContent.svelte'

  let ifr: HTMLIFrameElement
  let liquidVisible = false
  let sy

  $: isRotated = sy > 900

  onMount(async () => {
    await import('flux-init')
    Typewriter = (await import('svelte-typewriter')).default

    const appName = 'Noir Checkout',
      // themeColor = '#E1143D',
      // themeColor = '#F1071C',
      themeColor = '#fffc00',
      SnapWallet = new (window as any).Snap({
        appName,
        environment: 'sandbox',
        intent: 'buy',
        wallets: [],
        focus: false,
        theme: {
          modalBackground: 'rgba(0,0,0,.8)',
          modalPopupBackground: 'rgba(10,10,10,.85)',
          color: 'rgba(0,0,0,.9)',
          badgeTextColor: '#333',
          colorLightened: 'rgba(255,252,0,.3)',
          shadowBottomColor: 'rgba(0,0,0,.25)',
          colorInverse: '#fff',
          buttonColor: themeColor,
          buttonTextColor: '#000',
          buttonGlowColor: '255, 255, 255',
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
      liquidVisible = true
      ifr.classList.add('loaded')
    }
    ifr.src = SnapWallet.generateURL()
    ifr.classList.add('snapWallet')

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
    <h1 class="blur">Welcome to <b>Snap Wallet</b></h1>
    <h2 class="blur">
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
    <article class="blur">
      The fastest wallet connecting crypto currency to everything— in a snap!
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

<Overview />
<NFT />
<Donation />
<Buy />
<Footer />

<svelte:window bind:scrollY={sy} />

<span
  class="top-bg"
  style={`transform: rotate(${isRotated ? '180deg' : 0}) translate(0 ,${
    sy * (isRotated ? -0.3 : 0.3)
  }px)`}
/>

{#if liquidVisible}
  <span
    in:fly={{ duration: 1000 }}
    out:fly={{ duration: 5000, x: -100, y: 25, opacity: 0.1 }}
  >
    <LiquidContent />
  </span>
{/if}

<style lang="scss">
  @import '../../../widget/src/styles/animations.scss';
  $textColor: #333;
  $easeOutExpo: cubic-bezier(0.16, 1, 0.3, 1);
  $easeOutBack: cubic-bezier(0.34, 1.25, 0.64, 1);

  :global(iframe.snapWallet) {
    backdrop-filter: blur(2px) brightness(75%) grayscale(25%) !important;
    -webkit-backdrop-filter: blur(2px) brightness(75%) grayscale(25%) !important;
  }
  :global(.blur) {
    backdrop-filter: blur(8px) !important;
    -webkit-backdrop-filter: blur(8px) !important;
  }
  main {
    position: relative;
    z-index: 1;
    display: flex;
    max-width: 960px;
    height: 100vh;
    margin: 0 auto;
    transform: translateY(20%);
    .col {
      position: relative;
      max-width: 50%;
    }
    h1,
    h2 {
      white-space: nowrap;
      font-size: 2.5rem;
      line-height: 1.1;
      margin: 0.5rem 0;
    }
    h1 {
      font-weight: 300;
      font-size: 2.5rem;
      margin-left: -6rem;
      margin-bottom: 1.15rem;
      b {
        font-weight: bold;
      }
    }
    h2 {
      margin: 0 0 1.25rem;
      margin-left: -6rem;
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
      font-size: 1.25rem;
      margin: 0;
      margin-left: -6rem;
      max-width: 400px;
      line-height: 1.35;
    }
    :global(iframe) {
      position: relative;
      border-radius: 20px;
      top: -5%;
      right: -20%;
      opacity: 0;
    }
    :global(iframe.loaded) {
      opacity: 1;
      box-shadow: 5px 5px 18px 5px rgba(0, 0, 0, 0.4);
      transition: opacity 1s $easeOutExpo, box-shadow 0.3s $easeOutExpo,
        height 0.3s $easeOutBack;
      will-change: opacity, box-shadow, height;
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
    // transform: translateY(0);
    // bottom: -500px;
    // bottom: 0;
    // margin-top: -210px;
    bottom: -5px;
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
        h1 > b {
          display: block;
        }
      }
    }
  }
  @media (max-width: 1000px) {
    :global(body:before) {
      display: none;
    }
    :global(body) {
      background-color: #fffc00 !important;
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
          h1 {
            margin-left: 0;
            font-size: 1.7rem;
          }
          h2 {
            margin-left: 0;
            font-size: 1rem;
          }
          article {
            max-width: 75%;
            margin: 0 2rem 6rem 0;
          }
        }
        .wallet {
          position: relative;
          max-width: inherit;
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
