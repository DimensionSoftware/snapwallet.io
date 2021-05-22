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
  <Button on:mousedown={contactUs}>Contact Us</Button>
</div>
<hr />
<footer>
  <ol>
    <li><h2>Snap Wallet</h2></li>
    <li>
      <h4>
        <a href="https://snapwallet.io/docs/guide" target="_blank"
          >API Documentation</a
        >
      </h4>
    </li>
    <li>
      <h4><a href="mailto:support@snapwallet.io">support@snapwallet.io</a></h4>
    </li>
  </ol>
  <p title="Scroll to Top" on:mousedown={scrollToTop}>
    <big
      >Snap Wallet
      <img height="24px" width="24px" title="Love" alt="Love" src="/love.svg" />
      Silicon Beach, CA
    </big>
    <br /><small
      ><a
        href="https://dimensionsoftware.com"
        on:mousedown|stopPropagation
        target="_blank">Dimension Software since 1998</a
      ></small
    >
  </p>
</footer>

<style lang="scss">
  .cards {
    display: flex;
    position: relative;
    flex-direction: row;
    h4 {
      font-weight: 300;
      text-align: center;
      margin-top: -1rem;
      font-size: 1.2rem;
    }
    article:first-child > span {
      right: -25%;
    }
    article {
      margin: 0 2rem;
    }
    article > span {
      position: absolute;
      height: 2px;
      border-radius: 1rem;
      right: -25%;
      background: linear-gradient(to right, transparent, #ddd, transparent);
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
    left: auto;
    right: auto;
    margin: 0 auto;
    text-align: center;
    border-radius: 0.25rem;
  }
  footer {
    cursor: pointer;
    margin-top: 0.05rem;
    padding: 1.5rem 0 1rem;
    background: #fffc00;
    color: rgba(25, 25, 25, 0.8);
    ol {
      list-style: none;
      margin: 7rem 0 5rem 5rem;
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
        color: rgba(0, 0, 0, 0.8);
        text-decoration: none;
      }
    }
    p {
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
