<script lang="ts">
  export let title
  export let description
  export let docLink
  export let right
  export let name
  export let hasBackground
</script>

<a id={name} />
<section class={name} class:hasBackground>
  <article>
    <h2 class:right>
      {title}
    </h2>
    {#if description}
      <h3>{description}</h3>
    {/if}
    <div class="flex">
      <slot name="left" />
      <div class="relative">
        <slot name="right" />
      </div>
    </div>
  </article>
  {#if docLink}
    <a class="bottom docs-link" href={docLink} target="_blank">
      See Docs
      <br />
      <img
        height="25px"
        width="25px"
        title="Get Started with Code Snippets!"
        alt="Made"
        src="/made.svg"
      />
    </a>
  {/if}
</section>

<style lang="scss">
  @import '../../../widget/src/styles/animations.scss';
  .flex {
    display: flex;
    grid-gap: 2rem;
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
    padding: 10rem 10rem 0;
    &.hasBackground article {
      padding: 3rem;
      border-radius: 2rem;
      background: linear-gradient(#fafafa, #f0f0f0);
    }
    .docs-link {
      margin: 1rem auto 0;
    }
    article {
      position: relative;
      padding: 3rem 0;
      max-width: 900px;
      width: 100%;
      margin: 0 auto;

      h2 {
        position: absolute;
        top: -6.25rem;
        background: linear-gradient(#333, #000);
        background-clip: text;
        -webkit-background-clip: text;
        color: transparent;
        font-size: 2rem;
        font-weight: bold;
        &.right {
          right: 0.75rem;
          left: inherit;
        }
      }
      h3 {
        position: absolute;
        top: -3rem;
        font-size: 1.25rem;
        font-weight: 100;
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
      margin: 1.5rem 0 0 0;
      color: #222;
      &.bottom {
        display: block;
        margin: 1.5rem auto 0;
        img {
          margin: 0.5rem 0;
        }
      }
    }
  }
  @media (max-width: 1000px) {
    section {
      display: none;
    }
  }
</style>
