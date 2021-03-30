export const exportTransactionsAsCSV = (transactions: any[]) => {
  const rows = transactions
    .filter(txn => txn)
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
  window.location.href = encodeURI(dataURI)
}
