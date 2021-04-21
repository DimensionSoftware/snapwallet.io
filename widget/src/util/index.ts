import { getContext } from 'svelte'
import nodeDebug from 'debug'
import * as Icons from './icons'
import { ParentMessages, CACHED_PRIMARY_PAYMENT_METHOD_KEY } from '../constants'

export const CryptoIcons = Icons

// pure fns
// ---------
export const focus = (e: HTMLElement, duration = 150) => {
  setTimeout(() => e?.focus(), duration)
}

export const focusFirstInput = (duration = 400) => {
  focus(document.querySelector('input:first-child'), duration)
}

export const onEnterPressed = (e, cb) => {
  if (onKeysPressed(e, ['Enter'])) cb()
}

export const onEscPressed = (e, cb) => {
  if (onKeysPressed(e, ['Escape'])) cb()
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
  const error = nodeDebug('SnapWallet:error')
  const warn = nodeDebug('SnapWallet:warn')
  const debug = nodeDebug('SnapWallet:debug')
  const info = nodeDebug('SnapWallet:info')

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

export const isValidDate = (d: Date) => {
  return d instanceof Date && !isNaN(+d)
}

export const formatHumanDate = (apiDate: string) =>
  formatDate(apiDate, {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  } as any)

export const formatDate = (
  apiDate: string,
  options = { year: 'numeric', month: 'numeric', day: 'numeric' },
) => {
  Logger.debug('date', apiDate)
  const date = new Date(Date.parse(apiDate))
  return date.toLocaleDateString(
    date.getTimezoneOffset() as any,
    options as any,
  )
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

export const dropEndingZeros = (str: string): string => {
  return str.replace(/\.0+0?$/g, '')
}

export const capitalize = (s: string) => s[0].toUpperCase() + s.substr(1)

export const cachePrimaryPaymentMethodID = (pmId: string) => {
  try {
    window.localStorage.setItem(CACHED_PRIMARY_PAYMENT_METHOD_KEY, pmId)
  } catch (e) {
    Logger.warn('Could not cache pm id', e)
  }
}

export const getPrimaryPaymentMethodID = (): string => {
  try {
    return window.localStorage.getItem(CACHED_PRIMARY_PAYMENT_METHOD_KEY)
  } catch (e) {
    Logger.warn('Could not get cached pm id', e)
    return ''
  }
}

export const resizeWidget = (height: number, appName: string) => {
  window.dispatchEvent(
    new CustomEvent(ParentMessages.RESIZE, {
      detail: { height: `${height}px`, appName },
    }),
  )
}
