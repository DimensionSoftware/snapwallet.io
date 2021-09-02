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
  <svg on:click={copy} viewBox="0 0 20 20" height="20" width="20">
    <path
      d="M17.391,2.406H7.266c-0.232,0-0.422,0.19-0.422,0.422v3.797H3.047c-0.232,0-0.422,0.19-0.422,0.422v10.125c0,0.232,0.19,0.422,0.422,0.422h10.125c0.231,0,0.422-0.189,0.422-0.422v-3.797h3.797c0.232,0,0.422-0.19,0.422-0.422V2.828C17.812,2.596,17.623,2.406,17.391,2.406 M12.749,16.75h-9.28V7.469h3.375v5.484c0,0.231,0.19,0.422,0.422,0.422h5.483V16.75zM16.969,12.531H7.688V3.25h9.281V12.531z"
    />
  </svg>
</span>

<style>
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
    opacity: 0.8;
    transform: scale(1);
    transition: transform 0.2s ease-out;
  }
  svg:hover {
    transform: scale(1.05);
    opacity: 1;
    transition: none;
  }
</style>
