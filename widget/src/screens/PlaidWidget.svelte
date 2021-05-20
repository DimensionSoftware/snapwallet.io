<script lang="ts">
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import { cachePrimaryPaymentMethodID, Logger } from '../util'
  import { onDestroy, onMount } from 'svelte'
  import { paymentMethodStore } from '../stores/PaymentMethodStore'
  import { transactionStore } from '../stores/TransactionStore'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import FaIcon from 'svelte-awesome'
  import {
    faLink,
    faShieldAlt,
    faUniversity,
  } from '@fortawesome/free-solid-svg-icons'

  let handler
  let isCreatingPaymentMethod = false

  type WyreConfig = {
    plaidEnvironment: string
    plaidPublicKey: string
    plaidProducts: string[]
    plaidWebhook: string
  }

  const WYRE_CONFIG_URL = `${__ENV.WYRE_BASE_URL}/v2/client/config/plaid`

  const fetchWyreConfig = async () => {
    const res = await fetch(WYRE_CONFIG_URL)
    return (await res.json()) as WyreConfig
  }

  // async function getLinkToken(): Promise<string> {
  //   const resp = await window.API.fluxPlaidCreateLinkToken({})
  //   return resp.linkToken
  // }

  // Used for connecting SW Plaid accounts
  // async function connectAccounts(
  //   plaidPublicToken: string,
  //   institution: PlaidInstitution,
  //   accounts: PlaidAccount[],
  // ): Promise<void> {
  //   await window.API.fluxPlaidConnectBankAccounts({
  //     plaidPublicToken,
  //     institution,
  //     accounts,
  //   })
  // }

  function connectWyrePaymentMethod(
    plaidPublicToken: string,
    plaidAccountId: string,
  ): Promise<any> {
    return window.API.fluxWyreConnectBankAccount({
      plaidPublicToken,
      plaidAccountId,
    })
  }

  const onComplete = () => {
    paymentMethodStore.fetchWyrePaymentMethods().then(() => {
      const PMs = $paymentMethodStore.wyrePaymentMethods

      if (PMs.length) {
        PMs.sort(pm => (pm.status.toLowerCase() === 'active' ? -1 : 1))
        if (!$transactionStore.selectedSourcePaymentMethod) {
          cachePrimaryPaymentMethodID(PMs[0]?.id)
          transactionStore.setSelectedSourcePaymentMethod(PMs[0])
        }
      }

      isCreatingPaymentMethod = false
      push(Routes.ROOT)
    })
  }

  function initializePlaid() {
    return fetchWyreConfig().then(wyreConfig => {
      handler = window.Plaid.create({
        key: wyreConfig.plaidPublicKey,
        product: wyreConfig.plaidProducts,
        webhook: wyreConfig.plaidWebhook,
        env: wyreConfig.plaidEnvironment,
        selectAccount: true,
        onSuccess: async (_publicToken: string, metadata: any) => {
          Logger.debug(metadata)
          isCreatingPaymentMethod = true
          await connectWyrePaymentMethod(
            metadata.public_token,
            metadata.account_id,
          )
          onComplete()
        },
        onExit: (_err, _metadata) => {
          push(Routes.ROOT)
        },
        onEvent: (eventName, metadata) => {
          const event = eventName.toLowerCase()
          if (event === 'error') {
            Logger.error('Plaid error', metadata)
            throw new Error('A Plaid error occurred. Please contact support.')
          }
          Logger.debug(eventName, metadata)
        },
        // Required to be true for RN
        isWebview: false,
      })

      return Promise.resolve(handler)
    })
  }

  onDestroy(() => handler?.destroy())
  onMount(() => {
    initializePlaid().then(h => h.open())
  })
</script>

{#if isCreatingPaymentMethod}
  <ModalContent>
    <ModalHeader hideBackButton>Linking Bank</ModalHeader>
    <ModalBody>
      <div class="icon-container">
        <FaIcon scale="3" data={faUniversity} />
        <div class="connection">
          <FaIcon scale="1" data={faLink} />
        </div>
        <FaIcon scale="3" data={faShieldAlt} />
      </div>
      <p class="content-txt">
        We're linking your bank account. This should only take a few seconds.
      </p>
    </ModalBody>
  </ModalContent>
{/if}

<style lang="scss">
  @import '../styles/_vars.scss';

  .icon-container {
    margin-top: 3rem;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .content-txt {
    text-align: center;
    padding: 0 0.5rem;
    margin-top: 3rem;
  }

  .connection {
    transform: scale(1);
    animation: pulse 2s infinite;
    height: 2rem;
    width: 2rem;
    border-radius: 50%;
    margin: 0 1rem;
    padding: 0.5rem;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  @keyframes pulse {
    0% {
      transform: scale(0.95);
      box-shadow: 0 0 0 0 rgba(var(--theme-button-glow-color), 0.5);
    }

    70% {
      transform: scale(1);
      box-shadow: 0 0 0 10px rgba(0, 0, 0, 0);
    }

    100% {
      transform: scale(0.95);
      box-shadow: 0 0 0 0 rgba(0, 0, 0, 0);
    }
  }
</style>
