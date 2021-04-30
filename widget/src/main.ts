import App from './App.svelte'
import { AuthManager, genAPIClient } from './auth'
import { PusherUtil } from './util/pusher_util'
import { userStore } from './stores/UserStore'
import { Logger } from './util'
import { configStore } from './stores/ConfigStore'

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
let config = {} as any
try {
  config = JSON.parse(decodeURI(queryParams.get('config')))
} catch {
  config = {}
}

configStore.setInitial({
  ...config,
  sourceAmount: 100.0,
  defaultDestinationAsset: 'eth',
  theme: {
    modalBackground: 'rgba(40,40,40,.9)',
    modalPopupBackground: 'rgba(50,50,50,.95)',
    color: 'rgba(0,0,0,.9)',
    badgeTextColor: '#333',
    colorLightened: 'rgba(5,5,5,.8)',
    shadowBottomColor: 'rgba(0,0,0,.25)',
    colorInverse: '#fff',
    // buttonColor: '#fffc00',
    buttonColor: 'rgb(247, 127, 26)',
    buttonTextColor: '#000',
    // successColor: '#fffc00',
    successColor: 'rgb(247, 127, 26)',
    textColor: '#fff',
    inputTextColor: '#333',
  },
})

const app = new App({
  target: document.body,
  // NOTE: ConfigStore now handles external props
  props: {},
})

export default app
