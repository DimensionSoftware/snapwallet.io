# Onramp

Onramp is a common use case where the Snap Wallet widget is configured for buying cryptocurrencies using a debit card, credit card or bank account. Simply provide a destination wallet address using the `wallets` parameter in the widget configuration and Snap Wallet will handle the rest:

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
    <button id="buy-button">Buy Crypto</button>

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

      // Open using a button
      const btn = document.getElementById('buy-button')
      btn.onclick = snap.openWeb

      // Open using a QR code
      const canvas = document.getElementById('qr-canvas')
      snap.createQR({ element: canvas, pixelSize: 200 })
    </script>
  </body>
</html>
```
