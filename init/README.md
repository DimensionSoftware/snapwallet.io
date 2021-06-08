# Snap Wallet Init

The Snap Wallet fiat to cryptocurrency widget init script.

## Up and Running

```typescript
import SnapWallet, { WidgetEnvironments } from '@snapwallet/init'

// Base configuration
const snap = new SnapWallet({
  environment: WidgetEnvironments.SANDBOX,
  appName: 'Some App',
  wallets: [],
})

// Create a Snap Wallet URI for a React Native WebView
const uri = snap.generateURL()

// or open the iframe for web
snap.openWeb()
```
