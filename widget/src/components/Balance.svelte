<script>
  import { web3, connected } from 'svelte-web3'
  export let address
  export let amount
  export let pending = '...'
  // todo format + symbol
  $: balance = $connected && address ? $web3.eth.getBalance(address) : ''

  const humanBalance = balance => {
    console.log('bal', balance)
    if (!balance || balance === '0')
      return 'Please select another wallet with funds.'
    return balance < amount
      ? `Insufficient funds in connected wallet.  You may connect and transfer balances from multiple wallets.`
      : `Remaining wallet balance after transfer:  ${
          (amount - balance) / 1e18
        } ETH`
    // return $web3.fromWei(balance, 'ether')
  }
</script>

{#if address}{#await balance}{pending}{:then value}{humanBalance(
      value,
    )}{/await}{/if}
