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

import { LifecycleStatus } from './LifecycleStatus';
import { HttpFile } from '../http/http';

export class WyrePaymentMethod {
    'lifecyleStatus'?: LifecycleStatus;
    'id'?: string;
    'status'?: string;
    'name'?: string;
    'last4'?: string;
    'chargeableCurrencies'?: Array<string>;
    'depositableCurrencies'?: Array<string>;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "lifecyleStatus",
            "baseName": "lifecyleStatus",
            "type": "LifecycleStatus",
            "format": ""
        },
        {
            "name": "id",
            "baseName": "id",
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
            "name": "name",
            "baseName": "name",
            "type": "string",
            "format": ""
        },
        {
            "name": "last4",
            "baseName": "last4",
            "type": "string",
            "format": ""
        },
        {
            "name": "chargeableCurrencies",
            "baseName": "chargeableCurrencies",
            "type": "Array<string>",
            "format": ""
        },
        {
            "name": "depositableCurrencies",
            "baseName": "depositableCurrencies",
            "type": "Array<string>",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return WyrePaymentMethod.attributeTypeMap;
    }
    
    public constructor() {
    }
}
