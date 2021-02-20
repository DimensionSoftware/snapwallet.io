<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()
  export let type: string
  export let placeholder: string
  export let label: string

  let isActive: boolean = false
</script>

<div class="input-container">
  {#if label}
    <div class="input-label">{isActive ? label : ''}</div>
  {/if}
  <input
    {type}
    {placeholder}
    on:input={(e) => {
      isActive = Boolean(e.currentTarget?.value)
      dispatch('change', e)
    }}
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

  .input-label {
    font-size: 0.7rem;
    min-height: 20%;
    font-weight: 600;
  }

  input {
    padding: 0;
    padding-top: 0.5rem;
    margin: 0;
    color: $textColor1;
    outline: none;
    border: none;
    width: 100%;
  }
</style>
