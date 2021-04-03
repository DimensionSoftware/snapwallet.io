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
    wallets: config?.wallets || [],
    product: {
      // imageURL:
      //   'https://lh3.googleusercontent.com/NpXUf_nwxn9yhHk_1AwFxRE7Mg2Lb7_rZoxKRuhf5Tca9MKm0Fh1MXuUAlJNJooO34l6YX3d-2MEZ1kpZvQ18JtrQbQw8CHnBdnRUV8=s992',
      videoURL:
        'https://mkpcdn.com/videos/d3a277f4e6f1212c900a1da4ec915aa9_675573.mp4',
      destinationAmount: 0.0651,
      destinationTicker: 'ETH',
      destinationAddress: '0xf636B6aA45C554139763Ad926407C02719bc22f7',
      title: 'The Crown (Patrick Mahomes)',
    },
    theme: {
      /* NOTE: each attribute maps to a css variable which is prefixed by --theme- */
      // color: '#7f8c8d',
      // colorDarkened: '#2c3e50',
      // colorLightened: '#95a5a6',
      // Coinbase theme
      // modalBackground: '#070F15',
      // color: '#2187FF',
      // textColor: '#fff',
      // TODO: handle numbers in theme vars
      // 'text-color-3': '#fff',
      // colorInverse: '#263543',
      // warningColor: '#FFBD4A',
      // successColor: '#83E068',
      // errorColor: '#E7693C',
      // ...{
      //   // dark "snap" theme
      //   modalBackground: '#222',
      //   modalPopupBackground: '#444',
      //   color: 'rgb(100,100,100)',
      //   colorLightened: 'rgba(100,100,100,.8)',
      //   colorInverse: '#ddd',
      //   textColor: '#fff',
      //   inputTextColor: '#333',
      //   infoColor: '#74b9ff',
      // },
      ...config?.theme,
    },
  },
})

export default app
