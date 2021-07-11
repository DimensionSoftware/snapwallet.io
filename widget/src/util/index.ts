import { getContext } from 'svelte'
import nodeDebug from 'debug'
import * as Icons from './icons'
import { ParentMessages, CACHED_PRIMARY_PAYMENT_METHOD_KEY } from '../constants'
import { isValidMaskInput } from '../masks'

export const CryptoIcons = Icons

// pure fns
// ---------
export const focus = (e: HTMLElement, duration = 150) => {
  setTimeout(() => e?.focus(), duration)
}

export const focusFirstInput = (duration = 400) => {
  focus(document.querySelector('input:first-child'), duration)
}

export const onFocusSelect = node => {
  const handleFocus = event => {
    node && typeof node.select === 'function' && node.select()
  }
  node.addEventListener('focus', handleFocus)
  return {
    destroy() {
      node.removeEventListener('focus', handleFocus)
    },
  }
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

export const isValidKeyForMask = (e, mask, defaultValue) => {
  const newVal = defaultValue + String.fromCharCode(e.keyCode)
  const isValLongerThanMask = newVal.length > mask.length
  // Uses codes from the following table https://keycode.info/
  const isAltering =
    [8, 9, 12, 13, 16, 17, 18, 20, 41, 46].includes(e.keyCode) ||
    e.metaKey ||
    ['ArrowLeft', 'ArrowRight', 'ArrowUp', 'ArrowDown'].includes(e.key)

  const isInputValid = isValidMaskInput(newVal, mask) && !isValLongerThanMask

  if (!isInputValid && !isAltering) {
    e.preventDefault()
    return false
  }
}

// closest to n and divisible by m
export const closestNumber = (n: number, m: number) => {
  // find the quotient
  let q = parseInt('' + n / m)

  // 1st possible closest number
  let n1 = m * q

  // 2nd possible closest number
  let n2 = n * m > 0 ? m * (q + 1) : m * (q - 1)

  // if true, then n1 is the
  // required closest number
  if (Math.abs(n - n1) < Math.abs(n - n2)) return n1

  // else n2 is the required
  // closest number
  return n2
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
  })

export const formatDate = (
  apiDate: string,
  options = { year: 'numeric', month: 'numeric', day: 'numeric' },
) => {
  Logger.debug('date', apiDate)
  const date = new Date(Date.parse(apiDate))
  return date.toLocaleDateString(date.getTimezoneOffset(), options)
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
