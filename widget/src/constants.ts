export const supportedCurrencyPairs = {
  wyre: [
    'USD_BTC',
    'BTC_USD',
    'USD_ETH',
    'ETH_USD',
    'USD_USDC',
    'USDC_USD',
    'USD_USDT',
    'USDT_USD',
    'USD_DAI',
    'DAI_USD',
    'USD_MKR',
    'MKR_USD',
    'USD_GUSD',
    'GUSD_USD',
    'USD_PAX',
    'PAX_USD',
    'USD_LINK',
    'LINK_USD',
  ],
}

export const JWT_REFRESH_TOKEN_KEY = '__SNAP_REFRESH'
export const JWT_ACCESS_TOKEN_KEY = '__SNAP_ACCESS'

export enum Routes {
  ROOT = '/',
  SELECT_PAYMENT = '/select-payment',
  PROFILE_SEND_SMS = '/profile-send-sms',
  PROFILE_VERIFY_SMS = '/profile-verify-sms',
  PROFILE = '/profile',
  SEND_OTP = '/send-otp',
  VERIFY_OTP = '/verify-otp',
  PLAID_LINK = '/link-bank',
  CHECKOUT_OVERVIEW = '/checkout-overview',
  ADDRESS = '/address',
  FILE_UPLOAD = '/file-upload',
}

export enum APIErrors {
  UNAUTHORIZED = 16,
}

export enum ParentMessages {
  EXIT = '__SNAP_EXIT',
}

export enum PusherServerMessages {
  WYRE_PM_UPDATED = 'WYRE_PAYMENT_METHODS_UPDATED',
}

export enum PusherClientMessages {}
