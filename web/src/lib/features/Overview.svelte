<script lang="ts">
  import { fly } from 'svelte/transition'
  import { expoOut } from 'svelte/easing'
  import Feature from '$lib/Feature.svelte'
  import Visibility from '$lib/Visibility.svelte'

  const cards = [
    {
      title: 'Cart Checkout',
      name: 'cartcheckout',
      desc: 'All-in-One cart checkout for multiple products & services!',
      icon: 'Crypto_Checkout_final.png',
    },
    {
      title: 'Donations & Tips',
      name: 'donations',
      desc: 'Accept crypto donations and tips!  Snap Wallet makes that simple.',
      icon: 'Crypto_Tip_Jar_final.png',
    },
    {
      title: 'NFT Checkout',
      name: 'nft',
      desc: "Need a safe and reliable wallet to handle NFTs?  Snap Wallet's got you.",
      icon: 'nft_checkout_icon.png',
    },
  ]

  function scrollTo(name) {
    document.getElementById(name).scrollIntoView()
  }
</script>

<Feature
  name="overview"
  title="SW DOES IT ALL"
  description="Crypto doesn't have to be difficult to spend and buy."
>
  <div class="flex" slot="left">
    {#each cards as card, i}
      <article on:mousedown={_ => scrollTo(card.name)}>
        <Visibility steps={100} let:percent>
          {#if percent > 50}
            <img
              in:fly={{ easing: expoOut, duration: 500 + i * 500, x: -75 }}
              height="300"
              width="300"
              src={`/images/${card.icon}`}
              alt={card.title}
            />
          {:else}
            <img
              height="300"
              width="300"
              style="opacity: 0"
              src={`/images/${card.icon}`}
              alt={card.title}
            />
          {/if}
        </Visibility>
        <h4>{card.title}</h4>
        <p>{card.desc || ''}</p>
      </article>
    {/each}
  </div>
</Feature>

<style lang="scss">
  @import '../../../../widget/src/styles/animations.scss';

  .flex {
    display: flex;
  }
  article {
    max-width: 300px;
    margin: 0 2.5rem 1rem;
    box-shadow: 8px 10px 20px rgba(0, 0, 0, 0.15);
    transition: box-shadow 0.3s ease-out, transform 0.04s ease-in;
    border-radius: 1rem;
    padding: 0 1.5rem;
    cursor: pointer;
    filter: grayscale(50%);
    &:hover {
      filter: grayscale(0%);
      box-shadow: 0 0 0 3px #fff600, 5px 7px 15px rgba(0, 0, 0, 0.18),
        0 0 75px 5px rgba(#fff600, 0.2);
      transform: scale(0.995);
      transition: none;
    }
  }

  img {
    margin-bottom: -1rem;
  }
  :global(.overview h2),
  :global(.overview h3) {
    left: -12% !important;
  }
  h4 {
    text-align: center;
    font-size: 1.25rem;
    margin: 0 0 2rem 0;
  }
  p {
    font-size: 1rem;
    text-align: center;
    line-height: 1.4rem;
    padding: 0 1.5rem;
    font-weight: 200;
    margin: -1.25rem 0 5.5rem 0;
  }

  @media (max-width: 1250px) {
    article {
      margin-right: 0;
    }
    :global(.overview > article > h2),
    :global(.overview > article > h3) {
      left: -5% !important;
    }
  }
</style>
