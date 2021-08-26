const Wyre = require('../clients/wyre')
const { UnprocessableEntityError } = require('../error'),
  { createEvent, JOB_PUBLISHER } = require('../db')

const payoutTask = async (data) => {
  try {
    const userId = data.userId
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

    const internalBusinessTxn = {
      ...baseParams,
      sourceAmount: swAmount,
      dest: `wallet:${process.env.SNAP_WALLET_WYRE_SAVINGS_WALLET}`,
    }

    // NOTE: make idempotent
    await wyre.createTransfer(internalBusinessTxn)

    await createEvent({
      type: 'transferred_to_internal_wallet',
      meta: internalBusinessTxn,
      entity: { id: userId, kind: 'USER' },
    })

    const externalBusinessTxn = {
      ...baseParams,
      sourceAmount: remainingAmount,
      dest: `wallet:${data.destination}`,
    }

    await wyre.createTransfer(externalBusinessTxn)

    await createEvent({
      type: 'transferred_to_business_customer',
      meta: externalBusinessTxn,
      entity: { id: userId, kind: 'USER' },
    })
  } catch (e) {
    await JOB_PUBLISHER.publish({ worker: 'payoutTask', config: data })
    throw e
  }
}

module.exports = {
  payoutTask,
}
