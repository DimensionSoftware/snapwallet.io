import { paymentMethodStore } from '../stores/PaymentMethodStore'
import { Logger } from '.'
import { PusherServerMessages, PusherClientMessages } from '../constants'
import { userStore } from '../stores/UserStore'
import { transactionsStore } from '../stores/TransactionsStore'

const handleWyrePaymentMethodUpdates = data => {
  Logger.debug(data)
  paymentMethodStore.fetchWyrePaymentMethods()
}

const handleWyreAccountUpdates = async data => {
  Logger.debug(data)
  await userStore.fetchFlags()
  userStore.fetchUserProfile()
  paymentMethodStore.fetchWyrePaymentMethods()
}

const handleWyreTransferUpdate = async data => {
  Logger.debug(data)
  transactionsStore.fetchUserTransactions()
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
    channel.bind(
      PusherServerMessages.WYRE_ACCOUNT_UPDATED,
      handleWyreAccountUpdates,
    )
    channel.bind(
      PusherServerMessages.WYRE_TRANSFER_UPDATED,
      handleWyreTransferUpdate,
    )
    Logger.debug('PUSHER LOADED')
  }
}

export const PusherUtil = (() => {
  /**
   * Initial setup of pusher.
   */
  const setup = () => {
    window.tryInitializePusher = tryInitializePusher
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

  return { setup, send }
})()
