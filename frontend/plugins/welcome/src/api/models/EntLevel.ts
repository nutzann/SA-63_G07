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
    EntLevelEdges,
    EntLevelEdgesFromJSON,
    EntLevelEdgesFromJSONTyped,
    EntLevelEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntLevel
 */
export interface EntLevel {
    /**
     * 
     * @type {EntLevelEdges}
     * @memberof EntLevel
     */
    edges?: EntLevelEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntLevel
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntLevel
     */
    name?: string;
}

export function EntLevelFromJSON(json: any): EntLevel {
    return EntLevelFromJSONTyped(json, false);
}

export function EntLevelFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntLevel {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntLevelEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
    };
}

export function EntLevelToJSON(value?: EntLevel | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntLevelEdgesToJSON(value.edges),
        'id': value.id,
        'name': value.name,
    };
}

