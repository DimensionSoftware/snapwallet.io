import nodeDebug from 'debug'
import { JWT_ACCESS_TOKEN_KEY, JWT_REFRESH_TOKEN_KEY } from '../constants'
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
import { FluxBearerAuthentication } from '../auth'

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
  try {
    window.localStorage.setItem('debug', __ENV.DEBUG)
    // These are needed for chrome
    window.localStorage.debug = __ENV.DEBUG
    nodeDebug.log = console.log.bind(console)
  } catch {
    console.warn('Unable to enable logger. Incognito?')
  }
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
  return window.AUTH_MANAGER.viewerIsLoggedIn()
}

// Authenticated route common configuration
export const authedRouteOptions = (component: any) => ({
  conditions: [isJWTValid],
  component,
})

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

export const dropEndingZeros = (str: string): string => {
  return str.replace(/\.0+0?$/g, '')
}
