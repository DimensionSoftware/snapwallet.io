<script lang="ts">
  import { onMount, createEventDispatcher } from 'svelte'
  import type { Masks } from '../../types'
  import { withMaskOnInput, isValidMaskInput } from '../../masks'
  import { focus } from '../../util'
  import FaIcon from 'svelte-awesome'
  import { faCaretDown } from '@fortawesome/free-solid-svg-icons'
  import { userStore } from '../../stores/UserStore'
  import { SVG_FLAG_ICON_PATH } from '../../constants'

  const dispatch = createEventDispatcher()
  export let type: string = 'text'
  export let placeholder: string = ''
  export let inputmode: string = 'text'
  export let autocapitalize: string = ''
  export let defaultValue: string | number = ''
  export let autocomplete: string = 'on'
  export let autofocus: boolean
  export let required: boolean
  export let pattern: string = ''
  export let mask: Masks
  export let id: string

  let isActive: boolean = Boolean(defaultValue)

  onMount(function () {
    focus(document.querySelector('input[autofocus]'), 200)
  })
</script>

<div class:active={isActive} class="input-container">
  <div on:mousedown={() => dispatch('select')} class="country-select">
    <img
      alt={$userStore.phoneNumberCountry.code.toUpperCase()}
      width="32"
      src={`${SVG_FLAG_ICON_PATH}/${
        $userStore.phoneNumberCountry.code[0] +
        $userStore.phoneNumberCountry.code[1]
      }.svg`.toLowerCase()}
    />
    {$userStore.phoneNumberCountry.dial_code}
    <FaIcon data={faCaretDown} />
  </div>
  <input
    {id}
    {type}
    {inputmode}
    {autocapitalize}
    {placeholder}
    {autocomplete}
    {autofocus}
    {pattern}
    {required}
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

  :global(.phone .input-label) {
    margin-left: 11px !important;
  }
  .input-container {
    display: flex;
    padding-bottom: 0;
    margin-bottom: 0rem;
    position: relative;
    .country-select {
      display: flex;
      position: relative;
      z-index: 2;
      background-color: var(--theme-color-inverse);
      border-top-left-radius: 0.5em;
      border-bottom-left-radius: 0.5em;
      text-transform: uppercase;
      width: 40%;
      padding: 1.6em 0px 15px 12px !important;
      vertical-align: middle;
      color: var(--theme-input-text-color);
      border: none;
      border-bottom: 1px solid transparent;
      transform: scale(1);
      transition-duration: 0.3s;
      transition-property: transform;
      font-size: 1.8em;
      font-weight: initial;
      align-items: center;
      justify-content: space-between;
    }
    @import '../../styles/input.scss';
  }
</style>
