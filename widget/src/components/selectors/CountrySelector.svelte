<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import { countries } from '../../util/country'
  import PopupSelector from '../inputs/PopupSelector.svelte'
  import CountryCard from '../cards/CountryCard.svelte'
  import * as Flags from 'svelte-flagicon'
  import type { ICountry } from '../../types'

  export let visible = false

  let filteredCountries: ICountry[] = Object.values(countries)
  let searchTimeout

  const dispatch = createEventDispatcher()

  const searchCountries = val => {
    const searchTerm = val?.toLowerCase()
    if (!searchTerm) {
      clearTimeout(searchTimeout)
      filteredCountries = Object.values(countries)
      return
    }
    if (searchTimeout) clearTimeout(searchTimeout)
    searchTimeout = debounceSearch(searchTerm)
  }

  const debounceSearch = searchTerm => {
    return setTimeout(() => {
      filteredCountries = Object.values(countries).filter(c => {
        const terms = [c.name, c.code, c.dial_code].join(',').toLowerCase()
        return terms.includes(searchTerm)
      })
    }, 400)
  }
</script>

<PopupSelector
  {visible}
  on:close={() => dispatch('close')}
  headerTitle="Select Country"
>
  <div class="scroll-y selector-container">
    <input
      placeholder="Search..."
      class="search-input"
      on:input={e => {
        searchCountries(e.target?.value)
      }}
    />
    <CountryCard
      on:click={() => dispatch('select', { country: countries['US'] })}
    >
      <div style="display:flex;align-items:center;">
        <Flags.Us />
        <span style="margin-left:1rem;">United States</span>
      </div>
    </CountryCard>
    <CountryCard
      on:click={() => dispatch('select', { country: countries['GB'] })}
    >
      <div style="display:flex;align-items:center;">
        <Flags.Gb />
        <span style="margin-left:1rem;">United Kingdom</span>
      </div>
    </CountryCard>

    <h5>Countries</h5>
    {#if filteredCountries.length}
      {#each filteredCountries as country}
        <CountryCard on:click={() => dispatch('select', { country })}>
          <div style="display:flex;align-items:center;">
            <svelte:component
              this={Flags[country.code[0] + country.code[1].toLowerCase()]}
            />
            <span style="margin-left:1rem;">{country.name}</span>
          </div>
        </CountryCard>
      {/each}
    {:else}
      No countries were found
    {/if}
    <div class="spacer" />
  </div>
</PopupSelector>

<style lang="scss">
  @import '../../styles/selectors.scss';
  .search-input {
    appearance: none;
    width: 100%;
    outline: none;
    border: none;
    border-bottom: 1px solid var(--theme-text-color);
    color: var(--theme-text-color);
    margin-bottom: 0.75rem;
    padding-left: 0;
    background: var(--theme-modal-popup-background);
  }
  :global(.country-select > .fa-icon) {
    position: relative;
    left: -0.4rem;
  }
</style>
