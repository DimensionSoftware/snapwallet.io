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

import { Address } from './Address';
import { HttpFile } from '../http/http';

export class WyreDebitCardInfo {
    'firstName'?: string;
    'lastName'?: string;
    'phoneNumber'?: string;
    'number'?: string;
    'expirationMonth'?: string;
    'expirationYear'?: string;
    'verificationCode'?: string;
    'address'?: Address;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "firstName",
            "baseName": "firstName",
            "type": "string",
            "format": ""
        },
        {
            "name": "lastName",
            "baseName": "lastName",
            "type": "string",
            "format": ""
        },
        {
            "name": "phoneNumber",
            "baseName": "phoneNumber",
            "type": "string",
            "format": ""
        },
        {
            "name": "number",
            "baseName": "number",
            "type": "string",
            "format": ""
        },
        {
            "name": "expirationMonth",
            "baseName": "expirationMonth",
            "type": "string",
            "format": ""
        },
        {
            "name": "expirationYear",
            "baseName": "expirationYear",
            "type": "string",
            "format": ""
        },
        {
            "name": "verificationCode",
            "baseName": "verificationCode",
            "type": "string",
            "format": ""
        },
        {
            "name": "address",
            "baseName": "address",
            "type": "Address",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return WyreDebitCardInfo.attributeTypeMap;
    }
    
    public constructor() {
    }
}
