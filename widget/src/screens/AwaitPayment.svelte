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
  import { CryptoIcons, formatLocaleCurrency, dropEndingZeros } from '../util'
  import { Routes } from '../constants'
  import { transactionStore } from '../stores/TransactionStore'
  import { configStore } from '../stores/ConfigStore'
  import Balance from '../components/Balance.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import Metamask from '../components/Metamask.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Surround from '../components/cards/Surround.svelte'
  import Clipboard from '../components/Clipboard.svelte'
  import Address from '../components/Address.svelte'
  import { TransactionMediums } from '../types'
  import AccountSelector from '../components/selectors/AccountSelector.svelte'
  import { formatExpiration } from '../util/transactions'

  let isPaymentSelectorVisible = false

  $: checkAccount = $selectedAccount
  $: isDebitCard = $transactionStore.inMedium === TransactionMediums.DEBIT_CARD
  $: formattedExpiration = formatExpiration(
    $transactionStore.transactionExpirationSeconds,
  )
  $: isMetamask = false

  const //{ destinationCurrency } = $transactionStore,
    { destCurrency, destAddress, destAmount } =
      $transactionStore.wyrePreview || {
        destCurrency: 'BTC',
        destAddress: '0xYOUR_WALLET',
        destAmount: 0,
      },
    Icon = CryptoIcons[destCurrency ?? 'BTC']

  onMount(() => {
    // render qrcode
    QR.render(
      {
        text: destAddress ?? '',
        radius: 0.0, // 0.0 to 0.5
        ecLevel: 'H', // L, M, Q, H
        fill: '#111',
        size: 115, // in pixels
      },
      document.getElementById('qrcode'),
    )
    // enable web3
    isMetamask = typeof web3 !== 'undefined'
    // // TODO use metamask.svelte
    // try {
    //   defaultEvmStores.setProvider()
    // } catch (e) {
    //   console.log(e)
    // }
  })

  const doSuccess = e => {
    e.preventDefault()
    push(Routes.CART_SUCCESS)
    return false
  }
</script>

<ModalContent>
  <ModalHeader>Awaiting Payment</ModalHeader>
  <ModalBody klass="awaiting-payment">
    <div class="expires">
      <FaIcon data={faClock} />
      <div style="margin-right:0.35rem;" />
      <b>{formattedExpiration}</b>
    </div>
    <Surround glow>
      {#if isMetamask}
        <Metamask />
      {:else}
        <h2>Scan to Send</h2>
        <div class="row">
          <h4 class="amount">
            {formatLocaleCurrency(destCurrency, destAmount)}
          </h4>
          <Clipboard value={destAmount} />
        </div>
        <div id="qrcode" class="qrcode" title="Scan to Send Payment">
          <div class="crypto-icon">
            <Icon size="50" height="50" width="50" viewBox="-4 0 40 40" />
          </div>
        </div>
        <div class="row">
          <Address address={destAddress} />
          <Clipboard value={destAddress} />
        </div>
      {/if}
    </Surround>
    <div
      class="payment"
      title="Click to Change Payment Method"
      on:click={() => (isPaymentSelectorVisible = true)}
    >
      <b>
        <!-- Multiple PMs will be possible for buy and bank account is only option for sell atm -->
        {#if $transactionStore.selectedSourcePaymentMethod}
          {$transactionStore.selectedSourcePaymentMethod.name}
        {:else if isDebitCard}
          Pay with Debit Card
        {:else}
          Change Payment Method
        {/if}
      </b>
    </div>
  </ModalBody>
  {#if isPaymentSelectorVisible}
    <AccountSelector
      visible
      on:close={() => (isPaymentSelectorVisible = false)}
    />
  {/if}
  <h3 class="test">
    {#if $configStore.environment === 'sandbox'}
      <a on:click={doSuccess}>Test Success</a>
    {/if}
  </h3>
</ModalContent>

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
