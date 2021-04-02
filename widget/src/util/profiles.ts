import type {
  ProfileDataItemKind,
  ProfileDataItemRemediation,
} from 'api-client'
import { UserProfileFieldTypes } from '../constants'
import type { RemediationGroups } from '../types'

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
  let message = 'Your personal information requires an update.'
  let fields = []

  remediations.forEach(r => {
    if (UserProfileFieldTypes.LEGAL_NAME === r.kind) fields.push('legal name')
    if (UserProfileFieldTypes.DATE_OF_BIRTH === r.kind) fields.push('birthdate')
    if (UserProfileFieldTypes.US_SSN === r.kind)
      fields.push('social security number')
  })

  // An unsupported field made it here.
  if (!fields.length) return `${message} Please contact support.`
  if (fields.length >= 2) {
    const fIdx = fields.length - 1
    fields[fIdx] = `and ${fields[fIdx]}`
  }
  return `${message} Fields include ${fields.join(', ')}.`
}
