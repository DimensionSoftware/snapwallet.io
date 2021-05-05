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
// woogly boogly.... woodey woobey
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
})

const app = new App({
  target: document.body,
  // NOTE: ConfigStore now handles external props
  props: {},
})

export default app

// ..
