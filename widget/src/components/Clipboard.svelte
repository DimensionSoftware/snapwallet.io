<script>
  import { tick } from 'svelte'

  export let value = null
  let hiddenInput

  async function copy() {
    // copy-to-clipboard
    // - TODO there are more portable ways
    await tick()
    hiddenInput.focus()
    hiddenInput.select()
    try {
      if (!document.execCommand('copy'))
        console.warn(`Unable to copy: ${value}`)
    } catch (err) {
      console.warn(`Error copying "${value}": ${err}`)
    }
  }
</script>

{#if value != null}
  <textarea bind:this={hiddenInput}>{value}</textarea>
{/if}
<span title={`Copy ${value}`}>
  <svg
    on:click={copy}
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 22.5 22.5"
    height="16"
    width="16"
    ><defs
      ><style>
        .cls-1 {
          fill: #fff;
        }
      </style></defs
    ><g id="Layer_2" data-name="Layer 2"
      ><g id="copy_icon"
        ><path
          class="cls-1"
          d="M20.11,4.18H6.57A2.39,2.39,0,0,0,4.18,6.57V20.11A2.39,2.39,0,0,0,6.57,22.5H20.11a2.39,2.39,0,0,0,2.39-2.39V6.57A2.39,2.39,0,0,0,20.11,4.18Zm2,15.93a2,2,0,0,1-2,2H6.57a2,2,0,0,1-2-2V6.57a2,2,0,0,1,2-2H20.11a2,2,0,0,1,2,2Z"
        /><path
          class="cls-1"
          d="M2.59,17.52h-.4A1.79,1.79,0,0,1,.4,15.73V2.19A1.79,1.79,0,0,1,2.19.4H15.73a1.79,1.79,0,0,1,1.79,1.79v.4a.2.2,0,0,0,.4,0v-.4A2.19,2.19,0,0,0,15.73,0H2.19A2.19,2.19,0,0,0,0,2.19V15.73a2.19,2.19,0,0,0,2.19,2.19h.4a.2.2,0,0,0,0-.4Z"
        /><path
          class="cls-1"
          d="M16.26,13.14H13.54V10.42a.2.2,0,1,0-.4,0v2.72H10.42a.2.2,0,1,0,0,.4h2.72v2.72a.2.2,0,0,0,.4,0V13.54h2.72a.2.2,0,0,0,0-.4Z"
        /></g
      ></g
    ></svg
  >
</span>

<style lang="scss">
  textarea {
    position: absolute;
    top: 0;
    left: 0;
    width: 0;
    height: 0;
    padding: 0;
    border: none;
    outline: none;
    box-shadow: none;
    background: transparent;
  }

  svg {
    position: relative;
    cursor: pointer;
    top: 2px;
    opacity: 1;
    transform: scale(1);
    transition: transform 0.2s ease-out;
    &:hover {
      transform: scale(1.05);
      opacity: 1;
      transition: none;
    }
    path {
      fill: var(--theme-text-color);
    }
  }
</style>
