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

import { TokenMaterial } from './TokenMaterial';
import { User } from './User';
import { HttpFile } from '../http/http';

/**
* response
*/
export class OneTimePasscodeVerifyResponse {
    'tokens'?: TokenMaterial;
    'user'?: User;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "tokens",
            "baseName": "tokens",
            "type": "TokenMaterial",
            "format": ""
        },
        {
            "name": "user",
            "baseName": "user",
            "type": "User",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return OneTimePasscodeVerifyResponse.attributeTypeMap;
    }
    
    public constructor() {
    }
}

