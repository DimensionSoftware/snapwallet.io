<script lang="ts">
  import { onMount } from 'svelte'
  import Feature from '../Feature.svelte'
  import Button from '../Button.svelte'

  const config = {
    appName: 'Donation',
    environment: 'sandbox',
    intent: 'donate',
    payee: 'Donate to Snap',
    focus: true,
    wallets: [{ asset: 'btc', address: 'ms6k9Mdsbq5ZkoXakJexxjGjpH2PbSQdWK' }],
    theme: {
      color: 'rgb(222, 49, 45)',
      buttonColor: 'rgb(241, 7, 28)',
    },
  }

  let snap: any = {}

  const donateUSDAmount = sourceAmount => {
    return () => snap.openWeb({ sourceAmount })
  }

  const donateAsset = defaultDestinationAsset => {
    return () => snap.openWeb({ defaultDestinationAsset })
  }

  onMount(async () => {
    await import('flux-init')
    snap = new (window as any).Snap({
      ...config,
      onMessage: msg => {
        switch (msg.event) {
          case snap.events.EXIT:
            snap.closeWeb()
            break
          case snap.events.SUCCESS:
            console.log('Success!', msg)
          case snap.events.RESIZE:
          default:
            // resize iframe/viewport happened
            break
        }
      },
    })

    // Open using a QR code
    const canvas = document.getElementById('donation-qr-canvas')
    snap.createQR({
      foregroundColor: '#333',
      backgroundColor: null, // transparent default
      element: canvas,
      pixelSize: 100,
    })
  })
</script>

<Feature
  name="donations"
  title="Donations & Tips"
  description="Accept crypto donations and tips to your business.  Integration is easy and fully customizable."
  hasBackground={true}
  docLink="https://snapwallet.io/docs/guide/use-cases/donations.html"
  icon="/images/coin2.png"
>
  <div class="relative" slot="left">
    <div>
      <h3>Donate</h3>
      <small>to Snap Wallet</small>
      <aside>
        <Button id="btc" on:mousedown={donateAsset('btc')}>Send BTC</Button>
        <Button id="eth" on:mousedown={donateAsset('eth')}>Send ETH</Button>
      </aside>
      <input placeholder="Your Name" type="text" />
      <textarea maxlength="250" placeholder="Message" />
    </div>
  </div>
  <div class="relative" slot="right">
    <div>
      <h3>Tip</h3>
      <small>to Snap Wallet</small>
      <aside class="small">
        <Button id="1" on:mousedown={donateUSDAmount(1.0)}>$1</Button>
        <Button id="5" on:mousedown={donateUSDAmount(5.0)}>$5</Button>
        <Button id="10" on:mousedown={donateUSDAmount(10.0)}>$10</Button>
      </aside>
    </div>
    <div class="qr" on:mousedown={snap.openWeb}>
      <canvas id="donation-qr-canvas" />
    </div>
  </div>
</Feature>

<style lang="scss">
  @import '../../../../widget/src/styles/animations.scss';

  .relative {
    color: #222;
    position: relative;
    height: 100%;
    &:first-child {
      margin-right: 3rem;
    }
    > div {
      cursor: pointer;
    }
  }
  :global(.donations h2) {
    margin-top: -4rem;
  }
  :global(.donations h3) {
    margin-top: -4rem;
    font-size: 1.5rem;
  }
  :global(.donations > article) {
    border: 1px solid #eee;
    position: relative;
    width: 800px;
    padding: 2rem 3rem 3rem 3rem !important;
    margin-top: 4.55rem !important;
    background: transparent !important;
    box-shadow: 0 0 0 1px transparent, 5px 7px 10px rgba(0, 0, 0, 0.15);
    transition: box-shadow 0.3s ease-out, transform 0.04s ease-in !important;
  }
  :global(.donations > article:hover) {
    box-shadow: 0 0 0 1px rgba(#ff6b00, 0.25), 4px 6px 9px rgba(0, 0, 0, 0.15) !important;
    transition: none !important;
  }
  .docs-link {
    margin: 1rem auto 0;
  }
  h3 {
    margin: 0.5rem 0 0 0;
  }
  small {
    font-weight: 300;
    margin: 0 0 0 0.1rem;
  }
  aside {
    font-size: 0.9rem;
    display: flex;
    margin-top: 2rem;
    &.small :global(button) {
      font-size: 1rem;
      padding: 0.65rem 2rem;
      margin-top: 0.25rem;
    }
  }
  h4 {
    margin: 1.5rem 0 2rem 0;
  }
  p {
    white-space: nowrap;
    font-size: 1.25rem;
    line-height: 1.75rem;
    padding: 0;
    margin: -0.25rem 0 1.5rem;
  }
  :global(button) {
    font-size: 1rem;
    margin-right: 0.75rem !important;
    border: none !important;
    box-shadow: none !important;
    transition: box-shadow 0.3s ease-out !important;
  }
  :global(button:hover) {
    transition: none !important;
  }
  input.name,
  input,
  textarea {
    outline: none;
    font-family: inherit;
    font-size: 1rem;
    min-width: 100px;
    min-height: 75px;
    max-width: 225px;
    max-height: 225px;
    padding: 0.75rem 1.5rem 0.75rem 1rem;
    border-radius: 0.85rem;
    border: 1px solid rgba(0, 0, 0, 0.025);
    margin: 1.5rem 0 1rem;
    transition: box-shadow 0.1s ease-in, border 0.2s ease-in;
    box-shadow: 1px 2px 3px rgba(0, 0, 0, 0.15);
    &:hover {
      box-shadow: 1px 3px 5px rgba(0, 0, 0, 0.18);
      transition: none;
    }
    &:focus,
    &:active {
      transition: none;
      box-shadow: 0 1px 1px 0 #555, 0 0 0 2px #ff6b00,
        1px 3px 5px rgba(0, 0, 0, 0.18);
      animation: focus 0.16s;
    }
  }
  input[type='text'] {
    min-height: 1rem;
    margin-top: 2rem;
    margin-bottom: 0;
  }
  textarea {
    margin-top: 1rem;
    width: 75%;
  }
  .qr {
    cursor: pointer;
    position: absolute;
    bottom: 0;
    right: -55px;
    display: block;
    border-radius: 100%;
    padding: 20px;
    overflow: hidden;
    #donation-qr-canvas {
      height: 75px;
      width: 75px;
    }
  }

  :global(.donations aside > button) {
    color: #fff !important;
    font-size: 1.15rem !important;
    background: rgb(241, 7, 28) !important;
    // background: linear-gradient(63.26deg, #ff6b00, #f1071c 75%) !important;
    background: linear-gradient(
      63.26deg,
      #ff6b00,
      rgb(241, 7, 28) 75%
    ) !important;
    border-color: rgb(222, 49, 45) !important;
    z-index: 4;
    text-align: center;
    border-radius: 4.25rem;
    transition: box-shadow 0.3s ease-out;
  }
  :global(.donations aside > button:hover) {
    box-shadow: 0 0 0 1px red, 0 0 0 2px #ff6b00 !important;
    background: linear-gradient(
      63.26deg,
      #ff6b00,
      rgb(241, 7, 28) 75%
    ) !important;
  }
</style>
