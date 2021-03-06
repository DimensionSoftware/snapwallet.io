<script lang="ts">
  import { push } from 'svelte-spa-router'
  import { Routes } from '../constants'
  async function getLinkToken(): Promise<string> {
    const resp = await window.API().fluxPlaidCreateLinkToken({})

    return resp.linkToken
  }

  async function connectAccounts(
    plaidPublicToken: string,
    plaidAccountIds: string[],
  ): Promise<void> {
    await window.API().fluxPlaidConnectBankAccounts({
      plaidPublicToken,
      plaidAccountIds,
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

  function initializePlaid() {
    getLinkToken().then(token => {
      const handler = window.Plaid.create({
        token,
        onSuccess: (
          publicToken: string,
          metadata: PlaidSuccessCallbackMetadata,
        ) => {
          console.log(metadata)
          connectAccounts(
            publicToken,
            metadata.accounts.map(a => a.id),
          ).then(() => {
            setTimeout(() => push(Routes.PROFILE), 700)
          })
        },
        onLoad: () => {},
        onExit: (err, metadata) => {
          handler.destroy()
          push('/')
        },
        onEvent: (eventName, metadata) => {},
        receivedRedirectUri: null,
      })
      handler.open()
    })

    /* handler.destroy() <-- cleanup function for plaid */
  }
</script>

<svelte:head>
  <script
    src="https://cdn.plaid.com/link/v2/stable/link-initialize.js"
    on:load={initializePlaid}></script>
</svelte:head>

<style lang="scss">
  @import '../styles/_vars.scss';
</style>
