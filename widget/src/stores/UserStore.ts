import type { UserFlags } from 'api-client'
import { writable } from 'svelte/store'
import { Routes, UserProfileFieldTypes } from '../constants'

type ViewerFlags = UserFlags & { hasEmail: boolean; hasPhone: boolean }

type VirtualProfile = {
  fullName?: string
  socialSecurityNumber?: string
  birthDate?: string
}

type UserStoreState = {
  emailAddress: string
  firstName: string
  lastName: string
  socialSecurityNumber: string
  phoneNumber: string
  lastKnownRoute: Routes
  flags: ViewerFlags
  address: {
    street1: string
    street2: string
    city: string
    state: string
    postalCode: string
    country: string
  }
  isLoggedIn: boolean
  virtual: VirtualProfile
  isProfileComplete: boolean
}

const initialAddress = {
  street1: '',
  street2: '',
  city: '',
  state: '',
  postalCode: '',
  country: '',
}

function createStore() {
  const defaultUser = {
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
      isLoggedIn: false,
      virtual: {
        fullName: '',
        birthDate: '',
        socialSecurityNumber: '',
      },
      isProfileComplete: false,
    },
    { subscribe, update } = writable<UserStoreState>(defaultUser)

  return {
    subscribe,
    fetchUserProfile: async () => {
      const { profile: userProfile } = await window.API.fluxViewerProfileData()
      const virtual: VirtualProfile = {}
      userProfile.forEach(item => {
        if (item.kind === UserProfileFieldTypes.LEGAL_NAME) {
          virtual.fullName = [...new Array(item.length)].join('*')
        }

        // these fields must always be the same length, so--
        if (item.kind === UserProfileFieldTypes.DATE_OF_BIRTH) {
          virtual.birthDate = '**-**-****'
        }

        if (item.kind === UserProfileFieldTypes.US_SSN) {
          virtual.socialSecurityNumber = '***-**-****'
        }
      })

      const isProfileComplete = Boolean(
        virtual.birthDate && virtual.fullName && virtual.socialSecurityNumber,
      )

      update(s => ({ ...s, virtual, isProfileComplete }))
    },
    setVirtual: (virtual: VirtualProfile) => {
      update(s => ({ ...s, virtual }))
    },
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
    setIsLoggedIn: (isLoggedIn: boolean) => {
      update(s => (isLoggedIn ? { ...s, isLoggedIn } : defaultUser))
    },
  }
}

export const userStore = createStore()
