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

import { PricingRate } from './PricingRate';
import { HttpFile } from '../http/http';

/**
* response
*/
export class PricingDataResponse {
    'rates'?: { [key: string]: PricingRate; };

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "rates",
            "baseName": "rates",
            "type": "{ [key: string]: PricingRate; }",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return PricingDataResponse.attributeTypeMap;
    }
    
    public constructor() {
    }
}

