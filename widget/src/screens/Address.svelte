<script lang="ts">
    import vld8 from 'validator'
    import { push } from 'svelte-spa-router'
    import ModalBody from '../components/ModalBody.svelte'
    import ModalContent from '../components/ModalContent.svelte'
    import ModalFooter from '../components/ModalFooter.svelte'
    import Button from '../components/Button.svelte'
    import Input from '../components/inputs/Input.svelte'
    import Label from '../components/inputs/Label.svelte'
    import ModalHeader from '../components/ModalHeader.svelte'
    import { userStore } from '../stores/UserStore'
    import { toaster } from '../stores/ToastStore'
    import { Logger, onEnterPressed } from '../util'
    import { APIErrors, Routes } from '../constants'
  
    let animation = 'left'
    let isMakingRequest = false
  
    const timeout = 700
  
    const handleNextStep = async () => {
        window.API.fluxSaveProfileData({
            address: {
                street1:"",
                street2:"",
                city:"",
                state:"",
                postalCode:"",
                country:"",
            }

        })
    }
  </script>
  
  <ModalContent {animation}>
    <ModalBody>
      <ModalHeader hideBackButton>Hi, we need your address dawg</ModalHeader>
      <Label label="Street 1">
        <Input type="text"  placeholder="street1" />
      </Label>
      <Label label="Street 2">
        <Input type="text"  placeholder="street2" />
      </Label>
      <Label label="City">
        <Input type="text"  placeholder="city" />
      </Label>
      <Label label="State">
        <Input type="text"  placeholder="state" />
      </Label>
      <Label label="PostalCode">
        <Input type="text"  placeholder="postalCode" />
      </Label>
      <Label label="Country">
        <Input type="text"  placeholder="country" />
      </Label>
    </ModalBody>
  </ModalContent>
  
  <style lang="scss">
    @import '../styles/_vars.scss';
  </style>
  