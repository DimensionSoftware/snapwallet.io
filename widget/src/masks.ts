import { MaskTypes } from './types'

enum Masks {
  PHONE = '+x (xxx) xxx-xxxx',
  SSN = 'xxx-xx-xxxx',
  INTL_DATE = 'xxxx-xx-xx',
}

/**
 * Masks any string value.
 */
const maskValue = (val: string, mask: Masks, maskChar: string): string => {
  const minChars = mask.match(/x/g).length
  const splitVal = val.split('')
  const filledChars = splitVal.length
  return [...new Array(minChars)].reduce((acc, _v, index) => {
    if (index > filledChars - 1) return acc.replace(/x/g, maskChar)
    return acc.replace(/x/, splitVal[index])
  }, mask)
}

/**
 *
 * Unmasks a previously masked value.
 *
 */
const unMaskValue = (val: string, mask: Masks): string => {
  const minChars = mask.match(/x/g).length
  const splitVal = val.split('')
  return [...new Array(minChars)].reduce((acc, v, index) => {
    if (mask[index] === 'x' && splitVal[index])
      return `${acc}${splitVal[index]}`
    return acc
  }, '')
}

/**
 * Handles masking values on the fly.
 */
const maskOnType = (val: string, mask: Masks) => {
  val = unMaskValue(val, mask)
  const splitVal = val.split('')
  return splitVal
    .reduce((acc, v) => {
      return acc.replace(/x/, v)
    }, mask)
    .split('x')[0]
    .trim()
}

/**
 * Displays a mask for any string value.
 * The entire mask will always be visible.
 *
 * @param val The value to be masked.
 * @param maskType The mask type to use for the given value.
 */
export const withMask = (val: string = '', maskType: MaskTypes) => {
  if (maskType === MaskTypes.INTL_DATE)
    return maskValue(val, Masks.INTL_DATE, ' ')
  if (maskType === MaskTypes.PHONE) return maskValue(val, Masks.PHONE, ' ')
  if (maskType === MaskTypes.SSN) return maskValue(val, Masks.SSN, ' ')
  return val
}

/**
 * Mask any string value.
 *
 * @param val The value to be masked.
 * @param maskType The mask type to use for the given value.
 */
export const stripMask = (val: string = '', maskType: MaskTypes) => {
  if (maskType === MaskTypes.INTL_DATE) return unMaskValue(val, Masks.INTL_DATE)
  if (maskType === MaskTypes.PHONE) return unMaskValue(val, Masks.PHONE)
  if (maskType === MaskTypes.SSN) return unMaskValue(val, Masks.SSN)
  return val
}

/**
 * Mask any string value while typing.
 * Only the entered values and masking up
 * to that point will be visible.
 *
 * @param val The value to be masked.
 * @param maskType The mask type to use for the given value.
 */
export const withMaskOnInput = (val: string = '', maskType: MaskTypes) => {
  if (maskType === MaskTypes.INTL_DATE) return maskOnType(val, Masks.INTL_DATE)
  if (maskType === MaskTypes.PHONE) return maskOnType(val, Masks.PHONE)
  if (maskType === MaskTypes.SSN) return maskOnType(val, Masks.SSN)
  return val
}
