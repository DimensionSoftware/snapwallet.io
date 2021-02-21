<script lang="ts">
  export let hidden: boolean
</script>

<div class="input-label">
  {#if !hidden}
    <slot />
  {/if}
</div>

<style lang="scss">
  .input-label {
    font-size: 0.7rem;
    height: 25%;
    font-weight: 600;
  }
</style>
