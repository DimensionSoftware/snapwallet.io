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

/**
* Represents the wyre user account status (or other tps)
*/
export class ThirdPartyUserAccount {
    'lifecycleStatus'?: LifecycleStatus;
    'status'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "lifecycleStatus",
            "baseName": "lifecycleStatus",
            "type": "LifecycleStatus",
            "format": ""
        },
        {
            "name": "status",
            "baseName": "status",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return ThirdPartyUserAccount.attributeTypeMap;
    }
    
    public constructor() {
    }
}

