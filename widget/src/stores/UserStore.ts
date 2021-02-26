import { writable } from 'svelte/store'

function createStore() {
  const { subscribe, update } = writable({
    intent: 'buy',
    firstName: '',
    lastName: '',
    socialSecurityNumber: '',
    birthDate: '',
  })

  return {
    subscribe,
    setIntent: (intent: string) => update(s => ({ ...s, intent })),
    setFirstName: (firstName: string) => update(s => ({ ...s, firstName })),
    setLastName: (lastName: string) => update(s => ({ ...s, lastName })),
    setSocialSecurityNumber: (socialSecurityNumber: string) =>
      update(s => ({ ...s, socialSecurityNumber })),
    setBirthDate: (birthDate: string) => update(s => ({ ...s, birthDate })),
  }
}

export const userStore = createStore()
