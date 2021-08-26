const Wyre = require('../clients/wyre')
const { UnprocessableEntityError } = require('../error'),
  { createEvent, JOB_PUBLISHER, EVENTS } = require('../db'),
  { EVENT_KINDS } = require('../constants')

/**
 * Transfer a percentage of the transaction funds
 * to our internal business wallet.
 */
const processInternalBusinessTransaction = async ({ params, userId }) => {
  try {
    const wyre = new Wyre()
    await wyre.createTransfer(params)
    await createEvent({
      type: EVENT_KINDS.transferred_to_internal_wallet,
      meta: params,
      entity: { id: userId, kind: 'USER' },
    })
  } catch (e) {
    throw e
  }
}

/**
 * Transfer transaction funds to our business customer.
 */
const processExternalBusinessTransaction = async ({ params, userId }) => {
  try {
    const wyre = new Wyre()
    await wyre.createTransfer(params)
    await createEvent({
      type: EVENT_KINDS.transferred_to_internal_wallet,
      meta: params,
      entity: { id: userId, kind: 'USER' },
    })
  } catch (e) {
    throw e
  }
}

const payoutTask = async (data, logger) => {
  try {
    const userId = data.userId
    const events = await EVENTS.listByEntityID(userId)

    const { internalTxnRequired, externalTxnRequired } = events.reduce(
      (acc, e) => {
        if (e.kind === EVENT_KINDS.transferred_to_business_customer)
          return { ...acc, externalTxnRequired: false }
        if (e.kind === EVENT_KINDS.transferred_to_internal_wallet)
          return { ...acc, internalTxnRequired: false }

        return acc
      },
      { internalTxnRequired: true, externalTxnRequired: true }
    )

    if (!internalTxnRequired && !externalTxnRequired) {
      logger.info({ msg: 'Transaction already processed. Skipping.' })
    }

    const wyre = new Wyre()
    const wallet = await wyre.getWallet({ id: data.source })
    const walletBalances = wallet.availableBalances
    // Any currency not in the map will fall under ERC20 so fallback to eth address
    const walletBalance =
      walletBalances[data.sourceCurrency] || walletBalances['ETH']

    if (!walletBalance || walletBalance < data.sourceAmount) {
      throw new UnprocessableEntityError(
        `Wallet has insufficient funds for processing.`
      )
    }

    // TODO: figure out what fee should be and take from environment.
    const swAmount = 0.0025 * data.sourceAmount
    const remainingAmount = data.sourceAmount - swAmount
    const baseParams = {
      preview: false,
      autoConfirm: true,
      sourceCurrency: data.sourceCurrency,
      destCurrency: data.destCurrency,
      source: `wallet:${data.source}`,
    }

    // Collect our percentage
    if (internalTxnRequired) {
      const params = {
        ...baseParams,
        sourceAmount: swAmount,
        dest: `wallet:${process.env.SNAP_WALLET_WYRE_SAVINGS_WALLET}`,
      }
      await processInternalBusinessTransaction({ params, userId })
    }

    // Transfer remaining to our business customer
    if (externalTxnRequired) {
      const params = {
        ...baseParams,
        sourceAmount: remainingAmount,
        dest: `wallet:${data.destination}`,
      }
      await processExternalBusinessTransaction({ params, userId })
    }
  } catch (e) {
    await JOB_PUBLISHER.publish({ worker: 'payoutTask', config: data })
    throw e
  }
}

module.exports = {
  payoutTask,
}
