import type { UserFlags } from 'api-client'
import { writable } from 'svelte/store'
import { Routes } from '../constants'

function createStore() {
  const { subscribe, update } = writable({
    intent: 'buy',
    emailAddress: '',
    firstName: '',
    lastName: '',
    socialSecurityNumber: '',
    birthDate: '',
    // Used for routing to last position
    // when auth kicks in.
    lastKnownRoute: Routes.ROOT,
    flags: {} as UserFlags,
    address: {
      street1: '',
      street2: '',
      city: '',
      state: '',
      postalCode: '',
      country: '',
    },
  })

  return {
    subscribe,
    setIntent: (intent: string) => update(s => ({ ...s, intent })),
    setEmailAddress: (emailAddress: string) =>
      update(s => ({ ...s, emailAddress })),
    setFullName: (fullName: string) => {
      const name = fullName.trim().split(/[\s+]/)
      // extract first & last name
      // - could easily add middle, later
      const [firstName, lastName] =
        name.length === 2
          ? [name[0], name[1]]
          : [name[0], name[name.length - 1]]
      return update(s => ({
        ...s,
        firstName,
        lastName,
      }))
    },
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
  }
}

export const userStore = createStore()
