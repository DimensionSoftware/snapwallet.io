<script lang="ts">
  import FaIcon from 'svelte-awesome'
  import { faClock } from '@fortawesome/free-solid-svg-icons'
  export let time: string = ''
</script>

<div class="time-ticker">
  <FaIcon data={faClock} />
  <div style="margin-right:0.35rem;" />
  <div>{time}</div>
</div>

<style lang="scss">
  .time-ticker {
    display: flex;
    margin-bottom: 1.25rem;
    display: flex;
    align-items: center;
    justify-content: flex-end;
    color: var(--theme-text-color-3);
  }
</style>
