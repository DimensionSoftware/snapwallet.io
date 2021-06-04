<script>
  import { onMount, tick } from 'svelte'

  import { elasticOut } from 'svelte/easing'
  import { spring, tweened } from 'svelte/motion'

  export let width = 500
  export let height = 500
  export let step = 10
  export let position
  export let triggerValue = 350
  export let precision = 5
  export let slideDuration = 1000
  export let slideEasing = elasticOut
  export let buttonSize = 10
  export let veilExpansionStiffness = 0.03
  export let veilExpansionDamping = 0.1
  export let buttonExpansionStiffness = 0.05
  export let buttonExpansionDamping = 0.2
  export let onComplete
  export let onButtonMove

  function uuid(short = false) {
    let dt = new Date().getTime()
    const BLUEPRINT = short
      ? 'xyxxyxyx'
      : 'xxxxxxxx-xxxx-yxxx-yxxx-xxxxxxxxxxxx'
    const RESULT = BLUEPRINT.replace(/[xy]/g, function (c) {
      let r = (dt + Math.random() * 16) % 16 | 0
      dt = Math.floor(dt / 16)
      return (c == 'x' ? r : (r & 0x3) | 0x8).toString(16)
    })
    return RESULT
  }

  let id = `clip-${uuid()}`

  const IDLE = 0,
    START_SLIDING = 1,
    SLIDING = 2,
    DONE = 3
  let status = IDLE
  let active = false

  let btn
  let slide
  let tip
  let size
  let originalTipPosition

  let path = ''
  let mounted = true

  let btnOpacity = 1.0

  async function onBoundsChange(width, height) {
    mounted = false
    await tick()
    mounted = true
  }

  $: onBoundsChange(width, height)

  function init() {
    window.addEventListener('closeLiquid', closeLiquid)

    if (!position) position = { x: 70, y: height - 200 }

    if (!position.x) position.x = 70

    if (!position.y) position.y = height - 200

    originalTipPosition = tweened(
      { x: 0, y: position.y },
      {
        duration: slideDuration,
        easing: slideEasing,
      },
    )

    originalTipPosition.set(position)

    slide = tweened(0, {
      duration: slideDuration,
      easing: slideEasing,
    })

    size = spring(
      { buttonSize },
      {
        stiffness: veilExpansionStiffness,
        damping: veilExpansionDamping,
      },
    )

    tip = spring($originalTipPosition, {
      stiffness: buttonExpansionStiffness,
      damping: buttonExpansionDamping,
    })

    requestAnimationFrame(render)
  }

  let deltaLeft = {
    slide: 0,
    base: 0,
  }
  let lastX, lastY
  let rendering = false
  function render() {
    if (DONE === status) return
    let x = $tip.x < $originalTipPosition.x ? $originalTipPosition.x : $tip.x
    let y = height + x - $tip.y
    rendering = true
    let s

    if (status < START_SLIDING && x > triggerValue) status = START_SLIDING

    if (status === START_SLIDING) {
      status = SLIDING
      slide.set(width)
    }

    if (status >= START_SLIDING) {
      size.set({
        buttonSize: $size.buttonSize + x,
      })
      deltaLeft.slide = x
      deltaLeft.base = $slide
    }

    if (deltaLeft.base >= width) {
      status = DONE
      if (onComplete) onComplete()
    } else if (x !== lastX || y !== lastY) {
      if (onButtonMove) onButtonMove(x, y)
    }

    s = $size.buttonSize + x * 0.5

    btnOpacity = 1 + ($originalTipPosition.x - x) / 100

    // start drawing
    let nx = 0
    let ny = height + x
    let coords = `M-${step},0`
    for (ny = height + x; ny >= 0; ny -= step) {
      nx = x / Math.pow(Math.E, Math.pow(ny - y, 2) / (2 * s * s))
      coords += ` L${(deltaLeft.base + nx).toFixed(precision)},${(
        height +
        x -
        ny
      ).toFixed(precision)}`
    }
    coords += ` L-${step},${height + x}`
    //finish drawing

    path = coords
    rendering = false
    lastX = x
    lastY = y
    requestAnimationFrame(render)
  }

  function activate(e) {
    if (e.target === btn) active = true
  }

  function deactivate(e) {
    active = false
    if (status === IDLE) {
      tip.set($originalTipPosition)
      slide.set(width)
    }
    // setTimeout(_ => requestAnimationFrame(render), 100)
  }

  function watch(e, x, y) {
    if (
      active &&
      status === IDLE &&
      !rendering &&
      (lastX !== x || lastY !== y)
    ) {
      tip.set({ x, y })
      lastX = x
      lastY = y
    }
  }
  function keydown(e) {
    if (e.key == 'Escape') closeLiquid(e)
  }

  function openLiquid(e) {
    status = SLIDING
    mounted = active = true
    tick().then(_ => requestAnimationFrame(render))
  }

  function closeLiquid(e) {
    mounted = false
    deactivate(e)
  }

  onMount(init)
  $: isOpen = active || status == DONE
</script>

<svelte:window on:keydown={keydown} />

{#if mounted}
  <div
    title="Double Click to Open!"
    class="wrapper"
    class:isOpen
    on:dblclick={openLiquid}
    on:mousedown={activate}
    on:mouseup={deactivate}
    on:mousemove={e => watch(e, e.clientX, e.clientY)}
    on:touchstart={activate}
    on:touchend={deactivate}
    on:touchcancel={deactivate}
    on:touchmove={e => watch(e, e.touches[0].clientX, e.touches[0].clientY)}
  >
    <svg {width} {height} viewBox="0 0 {width} {height}">
      <clipPath {id}>
        <path d={path} fill="rgba(0,0,0,0.5)" stroke="black" />
      </clipPath>
    </svg>

    <div class="page" style="clip-path: url(#{id});">
      <slot name="page" />
    </div>

    <div
      bind:this={btn}
      class="btn noselection"
      style="opacity:{btnOpacity};left:{deltaLeft.base +
        ($tip.x < $originalTipPosition.x ? $originalTipPosition.x : $tip.x) -
        55}px;top:{$tip.y - 35 - 2.5}px"
    >
      <svg
        version="1.1"
        xmlns="http://www.w3.org/2000/svg"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        x="0px"
        y="0px"
        viewBox="0 -125 385.343 385.343"
        style="enable-background:new 0 0 185.343 185.343;"
        xml:space="preserve"
      >
        <g>
          <g>
            <path
              style="fill:#010002;"
              d="M51.707,185.343c-2.741,0-5.493-1.044-7.593-3.149c-4.194-4.194-4.194-10.981,0-15.175
			l74.352-74.347L44.114,18.32c-4.194-4.194-4.194-10.987,0-15.175c4.194-4.194,10.987-4.194,15.18,0l81.934,81.934
			c4.194,4.194,4.194,10.987,0,15.175l-81.934,81.939C57.201,184.293,54.454,185.343,51.707,185.343z"
            />
          </g>
        </g>
        <g />
        <g />
        <g />
        <g />
        <g />
        <g />
        <g />
        <g />
        <g />
        <g />
        <g />
        <g />
        <g />
        <g />
        <g />
      </svg>
    </div>
  </div>
{/if}

<style>
  .noselection {
    -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
    -khtml-user-select: none; /* Konqueror HTML */
    -moz-user-select: none; /* Old versions of Firefox */
    -ms-user-select: none; /* Internet Explorer/Edge */
    user-select: none; /* Non-prefixed version, currently
									supported by Chrome, Edge, Opera and Firefox */
  }

  .wrapper {
    position: fixed;
    overflow: hidden;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    z-index: 11;
    max-width: 100px;
  }

  .isOpen {
    max-width: 100%;
    width: 100%;
  }

  .btn {
    display: grid;
    cursor: pointer;
    justify-items: center;
    align-items: center;
    position: absolute;
    width: 60px;
    height: 60px;
    border-radius: 50%;
    font-size: 4rem;
    position: relative;
  }

  .page {
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    position: absolute;
    display: grid;
  }
  .page > * {
    width: 100%;
    height: 100%;
  }
  svg {
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    z-index: 1;
    pointer-events: none;
  }
</style>
