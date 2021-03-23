import App from './App.svelte'
import { AuthManager, genAPIClient } from './auth'
import { Logger } from './util'
import { PusherUtil } from './util/pusher_util'

window.AUTH_MANAGER = new AuthManager()
window.AUTH_MANAGER.watch()
window.API = genAPIClient(window.AUTH_MANAGER)

PusherUtil.setup()

// a test
window.addEventListener('logout', () => {
  Logger.debug('viewer has logged out')
})

window.addEventListener('prelogout', () => {
  const expirationEpoch = window.AUTH_MANAGER.getSessionExpiration()
  const expirationDate = new Date(expirationEpoch)
  // TODO: use ^^ for countdown input

  Logger.debug(`viewer will be logged out at: ${expirationDate}`)

  // TODO pop inactive session dialog up
  // TODO either reactivate session via getAccessToken() or logout()
})

const queryParams = new URLSearchParams(window.location.search)

const app = new App({
  target: document.body,
  props: {
    apiKey: queryParams.get('apiKey') || '',
    appName: queryParams.get('appName') || 'Flux',
    intent: queryParams.get('intent') || 'buy',
    theme: {
      /* NOTE: each attribute maps to a css variable which is prefixed by --theme- */
      // color: '#7f8c8d',
      // colorDarkened: '#2c3e50',
      // colorLightened: '#95a5a6',
    },
  },
})

export default app
