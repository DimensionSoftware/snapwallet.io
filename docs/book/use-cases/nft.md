# NFT Checkouts

Configure Snap Wallet for NFT checkouts in a snap! Simply provide a `product` configuration parameter to get started. A user can purchase an NFT using a debit card or a bank account.

```html
<!DOCTYPE html5>
<html>
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
  </head>
  <body>
    <canvas id="qr-canvas"></canvas>
    <br />
    <button id="buy-button">Buy The Crown #1</button>

    <script src="https://snapwallet.io/widget/dist/init.js"></script>
    <script>
      const snap = new window.Snap({
        appName: 'Example App',
        // 'sandbox' or 'production' default: 'production'
        environment: 'sandbox',
        product: {
          videoURL:
            'https://mkpcdn.com/videos/d3a277f4e6f1212c900a1da4ec915aa9_675573.mp4',
          // Optionally provide an image URL instead
          // imageURL: '',
          destinationAmount: 20,
          destinationTicker: 'ETH',
          destinationAddress: '0xf636B6aA45C554139763Ad926407C02719bc22f7',
          title: 'The Crown #1',
        },
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
