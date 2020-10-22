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
    EntAreaEdges,
    EntAreaEdgesFromJSON,
    EntAreaEdgesFromJSONTyped,
    EntAreaEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntArea
 */
export interface EntArea {
    /**
     * 
     * @type {EntAreaEdges}
     * @memberof EntArea
     */
    edges?: EntAreaEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntArea
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntArea
     */
    name?: string;
}

export function EntAreaFromJSON(json: any): EntArea {
    return EntAreaFromJSONTyped(json, false);
}

export function EntAreaFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntArea {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntAreaEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
    };
}

export function EntAreaToJSON(value?: EntArea | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntAreaEdgesToJSON(value.edges),
        'id': value.id,
        'name': value.name,
    };
}

