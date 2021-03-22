<script lang="ts">
  export let items = []
</script>

<ul class="vertical-stepper">
  {#each items as item}
    <li>
      <svelte:component this={item} />
    </li>
  {/each}
</ul>

<style lang="scss">
  @import '../styles/_vars.scss';
  .vertical-stepper {
    position: relative;
    display: flex;
  }
</style>
