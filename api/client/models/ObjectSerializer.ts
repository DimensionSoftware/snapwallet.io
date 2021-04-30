export * from './Address';
export * from './ChangeViewerEmailRequest';
export * from './ChangeViewerPhoneRequest';
export * from './DocumentInput';
export * from './GetImageResponse';
export * from './GotoResponse';
export * from './ImageProcessingMode';
export * from './InlineResponse200';
export * from './LifecycleStatus';
export * from './OneTimePasscodeRequest';
export * from './OneTimePasscodeVerifyRequest';
export * from './OneTimePasscodeVerifyResponse';
export * from './PlaidAccount';
export * from './PlaidConnectBankAccountsRequest';
export * from './PlaidCreateLinkTokenResponse';
export * from './PlaidInstitution';
export * from './PricingDataResponse';
export * from './PricingRate';
export * from './ProfileDataInfo';
export * from './ProfileDataItemInfo';
export * from './ProfileDataItemKind';
export * from './ProfileDataItemRemediation';
export * from './ProfileDataItemStatus';
export * from './ProtobufAny';
export * from './RpcStatus';
export * from './SaveProfileDataRequest';
export * from './SnapWidgetConfig';
export * from './SnapWidgetProduct';
export * from './SnapWidgetWallet';
export * from './ThirdPartyUserAccount';
export * from './TokenExchangeRequest';
export * from './TokenExchangeResponse';
export * from './TokenMaterial';
export * from './UploadFileResponse';
export * from './UsGovernmentIdDocumentInput';
export * from './UsGovernmentIdDocumentInputKind';
export * from './User';
export * from './UserFlags';
export * from './ViewerDataResponse';
export * from './WidgetGetShortUrlResponse';
export * from './WyreConfirmTransferRequest';
export * from './WyreCreateDebitCardOrderRequest';
export * from './WyreCreateDebitCardOrderResponse';
export * from './WyreCreateTransferRequest';
export * from './WyreDebitCardInfo';
export * from './WyreGetDebitCardOrderAuthorizationsRequest';
export * from './WyreGetDebitCardOrderAuthorizationsResponse';
export * from './WyrePaymentMethod';
export * from './WyrePaymentMethods';
export * from './WyreTransfer';
export * from './WyreTransferDetail';
export * from './WyreTransfers';
export * from './WyreWalletOrderReservationQuote';
export * from './WyreWebhookRequest';

import { Address } from './Address';
import { ChangeViewerEmailRequest } from './ChangeViewerEmailRequest';
import { ChangeViewerPhoneRequest } from './ChangeViewerPhoneRequest';
import { DocumentInput } from './DocumentInput';
import { GetImageResponse } from './GetImageResponse';
import { GotoResponse } from './GotoResponse';
import { ImageProcessingMode } from './ImageProcessingMode';
import { InlineResponse200 } from './InlineResponse200';
import { LifecycleStatus } from './LifecycleStatus';
import { OneTimePasscodeRequest } from './OneTimePasscodeRequest';
import { OneTimePasscodeVerifyRequest } from './OneTimePasscodeVerifyRequest';
import { OneTimePasscodeVerifyResponse } from './OneTimePasscodeVerifyResponse';
import { PlaidAccount } from './PlaidAccount';
import { PlaidConnectBankAccountsRequest } from './PlaidConnectBankAccountsRequest';
import { PlaidCreateLinkTokenResponse } from './PlaidCreateLinkTokenResponse';
import { PlaidInstitution } from './PlaidInstitution';
import { PricingDataResponse } from './PricingDataResponse';
import { PricingRate } from './PricingRate';
import { ProfileDataInfo } from './ProfileDataInfo';
import { ProfileDataItemInfo } from './ProfileDataItemInfo';
import { ProfileDataItemKind } from './ProfileDataItemKind';
import { ProfileDataItemRemediation } from './ProfileDataItemRemediation';
import { ProfileDataItemStatus } from './ProfileDataItemStatus';
import { ProtobufAny } from './ProtobufAny';
import { RpcStatus } from './RpcStatus';
import { SaveProfileDataRequest } from './SaveProfileDataRequest';
import { SnapWidgetConfig } from './SnapWidgetConfig';
import { SnapWidgetProduct } from './SnapWidgetProduct';
import { SnapWidgetWallet } from './SnapWidgetWallet';
import { ThirdPartyUserAccount } from './ThirdPartyUserAccount';
import { TokenExchangeRequest } from './TokenExchangeRequest';
import { TokenExchangeResponse } from './TokenExchangeResponse';
import { TokenMaterial } from './TokenMaterial';
import { UploadFileResponse } from './UploadFileResponse';
import { UsGovernmentIdDocumentInput } from './UsGovernmentIdDocumentInput';
import { UsGovernmentIdDocumentInputKind } from './UsGovernmentIdDocumentInputKind';
import { User } from './User';
import { UserFlags } from './UserFlags';
import { ViewerDataResponse } from './ViewerDataResponse';
import { WidgetGetShortUrlResponse } from './WidgetGetShortUrlResponse';
import { WyreConfirmTransferRequest } from './WyreConfirmTransferRequest';
import { WyreCreateDebitCardOrderRequest } from './WyreCreateDebitCardOrderRequest';
import { WyreCreateDebitCardOrderResponse } from './WyreCreateDebitCardOrderResponse';
import { WyreCreateTransferRequest } from './WyreCreateTransferRequest';
import { WyreDebitCardInfo } from './WyreDebitCardInfo';
import { WyreGetDebitCardOrderAuthorizationsRequest } from './WyreGetDebitCardOrderAuthorizationsRequest';
import { WyreGetDebitCardOrderAuthorizationsResponse } from './WyreGetDebitCardOrderAuthorizationsResponse';
import { WyrePaymentMethod } from './WyrePaymentMethod';
import { WyrePaymentMethods } from './WyrePaymentMethods';
import { WyreTransfer } from './WyreTransfer';
import { WyreTransferDetail } from './WyreTransferDetail';
import { WyreTransfers } from './WyreTransfers';
import { WyreWalletOrderReservationQuote } from './WyreWalletOrderReservationQuote';
import { WyreWebhookRequest } from './WyreWebhookRequest';

/* tslint:disable:no-unused-variable */
let primitives = [
                    "string",
                    "boolean",
                    "double",
                    "integer",
                    "long",
                    "float",
                    "number",
                    "any"
                 ];

const supportedMediaTypes: { [mediaType: string]: number } = {
  "application/json": Infinity,
  "application/octet-stream": 0
}

                 
let enumsMap: Set<string> = new Set<string>([
    "ImageProcessingMode",
    "LifecycleStatus",
    "ProfileDataItemKind",
    "ProfileDataItemStatus",
    "UsGovernmentIdDocumentInputKind",
]);

let typeMap: {[index: string]: any} = {
    "Address": Address,
    "ChangeViewerEmailRequest": ChangeViewerEmailRequest,
    "ChangeViewerPhoneRequest": ChangeViewerPhoneRequest,
    "DocumentInput": DocumentInput,
    "GetImageResponse": GetImageResponse,
    "GotoResponse": GotoResponse,
    "InlineResponse200": InlineResponse200,
    "OneTimePasscodeRequest": OneTimePasscodeRequest,
    "OneTimePasscodeVerifyRequest": OneTimePasscodeVerifyRequest,
    "OneTimePasscodeVerifyResponse": OneTimePasscodeVerifyResponse,
    "PlaidAccount": PlaidAccount,
    "PlaidConnectBankAccountsRequest": PlaidConnectBankAccountsRequest,
    "PlaidCreateLinkTokenResponse": PlaidCreateLinkTokenResponse,
    "PlaidInstitution": PlaidInstitution,
    "PricingDataResponse": PricingDataResponse,
    "PricingRate": PricingRate,
    "ProfileDataInfo": ProfileDataInfo,
    "ProfileDataItemInfo": ProfileDataItemInfo,
    "ProfileDataItemRemediation": ProfileDataItemRemediation,
    "ProtobufAny": ProtobufAny,
    "RpcStatus": RpcStatus,
    "SaveProfileDataRequest": SaveProfileDataRequest,
    "SnapWidgetConfig": SnapWidgetConfig,
    "SnapWidgetProduct": SnapWidgetProduct,
    "SnapWidgetWallet": SnapWidgetWallet,
    "ThirdPartyUserAccount": ThirdPartyUserAccount,
    "TokenExchangeRequest": TokenExchangeRequest,
    "TokenExchangeResponse": TokenExchangeResponse,
    "TokenMaterial": TokenMaterial,
    "UploadFileResponse": UploadFileResponse,
    "UsGovernmentIdDocumentInput": UsGovernmentIdDocumentInput,
    "User": User,
    "UserFlags": UserFlags,
    "ViewerDataResponse": ViewerDataResponse,
    "WidgetGetShortUrlResponse": WidgetGetShortUrlResponse,
    "WyreConfirmTransferRequest": WyreConfirmTransferRequest,
    "WyreCreateDebitCardOrderRequest": WyreCreateDebitCardOrderRequest,
    "WyreCreateDebitCardOrderResponse": WyreCreateDebitCardOrderResponse,
    "WyreCreateTransferRequest": WyreCreateTransferRequest,
    "WyreDebitCardInfo": WyreDebitCardInfo,
    "WyreGetDebitCardOrderAuthorizationsRequest": WyreGetDebitCardOrderAuthorizationsRequest,
    "WyreGetDebitCardOrderAuthorizationsResponse": WyreGetDebitCardOrderAuthorizationsResponse,
    "WyrePaymentMethod": WyrePaymentMethod,
    "WyrePaymentMethods": WyrePaymentMethods,
    "WyreTransfer": WyreTransfer,
    "WyreTransferDetail": WyreTransferDetail,
    "WyreTransfers": WyreTransfers,
    "WyreWalletOrderReservationQuote": WyreWalletOrderReservationQuote,
    "WyreWebhookRequest": WyreWebhookRequest,
}

export class ObjectSerializer {
    public static findCorrectType(data: any, expectedType: string) {
        if (data == undefined) {
            return expectedType;
        } else if (primitives.indexOf(expectedType.toLowerCase()) !== -1) {
            return expectedType;
        } else if (expectedType === "Date") {
            return expectedType;
        } else {
            if (enumsMap.has(expectedType)) {
                return expectedType;
            }

            if (!typeMap[expectedType]) {
                return expectedType; // w/e we don't know the type
            }

            // Check the discriminator
            let discriminatorProperty = typeMap[expectedType].discriminator;
            if (discriminatorProperty == null) {
                return expectedType; // the type does not have a discriminator. use it.
            } else {
                if (data[discriminatorProperty]) {
                    var discriminatorType = data[discriminatorProperty];
                    if(typeMap[discriminatorType]){
                        return discriminatorType; // use the type given in the discriminator
                    } else {
                        return expectedType; // discriminator did not map to a type
                    }
                } else {
                    return expectedType; // discriminator was not present (or an empty string)
                }
            }
        }
    }

    public static serialize(data: any, type: string, format: string) {
        if (data == undefined) {
            return data;
        } else if (primitives.indexOf(type.toLowerCase()) !== -1) {
            return data;
        } else if (type.lastIndexOf("Array<", 0) === 0) { // string.startsWith pre es6
            let subType: string = type.replace("Array<", ""); // Array<Type> => Type>
            subType = subType.substring(0, subType.length - 1); // Type> => Type
            let transformedData: any[] = [];
            for (let index in data) {
                let date = data[index];
                transformedData.push(ObjectSerializer.serialize(date, subType, format));
            }
            return transformedData;
        } else if (type === "Date") {
            if (format == "date") {
                let month = data.getMonth()+1
                month = month < 10 ? "0" + month.toString() : month.toString()
                let day = data.getDate();
                day = day < 10 ? "0" + day.toString() : day.toString();

                return data.getFullYear() + "-" + month + "-" + day;
            } else {
                return data.toISOString();
            }
        } else {
            if (enumsMap.has(type)) {
                return data;
            }
            if (!typeMap[type]) { // in case we dont know the type
                return data;
            }
            
            // Get the actual type of this object
            type = this.findCorrectType(data, type);

            // get the map for the correct type.
            let attributeTypes = typeMap[type].getAttributeTypeMap();
            let instance: {[index: string]: any} = {};
            for (let index in attributeTypes) {
                let attributeType = attributeTypes[index];
                instance[attributeType.baseName] = ObjectSerializer.serialize(data[attributeType.name], attributeType.type, attributeType.format);
            }
            return instance;
        }
    }

    public static deserialize(data: any, type: string, format: string) {
        // polymorphism may change the actual type.
        type = ObjectSerializer.findCorrectType(data, type);
        if (data == undefined) {
            return data;
        } else if (primitives.indexOf(type.toLowerCase()) !== -1) {
            return data;
        } else if (type.lastIndexOf("Array<", 0) === 0) { // string.startsWith pre es6
            let subType: string = type.replace("Array<", ""); // Array<Type> => Type>
            subType = subType.substring(0, subType.length - 1); // Type> => Type
            let transformedData: any[] = [];
            for (let index in data) {
                let date = data[index];
                transformedData.push(ObjectSerializer.deserialize(date, subType, format));
            }
            return transformedData;
        } else if (type === "Date") {
            return new Date(data);
        } else {
            if (enumsMap.has(type)) {// is Enum
                return data;
            }

            if (!typeMap[type]) { // dont know the type
                return data;
            }
            let instance = new typeMap[type]();
            let attributeTypes = typeMap[type].getAttributeTypeMap();
            for (let index in attributeTypes) {
                let attributeType = attributeTypes[index];
                instance[attributeType.name] = ObjectSerializer.deserialize(data[attributeType.baseName], attributeType.type, attributeType.format);
            }
            return instance;
        }
    }


    /**
     * Normalize media type
     *
     * We currently do not handle any media types attributes, i.e. anything
     * after a semicolon. All content is assumed to be UTF-8 compatible.
     */
    public static normalizeMediaType(mediaType: string | undefined): string | undefined {
        if (mediaType === undefined) {
            return undefined;
        }
        return mediaType.split(";")[0].trim().toLowerCase();
    }

    /**
     * From a list of possible media types, choose the one we can handle best.
     *
     * The order of the given media types does not have any impact on the choice
     * made.
     */
    public static getPreferredMediaType(mediaTypes: Array<string>): string {
        /** According to OAS 3 we should default to json */
        if (!mediaTypes) {
            return "application/json";
        }

        const normalMediaTypes = mediaTypes.map(this.normalizeMediaType);
        let selectedMediaType: string | undefined = undefined;
        let selectedRank: number = -Infinity;
        for (const mediaType of normalMediaTypes) {
            if (supportedMediaTypes[mediaType!] > selectedRank) {
                selectedMediaType = mediaType;
                selectedRank = supportedMediaTypes[mediaType!];
            }
        }

        if (selectedMediaType === undefined) {
            throw new Error("None of the given media types are supported: " + mediaTypes.join(", "));
        }

        return selectedMediaType!;
    }

    /**
     * Convert data to a string according the given media type
     */
    public static stringify(data: any, mediaType: string): string {
        if (mediaType === "application/json") {
            return JSON.stringify(data);
        }

        throw new Error("The mediaType " + mediaType + " is not supported by ObjectSerializer.stringify.");
    }

    /**
     * Parse data from a string according to the given media type
     */
    public static parse(rawData: string, mediaType: string | undefined) {
        if (mediaType === undefined) {
            throw new Error("Cannot parse content. No Content-Type defined.");
        }

        if (mediaType === "application/json") {
            return JSON.parse(rawData);
        }

        throw new Error("The mediaType " + mediaType + " is not supported by ObjectSerializer.parse.");
    }
}
