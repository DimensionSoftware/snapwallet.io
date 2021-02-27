<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import Label from './Label.svelte'
  const dispatch = createEventDispatcher()
  export let type: string = 'text'
  export let placeholder: string = ''
  export let inputmode: string = 'text'
  export let label: string = ''
  export let autocapitalize: string = ''
  export let defaultValue: string | number = ''
  export let autocomplete: string = 'on'

  let isActive: boolean = Boolean(defaultValue)
</script>

<div class:label class:active={isActive} class="input-container">
  <Label hidden={!isActive || !label}>{label}</Label>
  <input
    {type}
    {inputmode}
    {autocapitalize}
    {placeholder}
    {autocomplete}
    on:input={e => {
      isActive = Boolean(e.currentTarget?.value)
      dispatch('change', e.target.value)
    }}
    min={type === 'number' ? 0.0 : null}
    value={defaultValue || ''}
  />
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';

  .input-container {
    padding-bottom: 0;
    height: 3rem;
    border-bottom: 1px solid;
    border-bottom-color: var(--theme-color);
    margin-bottom: 0.5rem;
    position: relative;
    &.label {
      height: 3.5rem;
    }
    &.active {
      border-bottom-width: 1.5px;
    }
  }

  input {
    position: absolute;
    bottom: 0.5rem;
    padding: 0;
    margin: 0;
    color: var(--theme-text-color);
    outline: none;
    border: none;
    width: 100%;
    box-shadow: none;
    &:required {
      box-shadow: none;
    }
    // HACK: remove yellow autofill background
    &:-webkit-autofill,
    &:-webkit-autofill:hover,
    &:-webkit-autofill:focus,
    &:-webkit-autofill:active {
      box-shadow: 0 0 0px 1000px var(--theme-modal-background-color) inset !important;
      background-color: var(--theme-modal-background-color) !important;
      background-clip: content-box !important;
      -webkit-text-fill-color: var(--theme-text-color) !important;
    }
  }
</style>
