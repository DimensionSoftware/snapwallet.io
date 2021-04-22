import type {
  ProfileDataItemInfo,
  ProfileDataItemRemediation,
  UserFlags,
} from 'api-client'
import { writable } from 'svelte/store'
import { Routes, UserProfileFieldTypes } from '../constants'
import type { ICountry } from '../types'
import { countries } from '../util/country'

type ViewerFlags = UserFlags & { hasEmail: boolean; hasPhone: boolean }

type VirtualProfile = {
  fullName?: string
  socialSecurityNumber?: string
  birthDate?: string
  address?: {
    street1: string
    street2?: string
    country: string
    city: string
    state: string
    postalCode: string
  }
}

type UserStoreState = {
  emailAddress: string
  firstName: string
  lastName: string
  socialSecurityNumber: string
  birthDate: string
  phoneNumber: string
  phoneNumberCountry: ICountry
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
  isProfilePending: boolean
  profileRemediations: ProfileDataItemRemediation[]
  profileItems: { [k: string]: ProfileDataItemInfo }
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
      phoneNumberCountry: countries['US'],
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
        address: initialAddress,
      },
      isProfileComplete: false,
      isProfilePending: false,
      profileRemediations: [],
      profileItems: {},
    },
    { subscribe, update, set } = writable<UserStoreState>(defaultUser)

  return {
    subscribe,
    reset: () => set(defaultUser),
    setProfilePending: (isProfilePending: boolean = true) => {
      update(s => ({ ...s, isProfilePending }))
    },
    fetchUserProfile: async () => {
      const {
        profile: userProfile,
        remediations = [],
        wyre,
      } = await window.API.fluxViewerProfileData()
      const virtual: VirtualProfile = {}
      const profileItems = {}
      userProfile.forEach(item => {
        profileItems[item.kind] = item
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

        if (item.kind === UserProfileFieldTypes.FULL_ADDRESS) {
          virtual.address = {
            street1: '***** ***********',
            street2: '**** ***',
            country: '**',
            city: '********',
            state: '**',
            postalCode: '*****',
          }
        }
      })

      const isProfileComplete = Boolean(
        virtual.birthDate && virtual.fullName && virtual.socialSecurityNumber,
      )

      const isProfilePending = wyre?.status === 'OPEN' && !remediations?.length

      update(s => ({
        ...s,
        virtual,
        profileItems,
        isProfilePending,
        isProfileComplete,
        profileRemediations: remediations,
      }))
    },
    setVirtual: (virtual: VirtualProfile) => {
      update(s => ({ ...s, virtual }))
    },
    setPhoneNumber: (phoneNumber: string) => {
      update(s => ({ ...s, phoneNumber }))
    },
    setPhoneNumberCountry: (phoneNumberCountry: ICountry) => {
      update(s => ({ ...s, phoneNumberCountry }))
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
    fetchFlags: async () => {
      const { flags: userFlags, user } = await window.API.fluxViewerData()
      const flags = {
        ...userFlags,
        hasEmail: Boolean(user.email),
        hasPhone: Boolean(user.phone),
      }
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
    clearProfile: () => {
      update(s => ({
        ...s,
        firstName: '',
        lastName: '',
        socialSecurityNumber: '',
        birthDate: '',
      }))
    },
    setIsLoggedIn: (isLoggedIn: boolean) => {
      update(s => (isLoggedIn ? { ...s, isLoggedIn } : defaultUser))
    },
  }
}

export const userStore = createStore()
