import { ResponseContext, RequestContext, HttpFile } from '../http/http';
import * as models from '../models/all';
import { Configuration} from '../configuration'

import { Address } from '../models/Address';
import { ChangeViewerEmailRequest } from '../models/ChangeViewerEmailRequest';
import { ChangeViewerPhoneRequest } from '../models/ChangeViewerPhoneRequest';
import { DocumentInput } from '../models/DocumentInput';
import { GetImageResponse } from '../models/GetImageResponse';
import { ImageProcessingMode } from '../models/ImageProcessingMode';
import { InlineResponse200 } from '../models/InlineResponse200';
import { LifecycleStatus } from '../models/LifecycleStatus';
import { OneTimePasscodeRequest } from '../models/OneTimePasscodeRequest';
import { OneTimePasscodeVerifyRequest } from '../models/OneTimePasscodeVerifyRequest';
import { OneTimePasscodeVerifyResponse } from '../models/OneTimePasscodeVerifyResponse';
import { PlaidConnectBankAccountsRequest } from '../models/PlaidConnectBankAccountsRequest';
import { PlaidCreateLinkTokenResponse } from '../models/PlaidCreateLinkTokenResponse';
import { PricingDataResponse } from '../models/PricingDataResponse';
import { PricingRate } from '../models/PricingRate';
import { ProfileDataInfo } from '../models/ProfileDataInfo';
import { ProfileDataItemInfo } from '../models/ProfileDataItemInfo';
import { ProfileDataItemKind } from '../models/ProfileDataItemKind';
import { ProfileDataItemStatus } from '../models/ProfileDataItemStatus';
import { ProtobufAny } from '../models/ProtobufAny';
import { RpcStatus } from '../models/RpcStatus';
import { SaveProfileDataRequest } from '../models/SaveProfileDataRequest';
import { ThirdPartyUserAccount } from '../models/ThirdPartyUserAccount';
import { TokenExchangeRequest } from '../models/TokenExchangeRequest';
import { TokenExchangeResponse } from '../models/TokenExchangeResponse';
import { TokenMaterial } from '../models/TokenMaterial';
import { UploadFileResponse } from '../models/UploadFileResponse';
import { UsGovernmentIdDocumentInput } from '../models/UsGovernmentIdDocumentInput';
import { UsGovernmentIdDocumentInputKind } from '../models/UsGovernmentIdDocumentInputKind';
import { User } from '../models/User';
import { UserAccountRemediation } from '../models/UserAccountRemediation';
import { UserFlags } from '../models/UserFlags';
import { ViewerDataResponse } from '../models/ViewerDataResponse';
import { WyreConfirmTransferRequest } from '../models/WyreConfirmTransferRequest';
import { WyreCreateTransferRequest } from '../models/WyreCreateTransferRequest';
import { WyrePaymentMethod } from '../models/WyrePaymentMethod';
import { WyrePaymentMethods } from '../models/WyrePaymentMethods';
import { WyreTransfer } from '../models/WyreTransfer';
import { WyreTransfers } from '../models/WyreTransfers';
import { WyreWebhookRequest } from '../models/WyreWebhookRequest';

import { ObservableFluxApi } from "./ObservableAPI";
import { FluxApiRequestFactory, FluxApiResponseProcessor} from "../apis/FluxApi";

export interface FluxApiFluxChangeViewerEmailRequest {
    /**
     * 
     * @type ChangeViewerEmailRequest
     * @memberof FluxApifluxChangeViewerEmail
     */
    body: ChangeViewerEmailRequest
}

export interface FluxApiFluxChangeViewerPhoneRequest {
    /**
     * 
     * @type ChangeViewerPhoneRequest
     * @memberof FluxApifluxChangeViewerPhone
     */
    body: ChangeViewerPhoneRequest
}

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

export interface FluxApiFluxSaveProfileDataRequest {
    /**
     * 
     * @type SaveProfileDataRequest
     * @memberof FluxApifluxSaveProfileData
     */
    body: SaveProfileDataRequest
}

export interface FluxApiFluxTokenExchangeRequest {
    /**
     * 
     * @type TokenExchangeRequest
     * @memberof FluxApifluxTokenExchange
     */
    body: TokenExchangeRequest
}

export interface FluxApiFluxUploadFileRequest {
    /**
     * The file to upload.
     * @type HttpFile
     * @memberof FluxApifluxUploadFile
     */
    file?: HttpFile
}

export interface FluxApiFluxViewerDataRequest {
}

export interface FluxApiFluxViewerProfileDataRequest {
}

export interface FluxApiFluxWyreConfirmTransferRequest {
    /**
     * 
     * @type string
     * @memberof FluxApifluxWyreConfirmTransfer
     */
    transferId: string
    /**
     * 
     * @type WyreConfirmTransferRequest
     * @memberof FluxApifluxWyreConfirmTransfer
     */
    body: WyreConfirmTransferRequest
}

export interface FluxApiFluxWyreCreateTransferRequest {
    /**
     * 
     * @type WyreCreateTransferRequest
     * @memberof FluxApifluxWyreCreateTransfer
     */
    body: WyreCreateTransferRequest
}

export interface FluxApiFluxWyreGetPaymentMethodsRequest {
}

export interface FluxApiFluxWyreGetTransferRequest {
    /**
     * 
     * @type string
     * @memberof FluxApifluxWyreGetTransfer
     */
    transferId: string
}

export interface FluxApiFluxWyreGetTransfersRequest {
}

export interface FluxApiFluxWyreWebhookRequest {
    /**
     * 
     * @type string
     * @memberof FluxApifluxWyreWebhook
     */
    hookId: string
    /**
     * 
     * @type WyreWebhookRequest
     * @memberof FluxApifluxWyreWebhook
     */
    body: WyreWebhookRequest
}


export class ObjectFluxApi {
    private api: ObservableFluxApi

    public constructor(configuration: Configuration, requestFactory?: FluxApiRequestFactory, responseProcessor?: FluxApiResponseProcessor) {
        this.api = new ObservableFluxApi(configuration, requestFactory, responseProcessor);
	}

    /**
     * requires an otp code and the desired email address change
     * Change users email (viewer based on jwt)
     * @param param the request object
     */
    public fluxChangeViewerEmail(param: FluxApiFluxChangeViewerEmailRequest, options?: Configuration): Promise<any> {
        return this.api.fluxChangeViewerEmail(param.body,  options).toPromise();
    }
	
    /**
     * requires an otp code and the desired phone change
     * Change users phone (viewer based on jwt)
     * @param param the request object
     */
    public fluxChangeViewerPhone(param: FluxApiFluxChangeViewerPhoneRequest, options?: Configuration): Promise<any> {
        return this.api.fluxChangeViewerPhone(param.body,  options).toPromise();
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
     * ...
     * SaveProfileData saves profile data items for the user
     * @param param the request object
     */
    public fluxSaveProfileData(param: FluxApiFluxSaveProfileDataRequest, options?: Configuration): Promise<ProfileDataInfo> {
        return this.api.fluxSaveProfileData(param.body,  options).toPromise();
    }
	
    /**
     * Exchange a refresh token for new token material; refresh tokens can only be used once If refresh tokens are used more than once RTR dictates that any access tokens which were created by it should be immediately revoked this is because this indicates an attack (something is wrong)
     * @param param the request object
     */
    public fluxTokenExchange(param: FluxApiFluxTokenExchangeRequest, options?: Configuration): Promise<TokenExchangeResponse> {
        return this.api.fluxTokenExchange(param.body,  options).toPromise();
    }
	
    /**
     * Uploads a file and returns a fileId.
     * @param param the request object
     */
    public fluxUploadFile(param: FluxApiFluxUploadFileRequest, options?: Configuration): Promise<InlineResponse200> {
        return this.api.fluxUploadFile(param.file,  options).toPromise();
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
     * Provides user (viewer) data associated with the access token
     * Get viewer profile data
     * @param param the request object
     */
    public fluxViewerProfileData(param: FluxApiFluxViewerProfileDataRequest, options?: Configuration): Promise<ProfileDataInfo> {
        return this.api.fluxViewerProfileData( options).toPromise();
    }
	
    /**
     * @param param the request object
     */
    public fluxWyreConfirmTransfer(param: FluxApiFluxWyreConfirmTransferRequest, options?: Configuration): Promise<WyreTransfer> {
        return this.api.fluxWyreConfirmTransfer(param.transferId, param.body,  options).toPromise();
    }
	
    /**
     * @param param the request object
     */
    public fluxWyreCreateTransfer(param: FluxApiFluxWyreCreateTransferRequest, options?: Configuration): Promise<WyreTransfer> {
        return this.api.fluxWyreCreateTransfer(param.body,  options).toPromise();
    }
	
    /**
     * @param param the request object
     */
    public fluxWyreGetPaymentMethods(param: FluxApiFluxWyreGetPaymentMethodsRequest, options?: Configuration): Promise<WyrePaymentMethods> {
        return this.api.fluxWyreGetPaymentMethods( options).toPromise();
    }
	
    /**
     * @param param the request object
     */
    public fluxWyreGetTransfer(param: FluxApiFluxWyreGetTransferRequest, options?: Configuration): Promise<WyreTransfer> {
        return this.api.fluxWyreGetTransfer(param.transferId,  options).toPromise();
    }
	
    /**
     * @param param the request object
     */
    public fluxWyreGetTransfers(param: FluxApiFluxWyreGetTransfersRequest, options?: Configuration): Promise<WyreTransfers> {
        return this.api.fluxWyreGetTransfers( options).toPromise();
    }
	
    /**
     * @param param the request object
     */
    public fluxWyreWebhook(param: FluxApiFluxWyreWebhookRequest, options?: Configuration): Promise<any> {
        return this.api.fluxWyreWebhook(param.hookId, param.body,  options).toPromise();
    }
	

}



