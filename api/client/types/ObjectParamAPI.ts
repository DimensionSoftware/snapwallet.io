import { ResponseContext, RequestContext, HttpFile } from '../http/http';
import * as models from '../models/all';
import { Configuration} from '../configuration'

import { Address } from '../models/Address';
import { KYCProfile } from '../models/KYCProfile';
import { OneTimePasscodeRequest } from '../models/OneTimePasscodeRequest';
import { OneTimePasscodeVerifyRequest } from '../models/OneTimePasscodeVerifyRequest';
import { OneTimePasscodeVerifyResponse } from '../models/OneTimePasscodeVerifyResponse';
import { Organization } from '../models/Organization';
import { OrganizationApplication } from '../models/OrganizationApplication';
import { PaymentMethod } from '../models/PaymentMethod';
import { PlaidCreateLinkTokenResponse } from '../models/PlaidCreateLinkTokenResponse';
import { PricingDataResponse } from '../models/PricingDataResponse';
import { PricingRate } from '../models/PricingRate';
import { ProtobufAny } from '../models/ProtobufAny';
import { RpcStatus } from '../models/RpcStatus';
import { ThirdPartyUserAccount } from '../models/ThirdPartyUserAccount';
import { User } from '../models/User';
import { UserDataResponse } from '../models/UserDataResponse';
import { WyreAddBankPaymentMethodsRequest } from '../models/WyreAddBankPaymentMethodsRequest';

import { ObservableFluxApi } from "./ObservableAPI";
import { FluxApiRequestFactory, FluxApiResponseProcessor} from "../apis/FluxApi";

export interface FluxApiFluxOneTimePasscodeRequest {
    /**
     * 
     * @type OneTimePasscodeRequest
     * @memberof FluxApifluxOneTimePasscode
     */
    body: OneTimePasscodeRequest
}

export interface FluxApiFluxOneTimePasscodeVerifyRequest {
    /**
     * 
     * @type OneTimePasscodeVerifyRequest
     * @memberof FluxApifluxOneTimePasscodeVerify
     */
    body: OneTimePasscodeVerifyRequest
}

export interface FluxApiFluxPlaidCreateLinkTokenRequest {
    /**
     * 
     * @type any
     * @memberof FluxApifluxPlaidCreateLinkToken
     */
    body: any
}

export interface FluxApiFluxPricingDataRequest {
}

export interface FluxApiFluxUserDataRequest {
}

export interface FluxApiFluxWyreAddBankPaymentMethodsRequest {
    /**
     * 
     * @type WyreAddBankPaymentMethodsRequest
     * @memberof FluxApifluxWyreAddBankPaymentMethods
     */
    body: WyreAddBankPaymentMethodsRequest
}


export class ObjectFluxApi {
    private api: ObservableFluxApi

    public constructor(configuration: Configuration, requestFactory?: FluxApiRequestFactory, responseProcessor?: FluxApiResponseProcessor) {
        this.api = new ObservableFluxApi(configuration, requestFactory, responseProcessor);
	}

    /**
     * Will cause your email or phone to receive a one time passcode. This can be used in the verify step to obtain a token for login
     * Post email or phone in exchange for a one time passcode
     * @param param the request object
     */
    public fluxOneTimePasscode(param: FluxApiFluxOneTimePasscodeRequest, options?: Configuration): Promise<any> {
        return this.api.fluxOneTimePasscode(param.body,  options).toPromise();
    }
	
    /**
     * The passcode received in either email or phone text message should be provided here in order to obtain on access token
     * Post one time passcode in exchange for an access token
     * @param param the request object
     */
    public fluxOneTimePasscodeVerify(param: FluxApiFluxOneTimePasscodeVerifyRequest, options?: Configuration): Promise<OneTimePasscodeVerifyResponse> {
        return this.api.fluxOneTimePasscodeVerify(param.body,  options).toPromise();
    }
	
    /**
     * PlaidCreateLinkToken implements this flow: https://plaid.com/docs/link/link-token-migration-guide/
     * @param param the request object
     */
    public fluxPlaidCreateLinkToken(param: FluxApiFluxPlaidCreateLinkTokenRequest, options?: Configuration): Promise<PlaidCreateLinkTokenResponse> {
        return this.api.fluxPlaidCreateLinkToken(param.body,  options).toPromise();
    }
	
    /**
     * Provides pricing data for all markets with rate maps
     * Get pricing data
     * @param param the request object
     */
    public fluxPricingData(param: FluxApiFluxPricingDataRequest, options?: Configuration): Promise<PricingDataResponse> {
        return this.api.fluxPricingData( options).toPromise();
    }
	
    /**
     * Provides user data associated with the access token
     * Get user data
     * @param param the request object
     */
    public fluxUserData(param: FluxApiFluxUserDataRequest, options?: Configuration): Promise<UserDataResponse> {
        return this.api.fluxUserData( options).toPromise();
    }
	
    /**
     * requires a plaid processor token which in turn requires a plaid widget interaction where the user selects the account id
     * Post chosen bank info from plaid in order to create a new ACH pyment method in wyre
     * @param param the request object
     */
    public fluxWyreAddBankPaymentMethods(param: FluxApiFluxWyreAddBankPaymentMethodsRequest, options?: Configuration): Promise<any> {
        return this.api.fluxWyreAddBankPaymentMethods(param.body,  options).toPromise();
    }
	

}



