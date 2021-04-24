<script lang="ts">
  import { createEventDispatcher, onMount } from 'svelte'
  import { countries } from '../../util/country'
  import PopupSelector from '../inputs/PopupSelector.svelte'
  import CountryCard from '../cards/CountryCard.svelte'
  import * as Flags from 'svelte-flagicon'
  import VirtualList from '../VirtualList.svelte'

  export let visible = false

  $: filteredCountries = Object.values(countries)

  let searchTimeout

  let listStart
  let listEnd

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
      filteredCountries = Object.values(countries)
        .filter(c => {
          const terms = [c.name, c.code, c.dial_code].join(',').toLowerCase()
          return terms.includes(searchTerm)
        })
        .sort(a => {
          // Sort by the closest text match first
          if (a.name.toLowerCase().startsWith(searchTerm)) return -1
          return 1
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
    {#if filteredCountries.length}
      <VirtualList
        bind:start={listStart}
        bind:end={listEnd}
        items={filteredCountries}
        let:item
      >
        <CountryCard on:click={() => dispatch('select', { country: item })}>
          <div style="display:flex;align-items:center;">
            <svelte:component
              this={Flags[item.code[0] + item.code[1].toLowerCase()]}
            />
            <span style="margin-left:1rem;">{item.name}</span>
          </div>
        </CountryCard>
      </VirtualList>
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
