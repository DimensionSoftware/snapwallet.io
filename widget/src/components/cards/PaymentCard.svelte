<script lang="ts">
  import Card from './Card.svelte'
  import FaIcon from 'svelte-awesome'
  import { faChevronRight } from '@fortawesome/free-solid-svg-icons'
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()

  export let icon
  export let label: string
</script>

<Card on:click={() => dispatch('click')}>
  <div class="payment-card-container">
    <div class="content-container">
      <FaIcon data={icon} />
      <div class="label">{label}</div>
    </div>
    <FaIcon data={faChevronRight} />
  </div>
</Card>

<style lang="scss">
  @import '../../styles/_vars.scss';

  .payment-card-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 0;
    & > .content-container {
      height: 100%;
      width: 100%;
      display: flex;
      justify-content: flex-start;
      align-items: center;
      & > .label {
        font-weight: 500;
        margin-left: 0.5rem;
      }
    }
  }
</style>
