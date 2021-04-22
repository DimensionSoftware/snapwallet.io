import type { ProfileDataItemRemediation } from 'api-client'

export enum TransactionMediums {
  ACH = 'ach',
  BLOCKCHAIN = 'blockchain',
}

export interface IAsset {
  name: string
  ticker: string
  color: string
  popular?: boolean
}

export enum Masks {
  PHONE = 'xxx xxx-xxxx',
  SSN = 'xxx-xx-xxxx',
  INTL_DATE = 'xxxx-xx-xx',
  US_DATE = 'xx-xx-xxxx',
}

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
  destinationAmount: number
  destinationTicker: string
  destinationAddress: string
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
