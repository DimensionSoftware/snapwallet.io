<script lang="ts">
  import { onMount } from 'svelte'
  import Feature from '../Feature.svelte'
  import { formatLocaleCurrency } from '../../../../widget/src/util'
  import Button from '../Button.svelte'
  import Heart from '../Heart.svelte'

  const config = {
    appName: 'NFT Checkout',
    focus: true,
    product: {
      // imageURL:
      //   'https://lh3.googleusercontent.com/NpXUf_nwxn9yhHk_1AwFxRE7Mg2Lb7_rZoxKRuhf5Tca9MKm0Fh1MXuUAlJNJooO34l6YX3d-2MEZ1kpZvQ18JtrQbQw8CHnBdnRUV8=s992',
      videoURL:
        'https://mkpcdn.com/videos/d3a277f4e6f1212c900a1da4ec915aa9_675573.mp4',
      destinationAmount: 0.04,
      destinationTicker: 'ETH',
      destinationAddress: '0xf636B6aA45C554139763Ad926407C02719bc22f7',
      title: 'The Crown',
      author: 'Patrick Mahomes',
    },
    theme: {
      color: 'rgb(222, 49, 45)',
      buttonColor: 'rgb(241, 7, 28)',
    },
    wallets: [{ asset: 'btc', address: 'ms6k9Mdsbq5ZkoXakJexxjGjpH2PbSQdWK' }],
  }
  let snap: any = {}

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
    const canvas = document.getElementById('nft-qr-canvas')
    snap.createQR({
      text,
      radius: 0.5, // 0.0 to 0.5
      ecLevel: 'H', // L, M, Q, H
      fill: '#F1071C',
      background: 'transparent', // transparent default
      size: 128, // in pixels
      element: canvas,
      pixelSize: 100,
    })
  })
</script>

<Feature
  name="nft"
  title="NFT Checkout"
  description="Snap Wallet supports your favorite platform.  Check out with ANY NFT in one step.  Your loot is automagically added to your collection."
  docLink="https://snapwallet.io/docs/guide/use-cases/nft.html"
  hasImage={true}
  hasBackground={true}
  icon="/images/coin 1.png"
>
  <div class="relative" slot="left">
    <div class="qr" on:mousedown={snap.openWeb}>
      <canvas id="nft-qr-canvas" />
    </div>
    <video loop playsinline autoplay muted on:mousedown={snap.openWeb}
      ><source
        src="https://mkpcdn.com/videos/d3a277f4e6f1212c900a1da4ec915aa9_675573.mp4"
        class="svelte-1pit40i"
      /></video
    >
  </div>
  <div class="relative" slot="right">
    <div on:mousedown={snap.openWeb}>
      <h3>{config.product.title}</h3>
      <small>by {config.product.author}</small>
      <h3>Price</h3>
      <small>
        {typeof navigator !== 'undefined'
          ? formatLocaleCurrency(
              config.product.destinationTicker,
              config.product.destinationAmount,
            )
          : '0'}
      </small>
      <Heart />
      <Button>Buy</Button>
    </div>
  </div>
</Feature>

<style lang="scss">
  @import '../../../../widget/src/styles/animations.scss';

  .relative {
    padding: 0 0 0 1rem;
    position: relative;
    height: 100%;
    width: 50%;
    > div {
      cursor: pointer;
    }
  }
  .docs-link {
    margin: 1rem auto 0;
  }
  :global(.nft h2) {
    margin-top: -4rem;
  }
  :global(.nft h3) {
    position: relative;
    overflow: hidden;
    display: inline-block;
    margin-top: -4rem;
  }
  h3 {
    margin: 0.5rem 0 0.2rem 0;
    font-size: 1.5rem;
  }
  small {
    margin: 0 0 1.5rem 0.1rem;
    font-size: 1rem;
    font-weight: 200;
    display: block;
  }
  h4 {
    margin: 0.25rem 0 2rem 0;
  }
  :global(.nft article) {
    width: 800px;
    margin-top: 4.55rem !important;
    background: linear-gradient(#fff, rgba(255, 255, 255, 0.9)) !important;
    box-shadow: 5px 3px 10px 0 rgba(0, 0, 0, 0.2);
    padding: 1.5rem;
  }
  p {
    font-size: 1.25rem;
    line-height: 1.75rem;
    padding: 0;
    margin: -0.25rem 0 1.5rem;
    align-self: flex-start;
  }
  video {
    cursor: pointer;
    align-self: flex-start;
    min-height: 400px;
    min-width: 400px;
    max-width: 50%;
  }
  :global(.nft button) {
    text-transform: capitalize !important;
    color: #fff !important;
    font-size: 1.15rem !important;
    background: rgb(241, 7, 28) !important;
    border-color: rgb(241, 7, 28) !important;
    top: -3.5rem;
    width: 100%;
    border-radius: 0.25rem;
  }
  :global(.nft button:hover) {
    box-shadow: 0 0 0 3px rgba(222, 49, 45, 0.25) !important;
  }
  .qr {
    cursor: pointer;
    position: absolute;
    top: -30px;
    left: -20px;
    background: rgba(0, 0, 0, 1);
    border-radius: 100%;
    padding: 1.25rem;
    overflow: hidden;
    #nft-qr-canvas {
      padding: 1px;
      height: 75px;
      width: 75px;
    }
  }
</style>
