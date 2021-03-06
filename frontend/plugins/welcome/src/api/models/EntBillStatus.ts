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
    EntBillStatusEdges,
    EntBillStatusEdgesFromJSON,
    EntBillStatusEdgesFromJSONTyped,
    EntBillStatusEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBillStatus
 */
export interface EntBillStatus {
    /**
     * BillStatus holds the value of the "BillStatus" field.
     * @type {string}
     * @memberof EntBillStatus
     */
    billStatus?: string;
    /**
     * 
     * @type {EntBillStatusEdges}
     * @memberof EntBillStatus
     */
    edges?: EntBillStatusEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntBillStatus
     */
    id?: number;
}

export function EntBillStatusFromJSON(json: any): EntBillStatus {
    return EntBillStatusFromJSONTyped(json, false);
}

export function EntBillStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBillStatus {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'billStatus': !exists(json, 'BillStatus') ? undefined : json['BillStatus'],
        'edges': !exists(json, 'edges') ? undefined : EntBillStatusEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntBillStatusToJSON(value?: EntBillStatus | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'BillStatus': value.billStatus,
        'edges': EntBillStatusEdgesToJSON(value.edges),
        'id': value.id,
    };
}


