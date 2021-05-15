export const transactionsAsDataURI = (transactions: any[]) => {
  const rows = transactions
    // TODO: allow only completed txns
    // .filter(txn => txn.status === 'COMPLETED')
    .map(txn => [
      txn.id,
      txn.createdAt,
      txn.source,
      txn.sourceCurrency,
      txn.sourceAmount,
      txn.dest,
      txn.destCurrency,
      txn.destAmount,
      txn.status,
    ])
  // Add CSV headers
  rows.unshift([
    'ID',
    'Created Timestamp',
    'Source',
    'Source Currency',
    'Source Amount',
    'Destination',
    'Destination Currency',
    'Destination Amount',
    'Status',
  ])
  const dataURI =
    'data:text/csv;charset=utf-8,' + rows.map(e => e.join(',')).join('\n')
  return encodeURI(dataURI)
}

export const computeTransactionExpiration = (expiresAt?: string) => {
  const expiration = new Date(expiresAt) as any
  const now = new Date() as any
  return (expiration - now) / 1000
}

export const formatExpiration = secondsUntilExpiration => {
  const mins: number = Math.floor(secondsUntilExpiration / 60)
  const seconds: number = Math.floor(secondsUntilExpiration % 60)
  const displaySeconds: string =
    seconds > 9 ? seconds.toString() : `0${seconds}`
  return mins > 1 ? `${mins}m ${displaySeconds}s` : `${displaySeconds}s`
}
