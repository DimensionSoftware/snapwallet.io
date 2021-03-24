import App from './App.svelte'
import { AuthManager, genAPIClient } from './auth'
import { PusherUtil } from './util/pusher_util'

// Auth
window.AUTH_MANAGER = new AuthManager()
window.AUTH_MANAGER.watch()
window.API = genAPIClient(window.AUTH_MANAGER)
window.AUTH_MANAGER.addEventListeners()

// Pusher
PusherUtil.setup()

const queryParams = new URLSearchParams(window.location.search)
const config = JSON.parse(queryParams.get('config'))

const app = new App({
  target: document.body,
  props: {
    apiKey: config?.apiKey,
    appName: config.appName,
    intent: config.intent,
    wallets: config.wallets || [],
    theme: {
      /* NOTE: each attribute maps to a css variable which is prefixed by --theme- */
      // color: '#7f8c8d',
      // colorDarkened: '#2c3e50',
      // colorLightened: '#95a5a6',
    },
  },
})

export default app
