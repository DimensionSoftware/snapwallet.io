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
  export let required: boolean

  let isActive: boolean = Boolean(defaultValue)

  onMount(function () {
    setTimeout(() => document.querySelector('input[autofocus]')?.focus(), 200)
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
    {required}
    on:input={e => {
      isActive = Boolean(e.currentTarget?.value)
      dispatch('change', e.target.value)
    }}
    min={type === 'number' ? 0.0 : null}
    value={defaultValue || ''}
  />
  <span class="fx" />
  <span class="bg" />
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';
  @import '../../styles/animations.scss';

  .input-container {
    padding-bottom: 0;
    margin-bottom: 0.75rem;
    position: relative;
    input[type='number'] {
      padding-right: 0.75em;
    }
    input {
      position: relative;
      z-index: 2;
      margin: 0;
      color: var(--theme-textColor);
      background-color: #fff;
      outline: none;
      box-shadow: 0 -1px 10px -5px var(--theme-shadow-color);
      width: 100%;
      appearance: none;
      backface-visibility: hidden;
      transform: translateZ(0);
      border-radius: 0.5em;
      text-indent: 10px;
      text-transform: lowercase;
      overflow: hidden;
      cursor: pointer;
      text-overflow: ellipsis;
      outline: none;
      padding: 35px 5px 15px 5px !important;
      vertical-align: middle;
      font-size: 1.8em;
      color: var(--theme-text-color);
      border: none;
      border-bottom: 1px solid lighten($themeColor, 35%);
      outline: none;
      width: 100%;
      transform: scale(1);
      transition-duration: 0.3s;
      transition-property: transform;
      transition: color 0.2s ease-out, border 0.3s ease-out 0.1s;
      &:valid {
        border-bottom: 1px solid lighten($themeColor, 25%);
      }
      ~ .bg {
        position: absolute;
        content: '';
        top: 0;
        bottom: -1px;
        left: -1px;
        right: -1px;
        border-radius: 0.8em;
        background: linear-gradient(transparent, var(--theme-color));
        opacity: 0;
        transform: opacity(0), scale(0);
        transition: opacity 0.3s ease-out 0.2s, transform 0.4s ease-in 0.1s;
      }
      &:valid ~ .bg {
        opacity: 0.5;
        transform: opacity(1), scale(1);
        transition: none;
      }
      &:hover,
      &:focus {
        // background-image: none;
        transition: none;
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
        left: 0;
        right: 0;
        bottom: 0;
        height: 1px;
        overflow: hidden;
        &:after {
          content: '';
          position: absolute;
          width: 100%;
          left: 0;
          right: 0;
          height: 2px;
          z-index: 9;
          background: linear-gradient(
            to right,
            transparent,
            $themeColor,
            transparent
          );
          transform: scale(0);
          transition: transform 0.5s $easeOutExpo 0.15s;
        }
      }
      &:focus ~ .bg {
        animation: focus 0.18s;
        animation-timing-function: ease-in;
      }
      &:hover + .fx,
      &:focus + .fx {
        &:after {
          transform: scale(1);
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
