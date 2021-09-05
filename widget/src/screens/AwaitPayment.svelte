<script lang="ts">
  import { onMount } from 'svelte'
  import { faArrowDown } from '@fortawesome/free-solid-svg-icons'
  import FaIcon from 'svelte-awesome'
  // @ts-ignore
  import QR from 'qr-creator'
  import { CryptoIcons, formatLocaleCurrency, dropEndingZeros } from '../util'
  import { TransactionIntents, TransactionMediums } from '../types'
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import { ParentMessenger } from '../util/parent_messenger'
  import { configStore } from '../stores/ConfigStore'
  import { transactionStore } from '../stores/TransactionStore'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Surround from '../components/cards/Surround.svelte'
  import Button from '../components/Button.svelte'
  import Clipboard from '../components/Clipboard.svelte'

  // TODO remove demo defaults and use real values
  const dstCurrency = 'BTC', //$transactionStore.destinationCurrency.name ?? 'BTC',
    dstAmount = $transactionStore.destinationAmount ?? 69420,
    dstAddress = '0xDEADBEEF',
    Icon = CryptoIcons[dstCurrency]

  onMount(() => {
    // TODO build-in
    QR.render(
      {
        dstAddress,
        radius: 0.0, // 0.0 to 0.5
        ecLevel: 'H', // L, M, Q, H
        fill: '#111',
        size: 128, // in pixels
      },
      document.getElementById('qrcode'),
    )
  })
</script>

<ModalContent>
  <ModalHeader>Awaiting Payment</ModalHeader>
  <ModalBody>
    <div class="glow" />
    <Surround>
      <h2>Send</h2>
      <div class="row">
        <div class="crypto-icon">
          <Icon size="30" height="30" width="30" viewBox="-4 0 40 40" />
        </div>
        <h4>
          {formatLocaleCurrency(dstCurrency, dstAmount)}
        </h4>
        <Clipboard value={dstAmount} />
      </div>
      <FaIcon class="down-arrow" data={faArrowDown} />
      <div id="qrcode" class="qrcode" title="Scan to Send Payment" />
      <small>Or, Copy & Paste</small>
      <div class="row">
        <p>{dstAddress}</p>
        <Clipboard value={dstAddress} />
      </div>
    </Surround>
  </ModalBody>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';
  :global(.surround) {
    display: flex;
    text-align: center;
  }
  :global(.down-arrow) {
    align-self: center;
  }
  .glow {
    position: relative;
    box-shadow: 0 0 0 0 rgba(var(--theme-button-glow-color), 0.5);
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
  .row {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    font-size: 1.1rem;
    grid-gap: 0.5rem;
    margin-bottom: 0.05rem;
    .crypto-icon {
      margin-top: 7px;
    }
    h4 {
      margin: 0;
    }
  }
  .qrcode {
    margin: 1rem auto;
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
  p {
    margin: 0;
    font-size: 0.8rem;
  }
</style>
