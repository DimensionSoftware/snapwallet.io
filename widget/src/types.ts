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
}
