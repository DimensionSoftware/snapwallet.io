<script lang="ts">
  import { onMount } from 'svelte'
  // @ts-ignore
  import QR from 'qr-creator'
  import { faClock } from '@fortawesome/free-solid-svg-icons'
  import { push } from 'svelte-spa-router'
  import FaIcon from 'svelte-awesome'
  import {
    connected,
    web3,
    selectedAccount,
    chainId,
    chainData,
    defaultEvmStores,
  } from 'svelte-web3'
  import {
    CryptoIcons,
    formatLocaleCurrency,
    dropEndingZeros,
    walletForTicker,
  } from '../util'
  import { Routes } from '../constants'
  import { transactionStore } from '../stores/TransactionStore'
  import { configStore } from '../stores/ConfigStore'
  import Balance from '../components/Balance.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Surround from '../components/cards/Surround.svelte'
  import Clipboard from '../components/Clipboard.svelte'
  import Address from '../components/Address.svelte'
  import { TransactionMediums } from '../types'
  import AccountSelector from '../components/selectors/AccountSelector.svelte'
  import { formatExpiration } from '../util/transactions'
  import Button from './Button.svelte'

  let isPaymentSelectorVisible = false

  $: checkAccount = $selectedAccount
  $: isDebitCard = $transactionStore.inMedium === TransactionMediums.DEBIT_CARD
  $: formattedExpiration = formatExpiration(
    $transactionStore.transactionExpirationSeconds,
  )

  const //{ destinationCurrency } = $transactionStore,
    { destCurrency, destAddress, destAmount } =
      $transactionStore.wyrePreview || {
        destCurrency: 'ETH',
        destAddress: '0xYOUR_WALLET',
        destAmount: 0,
      },
    Icon = CryptoIcons[destCurrency ?? 'ETH']

  onMount(async () => {
    // render qrcode
    // QR.render(
    //   {
    //     text: destAddress ?? '',
    //     radius: 0.0, // 0.0 to 0.5
    //     ecLevel: 'H', // L, M, Q, H
    //     fill: '#111',
    //     size: 115, // in pixels
    //   },
    //   document.getElementById('qrcode'),
    // )
    // automagically connect web3
    if (typeof web3 !== 'undefined') connect()
  })

  const transact = async e => {
    console.log($configStore.wallets)
    const tx = await $web3.eth.sendTransaction({
      // gasPrice: $web3.utils.toHex($web3.utils.toWei('30', 'gwei')),
      gasLimit: $web3.utils.toHex('21000'),
      from: $selectedAccount,
      to: walletForTicker($configStore.wallets, destCurrency, {
        isTest: $configStore.environment !== 'production',
      }),
      value: $web3.utils.toHex(destAmount * 1000000000000000000), // eth to wei
    })
    console.log('receipt from transaction', tx)
    doSuccess(e)
  }

  let pending = false,
    type = 'Browser'
  const connect = async () => {
    pending = true
    try {
      await defaultEvmStores.setProvider()
      console.log('$selectedAccount', defaultEvmStores.$selectedAccount)
      console.log(web3.$accounts, defaultEvmStores.web3.$accounts, web3.$signer)
      pending = false
    } catch (e) {
      console.log(e)
      pending = false
      throw new Error(e)
    }
  }

  const disconnect = async () => {
    await defaultEvmStores.disconnect()
    pending = false
  }

  const doSuccess = e => {
    e.preventDefault()
    push(Routes.CART_SUCCESS)
    return false
  }
</script>

<h2>
  {$connected
    ? `${formatLocaleCurrency(destCurrency, destAmount)}`
    : 'Connect a Wallet'}
</h2>
<div class="col">
  {#if !$connected}
    <Button title="Connect Button" on:mousedown={connect}>Connect</Button>
  {:else}
    <Balance address={checkAccount} amount={destAmount} />
    <div style="height: 2.5rem;" />
    <Button title="Pay" on:mousedown={connect}>Pay</Button>
  {/if}
</div>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';
  :global(.modal-body.awaiting-payment) {
    padding: 0.75rem 3.25rem !important;
  }
  :global(.modal-body.awaiting-payment .clipboard-copy) {
    position: absolute;
    top: 3px;
    right: 0.5rem;
  }
  :global(.modal-body.awaiting-payment .surround) {
    display: flex;
    text-align: center;
    border-width: 4px !important;
    margin-top: 25%;
    padding-bottom: 1.25rem !important;
  }
  :global(.down-arrow) {
    align-self: center;
  }
  .glow {
    position: relative;
    box-shadow: 0 0 0 0 rgba(var(--theme-button-glow-color), 0.75);
    animation: glow 1.5s linear;
    animation-iteration-count: infinite;
    border-radius: 100%;
    top: -5px;
    height: 4px;
    width: 4px;
    background: transparent;
    margin: 0.25rem auto;
    &:before {
      position: absolute;
      content: '';
      width: 4px;
      left: -1px;
      top: -1px;
      height: 4px;
      border-radius: 100%;
      border: 1px solid var(--theme-color);
      opacity: 0.5;
    }
  }
  h2 {
    font-size: 1.2rem;
  }
  h3.test {
    margin-bottom: 2rem !important;
  }
  .crypto-icon {
    position: absolute;
    top: 40px;
    left: 35px;
    filter: grayscale(100%) contrast(150%);
    margin-bottom: 0.5rem;
  }
  .col {
    display: flex;
    flex-direction: column;
    padding: 0 1rem;
  }
  :global(.surround .amount) {
    margin: 0;
  }
  .row {
    display: flex;
    position: relative;
    flex-direction: row;
    justify-content: center;
    align-items: flex-start;
    font-size: 1.1rem;
    grid-gap: 0.5rem;
    margin-bottom: 0.05rem;
    h4 {
      margin: 0 0 0.5rem;
      max-width: 160px;
      white-space: pre-wrap;
      line-height: 1.25rem;
      text-align: left;
      word-break: break-word;
      text-overflow: ellipsis;
      overflow: hidden;
      &.amount {
        max-width: 130px;
        margin: 0;
      }
    }
  }
  .qrcode {
    position: relative;
    background: var(--theme-color);
    padding: 3px;
    padding-bottom: 1px;
    border-radius: 3px;
    margin: 0.2rem auto 0.75rem;
  }
  small {
    margin: 1.5rem 0 0.5rem 0;
    font-weight: 600;
    position: relative;
    &:after,
    &:before {
      position: absolute;
      content: '';
      top: 10px;
      width: 20px;
      height: 1px;
      opacity: 0.5;
    }
    &:after {
      margin-left: 5px;
      background-image: linear-gradient(
        to left,
        transparent,
        var(--theme-color)
      );
    }
    &:before {
      margin-left: -25px;
      background-image: linear-gradient(
        to right,
        transparent,
        var(--theme-color)
      );
    }
  }
  h2 {
    font-size: 1.1rem;
    margin: 2rem;
  }
  p {
    margin: 0;
    font-size: 0.8rem;
  }

  .payment {
    display: flex;
    display: none;
    cursor: pointer;
    opacity: 0.8;
    font-size: 0.8rem;
    margin: 3rem auto;
    gap: 0.5rem;
  }

  .expires {
    display: flex;
    justify-content: center;
    align-items: center;
    opacity: 0.8;
  }
</style>
