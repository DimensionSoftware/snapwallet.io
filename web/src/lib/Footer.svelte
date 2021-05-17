<script lang="ts">
  import Feature from './Feature.svelte'
  import Button from './Button.svelte'
  const scrollToTop = _ => window.scrollTo(0, 0),
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
<footer on:mousedown={scrollToTop}>
  <p title="Scroll to Top">
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
      right: -14%;
    }
    article > span {
      position: absolute;
      height: 2px;
      border-radius: 1rem;
      right: -20%;
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
    background: #fff;
    color: rgba(25, 25, 25, 0.8);
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
  @media (max-width: 1250px) {
    :global(article h2),
    :global(article h3) {
      left: 1rem !important;
    }
    :global(.overview article) {
      margin-right: 0 !important;
    }
    :global(.overview article:first-child) {
      margin-left: 3rem !important;
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
