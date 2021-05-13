import vld8 from 'validator'

// Allows different variations of names
const NON_NUMERIC_STRING = /^[ '\p{L}-]+$/u
const EXP_MONTH = /[0-9]{2}/
const EXP_YEAR = /[0-9]{4}/
const CVC_REGEX = /^([0-9]{3}|[0-9]{4})$/

interface IValidationRules {
  [field: string]: {
    validate: (value?: any) => boolean
    errorMessage: (field: string, value?: any) => string
  }
}

interface IValidationValues {
  [field: string]: string
}

/**
 * Given rules and values this function validates a form.
 * Returns an error if available and a boolean to check validity.
 * @param rules The rules for validation.
 * @param values The values for the set rules.
 */
export const validateForm = (
  rules: IValidationRules = {},
  values: IValidationValues,
): { isValid: boolean; error: string | null } => {
  return Object.entries(rules).reduce(
    (acc, [field, opts]) => {
      // Exit on first error
      if (acc.error || !acc.isValid) return acc
      if (opts.validate) {
        if (!opts.validate(values[field])) {
          return {
            isValid: false,
            error: opts.errorMessage(field, values[field]),
          }
        }
      }
      return acc
    },
    { error: null, isValid: true },
  )
}

/**
 * The debit card form validation rules.
 */
export const debitCardValidationRules: IValidationRules = {
  firstName: {
    validate: firstName => NON_NUMERIC_STRING.test(firstName),
    errorMessage: () => 'Please enter a valid first name',
  },
  lastName: {
    validate: lastName => NON_NUMERIC_STRING.test(lastName),
    errorMessage: () => 'Please enter a valid last name',
  },
  phoneNumber: {
    validate: phoneNumber => vld8.isMobilePhone(phoneNumber),
    errorMessage: () => 'Please enter a valid phone number',
  },
  cardNumber: {
    validate: cardNumber => vld8.isCreditCard(cardNumber),
    errorMessage: () => 'Please enter a valid card number',
  },
  cardExpiration: {
    validate: cardExp => {
      const [month, year] = cardExp.split('/')
      const expYear = Number(`20${year}`)
      const expMonth = Number(month)
      const now = new Date()
      const yearNow = Number(now.getFullYear())
      const monthNow = Number(now.getMonth()) + 1

      const isMissing =
        !EXP_MONTH.test(expMonth.toString()) ||
        !EXP_YEAR.test(expYear.toString())
      const isYearExpired = yearNow > expYear
      const isMonthExpired = monthNow > expMonth && yearNow === expYear

      return !isMissing && !isYearExpired && !isMonthExpired
    },
    errorMessage: () => 'Please enter a valid card expiration date',
  },
  cardVerificationCode: {
    validate: cvc => CVC_REGEX.test(cvc),
    errorMessage: () => 'Please enter a valid card verification code.',
  },
}
