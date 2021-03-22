import App from './App.svelte'
import { AuthManager, genAPIClient } from './auth'
import { Logger } from './util'

window.AUTH_MANAGER = new AuthManager()
window.AUTH_MANAGER.watch()
window.API = genAPIClient(window.AUTH_MANAGER)

window.tryInitializePusher = function tryInitializePusher() {
  if (window.Pusher) {
    // Use Logger to control log output for builds
    // See env.example DEBUG variable
    window.Pusher.log = Logger.debug
    window.Pusher.logToConsole = true
  }

  const userID = window.AUTH_MANAGER.viewerUserID()
  if (userID && !window.__SOCKET) {
    window.__SOCKET = new window.Pusher('dd280d42ccafc24e19ff', {
      cluster: 'us3',
    })

    const channel = window.__SOCKET.subscribe(userID)
    channel.bind(function (data) {
      Logger.debug(data)
    })

    //Logger.debug('PUSHER LOADED :)')
    console.log('pusher loaded')
  }
}

// a test
window.addEventListener('logout', () => {
  console.log('viewer has logged out')
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
