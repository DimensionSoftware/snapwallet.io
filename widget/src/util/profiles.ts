import type {
  ProfileDataItemInfo,
  ProfileDataItemKind,
  ProfileDataItemRemediation,
} from 'api-client'
import { UserProfileFieldTypes } from '../constants'
import type { RemediationGroups } from '../types'

/**
 * Get difference between two Sets
 */
const getDiff = (required, submitted) =>
  new Set([...required].filter(x => !submitted.has(x)))

export const groupRemediations = (
  remediations: ProfileDataItemRemediation[],
): RemediationGroups => {
  const result = {
    personal: [],
    address: [],
    contact: [],
    document: [],
  }

  remediations.forEach(r => {
    if (isPersonalInfo(r.kind)) result.personal.push(r)
    if (isAddressInfo(r.kind)) result.address.push(r)
    if (isContactInfo(r.kind)) result.contact.push(r)
    if (isDocumentInfo(r.kind)) result.document.push(r)
  })

  return result
}

export const isPersonalInfo = (kind: ProfileDataItemKind): boolean => {
  return [
    UserProfileFieldTypes.DATE_OF_BIRTH,
    UserProfileFieldTypes.LEGAL_NAME,
    UserProfileFieldTypes.US_SSN,
  ].includes(kind as UserProfileFieldTypes)
}

export const isAddressInfo = (kind: ProfileDataItemKind): boolean => {
  return kind === UserProfileFieldTypes.FULL_ADDRESS
}

export const isContactInfo = (kind: ProfileDataItemKind): boolean => {
  return [UserProfileFieldTypes.EMAIL, UserProfileFieldTypes.PHONE].includes(
    kind as UserProfileFieldTypes,
  )
}

export const isDocumentInfo = (kind: ProfileDataItemKind): boolean => {
  return [
    UserProfileFieldTypes.PROOF_OF_ADDRESS_DOC,
    UserProfileFieldTypes.US_GOVT_DOCUMENT,
    UserProfileFieldTypes.ACH_AUTH_FORM,
  ].includes(kind as UserProfileFieldTypes)
}

export const reducePersonalInfoFields = (
  remediations: ProfileDataItemRemediation[],
) => {
  if (!remediations.length) return 'Identity is used for verification.'
  const message = 'Your personal information requires an update.'
  const fields = []

  remediations.forEach(r => {
    if (UserProfileFieldTypes.LEGAL_NAME === r.kind) fields.push('legal name')
    if (UserProfileFieldTypes.DATE_OF_BIRTH === r.kind) fields.push('birthdate')
    if (UserProfileFieldTypes.US_SSN === r.kind)
      fields.push('social security number')
  })

  // An unsupported field made it here.
  if (!fields.length) return `${message} Please contact support.`

  let fieldMsg = fields[0]
  if (fields.length > 2) {
    const fIdx = fields.length - 1
    fields[fIdx] = `and ${fields[fIdx]}`
    fieldMsg = fields.join(', ')
  }
  if (fields.length === 2) {
    fieldMsg = `${fields[0]} and ${fields[1]}`
  }
  return `${message} Fields include ${fieldMsg}.`
}

export const reduceDocumentFields = (
  remediations: ProfileDataItemRemediation[],
) => {
  if (!remediations?.length) return 'Documents verify your person.'
  const message = 'One or more of your documents require an update.'
  const fields = []

  remediations.forEach(r => {
    if (UserProfileFieldTypes.ACH_AUTH_FORM === r.kind)
      fields.push('bank authorization form')
    if (UserProfileFieldTypes.PROOF_OF_ADDRESS_DOC === r.kind)
      fields.push('proof of address')
    if (UserProfileFieldTypes.US_GOVT_DOCUMENT === r.kind)
      fields.push('U.S. government ID')
  })

  // An unsupported field made it here.
  if (!fields.length) return `${message} Please contact support.`

  let fieldMsg = fields[0]
  if (fields.length > 2) {
    const fIdx = fields.length - 1
    fields[fIdx] = `and ${fields[fIdx]}`
    fieldMsg = fields.join(', ')
  }
  if (fields.length === 2) {
    fieldMsg = `${fields[0]} and ${fields[1]}`
  }
  return `${message} Documents include ${fieldMsg}.`
}

export const reduceAddressFields = (
  remediations: ProfileDataItemRemediation[],
) => {
  if (!remediations.length) return 'Residence is used for verification.'
  return 'An address update is required. Please provide your current residential address.'
}

export const reduceContactFields = (
  remediations: ProfileDataItemRemediation[],
) => {
  if (!remediations.length) return 'Communication and security'
  return 'One or more contacts is insufficient. Please update your contact information.'
}

/**
 * Find missing fields and reduce a message for the user.
 */
export const getMissingFieldMessages = (profileItems: {
  [k: string]: ProfileDataItemInfo
}) => {
  const sections = {
    personal: {
      required: new Set(getRequiredPersonalFields()),
      submitted: new Set(),
      missing: new Set(),
      isComplete: false,
      message: '',
    },
    address: {
      required: new Set(getRequiredAddressFields()),
      submitted: new Set(),
      missing: new Set(),
      isComplete: false,
      message: '',
    },
    contact: {
      required: new Set(getRequiredContactFields()),
      submitted: new Set(),
      missing: new Set(),
      isComplete: false,
      message: '',
    },
    document: {
      required: new Set(getRequiredDocumentFields()),
      submitted: new Set(),
      missing: new Set(),
      isComplete: false,
      message: '',
    },
  }

  Object.values(profileItems).forEach(pi => {
    Object.values(sections).forEach(section => {
      if (section.required.has(pi.kind as UserProfileFieldTypes)) {
        section.submitted.add(pi.kind)
      }
    })
  })

  Object.entries(sections).forEach(([sectionName, section]) => {
    section.missing = getDiff(section.required, section.submitted)
    section.isComplete =
      // User hasn't submitted anything for this section
      section.submitted.size === 0 ||
      section.submitted.size === section.required.size
    if (!section.isComplete) {
      const remediations = [...section.missing].map(m => ({
        kind: m,
      })) as ProfileDataItemInfo[]

      if (sectionName === 'personal') {
        section.message = reducePersonalInfoFields(remediations)
      }
      if (sectionName === 'address') {
        section.message = reduceAddressFields(remediations)
      }
      if (sectionName === 'contact') {
        section.message = reduceContactFields(remediations)
      }
      if (sectionName === 'document') {
        section.message = reduceDocumentFields(remediations)
      }
    }
  })

  return sections
}

/**
 * The required personal fields for a complete profile.
 */
export const getRequiredPersonalFields = () => {
  return [
    UserProfileFieldTypes.LEGAL_NAME,
    UserProfileFieldTypes.DATE_OF_BIRTH,
    UserProfileFieldTypes.US_SSN,
  ]
}

/**
 * The required address fields for a complete profile.
 */
export const getRequiredAddressFields = () => {
  return [UserProfileFieldTypes.FULL_ADDRESS]
}

/**
 * The required contact fields for a complete profile.
 */
export const getRequiredContactFields = () => {
  return [UserProfileFieldTypes.EMAIL, UserProfileFieldTypes.PHONE]
}

/**
 * The required document fields for a complete profile.
 */
export const getRequiredDocumentFields = () => {
  return [UserProfileFieldTypes.US_GOVT_DOCUMENT]
}
