<script>
  import FluidGaussianWrapper from './FluidGaussianWrapper.svelte'
  let width
  let height
</script>

<svelte:window bind:innerWidth={width} bind:innerHeight={height} />
<div current-liquid-page {...$$restProps}>
  <slot name="current-page" />
</div>
<FluidGaussianWrapper
  position={{
    y: 515,
  }}
  {width}
  {height}
  onComplete={() => {
    // console.log('completed!')
  }}
  onButtonMove={(x, y) => {
    // console.log(x, y)
  }}
>
  <div slot="page" class="wrapper-page">
    <slot name="next-page" />
  </div>
</FluidGaussianWrapper>

<style>
  .wrapper-page {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    z-index: 9999;
    transform: translate3d(0, 0, 0);
    will-change: transform;
  }
</style>
