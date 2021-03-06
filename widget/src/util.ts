import nodeDebug from 'debug'
import { JWT_SESSION_KEY } from './constants'

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
  return Number(num) && !isNaN(num) && num !== Infinity
}

// Application logger module
export const Logger = (() => {
  window.localStorage.debug = __ENV.DEBUG
  const error = nodeDebug('flux:error')
  const debug = nodeDebug('flux:debug')
  const info = nodeDebug('flux:info')

  return {
    error,
    debug,
    info,
  }
})()

// Set a JWT in local storage.
export const setFluxSession = (jwt?: string) => {
  try {
    if (!jwt) throw new Error('No token provided')
    window.localStorage.setItem(JWT_SESSION_KEY, jwt)
  } catch (e) {
    Logger.error('Error setting flux session:', e)
    throw e
  }
}

// Get a JWT from local storage.
export const getFluxSession = (): string => {
  return window.localStorage.getItem(JWT_SESSION_KEY) || ''
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
  return userData.exp < Math.floor(Date.now() / 1000)
}

// Authenticated route common configuration
export const authedRouteOptions = (component: any) => ({
  conditions: [isJWTValid],
  component,
})
