export const supportedCurrencyPairs = {
  wyre: [
    'AAVE_USD',
    'BAT_USD',
    'BTC_USD',
    'BUSD_USD',
    'COMP_USD',
    'CRV_USD',
    'DAI_USD',
    'DAI_USD',
    'ETH_USD',
    'GUSD_USD',
    'GUSD_USD',
    'LINK_USD',
    'MKR_USD',
    'PAX_USD',
    'SNX_USD',
    'UMA_USD',
    'UNI_USDT',
    'USDS_USD',
    'USDT_USD',
    'USD_AAVE',
    'USD_BAT',
    'USD_BTC',
    'USD_BUSD',
    'USD_COMP',
    'USD_CRV',
    'USD_DAI',
    'USD_ETH',
    'USD_GUSD',
    'USD_LINK',
    'USD_MKR',
    'USD_PAX',
    'USD_SNX',
    'USD_UMA',
    'USD_UNI',
    'USD_USDC',
    'USD_USDS',
    'USD_USDT',
    'USD_WBTC',
    'USD_YFI',
    'WBTC_USD',
    'YFI_USD',
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
  SUCCESS = '/success',
}

export enum APIErrors {
  UNAUTHORIZED = 16,
}

export enum ParentMessages {
  EXIT = '__SNAP_EXIT',
  SUCCESS = '__SNAP_SUCCESS',
}

export enum PusherServerMessages {
  WYRE_PM_UPDATED = 'WYRE_PAYMENT_METHODS_UPDATED',
  WYRE_ACCOUNT_UPDATED = 'WYRE_ACCOUNT_UPDATED',
}

export enum PusherClientMessages {}
