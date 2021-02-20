<script>
  import ModalBody from '../components/ModalBody.svelte'
  import ModalContent from '../components/ModalContent.svelte'
  import ModalFooter from '../components/ModalFooter.svelte'
  import IntentSelector from '../components/IntentSelector.svelte'
  import { toaster } from '../stores/ToastStore'
  import Button from '../components/Button.svelte'
  import { userStore } from '../stores/UserStore'
  import Input from '../components/inputs/Input.svelte'
  import PopupSelector from '../components/inputs/PopupSelector.svelte'
  import CryptCard from '../components/cards/CryptoCard.svelte'
  import Card from '../components/cards/Card.svelte'

  let selectorVisible = false
  const handleClick = () => {
    // toaster.pop({ msg: 'Success', success: true })
    // TODO: remove
    selectorVisible = true
  }
</script>

<ModalContent>
  <ModalBody>
    <IntentSelector />
    <Input label="Name" placeholder="Name" />
    <Card>Stuff</Card>
    <PopupSelector
      on:close={() => (selectorVisible = false)}
      visible={selectorVisible}
      headerTitle="Select Cryptocurrency"
    >
      <CryptCard name="Bitcoin" ticker="BTC" />
      <CryptCard name="Ethereum" ticker="ETH" />
      <CryptCard name="Tether" ticker="USDT" />
      <CryptCard name="USDC" ticker="USDC" />
    </PopupSelector>
  </ModalBody>
  <ModalFooter>
    <Button on:click={handleClick}>{$userStore.intent}</Button></ModalFooter
  >
</ModalContent>
