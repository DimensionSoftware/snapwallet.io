<script lang="ts">
  export let title
  export let description
  export let docLink
  export let right
  export let name
  export let hasBackground
  export let hasImage
</script>

<a id={name} />
<section class={name} class:hasImage class:hasBackground>
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
    padding: 10rem 10rem 7rem 0;
    position: relative;
    &.hasImage {
      &:before {
        content: '';
        position: absolute;
        top: 0;
        right: 0;
        left: 0;
        bottom: 0;
        transform: rotate(180deg);
        background-image: url('/static/bg.png');
      }
    }
    &.hasBackground article {
      padding: 3rem;
      border-radius: 2rem;
      background: linear-gradient(#fafafa, #f0f0f0);
    }
    article {
      position: relative;
      padding: 3rem 0;
      max-width: 900px;
      width: 100%;
      margin: 0 auto;
      h2 {
        position: absolute;
        top: -5.5rem;
        left: -5%;
        margin-left: -2rem;
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
        top: -2.25rem;
        left: -5%;
        max-width: 65%;
        margin-left: -2rem;
        font-size: 1.25rem;
        font-weight: 100;
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
      box-shadow: 1px 2px 3px rgba(0, 0, 0, 0.1);
      border-radius: 2rem;
      padding: 0.25rem 2rem;
      text-align: center;
      font-size: 1.1rem;
      font-weight: 100;
      margin: 1.5rem 0 0 0;
      color: #222;
      transition: background 0.2s ease-out, box-shadow 0.2s ease-out;
      overflow: hidden;
      backdrop-filter: blur(5px) contrast(110%);
      &:hover {
        transition: none;
        background-color: rgba(255, 255, 255, 0.7);
        box-shadow: 1px 2px 3px rgba(0, 0, 0, 0.3);
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
  @media (max-width: 1000px) {
    section {
      display: none;
    }
  }
</style>
