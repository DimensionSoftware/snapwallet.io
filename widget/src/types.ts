export enum TransactionMediums {
  ACH = 'ach',
  BLOCKCHAIN = 'blockchain',
}

export interface IAsset {
  name: string
  ticker: string
}

export enum MaskTypes {
  PHONE = 'phone_number',
  INTL_DATE = 'intl_date',
  SSN = 'ssn',
}
