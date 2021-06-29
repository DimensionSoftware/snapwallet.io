<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()

  export let hidden: boolean = false
  export let label: string = ''
  export let error: string = ''
  export let fx: boolean = true
</script>

{#if !hidden}
  <label
    on:click={_ => dispatch('click')}
    class:fx
    class={$$props.class}
    style={$$props.style}
  >
    <span class="input-label">{label}</span>
    <slot />
    <div class="error-help">
      {#if error}
        {error}
      {/if}
    </div>
  </label>
{/if}

<style lang="scss">
  @import '../../styles/_vars.scss';
  @import '../../styles/animations.scss';
  label {
    cursor: pointer;
    position: relative;
    color: var(--theme-text-color-3);
    text-transform: uppercase;
    font-size: 0.7rem;
    font-weight: 600;
    margin-bottom: 1.25em;
    z-index: 25;
    &.fx:active {
      animation: focus 0.16s;
    }
    span {
      position: absolute;
      top: 4px;
      margin-left: 15px;
      margin-top: 8px;
      z-index: 99;
    }
  }
  :global(label .input-container > input) {
    padding-top: 1.5em !important;
  }
  .error-help {
    color: var(--theme-error-color) !important;
    font-size: 0.5rem;
    padding: 0.4em;
  }
</style>
