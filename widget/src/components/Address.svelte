<script lang="ts">
  export let address: string = ''
  export let isExpanded: Boolean = false

  const trunc = s =>
      `${s.substr(0, 7)} ... ${s.substr(s.length - 5, s.length - 1)}`,
    toggleExpanded = () => (isExpanded = !isExpanded)
</script>

<h4
  on:click={toggleExpanded}
  title={address}
  style={`
    transition: height ${
      isExpanded ? '0.2s' : '0.4s'
    } var(--theme-ease-out-expo);
  height: ${isExpanded ? 4 : 1}rem`}
>
  {isExpanded ? address : trunc(address)}
</h4>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';
  h4 {
    cursor: pointer;
    margin: 0 0 0.5rem;
    max-width: 160px;
    white-space: pre-wrap;
    line-height: 1.25rem;
    text-align: left;
    word-break: break-word;
    text-overflow: ellipsis;
    overflow: hidden;
    font-size: 0.8rem;
  }
</style>
