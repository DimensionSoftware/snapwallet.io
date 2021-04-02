import type { ProfileDataItemRemediation } from 'api-client'

export enum TransactionMediums {
  ACH = 'ach',
  BLOCKCHAIN = 'blockchain',
}

export interface IAsset {
  name: string
  ticker: string
}

export enum Masks {
  PHONE = '+x (xxx) xxx-xxxx',
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
