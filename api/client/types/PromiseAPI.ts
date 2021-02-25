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
import { PricingDataResponse } from '../models/PricingDataResponse';
import { PricingRate } from '../models/PricingRate';
import { ProtobufAny } from '../models/ProtobufAny';
import { RpcStatus } from '../models/RpcStatus';
import { ThirdPartyUserAccount } from '../models/ThirdPartyUserAccount';
import { User } from '../models/User';
import { UserDataResponse } from '../models/UserDataResponse';
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
     * Provides pricing data for all markets with rate maps
     * Get pricing data
     */
    public fluxPricingData(options?: Configuration): Promise<PricingDataResponse> {
    	const result = this.api.fluxPricingData(options);
        return result.toPromise();
    }
	
    /**
     * Provides user data associated with the access token
     * Get user data
     */
    public fluxUserData(options?: Configuration): Promise<UserDataResponse> {
    	const result = this.api.fluxUserData(options);
        return result.toPromise();
    }
	
    /**
     * requires a plaid processor token which in turn requires a plaid widget interaction where the user selects the account id
     * Post chosen bank info from plaid in order to create a new ACH pyment method in wyre
     * @param body 
     */
    public fluxWyreAddBankPaymentMethod(body: any, options?: Configuration): Promise<any> {
    	const result = this.api.fluxWyreAddBankPaymentMethod(body, options);
        return result.toPromise();
    }
	

}



