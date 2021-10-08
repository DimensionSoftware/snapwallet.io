<script lang="ts">
  import { scale, fly } from 'svelte/transition'
  import { backOut, expoIn, expoOut } from 'svelte/easing'
  import Visibility from '$lib/Visibility.svelte'
  export let title
  export let description
  export let docLink
  export let right
  export let name
  export let hasBackground
  export let hasImage
  export let center
  export let icon
</script>

<a id={name} />
<section class={name} class:center class:hasImage class:hasBackground>
  <article>
    {#if icon}
      <Visibility steps={100} let:percent>
        {#if percent > 80}
          <img
            in:scale={{ opacity: 1, easing: backOut, duration: 750 }}
            out:fly={{ opacity: 0, easing: expoOut, duration: 250, y: -50 }}
            class="icon"
            src={icon}
            alt="Snap Wallet"
          />
        {:else}
          <img
            class="icon"
            src={icon}
            alt="Snap Wallet"
            style="opacity: 0;height: 100px; width: 100px;"
          />
        {/if}
      </Visibility>
    {/if}
    <Visibility steps={100} let:percent>
      {#if percent > 90}
        <h2
          in:fly={{ opacity: 0.5, y: 50, easing: expoOut, duration: 1000 }}
          out:fly={{ opacity: 0, easing: expoOut, duration: 650, y: -25 }}
          class:right
          class:blur={hasImage}
        >
          {title}
        </h2>
      {:else}
        <h2 class:right class:blur={hasImage} />
      {/if}
    </Visibility>
    <Visibility steps={100} let:percent>
      {#if percent > 80}
        {#if description}
          <h3
            in:fly={{ opacity: 0.5, y: 75, easing: expoOut, duration: 750 }}
            out:fly={{ opacity: 0, easing: expoOut, duration: 350, y: -15 }}
            class:blur={hasImage}
          >
            {description}
          </h3>
        {:else}
          <h3 class:blur={hasImage} />
        {/if}
      {/if}
    </Visibility>
    <div class="flex">
      <slot name="left" />
      <div class="relative">
        <slot name="right" />
      </div>
    </div>
  </article>
  {#if docLink}
    <a class="bottom docs-link" href={docLink} target="_blank">
      <img
        height="25px"
        width="25px"
        title="Get Started with Code Snippets!"
        alt="Made"
        src="/made.svg"
      />
      <span>See Docs</span>
    </a>
  {/if}
</section>

<style lang="scss">
  @import '../../../widget/src/styles/animations.scss';
  .flex {
    display: flex;
    justify-content: center;
    grid-gap: 2rem;
    margin-top: 1rem;
  }
  .relative {
    position: relative;
    width: 50%;
  }
  section {
    --theme-button-color: #fff600;
    background: white;
    display: flex;
    flex-direction: column;
    padding: 12rem 0 7rem 0;
    margin-right: -15px;
    position: relative;
    z-index: 1;
    &.center {
      article {
        h2,
        h3 {
          left: inherit;
          text-align: center;
          margin: 0 auto;
          width: 100%;
        }
        h3 {
          // margin-top: 2rem;
          margin-left: 7rem;
        }
      }
    }
    &.hasImage {
      background: transparent;
      // &:before {
      //   content: '';
      //   position: absolute;
      //   top: 0;
      //   right: 0;
      //   left: 0;
      //   bottom: 0;
      //   transform: rotate(180deg);
      //   background-image: url('/bg.png');
      // }
    }
    &.hasBackground article {
      padding: 3rem;
      border-radius: 2rem;
      background: linear-gradient(#fafafa, #f0f0f0);
    }
    article {
      position: relative;
      padding: 3rem 0;
      max-width: 1000px;
      margin: 0 auto;
      .icon {
        position: absolute;
        top: -11.8rem;
        left: calc(-11% - 150px);
        height: 100px;
        width: 100px;
      }
      h2 {
        position: absolute;
        top: -6.75rem;
        left: -5%;
        margin-left: -5rem;
        color: #000;
        font-size: 2.25rem;
        font-weight: bold;
        &.right {
          right: 0.75rem;
          left: inherit;
        }
      }
      h3 {
        position: absolute;
        top: -3.75rem;
        left: -5%;
        max-width: 75%;
        opacity: 0.8;
        margin-left: -5rem;
        font-size: 1.4rem;
        font-weight: 200;
      }
      h4 {
        font-size: 2rem;
      }
    }
    a {
      position: relative;
      text-decoration: none;
      text-align: center;
      z-index: 1;
      img {
        margin: 0 1rem 0.25rem 1rem;
        display: inline-block;
        vertical-align: middle;
      }
    }
    .docs-link {
      position: relative;
      margin: 1rem auto 0;
      border-radius: 2rem;
      padding: 0.25rem 2rem;
      text-align: center;
      font-size: 1.1rem;
      font-weight: 200;
      margin: 1.5rem 0 0 0;
      color: #222;
      border: 1px solid rgba(0, 0, 0, 0.1);
      transition: background 0.3s ease-out, box-shadow 0.2s ease-in,
        border 0s ease-in 0.1s;
      overflow: hidden;
      // background: rgba(255, 255, 255, 0.4);
      &:hover {
        transition: none;
        background-color: rgba(255, 255, 255, 0.85);
        box-shadow: 1px 2px 3px rgba(0, 0, 0, 0.2);
        border: 1px solid transparent;
      }
      &:active {
        animation: scaleIn 0.25s ease-out;
      }
      img {
        display: inline-block;
      }
      span {
        margin: 0 0 0 0.5rem;
        vertical-align: middle;
      }
      &.bottom {
        display: block;
        margin: 1.5rem auto 0;
        img {
          margin: 0.5rem 0;
        }
      }
    }
  }
  @media (max-width: 1400px) {
    .icon {
      display: none;
    }
  }
  @media (max-width: 1000px) {
    section {
      display: none;
    }
  }
  @media (max-width: 1250px) {
    article {
      :global(> h2),
      :global(> h3) {
        left: 1rem !important;
      }
    }
  }
</style>
