import { ResponseContext, RequestContext, HttpFile } from '../http/http';
import * as models from '../models/all';
import { Configuration} from '../configuration'
import { Observable, of, from } from '../rxjsStub';
import {mergeMap, map} from  '../rxjsStub';

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
import { PlaidConnectBankAccountsRequest } from '../models/PlaidConnectBankAccountsRequest';
import { PlaidCreateLinkTokenResponse } from '../models/PlaidCreateLinkTokenResponse';
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
import { WyreConfirmTransferRequest } from '../models/WyreConfirmTransferRequest';
import { WyreCreateTransferRequest } from '../models/WyreCreateTransferRequest';
import { WyrePaymentMethod } from '../models/WyrePaymentMethod';
import { WyrePaymentMethods } from '../models/WyrePaymentMethods';
import { WyreTransfer } from '../models/WyreTransfer';
import { WyreTransfers } from '../models/WyreTransfers';
import { WyreWebhookRequest } from '../models/WyreWebhookRequest';

import { FluxApiRequestFactory, FluxApiResponseProcessor} from "../apis/FluxApi";
export class ObservableFluxApi {
    private requestFactory: FluxApiRequestFactory;
    private responseProcessor: FluxApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: FluxApiRequestFactory,
        responseProcessor?: FluxApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new FluxApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new FluxApiResponseProcessor();
    }

    /**
     * requires an otp code and the desired email address change
     * Change users email (viewer based on jwt)
     * @param body 
     */
    public fluxChangeViewerEmail(body: ChangeViewerEmailRequest, options?: Configuration): Observable<any> {
    	const requestContextPromise = this.requestFactory.fluxChangeViewerEmail(body, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxChangeViewerEmail(rsp)));
	    	}));
    }
	
    /**
     * requires an otp code and the desired phone change
     * Change users phone (viewer based on jwt)
     * @param body 
     */
    public fluxChangeViewerPhone(body: ChangeViewerPhoneRequest, options?: Configuration): Observable<any> {
    	const requestContextPromise = this.requestFactory.fluxChangeViewerPhone(body, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxChangeViewerPhone(rsp)));
	    	}));
    }
	
    /**
     * @param id 
     */
    public fluxGoto(id: string, options?: Configuration): Observable<GotoResponse> {
    	const requestContextPromise = this.requestFactory.fluxGoto(id, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxGoto(rsp)));
	    	}));
    }
	
    /**
     * Will cause your email or phone to receive a one time passcode. This can be used in the verify step to obtain a token for login
     * Post email or phone in exchange for a one time passcode
     * @param body 
     */
    public fluxOneTimePasscode(body: OneTimePasscodeRequest, options?: Configuration): Observable<any> {
    	const requestContextPromise = this.requestFactory.fluxOneTimePasscode(body, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxOneTimePasscode(rsp)));
	    	}));
    }
	
    /**
     * The passcode received in either email or phone text message should be provided here in order to obtain on access token
     * Post one time passcode in exchange for an access token
     * @param body 
     */
    public fluxOneTimePasscodeVerify(body: OneTimePasscodeVerifyRequest, options?: Configuration): Observable<OneTimePasscodeVerifyResponse> {
    	const requestContextPromise = this.requestFactory.fluxOneTimePasscodeVerify(body, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxOneTimePasscodeVerify(rsp)));
	    	}));
    }
	
    /**
     * requires a plaid processor token which in turn requires a plaid widget interaction where the user selects the account id
     * Post chosen bank info from plaid in order to create a new ACH pyment method in wyre
     * @param body 
     */
    public fluxPlaidConnectBankAccounts(body: PlaidConnectBankAccountsRequest, options?: Configuration): Observable<any> {
    	const requestContextPromise = this.requestFactory.fluxPlaidConnectBankAccounts(body, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxPlaidConnectBankAccounts(rsp)));
	    	}));
    }
	
    /**
     * https://plaid.com/docs/link/link-token-migration-guide/
     * @param body 
     */
    public fluxPlaidCreateLinkToken(body: any, options?: Configuration): Observable<PlaidCreateLinkTokenResponse> {
    	const requestContextPromise = this.requestFactory.fluxPlaidCreateLinkToken(body, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxPlaidCreateLinkToken(rsp)));
	    	}));
    }
	
    /**
     * Provides pricing data for all markets with rate maps
     * Get pricing data
     */
    public fluxPricingData(options?: Configuration): Observable<PricingDataResponse> {
    	const requestContextPromise = this.requestFactory.fluxPricingData(options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxPricingData(rsp)));
	    	}));
    }
	
    /**
     * ...
     * SaveProfileData saves profile data items for the user
     * @param body 
     */
    public fluxSaveProfileData(body: SaveProfileDataRequest, options?: Configuration): Observable<ProfileDataInfo> {
    	const requestContextPromise = this.requestFactory.fluxSaveProfileData(body, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxSaveProfileData(rsp)));
	    	}));
    }
	
    /**
     * Exchange a refresh token for new token material; refresh tokens can only be used once If refresh tokens are used more than once RTR dictates that any access tokens which were created by it should be immediately revoked this is because this indicates an attack (something is wrong)
     * @param body 
     */
    public fluxTokenExchange(body: TokenExchangeRequest, options?: Configuration): Observable<TokenExchangeResponse> {
    	const requestContextPromise = this.requestFactory.fluxTokenExchange(body, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxTokenExchange(rsp)));
	    	}));
    }
	
    /**
     * Uploads a file and returns a fileId.
     * @param file The file to upload.
     */
    public fluxUploadFile(file?: HttpFile, options?: Configuration): Observable<InlineResponse200> {
    	const requestContextPromise = this.requestFactory.fluxUploadFile(file, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxUploadFile(rsp)));
	    	}));
    }
	
    /**
     * Provides user (viewer) data associated with the access token
     * Get viewer data
     */
    public fluxViewerData(options?: Configuration): Observable<ViewerDataResponse> {
    	const requestContextPromise = this.requestFactory.fluxViewerData(options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxViewerData(rsp)));
	    	}));
    }
	
    /**
     * Provides user (viewer) data associated with the access token
     * Get viewer profile data
     */
    public fluxViewerProfileData(options?: Configuration): Observable<ProfileDataInfo> {
    	const requestContextPromise = this.requestFactory.fluxViewerProfileData(options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxViewerProfileData(rsp)));
	    	}));
    }
	
    /**
     * @param body 
     */
    public fluxWidgetGetShortUrl(body: SnapWidgetConfig, options?: Configuration): Observable<WidgetGetShortUrlResponse> {
    	const requestContextPromise = this.requestFactory.fluxWidgetGetShortUrl(body, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxWidgetGetShortUrl(rsp)));
	    	}));
    }
	
    /**
     * @param transferId 
     * @param body 
     */
    public fluxWyreConfirmTransfer(transferId: string, body: WyreConfirmTransferRequest, options?: Configuration): Observable<WyreTransfer> {
    	const requestContextPromise = this.requestFactory.fluxWyreConfirmTransfer(transferId, body, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxWyreConfirmTransfer(rsp)));
	    	}));
    }
	
    /**
     * @param body 
     */
    public fluxWyreCreateTransfer(body: WyreCreateTransferRequest, options?: Configuration): Observable<WyreTransfer> {
    	const requestContextPromise = this.requestFactory.fluxWyreCreateTransfer(body, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxWyreCreateTransfer(rsp)));
	    	}));
    }
	
    /**
     */
    public fluxWyreGetPaymentMethods(options?: Configuration): Observable<WyrePaymentMethods> {
    	const requestContextPromise = this.requestFactory.fluxWyreGetPaymentMethods(options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxWyreGetPaymentMethods(rsp)));
	    	}));
    }
	
    /**
     * @param transferId 
     */
    public fluxWyreGetTransfer(transferId: string, options?: Configuration): Observable<WyreTransfer> {
    	const requestContextPromise = this.requestFactory.fluxWyreGetTransfer(transferId, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxWyreGetTransfer(rsp)));
	    	}));
    }
	
    /**
     * @param page 
     */
    public fluxWyreGetTransfers(page?: string, options?: Configuration): Observable<WyreTransfers> {
    	const requestContextPromise = this.requestFactory.fluxWyreGetTransfers(page, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxWyreGetTransfers(rsp)));
	    	}));
    }
	
    /**
     * @param hookId 
     * @param body 
     */
    public fluxWyreWebhook(hookId: string, body: WyreWebhookRequest, options?: Configuration): Observable<any> {
    	const requestContextPromise = this.requestFactory.fluxWyreWebhook(hookId, body, options);

		// build promise chain
    let middlewarePreObservable = from<RequestContext>(requestContextPromise);
    	for (let middleware of this.configuration.middleware) {
    		middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
    	}

    	return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
	    	pipe(mergeMap((response: ResponseContext) => {
	    		let middlewarePostObservable = of(response);
	    		for (let middleware of this.configuration.middleware) {
	    			middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
	    		}
	    		return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.fluxWyreWebhook(rsp)));
	    	}));
    }
	

}



