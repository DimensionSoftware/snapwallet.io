<script lang="ts">
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import Button from '../components/Button.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import { transactionStore } from '../stores/TransactionStore'
  import { formatLocaleCurrency } from '../util'

  const transactions = []
</script>

<ModalContent>
  <ModalBody>
    <ModalHeader>Transactions</ModalHeader>
    <div class="line-items">
      <ol>
        {#each transactions as transaction}
          <li>
            {#if transaction.isBuy}
              <div class="line-item muted">
                <div>From</div>
                <div>source</div>
              </div>
              <div class="line-item muted">
                <div>To</div>
                <div>to deets</div>
              </div>
            {:else}
              <div class="line-item muted">
                <div>From</div>
                <div>from deets</div>
              </div>
              <div class="line-item muted">
                <div>To</div>
                <div>destination</div>
              </div>
            {/if}
          </li>
        {/each}
      </ol>
      <div class="line dashed" />
      <div class="line-item">
        <div><b>Total</b></div>
        <div>
          <!-- {formatLocaleCurrency(fiatTicker, total)} -->
        </div>
      </div>
    </div>
  </ModalBody>
  <ModalFooter>
    <Button on:click={close}>OK</Button>
  </ModalFooter>
</ModalContent>

<style lang="scss">
  @import '../styles/_vars.scss';
  @import '../styles/animations.scss';

  .checkout-item-box {
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    .checkout-item-name {
      margin-top: 1.4rem;
      height: 100%;
      display: flex;
      flex-direction: column;
      justify-content: flex-start;
      font-weight: bold;
      font-size: 1.3rem;
    }
  }

  .line {
    height: 1px;
    max-height: 1px;
    width: 100%;
    border-bottom: 0.5px solid $textColor4;
    margin: 0.5rem 0 0.5rem 0;
    &.dashed {
      border-bottom: 0.7px dashed $textColor4;
    }
  }

  .line-items {
    width: 90%;
    align-self: center;
    margin-top: 1rem;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    & > .line-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
      &.muted {
        color: $textColor4;
        font-weight: 300;
      }
    }
  }
</style>
