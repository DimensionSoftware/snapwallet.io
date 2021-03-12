import type { Masks } from './types'

/**
 * Displays a mask for any string value.
 * The entire mask will always be visible.
 *
 * @param val The value to be masked.
 * @param maskType The mask type to use for the given value.
 * @param maskChar The character to use for masking.
 */
export const maskValue = (
  val: string,
  mask: Masks,
  maskChar: string,
): string => {
  const minChars = mask.match(/x/g).length
  const splitVal = val.split('')
  const filledChars = splitVal.length
  return [...new Array(minChars)].reduce((acc, _v, index) => {
    if (index > filledChars - 1) return acc.replace(/x/g, maskChar)
    return acc.replace(/x/, splitVal[index])
  }, mask)
}

/**
 * Unmask a previously masked value.
 *
 * @param val The value to be unmasked.
 * @param maskType The mask to use for unmasking the given value.
 */
export const unMaskValue = (val: string, mask: Masks): string => {
  if (!val || !mask) return val
  const minChars = mask.match(/x/g).length
  const splitVal = val.split('')
  return [...new Array(minChars)].reduce((acc, v, index) => {
    if (mask[index] === 'x' && splitVal[index])
      return `${acc}${splitVal[index]}`
    return acc
  }, '')
}

/**
 * Mask any string value while typing.
 * Only the entered values and masking up
 * to that point will be visible.
 *
 * @param val The value to be masked.
 * @param maskType The mask type to use for the given value.
 */
export const withMaskOnInput = (val?: string, mask?: Masks) => {
  if (!mask || !val) return val
  return unMaskValue(val, mask)
    .split('')
    .reduce((acc, v) => {
      return acc.replace(/x/, v)
    }, mask)
    .split('x')[0]
    .trim()
}
