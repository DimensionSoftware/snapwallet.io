import { ResponseContext, RequestContext, HttpFile } from '../http/http';
import * as models from '../models/all';
import { Configuration} from '../configuration'

import { Address } from '../models/Address';
import { ApiHttpBody } from '../models/ApiHttpBody';
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
import { ObservableFluxApi } from './ObservableAPI';


import { FluxApiRequestFactory, FluxApiResponseProcessor} from "../apis/FluxApi";
export class PromiseFluxApi {
    private api: ObservableFluxApi

    public constructor(
        configuration: Configuration,
        requestFactory?: FluxApiRequestFactory,
        responseProcessor?: FluxApiResponseProcessor
    ) {
        this.api = new ObservableFluxApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * Will cause your email or phone to receive a one time passcode. This can be used in the verify step to obtain a token for login
     * Post email or phone in exchange for a one time passcode
     * @param body 
     */
    public fluxOneTimePasscode(body: OneTimePasscodeRequest, options?: Configuration): Promise<any> {
    	const result = this.api.fluxOneTimePasscode(body, options);
        return result.toPromise();
    }
	
    /**
     * The passcode received in either email or phone text message should be provided here in order to obtain on access token
     * Post one time passcode in exchange for an access token
     * @param body 
     */
    public fluxOneTimePasscodeVerify(body: OneTimePasscodeVerifyRequest, options?: Configuration): Promise<OneTimePasscodeVerifyResponse> {
    	const result = this.api.fluxOneTimePasscodeVerify(body, options);
        return result.toPromise();
    }
	
    /**
     * requires a plaid processor token which in turn requires a plaid widget interaction where the user selects the account id
     * Post chosen bank info from plaid in order to create a new ACH pyment method in wyre
     * @param body 
     */
    public fluxPlaidConnectBankAccounts(body: PlaidConnectBankAccountsRequest, options?: Configuration): Promise<any> {
    	const result = this.api.fluxPlaidConnectBankAccounts(body, options);
        return result.toPromise();
    }
	
    /**
     * https://plaid.com/docs/link/link-token-migration-guide/
     * @param body 
     */
    public fluxPlaidCreateLinkToken(body: any, options?: Configuration): Promise<PlaidCreateLinkTokenResponse> {
    	const result = this.api.fluxPlaidCreateLinkToken(body, options);
        return result.toPromise();
    }
	
    /**
     * Provides pricing data for all markets with rate maps
     * Get pricing data
     */
    public fluxPricingData(options?: Configuration): Promise<PricingDataResponse> {
    	const result = this.api.fluxPricingData(options);
        return result.toPromise();
    }
	
    /**
     * https://github.com/googleapis/googleapis/blob/master/google/api/httpbody.proto
     * @param body 
     */
    public fluxUploadFile(body: ApiHttpBody, options?: Configuration): Promise<ApiHttpBody> {
    	const result = this.api.fluxUploadFile(body, options);
        return result.toPromise();
    }
	
    /**
     * Provides user (viewer) data associated with the access token
     * Get viewer data
     */
    public fluxViewerData(options?: Configuration): Promise<ViewerDataResponse> {
    	const result = this.api.fluxViewerData(options);
        return result.toPromise();
    }
	
    /**
     * https://plaid.com/docs/link/link-token-migration-guide/
     * @param body 
     */
    public fluxWyreCreateAccount(body: WyreCreateAccountRequest, options?: Configuration): Promise<any> {
    	const result = this.api.fluxWyreCreateAccount(body, options);
        return result.toPromise();
    }
	

}



