# Buy Use Case


so bro just...
```html
<!DOCTYPE html5>
<html>
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
  </head>
  <body>
    <canvas id="qr-canvas"></canvas>
    <br />
    <button id="buy-button">Buy Crypto</button>

    <script src="./dist/index.js"></script>
    <script>
      const snap = new Snap({
        appName: 'Example App',
        focus: true,
        wallets: [
          { asset: 'btc', address: 'ms6k9Mdsbq5ZkoXakJexxjGjpH2PbSQdWK' },
        ],
        onMessage: (msg) => {
          const closeEvents = [snap.events.EXIT, snap.events.SUCCESS]
          if (closeEvents.includes(msg.event)) {
            console.log('Message Data', msg.data)
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
