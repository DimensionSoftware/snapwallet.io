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
            img: 'https://images.dutchie.com/category-stock-photos/flower/flower-1.png?auto=format&ixlib=react-9.0.2&w=4088',
            author: 'phpchris',
            destinationAmount: '.00000420',
            destinationTicker: 'ETH',
            destinationAddress: '0xCAFEBABE',
          },
          {
            title: 'Fire OG',
            subtitle:
              'With its high THC content and strong cerebral effects, Fire OG is generally recommended for more experienced users',
            img: 'https://images.dutchie.com/flower-stock-10-v1.jpg?auto=format&fit=fill&fill=solid&fillColor=%23fff&__typename=ImgixSettings&ixlib=react-9.0.2&h=100&w=100&q=50&dpr=2',
            author: 'dreamc0dez',

            destinationAmount: '.00000420',
            destinationTicker: 'ETH',
            destinationAddress: '0xBEEFBABE',
          },
          {
            title: 'Blueberry Dream',
            subtitle:
              'Blueberry Dream is a Sativa-dominant bud that powers up the mind with an energizing high.',
            img: 'https://images.dutchie.com/flower-stock-10-v1.jpg?auto=format&fit=fill&fill=solid&fillColor=%23fff&__typename=ImgixSettings&ixlib=react-9.0.2&h=100&w=100&q=50&dpr=2',
            author: 'sMURF0r',
            destinationAmount: '.00000420',
            destinationTicker: 'ETH',
            destinationAddress: '0xDEADBEEF',
          },
          {
            title: "Bao's Vegan Chocolate",
            subtitle:
              'Rich dark chocolate, laced with toasted hazelnuts, pecans and dried currants, rounded off with cinnamon & salt.',
            img: 'https://images.dutchie.com/edibles-stock-chocolate-v1.jpg?auto=format&fit=fill&fill=solid&fillColor=%23fff&__typename=ImgixSettings&ixlib=react-9.0.2&h=100&w=100&q=50&dpr=2',
            author: 'bao',
            destinationAmount: '.00000420',
            destinationTicker: 'ETH',
            destinationAddress: '0xCAFEBEEF',
          },
          {
            title: 'Snap OG Dream',
            subtitle:
              'The absolute best, full-flavor, dreamy and high in the leather palette...',
            img: 'https://images.dutchie.com/flower-stock-10-v1.jpg?auto=format&fit=fill&fill=solid&fillColor=%23fff&__typename=ImgixSettings&ixlib=react-9.0.2&h=100&w=100&q=50&dpr=2',
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
        if (!msg || typeof msg !== 'string') return // guard
        try {
          const { event, data } = JSON.parse(msg)
          if (event === snap.events.RESIZE && data && ifr) {
            if (appName === data.appName) {
              ifr.height = data.height
              ifr.width = data.width
            }
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
  name="cartcheckout"
  title="Cart Checkout"
  description="Expand Your Checkout with Crypto Currency Checkout"
  docLink="https://snapwallet.io/docs/guide/use-cases/onramp.html"
  hasImage={true}
  hasBackground={true}
  icon="/images/coin1.png"
>
  <div slot="left" style="width: 600px;">
    <iframe
      class="loaded"
      title="Snap Wallet"
      frameborder="0"
      height="608px"
      width="360px"
      bind:this={ifr}
    />
  </div>
  <div class="example relative" slot="right">
    <div on:mousedown={snap.openWeb}>
      <br />
      <h3>Example<small>:</small> Dispensary</h3>
      <p class="story">
        Give your customers full flexibility to checkout with their crypto
        currency, debit accounts and credit cards.
      </p>
      <h4 style="margin: 0 0 0.5rem;">Why Customers Choose SW</h4>
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
          <img
            height="16"
            src="data:image/svg+xml;base64,PHN2ZyBoZWlnaHQ9IjQ4IiB3aWR0aD0iMTI2IiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxtYXNrIGlkPSJhIiBmaWxsPSIjZmZmIj48cGF0aCBkPSJNMCA0Ny40NzNoMTI2VjBIMHoiIGZpbGwtcnVsZT0iZXZlbm9kZCIvPjwvbWFzaz48ZyBmaWxsPSIjMTExIiBmaWxsLXJ1bGU9ImV2ZW5vZGQiPjxwYXRoIGQ9Ik02Ni4yNDggMTYuMjY4Yy0xLjA1Ny0uODg5LTIuODYxLTEuMzMzLTUuNDEzLTEuMzMzaC01Ljc1NnYxNy43ODhoNC4zMDR2LTUuNTc1aDEuOTI4YzIuMzQgMCA0LjA1Ni0uNTE1IDUuMTQ4LTEuNTQ2IDEuMjMtMS4xNTUgMS44NDktMi42OTMgMS44NDktNC42MTMgMC0xLjk5MS0uNjg3LTMuNTY1LTIuMDYtNC43MjFtLTUuMDQ0IDYuODU1aC0xLjgyMVYxOC45NmgxLjYzNmMxLjk5IDAgMi45ODUuNjk4IDIuOTg1IDIuMDk0IDAgMS4zNzgtLjkzNCAyLjA2OC0yLjggMi4wNjhtMTQuNDY5LTguMTg4aC00LjQ4OHYxNy43ODhoOS42OXYtNC4wMjZoLTUuMjAyem0xMy45OTUgMC03LjA1IDE3Ljc4OGg0LjgzMmwuOTI0LTIuNTg2SDk0LjVsLjg0NSAyLjU4Nmg0Ljg4NmwtNy0xNy43ODh6bS0uMDUzIDExLjYwMSAxLjg0OS02LjA4IDEuODIgNi4wOHoiLz48cGF0aCBkPSJNMTAyLjQ3MyAzMi43MjJoNC40ODlWMTQuOTM0aC00LjQ4OXptMjEuOTE3LTE0LjQ1NGE3LjM3NiA3LjM3NiAwIDAgMC0yLjE0LTIuMDUzYy0xLjM1NS0uODU0LTMuMjA0LTEuMjgtNS41NDUtMS4yOGgtNS45MTR2MTcuNzg3aDYuOTE4YzIuNSAwIDQuNTA2LS44MTcgNi4wMi0yLjQ1MyAxLjUxNC0xLjYzNSAyLjI3LTMuODA1IDIuMjctNi41MDggMC0yLjE1LS41MzctMy45ODEtMS42MS01LjQ5M20tNy4xODIgMTAuNDI3aC0xLjkyN3YtOS43MzRoMS45NTRjMS4zNzMgMCAyLjQyOC40MyAzLjE2OCAxLjI4Ny43NC44NTcgMS4xMSAyLjA3MyAxLjExIDMuNjQ3IDAgMy4yLTEuNDM1IDQuOC00LjMwNSA0LjhNMTguNjM3IDAgNC4wOSAzLjgxLjA4MSAxOC40MzlsNS4wMTQgNS4xNDhMMCAyOC42NWwzLjc3MyAxNC42OTMgMTQuNDg0IDQuMDQ3IDUuMDk2LTUuMDY0IDUuMDE0IDUuMTQ3IDE0LjU0Ny0zLjgxIDQuMDA4LTE0LjYzLTUuMDEzLTUuMTQ2IDUuMDk1LTUuMDYzTDQzLjIzMSA0LjEzIDI4Ljc0NS4wODNsLTUuMDk0IDUuMDYzek05LjcxIDYuNjI0bDcuNjYzLTIuMDA4IDMuMzUxIDMuNDQtNC44ODcgNC44NTZ6bTE2LjgyMiAxLjQ3OCAzLjQwNS0zLjM4MyA3LjYzIDIuMTMyLTYuMjI3IDYuMTg3em0tMjEuODYgOS4xMzYgMi4xMTEtNy43MDUgNi4xMjUgNi4yODgtNC44ODYgNC44NTZ6bTI5LjU0Ny0xLjI0MyA2LjIyNy02LjE4OSAxLjk4NiA3Ljc0LTMuNDA0IDMuMzg0em0tMTUuNTAyLS4xMjcgNC44ODctNC44NTYgNC44MDcgNC45MzYtNC44ODYgNC44NTZ6bS03LjgxNCA3Ljc2NSA0Ljg4Ni00Ljg1NiA0LjgxIDQuOTM2LTQuODg4IDQuODU2em0xNS41MDMuMTI3IDQuODg2LTQuODU2TDM2LjEgMjMuODRsLTQuODg3IDQuODU2ek00LjU3IDI5LjkyN2wzLjQwNi0zLjM4NSA0LjgwNyA0LjkzNy02LjIyNSA2LjE4NnptMTQuMDIxIDEuNTk4IDQuODg3LTQuODU2IDQuODA4IDQuOTM2LTQuODg2IDQuODU2em0xNS41MDIuMTI4IDQuODg3LTQuODU2IDMuMzUxIDMuNDM5LTIuMTEgNy43MDV6bS0yNC42NTYgOC45NyA2LjIyNi02LjE4OSA0LjgxIDQuOTM2LTMuNDA2IDMuMzg1em0xNi44NDMtMS4yMDYgNC44ODYtNC44NTYgNi4xMjYgNi4yODktNy42NjIgMi4wMDd6IiBtYXNrPSJ1cmwoI2EpIi8+PC9nPjwvc3ZnPg=="
          />
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
  :global(.cartcheckout article) {
    width: 1000px;
    height: 700px;
    margin-top: 4.55rem !important;
    background: linear-gradient(#fff, rgb(255, 254, 232)) !important;
    // box-shadow: 5px 3px 10px 0 rgba(0, 0, 0, 0.2);
  }
  :global(.cartcheckout h2) {
    margin-top: -4rem;
  }
  :global(.cartcheckout h3) {
    margin-top: -4rem;
    font-size: 1.5rem;
  }
  :global(.cartcheckout .relative > h3) {
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
  .example {
    small {
      opacity: 0.8;
      margin: 0 0 0 0.1rem;
    }
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
    bottom: -125px;
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
    display: flex;
    align-self: center;
    margin: 0 auto;
    border-radius: 2rem;
    overflow: hidden;
    transition: box-shadow 0.8s $easeOutExpo 0.1s, height 0.3s $easeOutBack,
      width 0.4s $easeOutBack 0.301s;
    will-change: box-shadow, height, width;
  }
</style>
