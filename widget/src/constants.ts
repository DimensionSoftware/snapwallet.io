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

export const JWT_TOKENS_KEY = '__SNAP_TOKENS'
export const JWT_TOKENS_LOCK_KEY = '__SNAP_TOKENS_LOCK'
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
  DEMO_CURRENCY_SELECTED = '__SNAP_DEMO_CURRENCY_SELECTED',
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

export const SUPPORTED_CRYPTOCURRENCY_ASSETS = [
  { name: 'Aave', ticker: 'AAVE', color: '#EFB914' },
  { name: 'Basic Attention Token', ticker: 'BAT', color: '#FF5000' },
  { name: 'Binance USD', ticker: 'BUSD', color: '#EFB914' },
  { name: 'Bitcoin', ticker: 'BTC', popular: true, color: '#F7931A' },
  { name: 'Curve', ticker: 'CRV', color: '#EFB914' },
  { name: 'Compound', ticker: 'COMP', color: '#00D395' },
  { name: 'DAI', ticker: 'DAI', color: '#F4B731' },
  { name: 'Ethereum', ticker: 'ETH', popular: true, color: '#627EEA' },
  { name: 'Gemini Dollar', ticker: 'GUSD', color: '#00DCFA' },
  { name: 'Link', ticker: 'LINK', color: '#2A5ADA' },
  { name: 'MakerDAO', ticker: 'MKR', color: '#1AAB9B' },
  { name: 'Paxos Standard', ticker: 'PAX', color: '#EDE708' },
  { name: 'Stably Dollar', ticker: 'USDS', color: '#EFB914' },
  { name: 'Synthetix', ticker: 'SNX', color: '#EFB914' },
  { name: 'Tether', ticker: 'USDT', color: '#26A17B', popular: true },
  { name: 'UMA', ticker: 'UMA', color: '#FF4A4A' },
  { name: 'USDC', ticker: 'USDC', color: '#2775C9', popular: true },
  { name: 'Uniswap', ticker: 'UNI', color: '#FF007A' },
  { name: 'Wrapped Bitcoin', ticker: 'WBTC', color: '#323544' },
  { name: 'Yearn.Finance', ticker: 'YFI', color: '#006AE3' },
]
