export const isValidNumber = (num: any) => {
  return Number(num) && !isNaN(num) && num !== Infinity
}
