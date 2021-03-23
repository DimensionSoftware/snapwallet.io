import { paymentMethodStore } from '../stores/PaymentMethodStore'
import { Logger } from '.'
import { PusherServerMessages, PusherClientMessages } from '../constants'

const handleWyrePaymentMethodUpdates = data => {
  Logger.debug(data)
  paymentMethodStore.fetchWyrePaymentMethods()
}

const tryInitializePusher = () => {
  if (window.Pusher) {
    // Use Logger to control log output for builds
    // See env.example DEBUG variable
    window.Pusher.log = Logger.debug
    window.Pusher.logToConsole = true
  }

  const userID = window.AUTH_MANAGER.viewerUserID()
  if (userID && !window.__SOCKET) {
    window.__SOCKET = new window.Pusher('dd280d42ccafc24e19ff', {
      cluster: 'us3',
      encrypted: true,
    })
    const channel = window.__SOCKET.subscribe(userID)

    channel.bind(
      PusherServerMessages.WYRE_PM_UPDATED,
      handleWyrePaymentMethodUpdates,
    )
    Logger.debug('PUSHER LOADED')
  }
}

/**
 * Send messages to the server over the user channel.
 *
 * @param eventName The event name to be sent to the server.
 * @param data Any data expected by the server for the event
 */
const send = (eventName: PusherClientMessages, data: object = {}) => {
  if (!window.Pusher || !window.__SOCKET) return
  const userID = window.AUTH_MANAGER.viewerUserID()
  window.__SOCKET.trigger(`client-${userID}`, eventName, data)
}

export const PusherUtil = (() => {
  const setup = () => {
    window.tryInitializePusher = tryInitializePusher
  }

  return { setup, send }
})()
