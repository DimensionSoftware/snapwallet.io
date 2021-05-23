<script lang="ts">
  import { numberWithCommas } from '../../../widget/src/util'

  $: counter = 69420
  $: active = false
  const inc = _ => {
    active = false
    counter++
    setTimeout(_ => (active = true), 100)
  }
</script>

<div class:active class="heart-container">
  <div
    on:mousedown|stopPropagation|preventDefault={inc}
    class:active
    class="heart"
    title="<3 SnapWallet"
  >
    <span>{numberWithCommas(counter)}</span>
  </div>
</div>

<style lang="scss">
  @import '../../widget/src/styles/_vars.scss';

  .heart-container {
    display: inline-block;
    transition: 0.3s $easeInExpo;
    &:hover {
      transform: scale(0.985) translateX(-1px);
      transition: 0.1s $easeOutExpo;
      .heart {
        opacity: 1;
      }
    }
    &.active {
      transform: scale(1.1) translateX(2px);
      transition: 0.1s $easeOutExpo;
      transition: none;
    }
  }
  .heart {
    position: relative;
    width: 100px;
    opacity: 0.75;
    height: 100px;
    transform: translate(-50%, -50%);
    margin: 1.5rem 0 0 1rem;
    font-weight: 400;
    font-size: 1.1rem;
    background: url('/images/Heart-Animation.png') no-repeat;
    background-position: 0 0;
    cursor: pointer;
    animation: fave-heart 1.25s steps(28);
    transition: opacity 0.3s ease-out;
    &.active {
      opacity: 1;
      background-position: -2800px 0;
      transition: background 1.25s steps(28);
    }
  }
  span {
    position: absolute;
    top: 2.5rem;
    left: 4.5rem;
  }
  @keyframes fave-heart {
    0% {
      background-position: 0 0;
    }
    100% {
      background-position: -2800px 0;
    }
  }
</style>
