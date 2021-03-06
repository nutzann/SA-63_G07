/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntArea,
    EntAreaFromJSON,
    EntAreaFromJSONTyped,
    EntAreaToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntLevelEdges
 */
export interface EntLevelEdges {
    /**
     * Area holds the value of the area edge.
     * @type {Array<EntArea>}
     * @memberof EntLevelEdges
     */
    area?: Array<EntArea>;
}

export function EntLevelEdgesFromJSON(json: any): EntLevelEdges {
    return EntLevelEdgesFromJSONTyped(json, false);
}

export function EntLevelEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntLevelEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'area': !exists(json, 'area') ? undefined : ((json['area'] as Array<any>).map(EntAreaFromJSON)),
    };
}

export function EntLevelEdgesToJSON(value?: EntLevelEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'area': value.area === undefined ? undefined : ((value.area as Array<any>).map(EntAreaToJSON)),
    };
}


