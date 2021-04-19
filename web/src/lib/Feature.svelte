<script lang="ts">
  export let title
  export let docLink
  export let right
</script>

<section>
  <article>
    <h2 class:right>
      {#if right}
        <a class="docs-link" href={docLink} target="_blank">
          <img
            height="25px"
            width="25px"
            title="Get Started with Code Snippets!"
            alt="Made"
            src="/made.svg"
          />
        </a>
        {title}
      {:else}
        {title}
        <a class="docs-link" href={docLink} target="_blank">
          <img
            height="25px"
            width="25px"
            title="Get Started with Code Snippets!"
            alt="Made"
            src="/made.svg"
          />
        </a>
      {/if}
    </h2>
    <div class="flex">
      <slot name="left" />
      <div class="relative">
        <slot name="right" />
      </div>
    </div>
  </article>
</section>

<style lang="scss">
  @import '../../../widget/src/styles/animations.scss';
  .flex {
    display: flex;
    grid-gap: 2rem;
  }
  .relative {
    position: relative;
  }
  section {
    --theme-button-color: #fff600;
    background: white;
    display: flex;
    flex-direction: column;
    padding: 10rem 10rem 0;
    .docs-link {
      margin: 1rem auto 0;
    }
    article {
      position: relative;
      padding: 3rem;
      border-radius: 2rem;
      background: #f6f6f6;
      max-width: 800px;
      width: 100%;
      margin: 0 auto;
      h2 {
        &.right {
          right: 0.75rem;
          left: inherit;
        }
        position: absolute;
        top: -5rem;
        left: 0.75rem;
        font-size: 2rem;
        font-weight: bold;
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
    }
  }
  @media (max-width: 1000px) {
    section {
      display: none;
    }
  }
</style>
