<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import Label from './Label.svelte'
  const dispatch = createEventDispatcher()
  export let type: string
  export let placeholder: string
  export let label: string
  export let defaultValue: string | number

  let isActive: boolean = false
</script>

<div class:label class="input-container">
  <Label hidden={!isActive || !label}>{label}</Label>
  <input
    {type}
    {placeholder}
    on:input={(e) => {
      isActive = Boolean(e.currentTarget?.value)
      dispatch('change', e)
    }}
    min={type === 'number' ? 0.0 : null}
    value={defaultValue}
  />
</div>

<style lang="scss">
  @import '../../styles/_vars.scss';

  .input-container {
    padding-bottom: 0;
    height: 3rem;
    border-bottom: 1px solid $textColor;
    margin-bottom: 0.5rem;
    position: relative;
    &.label {
      height: 3.5rem;
    }
  }

  input {
    position: absolute;
    bottom: 0.5rem;
    padding: 0;
    margin: 0;
    color: $textColor;
    outline: none;
    border: none;
    width: 100%;
    box-shadow: none;
    &:required {
      box-shadow: none;
    }
  }
</style>
