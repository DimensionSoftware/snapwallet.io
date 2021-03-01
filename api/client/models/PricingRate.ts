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

/**
* the pricing rate map; the keys are currency symbols and the values are the rates
*/
export class PricingRate {
    'rate'?: { [key: string]: number; };

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "rate",
            "baseName": "rate",
            "type": "{ [key: string]: number; }",
            "format": "float"
        }    ];

    static getAttributeTypeMap() {
        return PricingRate.attributeTypeMap;
    }
    
    public constructor() {
    }
}

