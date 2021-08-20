const Wyre = require('../clients/wyre')

const payoutTask = async (data) => {
  try {
    const wyre = new Wyre()

    // TODO: figure out what fee should be and take from environment.
    // TODO: handle case where there's an amount left over (user sent too much)
    const swAmount = 0.0025 * data.sourceAmount
    const remainingAmount = data.sourceAmount - swAmount
    const baseParams = {
      preview: false,
      autoConfirm: true,
      sourceCurrency: data.sourceCurrency,
      destCurrency: data.destCurrency,
    }

    return await Promise.all([
      wyre.createTransfer({
        ...baseParams,
        source: `wallet:${data.source}`,
        sourceAmount: swAmount,
        dest: `wallet:${process.env.SNAP_WALLET_WYRE_SAVINGS_WALLET}`,
      }),
      wyre.createTransfer({
        ...baseParams,
        source: `wallet:${data.source}`,
        sourceAmount: remainingAmount,
        dest: `wallet:${data.destination}`,
      }),
    ])
  } catch (e) {
    throw e
  }
}

module.exports = {
  payoutTask,
}
