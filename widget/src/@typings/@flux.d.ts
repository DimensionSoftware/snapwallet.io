import type { PromiseFluxApi } from 'api-client/dist/types/PromiseAPI'

declare var __ENV: {
  API_BASE_URL: string
  DEBUG: string
  [k: string]: string
}

declare global {
  interface Window {
    Plaid: any
    API?: PromiseFluxApi
  }
}
