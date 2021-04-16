<script lang="ts">
  import { onMount } from 'svelte'
  import Button from '../../../widget/src/components/Button.svelte'

  onMount(async () => {
    await import('flux-init')
    const snap = new Snap({
      appName: 'NFT Checkout',
      focus: true,
      product: {
        // imageURL:
        //   'https://lh3.googleusercontent.com/NpXUf_nwxn9yhHk_1AwFxRE7Mg2Lb7_rZoxKRuhf5Tca9MKm0Fh1MXuUAlJNJooO34l6YX3d-2MEZ1kpZvQ18JtrQbQw8CHnBdnRUV8=s992',
        videoURL:
          'https://mkpcdn.com/videos/d3a277f4e6f1212c900a1da4ec915aa9_675573.mp4',
        destinationAmount: 0.0004,
        destinationTicker: 'ETH',
        destinationAddress: '0xf636B6aA45C554139763Ad926407C02719bc22f7',
        title: 'The Crown (Patrick Mahomes)',
      },
      wallets: [
        { asset: 'btc', address: 'ms6k9Mdsbq5ZkoXakJexxjGjpH2PbSQdWK' },
      ],
      onMessage: msg => {
        const closeEvents = [snap.events.EXIT, snap.events.SUCCESS]
        switch (msg.event) {
          case snap.events.EXIT:
          case snap.events.SUCCESS:
            snap.closeWeb()
            break
          case snap.events.RESIZE:
            // resize iframe/viewport happened
            break
        }
      },
    })

    // Open using a button
    const btn = document.getElementById('buy-section')
    btn.onclick = snap.openWeb

    // Open using a QR code
    const canvas = document.getElementById('qr-canvas')
    snap.createQR({ element: canvas, pixelSize: 200 })
  })
</script>

<section id="buy-section">
  <article>
    <h2>NFT Checkout</h2>
    <div class="flex">
      <video loop autoplay muted class="nft-video svelte-1pit40i"
        ><source
          src="https://mkpcdn.com/videos/d3a277f4e6f1212c900a1da4ec915aa9_675573.mp4"
          class="svelte-1pit40i"
        /></video
      >
      <div class="relative">
        <p>
          Easily checkout with any NFT. Your loot is automagically added to your
          Collection.
        </p>
        <Button id="buy-button">Buy Crypto</Button>
        <div class="qr">
          <canvas id="qr-canvas" />
        </div>
      </div>
    </div>
  </article>
</section>

<style lang="scss">
  @import '../../../widget/src/styles/animations.scss';
  .flex {
    display: flex;
    grid-gap: 2rem;
  }
  .relative {
    position: relative;
  }
  section {
    --theme-button-color: #fff600;
    background: white;
    display: flex;
    flex-direction: column;
    padding: 10rem 10rem 10rem;
    margin-bottom: 0.25rem;
    article {
      position: relative;
      padding: 3rem;
      border-radius: 2rem;
      background: #f6f6f6;
      max-width: 800px;
      margin: 0 auto;
      h2 {
        position: absolute;
        top: -5rem;
        left: 0.75rem;
        font-size: 2rem;
        font-weight: bold;
      }
      p {
        font-size: 1.25rem;
        line-height: 1.75rem;
        padding: 0;
        margin: -0.25rem 0 1.5rem;
      }
      video {
        max-width: 50%;
      }
      .qr {
        position: absolute;
        bottom: -0.25rem;
        left: 0;
        #qr-canvas {
          max-height: 100px;
          border: 0.5rem solid #fff;
        }
      }
    }
  }
  @media (max-width: 1000px) {
    section {
      display: none;
    }
  }
</style>