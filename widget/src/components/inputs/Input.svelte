<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import Label from './Label.svelte'
  const dispatch = createEventDispatcher()
  export let type: string
  export let placeholder: string
  export let label: string
  export let forceLabel: boolean
  export let defaultValue: string | number

  let isActive: boolean = Boolean(forceLabel)
</script>

<div class="input-container">
  <Label hidden={!forceLabel || !isActive || !label}>{label}</Label>
  <input
    {type}
    placeholder={forceLabel ? '' : placeholder}
    on:input={(e) => {
      isActive = forceLabel || Boolean(e.currentTarget?.value)
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
    height: 50px;
    border-bottom: 1px solid $textColor1;
    margin-bottom: 0.5rem;
  }

  input {
    padding: 0;
    padding-top: 0.5rem;
    margin: 0;
    color: $textColor1;
    outline: none;
    border: none;
    width: 100%;
    box-shadow: none;
    &:required {
      box-shadow: none;
    }
  }
</style>
