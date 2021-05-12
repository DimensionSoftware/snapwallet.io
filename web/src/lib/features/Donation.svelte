<script lang="ts">
  import { onMount } from 'svelte'
  import Feature from '../Feature.svelte'
  import Button from '../Button.svelte'

  const themeColor = '#fffc00',
    config = {
      appName: 'Donation',
      focus: true,
      wallets: [
        { asset: 'btc', address: 'ms6k9Mdsbq5ZkoXakJexxjGjpH2PbSQdWK' },
      ],
      theme: {
        modalBackground: 'rgba(10,10,10,1)',
        modalPopupBackground: 'rgba(25,25,25,.9)',
        color: 'rgba(0,0,0,.9)',
        badgeTextColor: '#333',
        colorLightened: 'rgba(5,5,5,.8)',
        shadowBottomColor: 'rgba(0,0,0,.25)',
        colorInverse: '#fff',
        buttonColor: themeColor,
        buttonTextColor: '#000',
        buttonGlowColor: themeColor,
        successColor: themeColor,
        textColor: '#fff',
        inputTextColor: '#333',
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

    // TODO: Open using a QR code
    // const canvas = document.getElementById('qr-canvas')
    // snap.createQR({ element: canvas, pixelSize: 100 })
  })
</script>

<Feature
  name="donations"
  title="Donations & Tips"
  description="Accept crypto donations and tips to your business.  SnapWallet's simple onramp makes integration easy.  It's fully customizable."
  hasBackground={true}
  docLink="https://snapwallet.io/docs/guide/use-cases/donations.html"
>
  <div class="relative" slot="left">
    <div>
      <h3>Donate</h3>
      <small>to Snap Wallet</small>
      <aside>
        <Button on:mousedown={donateAsset('btc')}>Send BTC</Button>
        <Button on:mousedown={donateAsset('eth')}>Send ETH</Button>
      </aside>
      <input class="name" placeholder="Name" type="text" />
      <textarea placeholder="Message" />
    </div>
  </div>
  <div class="relative" slot="right">
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
    color: #fff;
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
  :global(.donations article) {
    margin-top: 4.55rem !important;
    background: linear-gradient(#000, rgba(0, 0, 0, 0.9)) !important;
    box-shadow: 5px 3px 10px 0 rgba(0, 0, 0, 0.2);
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
    margin-right: 0.75rem !important;
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

  input.name,
  textarea {
    font-size: 1rem;
    padding: 0.75rem 1.5rem 0.75rem 1rem;
    border: none;
    border-radius: 0.25rem;
    margin: 1.5rem 0 1rem;
  }
  input.name {
    margin-top: 3rem;
  }
  textarea {
    margin-top: 0;
    width: 75%;
  }
</style>
