<script lang="ts">
  export let title
  export let docLink
  export let right
</script>

<section>
  <article>
    <h2 class:right>
      {title}
    </h2>
    <div class="flex">
      <slot name="left" />
      <div class="relative">
        <slot name="right" />
      </div>
    </div>
  </article>
  <a class="bottom docs-link" href={docLink} target="_blank">
    Get Started with an Example
    <br />
    <img
      height="25px"
      width="25px"
      title="Get Started with Code Snippets!"
      alt="Made"
      src="/made.svg"
    />
  </a>
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
    .docs-link {
      margin: 1rem auto 0;
    }
    article {
      position: relative;
      padding: 3rem;
      border-radius: 2rem;
      background: linear-gradient(#fcfcfc, #f3f3f3);
      max-width: 800px;
      width: 100%;
      margin: 0 auto;
      h2 {
        &.right {
          right: 0.75rem;
          left: inherit;
        }
        position: absolute;
        top: -8rem;
        left: 1rem;
        background: linear-gradient(
          var(--theme-button-color),
          rgb(247, 127, 26)
        );
        background-clip: text;
        color: transparent;
        font-size: 3rem;
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
