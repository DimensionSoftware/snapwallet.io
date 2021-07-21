<script lang="ts">
  import { onMount, createEventDispatcher } from 'svelte'
  import type { Masks } from '../../types'
  import { withMaskOnInput } from '../../masks'
  import { focus, isValidKeyForMask, onFocusSelect } from '../../util'

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
      if (mask) return isValidKeyForMask(e, mask, defaultValue)
    }}
    on:input={e => {
      // last chance to fix inputs
      const val = e.target.value,
        isValLongerThanMask = val.length > mask.length
      // truncate if necessary
      if (isValLongerThanMask) e.target.value = val.substr(0, val.length - 1)
      // force correct format
      e.target.value = withMaskOnInput(e.target.value, mask)
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
