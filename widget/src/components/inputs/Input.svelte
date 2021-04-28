<script lang="ts">
  import { onMount, createEventDispatcher } from 'svelte'
  import type { Masks } from '../../types'
  import { withMaskOnInput, isValidMaskInput } from '../../masks'
  import { focus, onFocusSelect } from '../../util'

  const dispatch = createEventDispatcher()
  export let type: string = 'text'
  export let placeholder: string = ''
  export let inputmode: string = 'text'
  export let autocapitalize: string = ''
  export let defaultValue: string | number = ''
  export let autocomplete: string = 'on'
  export let autofocus: boolean
  export let autoselect: boolean
  export let required: boolean
  export let maxlength: number
  export let pattern: string = ''
  export let mask: Masks
  export let id: string

  let isActive: boolean = Boolean(defaultValue)

  const selectOnFocus = autoselect ? onFocusSelect : _ => {}

  onMount(function () {
    focus(document.querySelector('input[autofocus]'), 200)
  })
</script>

<div class:active={isActive} class="input-container">
  <input
    {id}
    {type}
    {inputmode}
    {autocapitalize}
    {placeholder}
    {autocomplete}
    {autofocus}
    {maxlength}
    {pattern}
    {required}
    use:selectOnFocus
    on:keydown={e => {
      if (mask) {
        const newVal = defaultValue + String.fromCharCode(e.keyCode)
        const isValLongerThanMask = newVal.length > mask.length
        // Uses codes from the following table https://keycode.info/
        const isAltering =
          [8, 9, 12, 13, 16, 17, 18, 20, 41, 46].includes(e.keyCode) ||
          e.metaKey

        const isInputValid =
          isValidMaskInput(newVal, mask) && !isValLongerThanMask

        if (!isInputValid && !isAltering) {
          e.preventDefault()
          return false
        }
      }
    }}
    on:keydown
    on:click
    on:focus={() => dispatch('focus')}
    on:input={e => {
      isActive = Boolean(e.currentTarget?.value)
      dispatch('change', e.target.value)
    }}
    min={type === 'number' ? 0.0 : null}
    value={withMaskOnInput(defaultValue, mask)}
  />
  <span class="fx" />
  <span class="bg" />
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';
  @import '../../styles/animations.scss';

  .input-container {
    padding-bottom: 0;
    margin-bottom: 0rem;
    position: relative;
    input[type='number'] {
      padding-right: 0.75em;
    }
    input {
      position: relative;
      z-index: 2;
      margin: 0;
      background-color: var(--theme-color-inverse);
      outline: none;
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
      padding: 1.6em 5px 15px 5px !important;
      vertical-align: middle;
      font-size: 1.8em;
      color: var(--theme-input-text-color);
      border: none;
      border-bottom: 1px solid var(--theme-color-lightened);
      outline: none;
      width: 100%;
      transform: scale(1);
      transition-duration: 0.3s;
      transition-property: transform;
      &:valid {
        border-bottom: 1px solid var(--theme-color-lightened);
      }
      // .bg is the input surround
      ~ .bg {
        position: absolute;
        content: '';
        top: 0;
        bottom: 0px;
        left: -1px;
        right: -1px;
        border-radius: 0.9em;
        opacity: 0;
        transform: scale(0);
        transition: opacity 0.5s ease-out, left 0s ease-out 0.51s,
          right 0s ease-out 0.51s, top 0.4s ease-out 0.01s,
          bottom 0.4s ease-out 0.11s, background 0s ease-out 0.51s,
          transform 0s ease-out 0.51s;
      }
      &:valid ~ .bg {
        transform: scale(1);
      }
      &:hover,
      &:focus {
        z-index: 1;
        border-bottom: 1px solid var(--theme-color);
        transition: none;
      }
      &:invalid {
        box-shadow: none;
      }
      &:focus {
        cursor: text;
        color: var(--theme-input-text-color);
        transition: none;
      }

      &:required {
        box-shadow: none;
      }
      // .fx is the subtle bottom line
      & + .fx {
        position: absolute;
        left: 0;
        right: 0;
        bottom: 0;
        height: 1px;
        background: linear-gradient(
          to right,
          transparent,
          var(--theme-color),
          transparent
        );
        z-index: 11;
        opacity: 0;
        transform: scale(0);
        transition: opacity 0.5s ease-out 0.1s, transform 0.5s ease-out 0.1s;
      }
      &:active ~ .bg,
      &:focus ~ .bg {
        left: -5px;
        top: -4px;
        right: -5px;
        bottom: -4px;
        background: var(--theme-color-lightened);
        opacity: 0.5;
        transform: scale(1);
        transition: none;
      }
      &:active + .fx,
      &:hover + .fx,
      &:focus + .fx {
        opacity: 1;
        transform: scale(1) translateX(0);
        transition: none;
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
