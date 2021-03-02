import nodeDebug from 'debug'

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
