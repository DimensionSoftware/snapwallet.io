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

export class WyreTransfer {
    'id'?: string;
    'source'?: string;
    'dest'?: string;
    'sourceCurrency'?: string;
    'destCurrency'?: string;
    'sourceAmount'?: number;
    'destAmount'?: number;
    'sourceName'?: string;
    'destName'?: string;
    'message'?: string;
    'exchangeRate'?: number;
    'fees'?: { [key: string]: number; };
    'blockchainTxId'?: string;
    'status'?: string;
    'createdAt'?: string;
    'closedAt'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "id",
            "baseName": "id",
            "type": "string",
            "format": ""
        },
        {
            "name": "source",
            "baseName": "source",
            "type": "string",
            "format": ""
        },
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
            "name": "destAmount",
            "baseName": "destAmount",
            "type": "number",
            "format": "double"
        },
        {
            "name": "sourceName",
            "baseName": "sourceName",
            "type": "string",
            "format": ""
        },
        {
            "name": "destName",
            "baseName": "destName",
            "type": "string",
            "format": ""
        },
        {
            "name": "message",
            "baseName": "message",
            "type": "string",
            "format": ""
        },
        {
            "name": "exchangeRate",
            "baseName": "exchangeRate",
            "type": "number",
            "format": "double"
        },
        {
            "name": "fees",
            "baseName": "fees",
            "type": "{ [key: string]: number; }",
            "format": "double"
        },
        {
            "name": "blockchainTxId",
            "baseName": "blockchainTxId",
            "type": "string",
            "format": ""
        },
        {
            "name": "status",
            "baseName": "status",
            "type": "string",
            "format": ""
        },
        {
            "name": "createdAt",
            "baseName": "createdAt",
            "type": "string",
            "format": ""
        },
        {
            "name": "closedAt",
            "baseName": "closedAt",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return WyreTransfer.attributeTypeMap;
    }
    
    public constructor() {
    }
}

