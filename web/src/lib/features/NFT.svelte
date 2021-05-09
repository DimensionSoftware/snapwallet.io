<script lang="ts">
  import { onMount } from 'svelte'
  import Feature from '../Feature.svelte'
  import { formatLocaleCurrency } from '../../../../widget/src/util'
  import Button from '../Button.svelte'

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
    snap.createQR({ element: canvas, pixelSize: 100 })
  })
</script>

<Feature
  name="nft"
  title="NFT Checkout"
  docLink="https://snapwallet.io/docs/guide/use-cases/checkout.html"
  hasBackground={true}
>
  <video loop playsinline autoplay muted slot="left" on:mousedown={snap.openWeb}
    ><source
      src="https://mkpcdn.com/videos/d3a277f4e6f1212c900a1da4ec915aa9_675573.mp4"
      class="svelte-1pit40i"
    /></video
  >
  <div class="relative" slot="right">
    <p>
      Checkout with any NFT. Your loot is automagically added to your
      Collection.
    </p>
    <div on:mousedown={snap.openWeb}>
      <h3>{config.product.title}</h3>
      <small>by {config.product.author}</small>
      <h4>
        {typeof navigator !== 'undefined'
          ? formatLocaleCurrency(
              config.product.destinationTicker,
              config.product.destinationAmount,
            )
          : '0'}
      </h4>
      <Button>Buy</Button>
    </div>
    <div class="qr" on:mousedown={snap.openWeb}>
      <canvas id="nft-qr-canvas" />
    </div>
  </div>
</Feature>

<style lang="scss">
  @import '../../../../widget/src/styles/animations.scss';

  .relative {
    padding: 0 0 0 1rem;
    position: relative;
    height: 100%;
    > div {
      cursor: pointer;
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
    margin-right: 2rem;
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
    bottom: -0.5rem;
    left: 1rem;
    #nft-qr-canvas {
      height: 100px;
      width: 100px;
      border: 0.5rem solid #fff;
    }
  }
</style>
