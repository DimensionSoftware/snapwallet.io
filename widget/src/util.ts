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

export const setFluxSession = (jwt?: string) => {
  try {
    if (!jwt) throw new Error('No token provided')
    window.localStorage.setItem(JWT_SESSION_KEY, jwt)
  } catch (e) {
    Logger.error('Error setting flux session:', e)
    throw e
  }
}

export const getFluxSession = (): string => {
  return window.localStorage.getItem(JWT_SESSION_KEY) || ''
}

export const numberWithCommas = (s: string) =>
  s.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',')