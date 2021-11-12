import type { ProfileDataItemRemediation } from 'api-client'

export enum TransactionMediums {
  ACH = 'ach',
  DEBIT_CARD = 'debit_card',
  BLOCKCHAIN = 'blockchain',
}

export interface IAsset {
  name: string
  ticker: string
  color: string
  popular?: boolean
}

export enum Masks {
  CODE = 'x',
  PHONE = 'xxx xxx-xxxx',
  SSN = 'xxx-xx-xxxx',
  INTL_DATE = 'xxxx-xx-xx',
  US_DATE = 'xx-xx-xxxx',
  DEBIT_CARD = 'xxxx xxxx xxxx xxxx',
  DEBIT_CARD_EXPIRATION_DATE = 'xx/xxxx',
}

export enum WidgetEnvironments {
  // ** development ** is only an option for explicitness
  // Simply provide the environment variable INIT_API_BASE_URL for dev
  DEVELOPMENT = 'development',
  SANDBOX = 'sandbox',
  PRODUCTION = 'production',
}

export type UserIntent = 'buy' | 'sell' | 'donate' | 'cart'

export type SrcDst = 'source' | 'destination'

export enum TransactionIntents {
  BUY = 'buy',
  SELL = 'sell',
}

export enum FileUploadTypes {
  US_PASSPORT = 'GI_US_PASSPORT',
  US_DRIVER_LICENSE = 'GI_US_DRIVING_LICENSE',
  ACH_AUTHORIZATION_FORM = 'ACH_AUTHZ_FORM',
  PROOF_OF_ADDRESS = 'PROOF_OF_ADDRESS',
}

export type RemediationGroups = {
  personal: ProfileDataItemRemediation[]
  address: ProfileDataItemRemediation[]
  contact: ProfileDataItemRemediation[]
  document: ProfileDataItemRemediation[]
}

export type ProductType = {
  imageURL?: string
  videoURL?: string
  title: string
  author?: string
  subtitle?: string
  destinationAmount: number
  destinationTicker: string
  destinationAddress: string
  qty?: number
}

export type WalletType = {
  asset: string
  address: string
  default?: boolean
}

export interface ICountry {
  name: string
  dial_code: string
  code: string
}
