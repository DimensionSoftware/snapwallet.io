import { ParentMessages } from '../constants'
import type { IAsset } from '../types'

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

  const resize = (params, appName) => {
    const height = typeof params === 'number' ? params : params.height
    const width = typeof params === 'number' ? undefined : params.width
    send({
      event: ParentMessages.RESIZE,
      data: {
        height,
        width,
        appName,
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

  /**
   * Used for demo purposes. Alert the parent of the selected currency.
   */
  const currencySelected = (currency: IAsset) => {
    const userID = window.AUTH_MANAGER.viewerUserID()
    send({
      event: ParentMessages.DEMO_CURRENCY_SELECTED,
      data: {
        userID,
        currency,
      },
    })
  }

  return {
    exit,
    resize,
    success,
    currencySelected,
  }
})()
