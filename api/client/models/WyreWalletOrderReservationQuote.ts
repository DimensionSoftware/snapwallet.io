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

export class WyreWalletOrderReservationQuote {
    'sourceCurrency'?: string;
    'sourceAmount'?: number;
    'sourceAmountWithoutFees'?: number;
    'destCurrency'?: string;
    'destAmount'?: number;
    'exchangeRate'?: number;
    'equivelancies'?: { [key: string]: number; };
    'fees'?: { [key: string]: number; };
    'dest'?: string;
    'expiresAt'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "sourceCurrency",
            "baseName": "sourceCurrency",
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
            "name": "sourceAmountWithoutFees",
            "baseName": "sourceAmountWithoutFees",
            "type": "number",
            "format": "double"
        },
        {
            "name": "destCurrency",
            "baseName": "destCurrency",
            "type": "string",
            "format": ""
        },
        {
            "name": "destAmount",
            "baseName": "destAmount",
            "type": "number",
            "format": "double"
        },
        {
            "name": "exchangeRate",
            "baseName": "exchangeRate",
            "type": "number",
            "format": "double"
        },
        {
            "name": "equivelancies",
            "baseName": "equivelancies",
            "type": "{ [key: string]: number; }",
            "format": "double"
        },
        {
            "name": "fees",
            "baseName": "fees",
            "type": "{ [key: string]: number; }",
            "format": "double"
        },
        {
            "name": "dest",
            "baseName": "dest",
            "type": "string",
            "format": ""
        },
        {
            "name": "expiresAt",
            "baseName": "expiresAt",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return WyreWalletOrderReservationQuote.attributeTypeMap;
    }
    
    public constructor() {
    }
}

