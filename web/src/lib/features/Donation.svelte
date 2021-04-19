<script lang="ts">
  import { onMount } from 'svelte'
  import Feature from '../Feature.svelte'
  import Button from '../../../../widget/src/components/Button.svelte'

  const config = {
    appName: 'Donation',
    focus: true,
    wallets: [{ asset: 'btc', address: 'ms6k9Mdsbq5ZkoXakJexxjGjpH2PbSQdWK' }],
    theme: {
      badgeTextColor: '#333',
      inputColor: '#333',
      shadowBottomColor: 'rgba(0,0,0,.075)',
      colorInverse: '#fff',
      buttonColor: 'rgba(247,127,26,1)',
      buttonTextColor: '#fff',
      successColor: '#fffc00',
      colorLightened: 'rgba(247,127,26,.5)',
    },
  }

  let snap: any = {}

  const donateUSDAmount = sourceAmount => {
    snap.defaultDestinationAsset = undefined
    return () => snap.openWeb({ sourceAmount })
  }

  const donateAsset = defaultDestinationAsset => {
    snap.sourceAmount = undefined
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
    const canvas = document.getElementById('qr-canvas')
    snap.createQR({ element: canvas, pixelSize: 100 })
  })
</script>

<Feature
  right
  docLink="https://snapwallet.io/docs/guide/use-cases/donate.html"
  title="Donations & Tips"
>
  <div class="relative" slot="left">
    <p>Accept Crypto Donations & Tips, Simply.</p>
    <div>
      <h3>Donate</h3>
      <small>to Snap Wallet</small>
      <aside>
        <Button on:mousedown={donateAsset('btc')}>Send BTC</Button>
        <Button on:mousedown={donateAsset('eth')}>Send ETH</Button>
      </aside>
    </div>
  </div>
  <div class="relative" slot="right">
    <p>&nbsp;</p>
    <div>
      <h3>Tip</h3>
      <small>to Snap Wallet</small>
      <aside>
        <Button on:mousedown={donateUSDAmount(1.0)}>$1</Button>
        <Button on:mousedown={donateUSDAmount(5.0)}>$5</Button>
        <Button on:mousedown={donateUSDAmount(10.0)}>$10</Button>
      </aside>
    </div>
  </div>
</Feature>

<style lang="scss">
  @import '../../../../widget/src/styles/animations.scss';

  .relative {
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
  aside {
    display: flex;
    gap: 0.5rem;
    margin-top: 2rem;
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
</style>
