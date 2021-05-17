<script lang="ts">
  import VirtualList from '../VirtualList.svelte'
  import { createEventDispatcher } from 'svelte'
  import { SUPPORTED_CRYPTOCURRENCY_ASSETS } from '../../constants'
  import { fly } from 'svelte/transition'
  import CryptoCard from '../cards/CryptoCard.svelte'
  import Label from '../inputs/Label.svelte'
  import PopupSelector from '../inputs/PopupSelector.svelte'
  const dispatch = createEventDispatcher()

  export let visible = false

  const items = SUPPORTED_CRYPTOCURRENCY_ASSETS.sort((a, b) => {
    // float popular to top
    if (a.popular && b.popular) return 0
    if (a.popular) return -1
    if (b.popular) return 1
    // alpha
    if (a.name < b.name) return -1
    if (a.name > b.name) return 1
    return 0
  })
</script>

{#if visible}
  <PopupSelector
    on:close={() => dispatch('close')}
    headerTitle="Select Currency"
  >
    <div class="selector-container">
      <VirtualList {items} let:item>
        <Label fx={false}>
          <CryptoCard on:mousedown={() => dispatch('close')} crypto={item} />
        </Label>
      </VirtualList>
      <div class="spacer" />
    </div>
  </PopupSelector>
{/if}

<style lang="scss">
  @import '../../styles/selectors.scss';
  .selector-container {
    padding-top: 0;
  }
  :global(svelte-virtual-list-viewport) {
    height: 120% !important;
    padding-bottom: 50% !important;
  }
  :global(svelte-virtual-list-row) {
    padding-left: 7px;
  }
</style>
