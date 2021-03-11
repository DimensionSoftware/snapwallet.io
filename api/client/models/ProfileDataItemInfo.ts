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

import { ProfileDataItemKind } from './ProfileDataItemKind';
import { ProfileDataItemStatus } from './ProfileDataItemStatus';
import { HttpFile } from '../http/http';

/**
* represents a saved profile data fields' status
*/
export class ProfileDataItemInfo {
    'id'?: string;
    'fileIds'?: Array<string>;
    'kind'?: ProfileDataItemKind;
    'subKind'?: string;
    'status'?: ProfileDataItemStatus;
    'length'?: number;
    'createdAt'?: string;
    'updatedAt'?: string;
    'sealedAt'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "id",
            "baseName": "id",
            "type": "string",
            "format": ""
        },
        {
            "name": "fileIds",
            "baseName": "fileIds",
            "type": "Array<string>",
            "format": ""
        },
        {
            "name": "kind",
            "baseName": "kind",
            "type": "ProfileDataItemKind",
            "format": ""
        },
        {
            "name": "subKind",
            "baseName": "subKind",
            "type": "string",
            "format": ""
        },
        {
            "name": "status",
            "baseName": "status",
            "type": "ProfileDataItemStatus",
            "format": ""
        },
        {
            "name": "length",
            "baseName": "length",
            "type": "number",
            "format": "int32"
        },
        {
            "name": "createdAt",
            "baseName": "createdAt",
            "type": "string",
            "format": ""
        },
        {
            "name": "updatedAt",
            "baseName": "updatedAt",
            "type": "string",
            "format": ""
        },
        {
            "name": "sealedAt",
            "baseName": "sealedAt",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return ProfileDataItemInfo.attributeTypeMap;
    }
    
    public constructor() {
    }
}

