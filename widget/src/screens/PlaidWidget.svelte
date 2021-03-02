<script lang="ts">
  async function getLinkToken(): Promise<string> {
    const resp = await window.API().fluxPlaidCreateLinkToken({})

    return resp.linkToken
  }

  function initializePlaid() {
    getLinkToken().then(token => {
      const handler = window.Plaid.create({
        token,
        environment: 'sandbox',
        onSuccess: (public_token, metadata) => {},
        onLoad: () => {},
        onExit: (err, metadata) => {},
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
