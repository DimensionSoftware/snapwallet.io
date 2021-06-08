# Snap Wallet Init

The Snap Wallet fiat to cryptocurrency widget init script.

## Up and Running

```typescript
import SnapWallet, { WidgetEnvironments } from '@snapwallet/init'

// Base configuration
const snap = new SnapWallet({
  environment: WidgetEnvironments.PRODUCTION,
  appName: 'Donate',
  intent: 'donate',
  payee: 'Snap Wallet',
  wallets: [
    { asset: 'btc', address: '1BpnDtnUJk24P6XKEVu7XaYx5qPZx9AyKg' },
    { asset: 'eth', address: '0x6d47da1135b13d0068bd38c7236fcb59838cbbdd' },
  ],
})

// Create a Snap Wallet URI for a React Native WebView
const uri = snap.generateURL()

// or open the iframe for web
snap.openWeb()
```

[Donate to Snap Wallet](https://api.snapwallet.io/g/fQo62s6GR)
