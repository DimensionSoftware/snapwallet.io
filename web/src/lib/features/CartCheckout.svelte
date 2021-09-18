<script lang="ts">
  import { onMount } from 'svelte'
  import Feature from '../Feature.svelte'

  let ifr: HTMLIFrameElement,
    snap: any = {}

  onMount(async () => {
    await import('flux-init')

    const appName = 'Cart Checkout',
      themeColor = '#fffc00',
      snap = new (window as any).Snap({
        appName,
        environment: 'sandbox',
        intent: 'cart',
        wallets: [],
        focus: false,
        products: [
          {
            title: 'Headband OG',
            subtitle:
              'Headband OG is a combination of OG Kush, Master Kush, and Sour Diesel',
            img: 'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.UVKFWQKmlDNfkSpzeIHf-gHaFj%26pid%3DApi&f=1',
            author: 'phpchris',
            destinationAmount: '.00000420',
            destinationTicker: 'ETH',
            destinationAddress: '0xCAFEBABE',
          },
          {
            title: 'Fire OG',
            subtitle:
              'With its high THC content and strong cerebral effects, Fire OG is generally recommended for more experienced users',
            img: 'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.UVKFWQKmlDNfkSpzeIHf-gHaFj%26pid%3DApi&f=1',
            author: 'dreamc0dez',

            destinationAmount: '.00000420',
            destinationTicker: 'ETH',
            destinationAddress: '0xBEEFBABE',
          },
          {
            title: 'Blueberry Dream',
            subtitle:
              'Blueberry Dream is a Sativa-dominant bud that powers up the mind with an energizing high.',
            img: 'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.UVKFWQKmlDNfkSpzeIHf-gHaFj%26pid%3DApi&f=1',
            author: 'sMURF0r',
            destinationAmount: '.00000420',
            destinationTicker: 'ETH',
            destinationAddress: '0xDEADBEEF',
          },
          {
            title: 'Snap OG Dream',
            subtitle:
              'The absolute best, full-flavor, dreamy and high in the leather palette...',
            img: 'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.UVKFWQKmlDNfkSpzeIHf-gHaFj%26pid%3DApi&f=1',
            author: 'sMURF0r',
            destinationAmount: '.00000420',
            destinationTicker: 'ETH',
            destinationAddress: '0xDEADBEEF',
          },
        ],
        theme: {
          modalBackground: 'rgba(0,0,0,1)',
          modalPopupBackground: 'rgba(10,10,10,.75)',
          modalBackgroundColor: '#222',
          color: 'rgba(255,252,0,.9)',
          badgeTextColor: '#333',
          colorLightened: 'rgba(255,252,0,.3)',
          shadowBottomColor: 'rgba(0,0,0,.25)',
          colorInverse: '#fff',
          buttonColor: themeColor,
          buttonTextColor: '#000',
          buttonGlowColor: '255, 255, 255',
          successColor: themeColor,
          textColor: '#fff',
          textColorNoBackground: '#fff',
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
          if (event === snap.events.RESIZE && data && ifr) {
            if (appName === data.appName) ifr.height = data.height
          }
        } catch (e) {
          console.warn('Unable to parse message', msg, e)
        }
      },
      false,
    )
    ifr.src = snap.generateURL()
    ifr.classList.add('snapWallet')
    // Open using a QR code
    const canvas = document.getElementById('buy-qr-canvas')
    snap.createQR({
      foregroundColor: '#333',
      backgroundColor: null, // transparent default
      element: canvas,
      pixelSize: 100,
    })
  })
</script>

<Feature
  name="fiat"
  title="Cart Checkout"
  description="Expand Your Checkout with Crypto Currency Checkout"
  docLink="https://snapwallet.io/docs/guide/use-cases/onramp.html"
  hasImage={true}
  hasBackground={true}
  icon="/images/coin1.png"
>
  <div slot="left">
    <iframe
      class="loaded"
      title="Snap Wallet"
      frameborder="0"
      height="608px"
      width="360px"
      bind:this={ifr}
    />
  </div>
  <div class="relative" slot="right">
    <div on:mousedown={snap.openWeb}>
      <br />
      <h3>Dispensary Example</h3>
      <p class="story">
        Give your customers full flexibility to checkout with their crypto
        currency, debit accounts and credit cards.
      </p>
      <h4 style="margin: 0 0 0.5rem;">Why Customers Choose Us</h4>
      <ul style="margin: .75rem 0 0 1rem;">
        <li>Fast & Secure</li>
        <li>Supported by all platforms</li>
        <li>Dependency-free & Embeddable!</li>
      </ul>
      <ul>
        <div style="margin-top: 2rem;">
          <h4 style="margin: 0 0 0.5rem;">Also Supports</h4>
          <img height="14" src="/images/card_mastercard.png" />
          <img height="16" src="/images/card_visa.png" />
          <img height="16" src="/images/card_discover.png" />
        </div>
      </ul>
    </div>
    <br />
    <div class="qr" on:mousedown={snap.openWeb}>
      <canvas id="buy-qr-canvas" />
    </div>
  </div>
</Feature>

<style lang="scss">
  @import '../../../../widget/src/styles/animations.scss';

  .relative {
    position: relative;
    padding: 0 4rem 0 4rem;
    height: 100%;
    height: 550px !important;
  }
  .flex {
    display: flex;
    margin-top: 0.5rem;
    justify-content: space-between;
  }
  :global(.fiat article) {
    width: 800px;
    margin-top: 4.55rem !important;
    background: linear-gradient(#fff, rgb(255, 254, 232)) !important;
    // box-shadow: 5px 3px 10px 0 rgba(0, 0, 0, 0.2);
  }
  :global(.fiat h2) {
    margin-top: -4rem;
  }
  :global(.fiat h3) {
    margin-top: -4rem;
    font-size: 1.5rem;
  }
  :global(.fiat .relative > h3) {
    white-space: nowrap;
  }
  ul {
    list-style: none;
    margin: 0;
    padding: 0;
    li {
      padding: 0;
      line-height: 1.1rem;
      margin: 0.5rem 0 0.25rem 0;
      span {
        display: inline-block;
        width: 37px;
      }
    }
  }
  .docs-link {
    margin: 1rem auto 0;
  }
  h3 {
    margin: 0.5rem 0 0 0;
  }
  small {
    margin: 0 0 0 0.1rem;
  }
  h4 {
    margin: 1.5rem 0 2rem 0;
  }
  .tag {
    position: absolute;
    z-index: 1;
    font-size: 0.8rem;
    background: rgba(#fffc00, 0.7);
    border: 1px solid rgba(#fffc00, 0.2);
    width: inherit;
    padding: 0.1rem 0.75rem;
    margin-left: 0.5rem;
    border-radius: 5rem;
    transform: rotate(2deg);
    transition: box-shadow 0.2s ease-out, border 0.3s ease-out;
    &:hover {
      background: #fffc00;
      border: 1px solid #fffc00;
      box-shadow: 0 0 0 2px #fffc00;
      transition: none;
    }
  }
  .story {
    margin: 0.5rem 0 1.5rem 0;
    font-size: 1rem;
  }
  p {
    font-size: 1.25rem;
    line-height: 1.5rem;
    padding: 0;
    margin: -0.25rem 0 1.5rem;
  }
  video {
    cursor: pointer;
    min-height: 400px;
    min-width: 400px;
    max-width: 50%;
  }
  :global(button) {
    font-size: 1rem;
  }
  .qr {
    cursor: pointer;
    position: absolute;
    bottom: -25px;
    right: -25px;
    display: block;
    padding: 20px;
    overflow: hidden;
    #buy-qr-canvas {
      padding: 2px;
      height: 75px;
      width: 75px;
    }
  }
  $easeOutExpo: cubic-bezier(0.16, 1, 0.3, 1);
  $easeOutBack: cubic-bezier(0.34, 1.25, 0.64, 1);
  iframe {
    border-radius: 2rem;
    overflow: hidden;
    // box-shadow: 0 0 25px rgba(0, 0, 0, 0.4);
    transition: box-shadow 0.8s $easeOutExpo 0.1s, height 0.3s $easeOutBack;
    will-change: box-shadow, height;
  }
</style>
