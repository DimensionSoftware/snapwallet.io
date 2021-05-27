<script lang="ts">
  import Feature from './Feature.svelte'
  import Button from './Button.svelte'
  const scrollToTop = _ =>
      window.scrollTo({
        top: 0,
        left: 0,
        behavior: 'smooth',
      }),
    contactUs = _ => {
      window.location.href =
        'mailto:support@snapwallet.com?subject=Hello Snap Wallet Team!'
    }
  const cards = [
    {
      title: 'Enter Amount',
      name: 'nft',
      icon: 'SW Crypto Coin.png',
      alt: 'Simply Checkout with Anything',
    },
    {
      title: 'Select a Payment Method',
      name: 'donations',
      icon: 'Dr Crypto clean shaven.png',
      alt: 'Accepts Credit Cards and Debit Accounts',
    },
    {
      title: 'Transaction Completed!',
      name: 'fiat',
      icon: 'Coin Checkmark.png',
      alt: 'On and Off-Ramp Crypto Currencies',
    },
  ]
</script>

<Feature
  name="started"
  title="Get started in just a few steps"
  description="Snap Wallet is simple, secure and fast."
  center
>
  <div class="cards flex" slot="left">
    {#each cards as card, i}
      <article on:mousedown={_ => scrollTo(card.name)}>
        <img
          title={card.alt}
          width="100"
          src={`/images/${card.icon}`}
          alt={card.title}
        />
        <h4>{card.title}</h4>
        {#if i !== cards.length - 1}
          <span />
        {/if}
      </article>
    {/each}
  </div>
</Feature>
<div class="contact">
  <Button id="contact" on:mousedown={contactUs}>Contact Us</Button>
</div>
<hr />
<footer>
  <ol>
    <li><h2>Snap Wallet</h2></li>
    <li>
      <h4>
        <a
          title="Snap Wallet Integration Documentation"
          href="https://snapwallet.io/docs/guide"
          target="_blank">API Documentation</a
        >
      </h4>
    </li>
    <li>
      <h4>
        <a
          title="Dimension Software on Silicon Beach"
          href="https://dimensionsoftware.com"
          target="_blank">Company</a
        >
      </h4>
    </li>
    <li>
      <h4>
        <a title="Snap Wallet Support" href="mailto:support@snapwallet.io"
          >Support</a
        >
      </h4>
    </li>
    <li>
      <h4>
        <a
          title="Snap Wallet Never Shares Your Information"
          href="https://login.dimensionsoftware.com/privacy"
          target="_blank">Privacy</a
        >
      </h4>
    </li>
  </ol>
  <p title="Scroll to Top!" on:mousedown={scrollToTop}>
    <big
      >Snap Wallet
      <img height="24px" width="24px" title="Love" alt="Love" src="/love.svg" />
      Silicon Beach, CA
    </big>
    <small
      ><a
        title="Dimension Software on Silicon Beach, Los Angeles!"
        href="https://dimensionsoftware.com"
        on:mousedown|stopPropagation
        target="_blank">Dimension Software since 1998</a
      ></small
    >
  </p>
</footer>

<style lang="scss">
  @import '../../../widget/src/styles/animations.scss';
  @import '../../../widget/src/styles/_vars.scss';
  .cards {
    display: flex;
    position: relative;
    flex-direction: row;
    article {
      margin: 0 7rem;
      display: flex;
      flex-direction: column;
      align-items: center;
      width: 150px;
      h4 {
        font-weight: 300;
        text-align: center;
        margin-top: 1rem;
        font-size: 1.5rem;
        white-space: nowrap;
      }
      & > span {
        position: absolute;
        height: 1px;
        right: -100%;
        background: #ccc;
        width: 100px;
        top: 35%;
      }
    }
  }
  hr {
    height: 5rem;
    background: #fff;
    margin: 0;
    border: none;
    box-shadow: 0 50px 30px 0 rgba(0, 0, 0, 0.5);
  }
  :global(.started article) {
    position: relative;
    width: 100%;
  }
  .contact {
    position: absolute;
    z-index: 2;
    left: calc(50% - 119px);
  }
  :global(.contact > button) {
    color: #fff !important;
    font-size: 1.15rem !important;
    background: rgb(241, 7, 28) !important;
    background: linear-gradient(
      63.26deg,
      #ff6b00,
      rgb(241, 7, 28) 75%
    ) !important;
    border-color: rgb(222, 49, 45) !important;
    top: -4.25rem;
    position: relative;
    z-index: 1;
    padding: 1rem 4rem;
    left: auto;
    right: auto;
    margin: 0 auto;
    text-align: center;
    border-radius: 4.25rem;
    transition: box-shadow 0.3s ease-out;
    &:hover {
      transition: none;
      box-shadow: 0 0 0 3px rgba(222, 49, 45, 0.25) !important;
    }
  }
  :global(body) {
    margin-bottom: 433px !important;
  }
  footer {
    position: fixed;
    z-index: -1;
    bottom: 0;
    left: 0;
    right: 0;
    margin-top: 0.05rem;
    overflow: hidden;
    padding: 1.5rem 0 1rem;
    background: #000;
    opacity: 0;
    animation: backgroundFadeIn 0.5s ease-in 2s forwards;
    color: rgba(255, 255, 255, 0.8);
    &:before {
      content: '';
      position: absolute;
      top: 0;
      right: 0;
      left: 0;
      bottom: 0;
      opacity: 0.075;
      background-image: url('/static/bg.png');
    }
    ol {
      position: relative;
      z-index: 1;
      list-style: none;
      margin: 9rem 0 5rem 10rem;
      li {
        font-size: 0.9rem;
        margin: 0;
        padding: 0;
        h2 {
          padding: 0.25rem 0.75rem;
          color: #fff;
        }
        h4 {
          margin: 0 0 0.75rem 0;
        }
      }

      a {
        position: relative;
        display: inline-block;
        padding: 0.4rem 0.75rem;
        margin: -0.4rem 0 0 0;
        border-radius: 3rem;
        z-index: 1;
        color: rgba(255, 255, 255, 0.7);
        text-decoration: none;
        overflow: hidden;
        transition: transform 0.3s $easeOutExpo 0.1s,
          background-color 0.2s $easeInExpo;
        &:before {
          content: '';
          position: absolute;
          top: 0;
          right: 0;
          left: 0;
          bottom: 0;
          opacity: 0;
          z-index: -1;
          background: linear-gradient(
            63.26deg,
            #ff6b00,
            #fffc00 75%
          ) !important;
          transform: translateX(-100%) scale(0);
          border-radius: 1rem;
          transition: opacity 0.3s ease-out, transform 0.3s ease-in;
        }
        &:hover {
          color: #000;
          box-shadow: 0 0 0 3px rgba(222, 49, 45, 0.5);
          font-weight: bold;
          transform: translateX(2px) scale(1.05);
          transition: none;
          &:before {
            opacity: 1;
            transform: translateX(0) scale(1);
            transition: none;
          }
        }
      }
    }
    p {
      position: relative;
      z-index: 1;
      font-size: 0.85rem;
      margin: 0;
      padding: 0;
      text-align: center;
      vertical-align: middle;
      cursor: pointer;
      small {
        display: block;
      }
      big {
        display: block;
        margin-bottom: 0.75rem;
        font-size: 0.9rem;
        color: rgba(255, 255, 255, 0.8);
      }
      img {
        position: relative;
        top: 8px;
        margin: 0 0.25rem;
      }
      br {
        display: block;
        margin-top: 0.5rem;
      }
      a {
        text-decoration: none;
        color: rgba(255, 255, 255, 0.8);
      }
    }
  }
  @media (max-width: 1000px) {
    .contact,
    footer,
    hr {
      display: none !important;
    }
  }
</style>
