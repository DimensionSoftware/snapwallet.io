<script lang="ts">
  import LiquidSwipeablePage from '$lib/LiquidSwipeablePage.svelte'
  import Button from '../lib/Button.svelte'

  const contactUs = _ => {
    window.location.href =
      'mailto:support@snapwallet.com?subject=Hello Snap Wallet Team!'
  }
</script>

<LiquidSwipeablePage>
  <div slot="next-page" class="page2">
    <span
      class="close"
      on:mousedown={() => window.dispatchEvent(new Event('closeLiquid'))}
      >x</span
    >
    <div class="content2">
      <h2>Get in Touch</h2>
      <h3>Yes, Snap Wallet will customize your widget!</h3>
      <div class="contact">
        <Button id="liquid_contact" on:mousedown={contactUs}>Contact Us</Button>
      </div>
    </div>
  </div>
</LiquidSwipeablePage>

<style lang="scss">
  .close {
    position: absolute;
    padding: 1rem;
    right: 5rem;
    top: 2rem;
    font-weight: 200;
    font-size: 2.5rem;
    cursor: pointer;
  }
  .page1 {
    flex-direction: column;
    display: flex;
    justify-content: space-between;
    padding: 0vh 13px;
    padding-top: 7vh;
    height: 100%;
    width: 100%;
  }
  .page2 {
    opacity: 1 !important;
    background: #fffc00;
    text-align: center;
    font-size: 2rem;
    font-weight: bold;
    flex-direction: column;
    display: flex;
    justify-content: space-between;
    padding: 0vh 23px;
    padding-top: 20vh;
    height: 100%;
    width: 100%;
  }
  .content2 {
    display: flex;
    flex-direction: column;
    text-align: left;
    height: 30vh;
    width: 600px;
    margin: 0 auto;
    transform: translateX(50px);
    h2 {
      margin: 0 0 1rem;
    }
    h3 {
      font-size: 1.25rem;
      opacity: 0.8;
      margin: -0.5rem 0 5rem 0;
    }
    .contact {
      align-self: flex-start;
    }
  }

  .contentL1 {
    color: #fff;
    font-weight: 100;
    text-align: left;
    letter-spacing: -2px;
    line-height: 18px;
    opacity: 0.3;
  }
  .contentL2 {
    color: #fff;
    margin-bottom: 8px;
    text-align: left;
  }
  .contentL3 {
    color: #fff;
    text-align: left;
    font-size: 15px;
    font-weight: 100;
    word-spacing: 2px;
    letter-spacing: 1px;
    opacity: 0.5;
  }
  .content2 .contentL1 {
    color: #fff;
    font-weight: 100;
    text-align: left;
    letter-spacing: -2px;
    line-height: 18px;
    opacity: 0.3;
  }
  .content2 .contentL2 {
    color: #fff;
    margin-bottom: 8px;
    text-align: left;
  }
  .content2 .contentL3 {
    color: #fff;
    text-align: left;
    font-size: 15px;
    font-weight: 100;
    word-spacing: 2px;
    letter-spacing: 1px;
    opacity: 0.5;
  }
  :global(.content2 .contact > button) {
    color: #fff !important;
    font-size: 1.15rem !important;
    background: rgb(241, 7, 28) !important;
    background: linear-gradient(
      63.26deg,
      #ff6b00,
      rgb(241, 7, 28) 75%
    ) !important;
    border-color: rgb(222, 49, 45) !important;
    position: relative;
    z-index: 1;
    padding: 1rem 4rem;
    text-align: center;
    margin-top: -3rem;
    border-radius: 4.25rem;
    transition: box-shadow 0.3s ease-out;
    &:hover {
      transition: none;
      box-shadow: 0 0 0 3px rgba(222, 49, 45, 0.25) !important;
    }
  }
</style>
