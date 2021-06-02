# Donations

This use case can be used to accept fiat to cryptocurrency donations for your open source project, non profit organization, small business, school or anywhere else you would like to accept donations. Simply provide a destination wallet address for each crypto currency you'd like to accept donations for and optionally a source asset and/or source amount.

```html
<!DOCTYPE html5>
<html>
  <head>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="user-scalable=0, initial-scale=1, minimum-scale=1, maximum-scale=1, width=device-width"
    />
    <meta name="apple-mobile-web-app-capable" content="yes" />
  </head>
  <body>
    <canvas id="qr-canvas"></canvas>
    <br />
    <button id="donate-button">Donate $25</button>

    <script src="https://snapwallet.io/widget/dist/init.js"></script>
    <script>
      const snap = new Snap({
        appName: 'Example App',
        // 'sandbox' or 'production' default: 'production'
        environment: 'sandbox',
        wallets: [
          { asset: 'btc', address: 'ms6k9Mdsbq5ZkoXakJexxjGjpH2PbSQdWK' },
        ],
        onMessage: (msg) => {
          console.log('Msg', msg.event, 'Msg Data', msg.data)
          const closeEvents = [snap.events.EXIT, snap.events.SUCCESS]
          if (closeEvents.includes(msg.event)) {
            snap.closeWeb()
          }
        },
      })

      // Donate $25 using a button
      const btn = document.getElementById('donate-button')
      btn.onclick = () => snap.openWeb({ sourceAmount: 25 })

      // Donate $25 using a QR code
      const canvas = document.getElementById('qr-canvas')
      snap.createQR({ element: canvas, pixelSize: 200 }, { sourceAmount: 25 })
    </script>
  </body>
</html>
```
