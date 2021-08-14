const Router = require('koa-router')
const Wyre = require('../../clients/wyre')
const { createEvent, getBusinessByAPIKey } = require('../../db')
const { UnprocessableEntityError } = require('../../error')
const router = new Router()

router.post('/', async (ctx, _next) => {
  const { sourceCurrency, sourceAmount, destCurrency } = ctx.request.body

  const business = await getBusinessByAPIKey(
    ctx.request.headers['x-snap-wallet-api-key']
  )

  if (!business)
    throw new UnprocessableEntityError(
      'Please provide a valid business API key'
    )

  const wyre = new Wyre()
  const { data: txnWallet } = await wyre.createWallet({ type: 'SAVINGS' })

  const txnAddresses = txnWallet.depositAddresses
  let depositAddress = txnAddresses[sourceCurrency]
  // Any source currency that does not have a corresponding
  // address is going to be an ERC20 address.
  if (!depositAddress) depositAddress = txnAddresses['ETH']

  const params = {
    source: `wallet:${txnWallet.id}`,
    sourceCurrency,
    sourceAmount,
    dest: `wallet:${business.wallet.id}`,
    destCurrency,
    message: 'Snap Wallet Transaction :P',
    preview: true,
    autoConfirm: false,
  }

  const { data: txn } = await wyre.createTransfer(params)

  await createEvent({
    type: 'transaction',
    meta: {
      transaction_direction: 'wallet_wallet',
      source: txnWallet.id,
      sourceCurrency: txn.sourceCurrency,
      destination: business.wallet.id,
      // Use source amount from client intentionally.
      // Webhook handler should be able to know if
      // the user transferred less than required.
      sourceAmount,
      destAmount: txn.destAmount,
      destCurrency: txn.destCurrency,
    },
  })

  ctx.body = {
    depositAddress,
    preview: {
      fees: txn.fees,
      exchangeRate: txn.exchangeRate || 0,
      sourceAmount: txn.sourceAmount,
      sourceCurrency: txn.sourceCurrency,
      destAmount: txn.destAmount,
      destCurrency: txn.destCurrency,
    },
  }
})

module.exports = router
