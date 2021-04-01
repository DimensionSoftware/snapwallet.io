import App from './App.svelte'
import { AuthManager, genAPIClient } from './auth'
import { PusherUtil } from './util/pusher_util'
import { userStore } from './stores/UserStore'

// Auth
window.AUTH_MANAGER = new AuthManager()
window.AUTH_MANAGER.watch()
window.API = genAPIClient(window.AUTH_MANAGER)

// Pusher
PusherUtil.setup()

// Handle user
userStore.setIsLoggedIn(window.AUTH_MANAGER.viewerIsLoggedIn())

const queryParams = new URLSearchParams(window.location.search)
const config = JSON.parse(queryParams.get('config'))

const app = new App({
  target: document.body,
  props: {
    apiKey: config?.apiKey,
    appName: config?.appName,
    intent: config?.intent,
    focus: config?.focus,
    wallets: config?.wallets || [],
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
      ...config?.theme,
    },
  },
})

export default app
