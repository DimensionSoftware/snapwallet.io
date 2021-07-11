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
  export let isTranslucent: boolean

  let isActive: boolean = Boolean(defaultValue)

  const selectOnFocus = autoselect ? onFocusSelect : _ => {}

  onMount(function () {
    focus(document.querySelector('input[autofocus]'), 200)
  })
</script>

<div class:active={isActive} class:isTranslucent class="input-container">
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
          e.metaKey ||
          ['ArrowLeft', 'ArrowRight', 'ArrowUp', 'ArrowDown'].includes(e.key)

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
    &.isTranslucent input {
      background: transparent;
    }
    @import '../../styles/input.scss';
  }
</style>
