import type { PromiseFluxApi } from 'api-client/dist/types/PromiseAPI'

declare global {
  interface Window {
    Plaid: any
    Pusher: any
    API?: PromiseFluxApi
  }

  declare var __ENV: {
    API_BASE_URL: string
    DEBUG: string
    [k: string]: string
  }
}
