<script lang="ts">
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Button from '../components/Button.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { TransactionIntents } from '../types'
  import { ParentMessenger } from '../util/parent_messenger'

  export let product

  $: ({ intent, wyrePreview } = $transactionStore)

  $: ({
    sourceCurrency,
    destCurrency: destinationCurrency,
    destAmount: destinationAmount,
  } = wyrePreview)
  $: isBuy = intent === TransactionIntents.BUY
  $: cryptoTicker = isBuy ? destinationCurrency : sourceCurrency

  const done = () => {
    ParentMessenger.exit()
    // if within a model, let that close first
    setTimeout(() => push(Routes.ROOT), 250)
  }
</script>

<ModalContent>
  <ModalHeader hideBackButton>Success</ModalHeader>
  <ModalBody>
    <div class="icon-box">
      <svg
        id="successAnimation"
        class="animated"
        xmlns="http://www.w3.org/2000/svg"
        width="125"
        height="125"
        viewBox="0 0 70 70"
      >
        <path
          id="successAnimationResult"
          fill="#D8D8D8"
          d="M35,60 C21.1928813,60 10,48.8071187 10,35 C10,21.1928813 21.1928813,10 35,10 C48.8071187,10 60,21.1928813 60,35 C60,48.8071187 48.8071187,60 35,60 Z M23.6332378,33.2260427 L22.3667622,34.7739573 L34.1433655,44.40936 L47.776114,27.6305926 L46.223886,26.3694074 L33.8566345,41.59064 L23.6332378,33.2260427 Z"
        />
        <circle
          id="successAnimationCircle"
          cx="35"
          cy="35"
          r="24"
          stroke="#979797"
          stroke-width="2"
          stroke-linecap="round"
          fill="transparent"
        />
        <polyline
          id="successAnimationCheck"
          stroke="#979797"
          stroke-width="2"
          points="23 34 34 43 47 27"
          fill="transparent"
        />
      </svg>
    </div>
    <div class="text-center">
      {#if product}
        <p>
          You have successfully confirmed your transfer of <b
            >{destinationAmount} {cryptoTicker}</b
          > for
        </p>
        <p>
          <b>{product.title}</b>
        </p>
      {:else}
        <p>Your <b>{cryptoTicker}</b> checkout is confirmed!</p>
      {/if}
      <p>
        Please allow up to five (5) business days for your purchase to complete.
      </p>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button on:mousedown={done}>Done</Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';

  $circle-length: 151px;
  $check-length: 36px;

  .icon-box {
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    color: var(--theme-success-color);
    margin-bottom: 0.5rem;
  }
  .text-center {
    margin: 0 1rem;
  }
  b {
    text-transform: uppercase;
  }

  @keyframes scaleAnimation {
    0% {
      opacity: 0;
      transform: scale(3);
    }
    100% {
      opacity: 1;
      transform: scale(1);
    }
  }
  @keyframes drawCircle {
    0% {
      stroke-dashoffset: $circle-length;
    }
    100% {
      stroke-dashoffset: 0;
    }
  }
  @keyframes drawCheck {
    0% {
      stroke-dashoffset: $check-length;
    }
    100% {
      stroke-dashoffset: 0;
    }
  }
  @keyframes fadeOut {
    0% {
      opacity: 1;
    }
    100% {
      opacity: 0;
    }
  }
  @keyframes fadeIn {
    0% {
      transform: scale(1);
      opacity: 0;
    }
    75% {
      opacity: 1;
      transform: scale(2);
    }
    100% {
      opacity: 1;
      transform: scale(1);
    }
  }
  #successAnimationCircle {
    stroke-dasharray: $circle-length $circle-length;
    stroke: var(--theme-success-color);
  }
  #successAnimationCheck {
    stroke-dasharray: $check-length $check-length;
    stroke: var(--theme-success-color);
  }
  #successAnimationResult {
    fill: var(--theme-success-color);
    opacity: 0;
  }
  #successAnimation.animated {
    animation: 0.9s var(--theme-ease-out-expo) 0s 1 both scaleAnimation;
    #successAnimationCircle {
      animation: 1s cubic-bezier(0.77, 0, 0.175, 1) 0s 1 both drawCircle,
        0.3s ease 0.9s 1 both fadeOut;
    }
    #successAnimationCheck {
      animation: 1s cubic-bezier(0.77, 0, 0.175, 1) 0s 1 both drawCheck,
        0.3s ease 0.9s 1 both fadeOut;
    }
    #successAnimationResult {
      animation: 0.3s var(--theme-ease-out-back) 0.8s both fadeIn;
    }
  }
</style>
