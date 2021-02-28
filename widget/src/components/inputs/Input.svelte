<script lang="ts">
  import { onMount, createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
  export let type: string = 'text'
  export let placeholder: string = ''
  export let inputmode: string = 'text'
  export let autocapitalize: string = ''
  export let defaultValue: string | number = ''
  export let autocomplete: string = 'on'
  export let autofocus: boolean

  let isActive: boolean = Boolean(defaultValue)

  onMount(function () {
    setTimeout(() => document.querySelector('input[autofocus]')?.focus(), 300)
  })
</script>

<div class:active={isActive} class="input-container">
  <input
    {type}
    {inputmode}
    {autocapitalize}
    {placeholder}
    {autocomplete}
    {autofocus}
    on:input={e => {
      isActive = Boolean(e.currentTarget?.value)
      dispatch('change', e.target.value)
    }}
    min={type === 'number' ? 0.0 : null}
    value={defaultValue || ''}
  />
  <span class="fx" />
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';
  @import '../../styles/animations.scss';

  .input-container {
    padding-bottom: 0;
    margin-bottom: 0.75rem;
    position: relative;
    input {
      padding: 0;
      margin: 0;
      color: var(--theme-textColor);
      background-color: #fff;
      outline: none;
      box-shadow: 0 -1px 10px -5px var(--theme-shadow-color);
      width: 100%;
      appearance: none;
      backface-visibility: hidden;
      transform: translateZ(0);
      border-radius: 3px;
      text-indent: 0.75em;
      text-transform: lowercase;
      overflow: hidden;
      cursor: pointer;
      text-overflow: ellipsis;
      outline: none;
      padding: 15px;
      padding-left: 0;
      padding-right: 0;
      vertical-align: middle;
      font-size: 1.1em;
      color: var(--theme-text-color);
      border: none;
      border-bottom: 1px solid lighten($themeColor, 35%);
      outline: none;
      width: 100%;
      transform: scale(1);
      transition-duration: 0.3s;
      transition-property: transform;
      transition: color 0.2s ease-out, border 0.3s ease-out 0.1s;
      &:hover,
      &:focus {
        // background-image: none;
        transition: none;
      }
      &:hover {
        z-index: 1;
        border-bottom: 1px solid var(--theme-color);
        transition: none;
      }
      &:invalid {
        box-shadow: none;
      }
      &:focus {
        cursor: text;
        z-index: 10;
        color: var(--theme-text-color);
        animation: focus 0.15s;
        animation-timing-function: ease-out;
        transition: none;
      }

      &:required {
        box-shadow: none;
      }
      & + .fx {
        position: absolute;
        text-align: left;
        left: 0;
        right: 0;
        bottom: 0;
        height: 2px;
        overflow: hidden;
        &:after {
          content: '';
          position: absolute;
          width: 100%;
          height: 1px;
          background-color: $themeColor;
          transform: translateX(-500px);
          transition: transform 0.35s $easeInExpo 0.15s;
        }
      }
      &:focus + .fx {
        &:before {
          position: absolute;
          content: '';
          height: 2px;
          z-index: 1;
          background-color: $themeColor;
        }
        &:after {
          width: 100%;
          transform: translateX(0);
          transition: none;
        }
      }
    }
    // HACK: remove yellow autofill background
    &:-webkit-autofill,
    &:-webkit-autofill:hover,
    &:-webkit-autofill:focus,
    &:-webkit-autofill:active {
      box-shadow: none !important;
      background-color: transparent !important;
      background-clip: content-box !important;
      -webkit-text-fill-color: var(--theme-text-color) !important;
    }
  }
</style>
