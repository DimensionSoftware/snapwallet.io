/**
 * api.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { OrganizationApplication } from './OrganizationApplication';
import { User } from './User';
import { HttpFile } from '../http/http';

/**
* an organization containing users, a user is one to many to organizations
*/
export class Organization {
    'id'?: string;
    'name'?: string;
    'users'?: Array<User>;
    'applications'?: Array<OrganizationApplication>;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "id",
            "baseName": "id",
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
            "name": "users",
            "baseName": "users",
            "type": "Array<User>",
            "format": ""
        },
        {
            "name": "applications",
            "baseName": "applications",
            "type": "Array<OrganizationApplication>",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return Organization.attributeTypeMap;
    }
    
    public constructor() {
    }
}

