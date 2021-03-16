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
      <label label="Street 1">
        <input type="text"  placeholder="street1" />
      </label>
      <label label="Street 2">
        <input type="text"  placeholder="street2" />
      </label>
      <label label="City">
        <input type="text"  placeholder="city" />
      </label>
      <label label="State">
        <input type="text"  placeholder="state" />
      </label>
      <label label="PostalCode">
        <input type="text"  placeholder="postalCode" />
      </label>
      <label label="Country">
        <input type="text"  placeholder="country" />
      </label>
    </ModalBody>
    <button>Plz give me now</button>
  </ModalContent>
  
  <style lang="scss">
    @import '../styles/_vars.scss';
    label {
        size: smaller;
    }
  </style>
  