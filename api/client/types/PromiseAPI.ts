import { ResponseContext, RequestContext, HttpFile } from '../http/http';
import * as models from '../models/all';
import { Configuration} from '../configuration'

import { Address } from '../models/Address';
import { ChangeViewerEmailRequest } from '../models/ChangeViewerEmailRequest';
import { ChangeViewerPhoneRequest } from '../models/ChangeViewerPhoneRequest';
import { DocumentInput } from '../models/DocumentInput';
import { GetImageResponse } from '../models/GetImageResponse';
import { GotoResponse } from '../models/GotoResponse';
import { ImageProcessingMode } from '../models/ImageProcessingMode';
import { InlineResponse200 } from '../models/InlineResponse200';
import { LifecycleStatus } from '../models/LifecycleStatus';
import { OneTimePasscodeRequest } from '../models/OneTimePasscodeRequest';
import { OneTimePasscodeVerifyRequest } from '../models/OneTimePasscodeVerifyRequest';
import { OneTimePasscodeVerifyResponse } from '../models/OneTimePasscodeVerifyResponse';
import { PlaidAccount } from '../models/PlaidAccount';
import { PlaidConnectBankAccountsRequest } from '../models/PlaidConnectBankAccountsRequest';
import { PlaidCreateLinkTokenResponse } from '../models/PlaidCreateLinkTokenResponse';
import { PlaidInstitution } from '../models/PlaidInstitution';
import { PricingDataResponse } from '../models/PricingDataResponse';
import { PricingRate } from '../models/PricingRate';
import { ProfileDataInfo } from '../models/ProfileDataInfo';
import { ProfileDataItemInfo } from '../models/ProfileDataItemInfo';
import { ProfileDataItemKind } from '../models/ProfileDataItemKind';
import { ProfileDataItemRemediation } from '../models/ProfileDataItemRemediation';
import { ProfileDataItemStatus } from '../models/ProfileDataItemStatus';
import { ProtobufAny } from '../models/ProtobufAny';
import { RpcStatus } from '../models/RpcStatus';
import { SaveProfileDataRequest } from '../models/SaveProfileDataRequest';
import { SnapWidgetConfig } from '../models/SnapWidgetConfig';
import { SnapWidgetProduct } from '../models/SnapWidgetProduct';
import { SnapWidgetWallet } from '../models/SnapWidgetWallet';
import { ThirdPartyUserAccount } from '../models/ThirdPartyUserAccount';
import { TokenExchangeRequest } from '../models/TokenExchangeRequest';
import { TokenExchangeResponse } from '../models/TokenExchangeResponse';
import { TokenMaterial } from '../models/TokenMaterial';
import { UploadFileResponse } from '../models/UploadFileResponse';
import { UsGovernmentIdDocumentInput } from '../models/UsGovernmentIdDocumentInput';
import { UsGovernmentIdDocumentInputKind } from '../models/UsGovernmentIdDocumentInputKind';
import { User } from '../models/User';
import { UserFlags } from '../models/UserFlags';
import { ViewerDataResponse } from '../models/ViewerDataResponse';
import { WidgetGetShortUrlResponse } from '../models/WidgetGetShortUrlResponse';
import { WyreConfirmDebitCardQuoteRequest } from '../models/WyreConfirmDebitCardQuoteRequest';
import { WyreConfirmDebitCardQuoteResponse } from '../models/WyreConfirmDebitCardQuoteResponse';
import { WyreConfirmTransferRequest } from '../models/WyreConfirmTransferRequest';
import { WyreCreateDebitCardQuoteRequest } from '../models/WyreCreateDebitCardQuoteRequest';
import { WyreCreateDebitCardQuoteResponse } from '../models/WyreCreateDebitCardQuoteResponse';
import { WyreCreateTransferRequest } from '../models/WyreCreateTransferRequest';
import { WyreDebitCardInfo } from '../models/WyreDebitCardInfo';
import { WyreGetDebitCardOrderAuthorizationsRequest } from '../models/WyreGetDebitCardOrderAuthorizationsRequest';
import { WyreGetDebitCardOrderAuthorizationsResponse } from '../models/WyreGetDebitCardOrderAuthorizationsResponse';
import { WyrePaymentMethod } from '../models/WyrePaymentMethod';
import { WyrePaymentMethods } from '../models/WyrePaymentMethods';
import { WyreTransfer } from '../models/WyreTransfer';
import { WyreTransferDetail } from '../models/WyreTransferDetail';
import { WyreTransfers } from '../models/WyreTransfers';
import { WyreWalletOrderReservationQuote } from '../models/WyreWalletOrderReservationQuote';
import { WyreWebhookRequest } from '../models/WyreWebhookRequest';
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
     * requires an otp code and the desired email address change
     * Change users email (viewer based on jwt)
     * @param body 
     */
    public fluxChangeViewerEmail(body: ChangeViewerEmailRequest, options?: Configuration): Promise<any> {
    	const result = this.api.fluxChangeViewerEmail(body, options);
        return result.toPromise();
    }
	
    /**
     * requires an otp code and the desired phone change
     * Change users phone (viewer based on jwt)
     * @param body 
     */
    public fluxChangeViewerPhone(body: ChangeViewerPhoneRequest, options?: Configuration): Promise<any> {
    	const result = this.api.fluxChangeViewerPhone(body, options);
        return result.toPromise();
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
     * ...
     * SaveProfileData saves profile data items for the user
     * @param body 
     */
    public fluxSaveProfileData(body: SaveProfileDataRequest, options?: Configuration): Promise<ProfileDataInfo> {
    	const result = this.api.fluxSaveProfileData(body, options);
        return result.toPromise();
    }
	
    /**
     * Exchange a refresh token for new token material; refresh tokens can only be used once If refresh tokens are used more than once RTR dictates that any access tokens which were created by it should be immediately revoked this is because this indicates an attack (something is wrong)
     * @param body 
     */
    public fluxTokenExchange(body: TokenExchangeRequest, options?: Configuration): Promise<TokenExchangeResponse> {
    	const result = this.api.fluxTokenExchange(body, options);
        return result.toPromise();
    }
	
    /**
     * Uploads a file and returns a fileId.
     * @param file The file to upload.
     */
    public fluxUploadFile(file?: HttpFile, options?: Configuration): Promise<InlineResponse200> {
    	const result = this.api.fluxUploadFile(file, options);
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
     * Provides user (viewer) data associated with the access token
     * Get viewer profile data
     */
    public fluxViewerProfileData(options?: Configuration): Promise<ProfileDataInfo> {
    	const result = this.api.fluxViewerProfileData(options);
        return result.toPromise();
    }
	
    /**
     * @param body 
     */
    public fluxWidgetGetShortUrl(body: SnapWidgetConfig, options?: Configuration): Promise<WidgetGetShortUrlResponse> {
    	const result = this.api.fluxWidgetGetShortUrl(body, options);
        return result.toPromise();
    }
	
    /**
     * @param body 
     */
    public fluxWyreConfirmDebitCardQuote(body: WyreConfirmDebitCardQuoteRequest, options?: Configuration): Promise<WyreConfirmDebitCardQuoteResponse> {
    	const result = this.api.fluxWyreConfirmDebitCardQuote(body, options);
        return result.toPromise();
    }
	
    /**
     * @param transferId 
     * @param body 
     */
    public fluxWyreConfirmTransfer(transferId: string, body: WyreConfirmTransferRequest, options?: Configuration): Promise<WyreTransferDetail> {
    	const result = this.api.fluxWyreConfirmTransfer(transferId, body, options);
        return result.toPromise();
    }
	
    /**
     * @param body 
     */
    public fluxWyreCreateDebitCardQuote(body: WyreCreateDebitCardQuoteRequest, options?: Configuration): Promise<WyreCreateDebitCardQuoteResponse> {
    	const result = this.api.fluxWyreCreateDebitCardQuote(body, options);
        return result.toPromise();
    }
	
    /**
     * @param body 
     */
    public fluxWyreCreateTransfer(body: WyreCreateTransferRequest, options?: Configuration): Promise<WyreTransferDetail> {
    	const result = this.api.fluxWyreCreateTransfer(body, options);
        return result.toPromise();
    }
	
    /**
     */
    public fluxWyreGetPaymentMethods(options?: Configuration): Promise<WyrePaymentMethods> {
    	const result = this.api.fluxWyreGetPaymentMethods(options);
        return result.toPromise();
    }
	
    /**
     * @param transferId 
     */
    public fluxWyreGetTransfer(transferId: string, options?: Configuration): Promise<WyreTransferDetail> {
    	const result = this.api.fluxWyreGetTransfer(transferId, options);
        return result.toPromise();
    }
	
    /**
     * @param page 
     */
    public fluxWyreGetTransfers(page?: string, options?: Configuration): Promise<WyreTransfers> {
    	const result = this.api.fluxWyreGetTransfers(page, options);
        return result.toPromise();
    }
	
    /**
     * @param body 
     */
    public fluxWyreGetWalletOrderAuthorizations(body: WyreGetDebitCardOrderAuthorizationsRequest, options?: Configuration): Promise<WyreGetDebitCardOrderAuthorizationsResponse> {
    	const result = this.api.fluxWyreGetWalletOrderAuthorizations(body, options);
        return result.toPromise();
    }
	
    /**
     * @param hookId 
     * @param body 
     */
    public fluxWyreWebhook(hookId: string, body: WyreWebhookRequest, options?: Configuration): Promise<any> {
    	const result = this.api.fluxWyreWebhook(hookId, body, options);
        return result.toPromise();
    }
	

}



