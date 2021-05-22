<script lang="ts">
  import { onMount } from 'svelte'
  import Feature from '../Feature.svelte'

  let ifr: HTMLIFrameElement,
    snap: any = {}

  onMount(async () => {
    await import('flux-init')

    const appName = 'Buy Checkout',
      themeColor = '#fffc00',
      Wallet = new (window as any).Snap({
        appName,
        intent: 'buy',
        wallets: [],
        focus: false,
        theme: {
          modalBackground: 'rgba(40,40,40,.9)',
          modalPopupBackground: 'rgba(50,50,50,.9)',
          color: 'rgba(0,0,0,.9)',
          badgeTextColor: '#333',
          colorLightened: 'rgba(5,5,5,.8)',
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
          if (event === Wallet.events.RESIZE && data && ifr) {
            if (appName === data.appName) ifr.height = data.height
          }
        } catch (e) {
          console.warn('Unable to parse message', msg, e)
        }
      },
      false,
    )
    ifr.src = Wallet.generateURL()
  })
</script>

<Feature
  name="fiat"
  title="Fiat Onramp & Offramp"
  description="Buy and sell crypto currency fast, with the best rates & lowest fees.  Pay in fiat, receive crypto currency or vice-versa."
  docLink="https://snapwallet.io/docs/guide/use-cases/onramp.html"
  hasImage={true}
  hasBackground={true}
  icon="/images/coin 3.png"
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
    <br style="margin-top: .5rem;" />
    <div on:mousedown={snap.openWeb}>
      <br />
      <h3>Buy & Sell Crypto</h3>
      <ul>
        <li>Fiat &rlarr; Crypto</li>
        <li>Crypto &rlarr; Fiat</li>
        <li>Crypto &rlarr; Crypto</li>
      </ul>
      <br />
      <h3>Snap Wallet is</h3>
      <ul>
        <li>&nbsp; Web, iOS & Android</li>
        <li>&nbsp; Fast & Secure</li>
        <li>&nbsp; Dependency-free & Embeddable</li>
      </ul>
    </div>
  </div>
</Feature>

<style lang="scss">
  @import '../../../../widget/src/styles/animations.scss';

  .relative {
    position: relative;
    padding: 0 4rem 0 4rem;
    height: 100%;
    > div {
      cursor: pointer;
    }
  }
  :global(.fiat article) {
    margin-top: 4.55rem !important;
    background: linear-gradient(#fff, rgba(255, 255, 255, 0.9)) !important;
    box-shadow: 5px 3px 10px 0 rgba(0, 0, 0, 0.2);
  }
  :global(.fiat h2) {
    margin-top: -4rem;
  }
  :global(.fiat h3) {
    margin-top: -4rem;
    font-size: 1.5rem;
  }
  ul {
    margin: 0;
    padding: 0;
    li {
      padding: 0;
      margin: 0.5rem 0 0.75rem 1rem;
    }
  }
  .docs-link {
    margin: 1rem auto 0;
  }
  h3 {
    margin: 0.5rem 0 1rem 0;
  }
  small {
    margin: 0 0 0 0.1rem;
  }
  h4 {
    margin: 1.5rem 0 2rem 0;
  }
  p {
    font-size: 1.25rem;
    line-height: 1.75rem;
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
    bottom: -0.25rem;
    left: 0;
    #qr-canvas {
      height: 100px;
      width: 100px;
      border: 0.5rem solid #fff;
    }
  }
  $easeOutExpo: cubic-bezier(0.16, 1, 0.3, 1);
  $easeOutBack: cubic-bezier(0.34, 1.25, 0.64, 1);
  iframe {
    border-radius: 2rem;
    overflow: hidden;
    box-shadow: 0 0 25px rgba(0, 0, 0, 0.4);
    transition: box-shadow 0.8s $easeOutExpo 0.1s, height 0.3s $easeOutBack;
    will-change: box-shadow, height;
  }
</style>
