<script lang="ts">
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  import { cachePrimaryPaymentMethodID, Logger } from '../util'
  import { onDestroy, onMount } from 'svelte'
  import type { PlaidAccount, PlaidInstitution } from 'api-client'
  import { paymentMethodStore } from '../stores/PaymentMethodStore'
  import { transactionStore } from '../stores/TransactionStore'

  let handler

  async function getLinkToken(): Promise<string> {
    const resp = await window.API.fluxPlaidCreateLinkToken({})
    return resp.linkToken
  }

  async function connectAccounts(
    plaidPublicToken: string,
    institution: PlaidInstitution,
    accounts: PlaidAccount[],
  ): Promise<void> {
    await window.API.fluxPlaidConnectBankAccounts({
      plaidPublicToken,
      institution,
      accounts,
    })
  }

  interface PlaidSuccessCallbackMetadataAccount {
    id: string
    name: string // 'Plaid Checking'
    mask: string // '0000'
    type: string // 'depository'
    subtype: string // 'checking'
  }

  interface PlaidSuccessCallbackMetadataInstitution {
    institution_id: string
    name: string
  }

  interface PlaidSuccessCallbackMetadata {
    institution: PlaidSuccessCallbackMetadataInstitution
    accounts: PlaidSuccessCallbackMetadataAccount[]
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
    return getLinkToken().then(token => {
      handler = window.Plaid.create({
        token,
        onSuccess: (
          publicToken: string,
          metadata: PlaidSuccessCallbackMetadata,
        ) => {
          Logger.debug(metadata)
          connectAccounts(
            publicToken,
            {
              id: metadata.institution.institution_id,
              name: metadata.institution.name,
            },
            metadata.accounts.map(
              (pa: PlaidSuccessCallbackMetadataAccount) => ({
                id: pa.id,
                name: pa.name,
                mask: pa.mask,
                type: pa.type,
                subType: pa.subtype,
              }),
            ),
          ).then(onComplete)
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
