# Checkout Use Case

SnapWallet supports checkouts for your organization.

In order to embed the widget for checkouts, set a description on the configuration to describe what is being purchased and where the money is being sent to.

i.e. "Rooster John Wick #32 sold by CryptoRoosters via OpenSea"

The amount can be fixed or enterable by the user.

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
        focus: true,
        product: {
          videoURL:
            'https://mkpcdn.com/videos/d3a277f4e6f1212c900a1da4ec915aa9_675573.mp4',
          destinationAmount: 20,
          destinationTicker: 'ETH',
          destinationAddress: '0xf636B6aA45C554139763Ad926407C02719bc22f7',
          title: 'The Crown #1',
        },
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

<p class="codepen" data-height="661" data-theme-id="light" data-default-tab="js,result" data-user="dreamcodez" data-slug-hash="VwPbqbG" style="height: 608px; width: 360px; box-sizing: border-box; display: flex; align-items: center; justify-content: center; border: 2px solid; margin: 1em 0; padding: 1em;" data-pen-title="Snap Wallet Checkout Example">
  <span>See the Pen <a href="https://codepen.io/dreamcodez/pen/VwPbqbG">
  Snap Wallet Checkout Example</a> by Matthew Elder (<a href="https://codepen.io/dreamcodez">@dreamcodez</a>)
  on <a href="https://codepen.io">CodePen</a>.</span>
</p>
<script async src="https://cpwebassets.codepen.io/assets/embed/ei.js"></script>
