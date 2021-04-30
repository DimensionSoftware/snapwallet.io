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

export class WyreCreateDebitCardOrderRequest {
    'dest'?: string;
    'sourceCurrency'?: string;
    'destCurrency'?: string;
    'sourceAmount'?: number;
    'country'?: string;
    'amountIncludesFees'?: boolean;
    'lockFields'?: Array<string>;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "dest",
            "baseName": "dest",
            "type": "string",
            "format": ""
        },
        {
            "name": "sourceCurrency",
            "baseName": "sourceCurrency",
            "type": "string",
            "format": ""
        },
        {
            "name": "destCurrency",
            "baseName": "destCurrency",
            "type": "string",
            "format": ""
        },
        {
            "name": "sourceAmount",
            "baseName": "sourceAmount",
            "type": "number",
            "format": "double"
        },
        {
            "name": "country",
            "baseName": "country",
            "type": "string",
            "format": ""
        },
        {
            "name": "amountIncludesFees",
            "baseName": "amountIncludesFees",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "lockFields",
            "baseName": "lockFields",
            "type": "Array<string>",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return WyreCreateDebitCardOrderRequest.attributeTypeMap;
    }
    
    public constructor() {
    }
}

