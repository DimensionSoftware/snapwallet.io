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
  import CartCheckout from '$lib/features/CartCheckout.svelte'
  import Footer from '$lib/Footer.svelte'
  import LiquidContent from '$lib/LiquidContent.svelte'

  let ifr: HTMLIFrameElement
  let liquidVisible = false
  let topBg

  let dy, dyLast, isRotated, lastIsRotated, isScrolled
  $: {
    // init scroll fx
    if (dyLast !== dy) {
      // scrolled, so--
      dyLast = dy
      isRotated = dy > 900
      // trigger scrolled css events
      if (dy > 20) {
        // don't touch the DOM unless we must
        if (!isScrolled) {
          requestAnimationFrame(() => {
            document.body.classList.add('scrolled')
            isScrolled = true
          })
        }
      } else {
        if (isScrolled) {
          requestAnimationFrame(() => {
            document.body.classList.remove('scrolled')
            isScrolled = false
          })
        }
      }
      if (lastIsRotated !== isRotated) {
        // don't touch the DOM unless we must
        requestAnimationFrame(() => {
          // transform background
          //   let sy = ~~(dy * 0.3) * (isRotated ? 1 : -1)
          //   topBg.style = `transform: rotate(${isRotated ? '180deg' : 0}
          // ) translateY(${sy}px)`
          topBg.style = `transform: translateZ(0) rotate(${
            isRotated ? '180deg' : 0
          })`
          lastIsRotated = isRotated
        })
      }
    }
  }

  onMount(async () => {
    await import('flux-init')
    Typewriter = (await import('svelte-typewriter')).default

    const appName = 'Noir Checkout',
      // themeColor = '#E1143D',
      // themeColor = '#F1071C',
      themeColor = '#fffc00',
      SnapWallet = new (window as any).Snap({
        appName,
        apiKey: 'eacaa046-3b2a-4961-a47d-7125b4f09a2b',
        environment: 'sandbox',
        intent: 'buy',
        wallets: [],
        focus: false,
        theme: {
          modalBackground: 'rgba(0,0,0,1)',
          modalBackgroundColor: '#222',
          modalPopupBackground: 'rgba(10,10,10,.7)',
          color: 'rgba(255,252,0,.9)',
          badgeTextColor: '#333',
          colorLightened: 'rgba(255,252,0,.3)',
          shadowBottomColor: 'rgba(0,0,0,.25)',
          colorInverse: '#fff',
          buttonColor: themeColor,
          buttonTextColor: '#000',
          buttonGlowColor: '255, 252, 1',
          textColor: '#fff',
          textColorNoBackground: '#fff',
          inputTextColor: '#333',
        },
      })

    // respond to widget events
    window.addEventListener(
      'message',
      ({ data: msg }) => {
        if (!msg || typeof msg !== 'string') return // guard
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
      // if (window.screen.width > 1000)
      //   setTimeout(() => (liquidVisible = true), 3000)
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

  function scrollFirst() {
    // scroll to top
    requestAnimationFrame(() => {
      document.getElementById('overview').scrollIntoView()
    })
  }
</script>

<svelte:window bind:scrollY={dy} />

<main>
  <div class="intro col">
    <h1 class="blur"><span>Welcome to</span> <b>Snap Wallet</b></h1>
    <h2 class="blur">
      Connect Crypto to Your
      {#if Typewriter}
        <Typewriter interval={50} delay={0} loop={1800}>
          <span>Idea</span>
          <span>NFT</span>
          <span>App</span>
          <span>Exchange</span>
          <span>Company</span>
          <span>Donations</span>
        </Typewriter>
      {:else}
        Idea
      {/if}
    </h2>
    <article class="blur">
      Snap Wallet is the best crypto wallet that you don't need to download.
      It's a secure, easy-to-use, zero-dependency, embedable wallet that
      connects crypto currency to everything in a snap!
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
<span class="btc coin" />
<span class="eth coin" />
<span class="dog coin" />
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 320"
  ><path
    fill="#ffffff"
    fill-opacity="1"
    d="M0,320L240,288L480,192L720,160L960,96L1200,192L1440,0L1440,320L1200,320L960,320L720,320L480,320L240,320L0,320Z"
  /></svg
>

<Overview />
<CartCheckout />
<Donation />
<NFT />
<Footer />

<span
  on:click={scrollFirst}
  class="gg-chevron-double-down"
  title="Experience Snap Wallet!"
/>
<span on:click={scrollFirst} class="bottom" title="Experience Snap Wallet!" />
<span id="top-bg" bind:this={topBg} class="top-bg" />

{#if liquidVisible}
  <span
    id="liquid"
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
    position: relative;
    backdrop-filter: blur(2px) brightness(75%) grayscale(25%) !important;
    -webkit-backdrop-filter: blur(2px) brightness(75%) grayscale(25%) !important;
    z-index: 9999999;
  }
  :global(.blur) {
    // backdrop-filter: blur(4px) !important;
    // -webkit-backdrop-filter: blur(4px) !important;
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
      border-radius: 20px !important;
      border-top-left-radius: 20px !important;
      top: -5%;
      right: -20%;
      opacity: 0;
    }
    :global(iframe.loaded) {
      opacity: 1;
      transition: height 0.3s $easeOutBack, width 0.4s $easeOutBack 0.301s;
      will-change: opacity, height, width;
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
  // responsive
  @media (max-width: 375px) {
    :global(body),
    :global(html) {
      overflow-y: scroll !important;
      overflow-x: hidden !important;
    }
  }
  @media (max-width: 480px) {
    :global(.intro) {
      padding-left: 0 !important;
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

  @media (max-width: 550px) {
    h1 > span {
      display: none;
    }
  }
  @media (max-width: 1000px) {
    :global(body:before) {
      display: none;
    }
    :global(body),
    :global(html) {
      background-color: #fffc00 !important;
      overflow-y: scroll !important;
      overflow-x: hidden !important;
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
            font-size: 1rem;
            max-width: 75%;
            margin: 0 2rem 6rem 0;
          }
        }
        .wallet {
          position: relative;
          max-width: inherit;
          padding-bottom: 15rem;
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
  @media (max-width: 1450px) {
    :global(section.overview > article) {
      :global(h2) {
        left: 5.5rem !important;
      }
      :global(div + div > h3) {
        left: 0% !important;
        margin-left: 0.5rem !important;
      }
    }
    :global(section.nft > article),
    :global(section.donations > article) {
      :global(div + div > h2),
      :global(div + div > h3) {
        left: 0% !important;
        margin-left: 0.5rem !important;
      }
    }
    :global(.intro) {
      font-size: 0.8rem;
      margin-left: 5rem;
    }
    :global(section.cartcheckout > article) {
      left: 0 !important;
      :global(div + div > h2),
      :global(div + div > h3) {
        left: 0% !important;
        margin-left: 0.5rem !important;
      }
    }
  }
</style>
