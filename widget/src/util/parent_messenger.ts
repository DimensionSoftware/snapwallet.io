import { ParentMessages } from '../constants'

/**
 * Util module for sending parent application messages
 */
export const ParentMessenger = (() => {
  /**
   * General send method for publishing
   * JSON stringified events to the parent.
   *
   * @param event Any outgoing JSON event event.
   */
  const send = (ev: { event: ParentMessages }) => {
    const msg = JSON.stringify(ev)
    if (window.parent) {
      window.parent.postMessage(msg, '*')
    }
    if ((window as any).ReactNativeWebView) {
      ;(window as any).ReactNativeWebView?.postMessage(msg)
    }
  }

  /**
   * User exited application (clicked X)
   */
  const exit = () =>
    send({
      event: ParentMessages.EXIT,
    })

  return {
    exit,
  }
})()
