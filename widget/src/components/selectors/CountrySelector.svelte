<script lang="ts">
  import { afterUpdate, createEventDispatcher, onMount } from 'svelte'
  import { countries, getFilteredCountries } from '../../util/country'
  import PopupSelector from '../inputs/PopupSelector.svelte'
  import CountryCard from '../cards/CountryCard.svelte'
  import * as Flags from 'svelte-flagicon'
  import VirtualList from '../VirtualList.svelte'
  import { userStore } from '../../stores/UserStore'
  import { transactionStore } from '../../stores/TransactionStore'
  import { TransactionMediums } from '../../types'
  import { debitCardStore } from '../../stores/DebitCardStore'

  export let visible = false
  export let whiteList: string[] = []

  $: filteredCountries = getFilteredCountries(
    Object.values(countries),
    whiteList,
  )

  let isDebitCard = $transactionStore.inMedium === TransactionMediums.DEBIT_CARD
  let selectedCountry = isDebitCard
    ? $debitCardStore.address.country
    : $userStore.geo.country

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

  // focus search
  onMount(() => {
    setTimeout(() => search.focus(), 300)
  })

  afterUpdate(() => {
    // Move the selected country to top of list
    if (selectedCountry && !search?.value) {
      const idx = filteredCountries.findIndex(fc => fc.code === selectedCountry)
      if (idx > -1) {
        const elem = filteredCountries[idx]
        filteredCountries.splice(idx, 1)
        filteredCountries.unshift(elem)
      }
    }
  })
</script>

<PopupSelector
  {visible}
  on:close={() => dispatch('close')}
  headerTitle="Select Country"
>
  <div class="selector-container">
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
    border-radius: none;
  }
  :global(.country-select > .fa-icon) {
    position: relative;
    left: -0.4rem;
  }
</style>
