import App from './App.svelte'
import {
  FluxApi,
  createConfiguration,
  ServerConfiguration,
  ResponseContext,
} from 'api-client'

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

declare global {
  interface Window {
    __api: FluxApi
    API: (newToken?: string) => FluxApi
  }
}

function genAPIClient(token?: string): FluxApi {
  return new FluxApi(
    createConfiguration({
      baseServer: new ServerConfiguration(__ENV.API_BASE_URL, {}),
      authMethods: token
        ? {
            Bearer: `Bearer ${token}`,
          }
        : null,
    }),
  )
}

function getAPIClient(newToken?: string): FluxApi {
  if (!(window as any).__api || newToken) {
    return (window.__api = genAPIClient(newToken))
  } else {
    return window.__api
  }
}

// for testing, when needed, uncomment :D
window.API = getAPIClient

export default app
