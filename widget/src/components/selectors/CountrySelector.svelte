<script lang="ts">
  import { afterUpdate, createEventDispatcher, onMount } from 'svelte'
  import { countries, getFilteredCountries } from '../../util/country'
  import PopupSelector from '../inputs/PopupSelector.svelte'
  import CountryCard from '../cards/CountryCard.svelte'
  import VirtualList from '../VirtualList.svelte'
  import { userStore } from '../../stores/UserStore'
  import { transactionStore } from '../../stores/TransactionStore'
  import { TransactionMediums } from '../../types'
  import { debitCardStore } from '../../stores/DebitCardStore'
  import { SVG_FLAG_ICON_PATH } from '../../constants'

  export let visible = false
  export let whiteList: string[] = []
  export let selectedCountryCode

  $: filteredCountries = getFilteredCountries(
    Object.values(countries),
    whiteList,
  )

  let isDebitCard = $transactionStore.inMedium === TransactionMediums.DEBIT_CARD

  let searchTimeout

  let listStart
  let listEnd
  let search

  const dispatch = createEventDispatcher()

  const searchCountries = val => {
    const searchTerm = val?.toLowerCase()
    if (!searchTerm) {
      clearTimeout(searchTimeout)
      filteredCountries = getFilteredCountries(
        Object.values(countries),
        whiteList,
      )
      return
    }
    if (searchTimeout) clearTimeout(searchTimeout)
    searchTimeout = debouncedSearch(searchTerm)
  }

  const debouncedSearch = searchTerm => setTimeout(() => doSearch(searchTerm)),
    doSearch = searchTerm =>
      (filteredCountries = getFilteredCountries(
        Object.values(countries),
        whiteList,
      )
        .filter(c => {
          const terms = [c.name, c.code, c.dial_code].join(',').toLowerCase()
          return terms.includes(searchTerm)
        })
        .sort(a => {
          // Sort by the closest text match first
          if (a.name.toLowerCase().startsWith(searchTerm)) return -1
          return 1
        }))

  // Move selected item to top of list
  $: {
    if (selectedCountryCode && !search?.value) {
      const idx = filteredCountries.findIndex(
        fc => fc.code.toLowerCase() === selectedCountryCode.toLowerCase(),
      )
      if (idx > -1) {
        const elem = filteredCountries[idx]
        filteredCountries.splice(idx, 1)
        filteredCountries.unshift(elem)
      }
    }
  }

  // focus search
  onMount(() => {
    setTimeout(() => search.focus(), 300)
  })
</script>

<PopupSelector
  {visible}
  on:close={() => dispatch('close')}
  headerTitle="Select Country"
>
  <div class="selector-container" style="height:90%">
    <input
      placeholder="Search..."
      bind:this={search}
      type="search"
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
            <img
              alt={item.code.toUpperCase()}
              width="32"
              src={`${SVG_FLAG_ICON_PATH}/${
                item.code[0] + item.code[1]
              }.svg`.toLowerCase()}
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
    appearance: none !important;
    width: 100%;
    outline: none !important;
    border: none;
    border-bottom: 1px solid var(--theme-text-color);
    color: var(--theme-text-color);
    margin-bottom: 0.75rem;
    padding: 1rem 0.5rem;
    background: transparent;
    border-radius: none !important;
  }
  :global(.country-select > .fa-icon) {
    position: relative;
    left: -0.4rem;
  }
</style>
