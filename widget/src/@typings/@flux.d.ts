import type { PromiseFluxApi } from 'api-client/dist/types/PromiseAPI'
import type { AuthManager } from '../auth'
declare global {
  interface Window {
    Plaid: any
    Pusher: any
    API?: PromiseFluxApi
    AUTH_MANAGER?: AuthManager
    __SOCKET?: any
    tryInitializePusher: any
  }

  declare var __ENV: {
    API_BASE_URL: string
    DEBUG: string
    WYRE_BASE_URL: string
    [k: string]: string
  }
}
