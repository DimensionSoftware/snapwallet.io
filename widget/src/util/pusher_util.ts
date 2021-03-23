import { paymentMethodStore } from '../stores/PaymentMethodStore'
import { Logger } from '.'
import { PusherMessages } from '../constants'

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

    channel.bind(PusherMessages.WYRE_PM_UPDATED, handleWyrePaymentMethodUpdates)
    Logger.debug('PUSHER LOADED')
  }
}

export const PusherUtil = (() => {
  const setup = () => {
    window.tryInitializePusher = tryInitializePusher
  }

  return { setup }
})()
