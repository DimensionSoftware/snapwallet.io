<script lang="ts">
  import { numberWithCommas } from '../../../widget/src/util'

  $: counter = 69420
  const inc = _ => counter++
</script>

<div on:mousedown={inc} class="heart">
  <span>{numberWithCommas(counter)}</span>
</div>

<style lang="scss">
  .heart {
    position: relative;
    width: 100px;
    height: 100px;
    transform: translate(-50%, -50%);
    margin: 3rem 0 0 1rem;
    font-weight: bold;
    font-size: 1.1rem;
    background: url(https://cssanimation.rocks/images/posts/steps/heart.png)
      no-repeat;
    background-position: 0 0;
    cursor: pointer;
    animation: fave-heart 1s steps(28);
    &:hover {
      opacity: 1;
      background-position: -2800px 0;
      transition: background 1s steps(28);
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
