<script lang="ts">
  import { onMount } from 'svelte'
  import Feature from '../Feature.svelte'

  let ifr: HTMLIFrameElement,
    snap: any = {}

  onMount(async () => {
    await import('flux-init')

    const appName = 'Buy Checkout',
      Wallet = new (window as any).Snap({
        appName,
        intent: 'buy',
        wallets: [],
        focus: false,
        theme: {
          // Coinbase theme
          modalBackground: '#070F15',
          modalPopupBackground: 'rgba(2,13,19,.95)',
          colorLightened: 'rgba(102,113,119,.8)',
          color: '#2187FF',
          textColor: '#fff',
          colorInverse: '#fff',
          buttonColor: '#2187FF',
          warningColor: '#FFBD4A',
          successColor: '#83E068',
          errorColor: '#E7693C',
          shadowBottomColor: 'rgba(0,0,0,.35)',
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
  docLink="https://snapwallet.io/docs/guide/use-cases/checkout.html"
  title="Fiat On & Offramp"
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
    <p>Buy and Sell Crypto Currency fast, with the best rates & lowest fees.</p>
    <div on:mousedown={snap.openWeb}>
      <br />
      <h3>Snap Wallet is</h3>
      <ul>
        <li>&nbsp; Web, iOS & Android</li>
        <li>&nbsp; Fast, Secure, Embeddable</li>
        <li>&nbsp; Pay in fiat, receive crypto currency</li>
      </ul>
    </div>
  </div>
</Feature>

<style lang="scss">
  @import '../../../../widget/src/styles/animations.scss';

  .relative {
    position: relative;
    padding: 0 0 0 5rem;
    height: 100%;
    > div {
      cursor: pointer;
    }
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
