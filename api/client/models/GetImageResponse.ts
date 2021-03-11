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

export class GetImageResponse {
    'filename'?: string;
    'mimeType'?: string;
    'size'?: number;
    'body'?: string;
    'width'?: number;
    'height'?: number;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "filename",
            "baseName": "filename",
            "type": "string",
            "format": ""
        },
        {
            "name": "mimeType",
            "baseName": "mimeType",
            "type": "string",
            "format": ""
        },
        {
            "name": "size",
            "baseName": "size",
            "type": "number",
            "format": "int32"
        },
        {
            "name": "body",
            "baseName": "body",
            "type": "string",
            "format": "byte"
        },
        {
            "name": "width",
            "baseName": "width",
            "type": "number",
            "format": "int32"
        },
        {
            "name": "height",
            "baseName": "height",
            "type": "number",
            "format": "int32"
        }    ];

    static getAttributeTypeMap() {
        return GetImageResponse.attributeTypeMap;
    }
    
    public constructor() {
    }
}

