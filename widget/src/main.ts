import App from './App.svelte'
import { AuthManager, genAPIClient } from './auth'
import { PusherUtil } from './util/pusher_util'
import { userStore } from './stores/UserStore'
import { Logger } from './util'

Logger.debug('PHP eval executed... ;P')

// Auth
window.AUTH_MANAGER = new AuthManager()
window.AUTH_MANAGER.watch()
window.API = genAPIClient(window.AUTH_MANAGER)

// Pusher
PusherUtil.setup()

// Handle user
userStore.setIsLoggedIn(window.AUTH_MANAGER.viewerIsLoggedIn())

const queryParams = new URLSearchParams(window.location.search)
let config = {}
try {
  config = JSON.parse(decodeURI(queryParams.get('config')))
} catch {
  config = {}
}

const app = new App({
  target: document.body,
  props: {
    apiKey: config?.apiKey,
    appName: config?.appName,
    intent: config?.intent,
    focus: config?.focus,
    hideClose: config?.hideClose,
    wallets: config?.wallets || [],
    product: {
      // imageURL:
      //   'https://lh3.googleusercontent.com/NpXUf_nwxn9yhHk_1AwFxRE7Mg2Lb7_rZoxKRuhf5Tca9MKm0Fh1MXuUAlJNJooO34l6YX3d-2MEZ1kpZvQ18JtrQbQw8CHnBdnRUV8=s992',
      videoURL:
        'https://mkpcdn.com/videos/d3a277f4e6f1212c900a1da4ec915aa9_675573.mp4',
      destinationAmount: 18.1,
      destinationTicker: 'ETH',
      destinationAddress: '0xf636B6aA45C554139763Ad926407C02719bc22f7',
      title: 'The Crown (Patrick Mahomes)',
    },
    theme: {
      /* NOTE: each attribute maps to a css variable which is prefixed by --theme- */
      ...config?.theme || {},
    },
  },
})

export default app
