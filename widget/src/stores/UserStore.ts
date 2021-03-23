import type { UserFlags } from 'api-client'
import { writable } from 'svelte/store'
import { Routes } from '../constants'

type ViewerFlags = UserFlags & { hasEmail: boolean; hasPhone: boolean }

const initialAddress = {
  street1: '',
  street2: '',
  city: '',
  state: '',
  postalCode: '',
  country: '',
  phoneNumber: '',
}

function createStore() {
  const { subscribe, update } = writable({
    emailAddress: '',
    firstName: '',
    lastName: '',
    socialSecurityNumber: '',
    birthDate: '',
    phoneNumber: '',
    // Used for routing to last position
    // when auth kicks in.
    lastKnownRoute: Routes.ROOT,
    flags: {} as ViewerFlags,
    address: { ...initialAddress },
  })

  return {
    subscribe,
    setPhoneNumber: (phoneNumber: string) => {
      update(s => ({ ...s, phoneNumber }))
    },
    setEmailAddress: (emailAddress: string) =>
      update(s => ({ ...s, emailAddress })),
    setFirstName: (firstName: string) => update(s => ({ ...s, firstName })),
    setLastName: (lastName: string) => update(s => ({ ...s, lastName })),
    setSocialSecurityNumber: (socialSecurityNumber: string) =>
      update(s => ({ ...s, socialSecurityNumber })),
    setBirthDate: (birthDate: string) => update(s => ({ ...s, birthDate })),
    updateLastKnownRoute: (lastKnownRoute: Routes) =>
      update(s => ({ ...s, lastKnownRoute })),
    setFlags: (flags: ViewerFlags) => {
      update(s => ({ ...s, flags }))
    },
    setFullAddress: (address: any) => {
      update(s => ({
        ...s,
        address,
      }))
    },
    clearAddress: () => {
      update(s => ({ ...s, address: { ...initialAddress } }))
    },
  }
}

export const userStore = createStore()
