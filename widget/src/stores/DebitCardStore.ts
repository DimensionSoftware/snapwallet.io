import { writable } from 'svelte/store'
import type { ICountry } from '../types'
import { countries } from '../util/country'

type DebitCardAddress = {
  street1: string
  street2?: string
  country: string
  city: string
  state: string
  postalCode: string
}

type StoreState = {
  reservationId: string
  number: string
  expirationDate: string
  verificationCode: string
  lastName: string
  firstName: string
  phoneNumber: string
  phoneNumberCountry: ICountry
  address: DebitCardAddress
}

const initialState: StoreState = {
  reservationId: '',
  firstName: '',
  lastName: '',
  number: '',
  // Exp date format: 10/2099
  expirationDate: '',
  verificationCode: '',
  phoneNumber: '',
  phoneNumberCountry: countries['US'],
  address: {
    street1: '',
    street2: '',
    city: '',
    state: '',
    country: '',
    postalCode: '',
  },
}

function createStore() {
  const { subscribe, update, set } = writable<StoreState>(initialState)

  return {
    subscribe,
    update: (fields: Partial<StoreState>) => {
      update(s => ({ ...s, ...fields }))
    },
    updateAddress: (address: Partial<DebitCardAddress>) => {
      update(s => ({ ...s, address: { ...s.address, ...address } }))
    },
    clear: () => set(initialState),
  }
}

export const debitCardStore = createStore()
