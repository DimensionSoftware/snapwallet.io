<script lang="ts">
  import ModalContent from '../components/ModalContent.svelte'
  import ModalBody from '../components/ModalBody.svelte'
  import ModalHeader from '../components/ModalHeader.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import { CryptoIcons, formatLocaleCurrency, dropEndingZeros } from '../util'
  import { transactionDetailsStore } from '../stores/TransactionsStore'
  import Badge from '../components/Badge.svelte'

  let { transaction } = $transactionDetailsStore
  let status = transaction.status.toLowerCase()
  let createdAt = new Date(transaction.createdAt)

  const cryptoFee = transaction.fees
    ? transaction.fees[transaction.destCurrency] / transaction.exchangeRate
    : 0
  const trueSourceAmount = transaction.fees
    ? transaction.sourceAmount -
      cryptoFee -
      transaction.fees[transaction.sourceCurrency]
    : 0

  const Icon = CryptoIcons[transaction.destCurrency]
</script>

<ModalContent>
  <ModalHeader>Details</ModalHeader>
  <ModalBody>
    <div class="checkout-item-box">
      <div style="width:30%;" class="checkout-item-icon">
        <Icon size="100%" height="100%" width="100%" viewBox="-4 0 40 40" />
      </div>
      <div class="checkout-item-name">
        {dropEndingZeros(transaction.destAmount.toFixed(8))}
        {transaction.destCurrency}
      </div>
    </div>
    <div class="line-items">
      <!-- ACH -->
      <div class="line-item muted">
        <div>Status</div>
        <div style="text-transform:capitalize;">
          <Badge
            success={status === 'completed'}
            warning={status === 'pending'}
            error={status === 'failed'}
          >
            {status}
          </Badge>
        </div>
      </div>
      <div class="line dashed" />
      <div class="line-item muted">
        <div>ID</div>
        <div>{transaction.id}</div>
      </div>
      <div class="line-item muted">
        <div>Date</div>
        <div>
          {createdAt.toLocaleDateString()}
        </div>
      </div>
      <div class="line-item muted">
        <div>Time</div>
        <div>
          {createdAt.toLocaleTimeString()}
        </div>
      </div>
      <div class="line dashed" />
      <div class="line-item muted">
        <div>From</div>
        <div>{transaction.sourceName}</div>
      </div>
      <div class="line-item muted">
        <div>To</div>
        <div>
          {transaction.dest.substring(0, 6)}...{transaction.dest.substring(
            transaction.dest.length - 4,
          )}
        </div>
      </div>
      <div class="line dashed" />
      <div class="line-item muted">
        <div>Subtotal</div>
        <div>
          {formatLocaleCurrency(transaction.sourceCurrency, trueSourceAmount)}
        </div>
      </div>
      <div class="line-item muted">
        <div>Crypto Fee</div>
        <div>
          {formatLocaleCurrency(transaction.sourceCurrency, cryptoFee)}
        </div>
      </div>
      <div class="line-item muted">
        <div>Service Fee</div>
        <div>
          {formatLocaleCurrency(
            transaction.sourceCurrency,
            transaction.fees ? transaction.fees[transaction.sourceCurrency] : 0,
          )}
        </div>
      </div>
      <div class="line dashed" />
      <div class="line-item">
        <div><b>Total</b></div>
        <div>
          <b
            >{formatLocaleCurrency(
              transaction.sourceCurrency,
              transaction.sourceAmount,
            )}</b
          >
        </div>
      </div>
    </div>
  </ModalBody>
  <ModalFooter />
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
      margin-top: 0.5rem;
      height: 100%;
      display: flex;
      flex-direction: column;
      justify-content: flex-start;
      font-weight: bold;
      font-size: 1.5rem;
    }
  }

  .line {
    height: 1px;
    max-height: 1px;
    width: 100%;
    border-bottom: 0.5px solid var(--theme-text-color);
    margin: 0.5rem 0 0.5rem 0;
    &.dashed {
      border-bottom: 0.7px dashed var(--theme-text-color);
    }
  }

  .line-items {
    width: 100%;
    align-self: center;
    margin-top: 2.5rem;
    padding: 0 0.7rem;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    & > .line-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
      &.muted {
        color: var(--theme-color-muted);
        font-weight: 300;
      }
    }
  }
</style>
