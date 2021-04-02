<script context="module" lang="ts">
  export const prerender = true
</script>

<script lang="ts">
  const domain = 'https://snapwallet.io',
    config = {
      theme: {
        modalBackground: '#333',
        color: '#fff',
        textColor: '#fff',
        inputTextColor: '#333',
      },
    }
</script>

<main>
  <div class="col">
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
  </div>
  <div class="col wallet" style="margin: 0 auto;">
    <iframe
      title="Snap Wallet"
      frameborder="0"
      height="608px"
      width="360px"
      src={`${domain}/widget?config="${encodeURI(JSON.stringify(config))}"`}
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
    max-width: 50%;
  }
  iframe {
    position: relative;
    border-radius: 20px;
    top: -20%;
    right: -20%;
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
  @media screen and (max-width: 450px) {
    body {
      overflow-y: scroll;
      main {
        flex-direction: column;
        .wallet {
          max-width: inherit;
          margin-top: 5rem;
          iframe {
            top: inherit;
            right: inherit;
          }
        }
      }
    }
  }
</style>
