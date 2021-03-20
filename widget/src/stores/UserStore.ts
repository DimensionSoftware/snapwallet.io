import type { UserFlags } from 'api-client'
import { writable } from 'svelte/store'
import { Routes } from '../constants'

const initialAddress = {
  street1: '',
  street2: '',
  city: '',
  state: '',
  postalCode: '',
  country: '',
}

function createStore() {
  const { subscribe, update } = writable({
    emailAddress: '',
    firstName: '',
    lastName: '',
    socialSecurityNumber: '',
    birthDate: '',
    // Used for routing to last position
    // when auth kicks in.
    lastKnownRoute: Routes.ROOT,
    flags: {} as UserFlags,
    address: { ...initialAddress },
  })

  return {
    subscribe,
    setEmailAddress: (emailAddress: string) =>
      update(s => ({ ...s, emailAddress })),
    setFirstName: (firstName: string) => update(s => ({ ...s, firstName })),
    setLastName: (lastName: string) => update(s => ({ ...s, lastName })),
    setSocialSecurityNumber: (socialSecurityNumber: string) =>
      update(s => ({ ...s, socialSecurityNumber })),
    setBirthDate: (birthDate: string) => update(s => ({ ...s, birthDate })),
    updateLastKnownRoute: (lastKnownRoute: Routes) =>
      update(s => ({ ...s, lastKnownRoute })),
    setFlags: (flags: UserFlags) => {
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
