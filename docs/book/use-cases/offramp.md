# Offramp

Offramp is a common use case where the Snap Wallet widget is configured to sell cryptocurrencies from a wallet, back into fiat (credit cards, debit cards, bank accounts, etc...). Once your customer has confirmed their transaction, the Snap Wallet widget will send a callback message prompting the initiation of the wallet transfer.

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
    <button id="sell-button">Sell Crypto</button>

    <script src="https://snapwallet.io/widget/dist/init.js"></script>
    <script>
      const snap = new Snap({
        appName: 'Example App',
        // 'sandbox' or 'production' default: 'production'
        environment: 'sandbox',
        intent: 'sell',
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
      const btn = document.getElementById('sell-button')
      btn.onclick = snap.openWeb

      // Open using a QR code
      const canvas = document.getElementById('qr-canvas')
      snap.createQR({ element: canvas, pixelSize: 200 })
    </script>
  </body>
</html>
```
