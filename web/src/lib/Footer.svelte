<script lang="ts">
  const scrollToTop = _ => window.scrollTo(0, 0)
</script>

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
  hr {
    height: 5rem;
    background: #fff;
    margin: 0;
    border: none;
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
  @media (max-width: 1000px) {
    footer,
    hr {
      display: none !important;
    }
  }
</style>
