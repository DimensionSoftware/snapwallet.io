<script lang="ts">
  import { push } from 'svelte-spa-router'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/inputs/Input.svelte'
  import Label from '../components/inputs/Label.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import { Routes } from '../constants'
  import { debitCardStore } from '../stores/DebitCardStore'
  import { onMount } from 'svelte'

  let pollTimer
  let cardCode = ''
  let smsCode = ''
  let submittingAuth = false

  $: smsCodeRequired = false
  $: cardCodeRequired = false
  $: isOneCodeRequired = Boolean(smsCodeRequired) || Boolean(cardCodeRequired)

  const handleNextStep = async () => {
    try {
      submittingAuth = true
      await window.API.fluxWyreSubmitDebitCardAuthorizations({
        reservationId: $debitCardStore.reservationId,
        orderId: $debitCardStore.orderId,
        sms2faCode: smsCode,
        card2faCode: cardCode,
      })
      push(Routes.SUCCESS)
    } finally {
      submittingAuth = false
    }
  }

  /**
   * Fetch Wyre debit card authorization codes
   * Both SMS and Card (micro deposit) codes may be required.
   */
  const fetchAuthorizations = async () => {
    const {
      card2faNeeded,
      smsNeeded,
    } = await window.API.fluxWyreGetDebitCardAuthorizations(
      $debitCardStore.orderId,
    )
    smsCodeRequired = smsNeeded
    cardCodeRequired = card2faNeeded
  }

  /**
   * Fetch authorizations regularly
   * until codes are required.
   */
  const pollAuthorizations = () => {
    const t = setInterval(() => {
      // Only one of these may be required
      if (!smsCodeRequired || !cardCodeRequired) {
        fetchAuthorizations()
      } else {
        clearInterval(t)
      }
    }, 4000)
    return t
  }

  onMount(() => {
    pollTimer = pollAuthorizations()
    return () => clearInterval(pollTimer)
  })
</script>

<ModalContent>
  <ModalHeader>Card Authorization</ModalHeader>
  <ModalBody>
    {#if smsCodeRequired}
      <Label label="SMS Code">
        <Input
          id="autocomplete"
          defaultValue={smsCode}
          placeholder="123456"
          on:change={e => (smsCode = e?.detail)}
        />
      </Label>
    {/if}
    {#if cardCodeRequired}
      <Label label="Card Code">
        <Input
          id="autocomplete"
          defaultValue={cardCode}
          placeholder="123456"
          on:change={e => (cardCode = e?.detail)}
        />
      </Label>
    {/if}
    {#if !isOneCodeRequired}
      <p style="padding: 1.5rem">
        Please wait while we authorize your card. This will only take a minute.
      </p>
      <div class="flip-card">
        <div class="flip-card-inner">
          <div class="flip-card-front">
            <div class="cardholder-text">
              <p style="text-transform:capitalize;">
                **** **** **** {$debitCardStore.number.substring(
                  $debitCardStore.number.length - 4,
                )}
                <br />
                {$debitCardStore.firstName}
                {$debitCardStore.lastName}
              </p>
            </div>
            <div class="nfc-chip" />
          </div>
          <div class="flip-card-back">
            <div class="flip-card-stripe" />
            <div class="flip-card-signature" />
          </div>
        </div>
      </div>
    {/if}
  </ModalBody>
  <ModalFooter>
    <Button
      disabled={!isOneCodeRequired}
      isLoading={submittingAuth}
      on:mousedown={handleNextStep}
      >{submittingAuth ? 'Buying...' : 'Buy Now'}</Button
    >
  </ModalFooter>
</ModalContent>

<style lang="scss">
  .flip-card {
    height: 10rem;
    width: 17.5rem;
    perspective: 1000px;
    backface-visibility: hidden;
    border-radius: 0.5rem;
    & > .flip-card-inner {
      border-radius: 0.5rem;
      animation: flipHorizontal 1s infinite;
      position: relative;
      width: 100%;
      height: 100%;
      text-align: center;
      transition: transform 0.8s;
      transform-style: preserve-3d;
      & > .flip-card-front {
        font-family: Courier, monospace;
        background-color: red;
        color: white;
        border-radius: 0.5rem;
        & > .nfc-chip {
          position: absolute;
          height: 2rem;
          width: 3rem;
          background-color: goldenrod;
          top: 3rem;
          left: 1.5rem;
          border-radius: 0.2rem;
        }
        & > .cardholder-text {
          position: absolute;
          width: 100%;
          top: 5rem;
          left: 1rem;
          text-align: left;
        }
      }

      & > .flip-card-back {
        border-radius: 0.5rem;
        transform: rotateY(180deg);
        background-color: red;
        backface-visibility: hidden;
        & > .flip-card-stripe {
          background-color: black;
          position: absolute;
          top: 1.5rem;
          width: 100%;
          height: 20%;
        }
        & > .flip-card-signature {
          position: absolute;
          height: 15%;
          width: 60%;
          top: 4.5rem;
          left: 1rem;
          background-color: white;
        }
      }
    }
  }

  /* Position the front and back of card face */
  .flip-card-front,
  .flip-card-back {
    position: absolute;
    width: 100%;
    height: 100%;
    -webkit-backface-visibility: hidden; /* Safari */
    backface-visibility: hidden;
  }

  @keyframes flipHorizontal {
    from {
      transform: rotateY(0deg);
    }
    to {
      transform: rotateY(359deg);
    }
  }
</style>
