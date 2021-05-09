<script lang="ts">
  import { onMount } from 'svelte'
  import Feature from '../Feature.svelte'
  import { formatLocaleCurrency } from '../../../../widget/src/util'
  import Button from '../Button.svelte'

  const cards = [
    {
      title: 'NFT Checkout',
      name: 'nft',
      desc:
        "Need a safe and reliable wallet to handle your NFT purchases?  SnapWallet's got you.",
      icon: 'Crypto_Checkout_final.png',
    },
    {
      title: 'Donations & Tips',
      name: 'donations',
      desc:
        'Accept crypto donations and tips for your services!  Snap Wallet makes that simple.',
      icon: 'Crypto_Tip_Jar_final.png',
    },
    {
      title: 'Fiat On & Off Ramp',
      name: 'fiat',
      desc:
        'Buy and Sell your fiat, like USD, into crypto:  Bitcoin & Ethereum, ERC20 tokens, etc...',
      icon: 'Crypto_Fiat_final.png',
    },
  ]

  function scrollTo(name) {
    console.log('n', name)
    document.getElementById(name).scrollIntoView()
  }
</script>

<Feature
  name="overview"
  title="SW DOES IT ALL"
  description="Crypto doesn't have to be difficult to spend and buy."
>
  <div class="flex" slot="left">
    {#each cards as card}
      <article on:mousedown={_ => scrollTo(card.name)}>
        <img
          height="300"
          width="300"
          src={`/images/${card.icon}`}
          alt={card.title}
        />
        <h4>{card.title}</h4>
        <p>{card.desc}</p>
      </article>
    {/each}
  </div>
</Feature>

<style lang="scss">
  @import '../../../../widget/src/styles/animations.scss';

  .flex {
    display: flex;
  }
  :global(.overview h2),
  :global(.overview h3) {
    left: -10%;
    margin-left: -2rem;
  }
  article {
    max-width: 300px;
    margin: 0 2.5rem;
    margin-top: 0;
    box-shadow: 8px 10px 20px rgba(0, 0, 0, 0.15);
    transition: box-shadow 0.3s ease-out, transform 0.04s ease-in;
    border-radius: 1rem;
    padding: 0 1.5rem;
    cursor: pointer;
    &:first-child {
      margin-left: -10%;
    }
    &:hover {
      box-shadow: 0 0 0 3px #fff600, 5px 7px 15px rgba(0, 0, 0, 0.18);
      transform: scale(0.995);
      transition: none;
    }
  }

  img {
    margin-bottom: -1rem;
  }
  h4 {
    text-align: center;
    // padding-left: 1.5rem;
    font-size: 1.25rem;
    margin: 0 0 2rem 0;
  }
  p {
    font-size: 1.05rem;
    line-height: 1.4rem;
    padding: 0 1.5rem;
    font-weight: 100;
    margin: -1.25rem 0 5.5rem 0;
  }
</style>
