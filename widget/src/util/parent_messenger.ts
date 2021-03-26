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
  const send = (ev: { event: ParentMessages; data: object }) => {
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
   * Sends user ID to parent for reference.
   */
  const exit = () => {
    const userID = window.AUTH_MANAGER.viewerUserID()
    send({
      event: ParentMessages.EXIT,
      data: {
        userID,
      },
    })
  }

  /**
   * User successfully completed a transaction.
   * Sends user ID and transaction to parent for reference.
   */
  const success = (txnId: string) => {
    const userID = window.AUTH_MANAGER.viewerUserID()
    send({
      event: ParentMessages.SUCCESS,
      data: {
        userID,
        txnId,
      },
    })
  }

  return {
    exit,
    success,
  }
})()
