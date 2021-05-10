<div class="heart"><span>29</span></div>

<style lang="scss">
  .heart {
    position: relative;
    width: 100px;
    height: 100px;
    transform: translate(-50%, -50%);
    margin: 0.75rem 0 0 1rem;
    background: url(https://cssanimation.rocks/images/posts/steps/heart.png)
      no-repeat;
    background-position: 0 0;
    cursor: pointer;
    opacity: 0.5;
    animation: fave-heart 1s steps(28);
    &:hover {
      opacity: 0.9;
      background-position: -2800px 0;
      transition: background 1s steps(28);
    }
    span {
      position: absolute;
      top: 2.5rem;
      left: 4.5rem;
    }
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
