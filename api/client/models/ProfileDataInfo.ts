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

import { ProfileDataItemInfo } from './ProfileDataItemInfo';
import { HttpFile } from '../http/http';

/**
* represents all profile data for a user
*/
export class ProfileDataInfo {
    'legalName'?: ProfileDataItemInfo;
    'dateOfBirth'?: ProfileDataItemInfo;
    'ssn'?: ProfileDataItemInfo;
    'address'?: ProfileDataItemInfo;
    'email'?: ProfileDataItemInfo;
    'phone'?: ProfileDataItemInfo;
    'governmentIdDoc'?: ProfileDataItemInfo;
    'proofOfAddressDoc'?: ProfileDataItemInfo;
    'achAuthorizationFormDoc'?: ProfileDataItemInfo;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "legalName",
            "baseName": "legalName",
            "type": "ProfileDataItemInfo",
            "format": ""
        },
        {
            "name": "dateOfBirth",
            "baseName": "dateOfBirth",
            "type": "ProfileDataItemInfo",
            "format": ""
        },
        {
            "name": "ssn",
            "baseName": "ssn",
            "type": "ProfileDataItemInfo",
            "format": ""
        },
        {
            "name": "address",
            "baseName": "address",
            "type": "ProfileDataItemInfo",
            "format": ""
        },
        {
            "name": "email",
            "baseName": "email",
            "type": "ProfileDataItemInfo",
            "format": ""
        },
        {
            "name": "phone",
            "baseName": "phone",
            "type": "ProfileDataItemInfo",
            "format": ""
        },
        {
            "name": "governmentIdDoc",
            "baseName": "governmentIdDoc",
            "type": "ProfileDataItemInfo",
            "format": ""
        },
        {
            "name": "proofOfAddressDoc",
            "baseName": "proofOfAddressDoc",
            "type": "ProfileDataItemInfo",
            "format": ""
        },
        {
            "name": "achAuthorizationFormDoc",
            "baseName": "achAuthorizationFormDoc",
            "type": "ProfileDataItemInfo",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return ProfileDataInfo.attributeTypeMap;
    }
    
    public constructor() {
    }
}

