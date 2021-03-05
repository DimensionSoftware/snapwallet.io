import { ResponseContext, RequestContext, HttpFile } from '../http/http';
import * as models from '../models/all';
import { Configuration} from '../configuration'

import { Address } from '../models/Address';
import { OneTimePasscodeRequest } from '../models/OneTimePasscodeRequest';
import { OneTimePasscodeVerifyRequest } from '../models/OneTimePasscodeVerifyRequest';
import { OneTimePasscodeVerifyResponse } from '../models/OneTimePasscodeVerifyResponse';
import { PlaidConnectBankAccountsRequest } from '../models/PlaidConnectBankAccountsRequest';
import { PlaidCreateLinkTokenResponse } from '../models/PlaidCreateLinkTokenResponse';
import { PricingDataResponse } from '../models/PricingDataResponse';
import { PricingRate } from '../models/PricingRate';
import { ProtobufAny } from '../models/ProtobufAny';
import { RpcStatus } from '../models/RpcStatus';
import { User } from '../models/User';
import { UserFlags } from '../models/UserFlags';
import { ViewerDataResponse } from '../models/ViewerDataResponse';
import { WyreCreateAccountRequest } from '../models/WyreCreateAccountRequest';

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

export interface FluxApiFluxPlaidConnectBankAccountsRequest {
    /**
     * 
     * @type PlaidConnectBankAccountsRequest
     * @memberof FluxApifluxPlaidConnectBankAccounts
     */
    body: PlaidConnectBankAccountsRequest
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

export interface FluxApiFluxViewerDataRequest {
}

export interface FluxApiFluxWyreCreateAccountRequest {
    /**
     * 
     * @type WyreCreateAccountRequest
     * @memberof FluxApifluxWyreCreateAccount
     */
    body: WyreCreateAccountRequest
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
     * requires a plaid processor token which in turn requires a plaid widget interaction where the user selects the account id
     * Post chosen bank info from plaid in order to create a new ACH pyment method in wyre
     * @param param the request object
     */
    public fluxPlaidConnectBankAccounts(param: FluxApiFluxPlaidConnectBankAccountsRequest, options?: Configuration): Promise<any> {
        return this.api.fluxPlaidConnectBankAccounts(param.body,  options).toPromise();
    }
	
    /**
     * https://plaid.com/docs/link/link-token-migration-guide/
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
     * Provides user (viewer) data associated with the access token
     * Get viewer data
     * @param param the request object
     */
    public fluxViewerData(param: FluxApiFluxViewerDataRequest, options?: Configuration): Promise<ViewerDataResponse> {
        return this.api.fluxViewerData( options).toPromise();
    }
	
    /**
     * https://plaid.com/docs/link/link-token-migration-guide/
     * @param param the request object
     */
    public fluxWyreCreateAccount(param: FluxApiFluxWyreCreateAccountRequest, options?: Configuration): Promise<any> {
        return this.api.fluxWyreCreateAccount(param.body,  options).toPromise();
    }
	

}



