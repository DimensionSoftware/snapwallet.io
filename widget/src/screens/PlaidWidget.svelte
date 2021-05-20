<script lang="ts">
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import { cachePrimaryPaymentMethodID, Logger } from '../util'
  import { onDestroy, onMount } from 'svelte'
  import { paymentMethodStore } from '../stores/PaymentMethodStore'
  import { transactionStore } from '../stores/TransactionStore'

  let handler
  let wyreConfig

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
        // Required for RN
        isWebview: true,
      })

      return Promise.resolve(handler)
    })
  }

  onDestroy(() => handler?.destroy())
  onMount(() => {
    initializePlaid().then(h => h.open())
  })
</script>

<style lang="scss">
  @import '../styles/_vars.scss';
</style>
