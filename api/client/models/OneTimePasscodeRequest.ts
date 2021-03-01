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
* request
*/
export class OneTimePasscodeRequest {
    'emailOrPhone'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "emailOrPhone",
            "baseName": "emailOrPhone",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return OneTimePasscodeRequest.attributeTypeMap;
    }
    
    public constructor() {
    }
}

