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
        'mailto:support@snapwallet.com?subject=Hello SnapWallet Team!'
    }
  const cards = [
    {
      title: 'Built-In Compliance',
      name: 'nft',
      icon: 'SW Crypto Coin.png',
      alt: 'Simply Checkout with Any NFT',
    },
    {
      title: 'Diverse Payment Methods',
      name: 'donations',
      icon: 'Dr Crypto clean shaven.png',
      alt: 'Accepts Credit Cards',
    },
    {
      title: 'Fast & Secure',
      name: 'fiat',
      icon: 'Coin Checkmark.png',
      alt: 'On and Off-Ramp Crypto Currencies',
    },
  ]
</script>

<Feature
  name="started"
  title="Get started in just a few minutes"
  description="SnapWallet integrates with your platform and is fully customizable"
  center
>
  <div class="cards flex" slot="left">
    {#each cards as card, i}
      <article on:mousedown={_ => scrollTo(card.name)}>
        <img
          title={card.alt}
          width="300"
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
    <br /><small
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
  $themeColor: #fffc00;
  .cards {
    display: flex;
    position: relative;
    flex-direction: row;
    h4 {
      font-weight: 200;
      text-align: center;
      margin-top: -1rem;
      font-size: 1.5rem;
    }
    article:first-child > span {
      right: -25%;
    }
    article {
      margin: 0 2rem;
    }
    article > span {
      position: absolute;
      height: 1px;
      border-radius: 1rem;
      right: -25%;
      background: linear-gradient(
        to right,
        transparent,
        darken($themeColor, 10%),
        transparent
      );
      width: 100px;
      top: 50%;
    }
  }
  hr {
    height: 5rem;
    background: #fff;
    margin: 0;
    border: none;
  }
  :global(.started article) {
    position: relative;
    width: 100%;
  }
  .contact {
    position: absolute;
    z-index: 2;
    left: calc(50% - 82px);
  }
  :global(.contact > button) {
    color: #fff !important;
    font-size: 1.15rem !important;
    background: rgb(222, 49, 45) !important;
    border-color: rgb(222, 49, 45) !important;
    top: -4.25rem;
    position: relative;
    z-index: 1;
    padding: 0.75rem 4rem;
    eft: auto;
    right: auto;
    margin: 0 auto;
    text-align: center;
    border-radius: 0.25rem;
  }
  footer {
    cursor: pointer;
    margin-top: 0.05rem;
    overflow: hidden;
    padding: 1.5rem 0 1rem;
    background: $themeColor;
    color: rgba(25, 25, 25, 0.8);
    position: relative;
    &:before {
      content: '';
      position: absolute;
      top: 0;
      right: 0;
      left: 0;
      bottom: 0;
      background-image: url('/static/bg.png');
      filter: blur(30px) contrast(120%);
      z-index: 0;
    }
    &:after {
      content: '';
      top: 0;
      left: 0;
      height: 1px;
      right: 0;
      position: absolute;
      background: linear-gradient(
        to right,
        $themeColor,
        darken($themeColor, 10%),
        $themeColor
      );
    }
    ol {
      position: relative;
      z-index: 1;
      list-style: none;
      margin: 7rem 0 5rem 10rem;
      li {
        font-size: 0.9rem;
        margin: 0;
        padding: 0;
        h2 {
          color: #000;
        }
        h4 {
          margin: 0.25rem 0 0.75rem 0;
        }
      }
      a {
        color: rgba(0, 0, 0, 0.7);
        text-decoration: none;
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
      big {
        font-size: 0.9rem;
        color: #222;
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
        color: #222;
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
