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
export const CACHED_PRIMARY_PAYMENT_METHOD_KEY = '__SNAP_PRIMARY_PM_ID'

export enum Routes {
  ROOT = '/',
  SEND_PAYMENT = '/send-payment',
  SELECT_PAYMENT = '/select-payment',
  TRANSACTIONS = '/transactions',
  PROFILE_SEND_SMS = '/profile-send-sms',
  PROFILE_VERIFY_SMS = '/profile-verify-sms',
  PROFILE = '/profile',
  PROFILE_STATUS = '/profile-status',
  PROFILE_UPDATE = '/profile-update',
  SEND_OTP = '/send-otp',
  VERIFY_OTP = '/verify-otp',
  PLAID_LINK = '/link-bank',
  CHECKOUT_OVERVIEW = '/checkout-overview',
  ADDRESS = '/address',
  ADDRESS_UPDATE = '/address-update',
  FILE_UPLOAD = '/file-upload',
  FILE_UPLOAD_UPDATE = '/file-upload-update',
  SUCCESS = '/success',
}

export enum APIErrors {
  UNAUTHORIZED = 16,
}

export enum ParentMessages {
  EXIT = '__SNAP_EXIT',
  RESIZE = '__SNAP_RESIZE',
  SUCCESS = '__SNAP_SUCCESS',
}

export enum PusherServerMessages {
  WYRE_PM_UPDATED = 'WYRE_PAYMENT_METHODS_UPDATED',
  WYRE_ACCOUNT_UPDATED = 'WYRE_ACCOUNT_UPDATED',
}

export enum PusherClientMessages {}

export enum UserProfileFieldTypes {
  EMAIL = 'K_EMAIL',
  LEGAL_NAME = 'K_LEGAL_NAME',
  DATE_OF_BIRTH = 'K_DATE_OF_BIRTH',
  US_SSN = 'K_US_SSN',
  FULL_ADDRESS = 'K_ADDRESS',
  PHONE = 'K_PHONE',
  US_GOVT_DOCUMENT = 'K_US_GOVERNMENT_ID_DOC',
  PROOF_OF_ADDRESS_DOC = 'K_PROOF_OF_ADDRESS_DOC',
  ACH_AUTH_FORM = 'K_ACH_AUTHORIZATION_FORM_DOC',
  UNKNOWN = 'K_UNKNOWN',
}
