<script context="module" lang="ts">
  export const prerender = true
</script>

<script lang="ts">
  import { onMount } from 'svelte'
  const domain = 'https://www.snapwallet.io'

  let ifr: HTMLIFrameElement

  onMount(async () => {
    const { default: Snap } = await import('flux-init')

    const SnapWallet = new Snap({
      appName: 'Snap Wallet',
      intent: 'buy',
      wallets: [],
      focus: true,
      theme: {
        modalBackground: '#222',
        modalPopupBackground: '#444',
        color: 'rgb(100,100,100)',
        colorLightened: 'rgba(100,100,100,.8)',
        colorInverse: '#ddd',
        textColor: '#fff',
        inputTextColor: '#333',
      },
    })

    ifr.src = SnapWallet.generateURL()
  })
</script>

<main>
  <div class="intro col">
    <h1>Snap Wallet</h1>
    <h2>Connect Crypto to Your Idea, Simply.</h2>
    <article>
      The "Add Money" button for Crypto Currency, Snap Wallet offers a fully
      configurable, gorgeous interface that delights your customers with a
      single line of code.
    </article>
    <div class="buttons col">
      <a class="button" href={`${domain}/docs/guide`} target="_blank"
        >Get Started</a
      >
      <a href={`https://api.snapwallet.io/swagger`} target="_blank"
        >API Documentation</a
      >
    </div>
    <!-- Features, Advantages, Benefits -->
    <!--ul>
      <li>Fast</li>
      <li>Secure</li>
      <li>Embeddable</li>
      <li>Customs pay in USD; you receive your preferred currency</li>
    </ul-->
  </div>
  <div class="col wallet" style="margin: 0 auto;">
    <iframe
      title="Snap Wallet"
      frameborder="0"
      height="608px"
      width="360px"
      bind:this={ifr}
    />
  </div>
</main>

<style lang="scss">
  @import '../../../widget/src/styles/animations.scss';
  $textColor: #333;
  main {
    display: flex;
    padding: 1em;
    max-width: 960px;
    margin: 15% auto 0;
    .col {
      position: relative;
      max-width: 50%;
      .buttons {
        margin: 5rem 0 0;
        max-width: 100%;
        .button {
          display: inline-block;
          font-weight: 600;
          border-radius: 3px;
          background: rgba($textColor, 0.9);
          color: #fff;
          padding: 1rem 3rem;
          margin-right: 1.5rem;
          transition: background 0.3s ease-out;
          &:hover {
            animation: scaleIn 0.5s ease-out forwards;
            background: $textColor;
            transition: none;
          }
        }
      }
    }
    h1,
    h2 {
      font-size: 2rem;
      font-weight: 500;
      line-height: 1.1;
      margin: 0.5rem 0;
    }
    h2 {
      margin: 2rem 0 0.75rem 0;
      font-size: 1.25rem;
    }
    a {
      color: $textColor;
      text-decoration: none;
      white-space: nowrap;
      margin-bottom: 1rem;
    }
    article {
      font-size: 1.1rem;
      margin: 0;
      line-height: 1.35;
    }
    iframe {
      position: relative;
      border-radius: 20px;
      top: -20%;
      right: -20%;
    }
  }
  @media (min-width: 480px) {
    h1,
    h2 {
      max-width: none;
    }
    article {
      max-width: none;
    }
  }

  // responsive
  @media (max-width: 375px) {
    :global(body) {
      overflow-y: scroll;
      main {
        padding: 0;
      }
    }
  }
  @media (max-width: 850px) {
    :global(body) {
      overflow-y: scroll;
      main {
        flex-direction: column;
        > .col {
          max-width: 100%;
        }
        .intro {
          padding-left: 2rem;
          padding-right: 1rem;
          article {
            max-width: none;
          }
        }
        .wallet {
          max-width: inherit;
          padding-top: 5rem;
          padding-bottom: 2rem;
          iframe {
            top: inherit;
            right: inherit;
          }
        }
      }
    }
  }
</style>
