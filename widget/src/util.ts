import nodeDebug from 'debug'
import { JWT_ACCESS_TOKEN_KEY, JWT_REFRESH_TOKEN_KEY } from './constants'
import * as CRYPTO_SVGS from 'svelte-cryptoicon'

// HACK: this lib. does not offer a good
// way to import icons dynamically, so convert all tickers to uppercase.
export const CryptoIcons = {}
Object.entries(CRYPTO_SVGS).forEach(([k, v]) => {
  CryptoIcons[k.toUpperCase()] = v
})

import {
  FluxApi,
  createConfiguration,
  ServerConfiguration,
  SecurityAuthentication,
  RequestContext,
} from 'api-client'

// pure fns
// ---------
export const onEnterPressed = (e, cb) => {
  if (onKeysPressed(e, ['Enter'])) cb()
}

export function onKeysPressed(e: Event, keys: Array<string>) {
  if (e instanceof KeyboardEvent) if (keys.includes(e.key)) return true
  return false
}

export const isValidNumber = (num: any) => {
  return isFinite(num) && !isNaN(num) && Number(num)
}

// Application logger module
export const Logger = (() => {
  window.localStorage.setItem('debug', __ENV.DEBUG)
  const error = nodeDebug('flux:error')
  const warn = nodeDebug('flux:warn')
  const debug = nodeDebug('flux:debug')
  const info = nodeDebug('flux:info')

  return {
    error,
    warn,
    debug,
    info,
  }
})()

// Set a JWT in local storage.
export const setFluxSession = (refreshToken: string | null, accessToken: string | null) => {
  try {
    if (refreshToken != null) {
      if (refreshToken) {
        window.localStorage.setItem(JWT_REFRESH_TOKEN_KEY, refreshToken)
      } else {
        window.localStorage.removeItem(JWT_REFRESH_TOKEN_KEY)
      }
    }

    if (accessToken != null) {
      if (accessToken) {
        window.localStorage.setItem(JWT_ACCESS_TOKEN_KEY, accessToken)
      } else {
        window.localStorage.removeItem(JWT_ACCESS_TOKEN_KEY)
      }
    }
  } catch (e) {
    Logger.error('Error setting flux session:', e)
    throw e
  }
}

// Get a JWT from local storage.
export const getFluxSession = (): string => {
  return window.localStorage.getItem(JWT_ACCESS_TOKEN_KEY) || ''
}


// Get a JWT from local storage.
export const getFluxRefreshToken = (): string => {
  return window.localStorage.getItem(JWT_REFRESH_TOKEN_KEY) || ''
}

export const numberWithCommas = (s: string) =>
  s.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',')

// Parse a JWT's data
export const parseJwt = token => {
  try {
    return JSON.parse(atob(token.split('.')[1]))
  } catch (e) {
    return null
  }
}

// Test for JWT expiration and existence.
export const isJWTValid = () => {
  const jwt = getFluxSession()
  const userData = parseJwt(jwt)
  if (!userData) return false
  const isTimeLeft = userData.exp > Math.floor(Date.now() / 1000)
  if (!isTimeLeft) setFluxSession(null, '')
  return isTimeLeft
}

// Authenticated route common configuration
export const authedRouteOptions = (component: any) => ({
  conditions: [isJWTValid],
  component,
})

export const genAPIClient = (): FluxApi => {
  const config = createConfiguration({
    baseServer: new ServerConfiguration(__ENV.API_BASE_URL, {}),
  })
  config.authMethods.Bearer = new FluxBearerAuthentication()

  return new FluxApi(config)
}

class FluxBearerAuthentication implements SecurityAuthentication {
  public constructor() {}

  public getName(): string {
    return 'Bearer'
  }

  public applySecurityAuthentication(context: RequestContext) {
    const token = getFluxSession()

    if (token) {
      context.setHeaderParam('Authorization', `Bearer ${token}`)
    }
  }
}

export const formatLocaleCurrency = (ticker: string, amount: number) => {
  amount = isValidNumber(amount) ? amount : 0
  const locale =
    (navigator?.languages || [])[0] || navigator?.language || 'en-US'
  return new Intl.NumberFormat(locale, {
    style: 'currency',
    currency: ticker,
  }).format(amount)
}

export const fileToBase64 = (file): Promise<string> =>
  new Promise((resolve, reject) => {
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = () => resolve(reader.result as string)
    reader.onerror = error => reject(error)
  })
