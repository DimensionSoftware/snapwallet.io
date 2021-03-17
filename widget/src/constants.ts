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
  PROFILE = '/profile',
  SEND_OTP = '/send-otp',
  VERIFY_OTP = '/verify-otp',
  PLAID_LINK = '/link-bank',
  CHECKOUT_OVERVIEW = '/checkout-overview',
  ADDRESS = '/address',
  ADDRESS_2 = '/address-2',
  FILE_UPLOAD = '/file-upload',
}

export enum APIErrors {
  UNAUTHORIZED = 16,
}
