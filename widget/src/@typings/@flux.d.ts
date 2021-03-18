import type { PromiseFluxApi } from 'api-client/dist/types/PromiseAPI'
import type { AuthManager } from '../auth'
declare global {
  interface Window {
    Plaid: any
    Pusher: any
    API?: PromiseFluxApi
    AUTH_MANAGER?: AuthManager
  }

  declare var __ENV: {
    API_BASE_URL: string
    DEBUG: string
    [k: string]: string
  }
}
