/**
 * Flux API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: evergreen
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { HttpFile } from '../http/http';

export class WyreSubmitDebitCardOrderAuthorizationsRequest {
    'orderId'?: string;
    'reservationId'?: string;
    'sms2faCode'?: string;
    'card2faCode'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "orderId",
            "baseName": "orderId",
            "type": "string",
            "format": ""
        },
        {
            "name": "reservationId",
            "baseName": "reservationId",
            "type": "string",
            "format": ""
        },
        {
            "name": "sms2faCode",
            "baseName": "sms2faCode",
            "type": "string",
            "format": ""
        },
        {
            "name": "card2faCode",
            "baseName": "card2faCode",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return WyreSubmitDebitCardOrderAuthorizationsRequest.attributeTypeMap;
    }
    
    public constructor() {
    }
}

