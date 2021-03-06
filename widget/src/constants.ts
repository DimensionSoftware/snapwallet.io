export const supportedCurrencyPairs = {
  wyre: [
    'USD_BTC',
    'USD_ETH',
    'USD_USDC',
    'USD_USDT',
    'USD_DAI',
    'USD_MKR',
    'USD_GUSD',
    'USD_PAX',
    'USD_LINK',
  ],
}

export const JWT_SESSION_KEY = '__FLUX_SESSION'

export enum Routes {
  ROOT = '/',
  SELECT_PAYMENT = '/select-payment',
  SEND_OTP = '/checkout',
  PROFILE = '/profile',
  VERIFY_OTP = '/verify-otp',
  PLAID_LINK = '/link-bank',
}
