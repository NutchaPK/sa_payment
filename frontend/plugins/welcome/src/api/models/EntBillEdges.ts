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
    EntBillStatus,
    EntBillStatusFromJSON,
    EntBillStatusFromJSONTyped,
    EntBillStatusToJSON,
    EntPayment,
    EntPaymentFromJSON,
    EntPaymentFromJSONTyped,
    EntPaymentToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBillEdges
 */
export interface EntBillEdges {
    /**
     * 
     * @type {EntBillStatus}
     * @memberof EntBillEdges
     */
    billstatus?: EntBillStatus;
    /**
     * 
     * @type {EntPayment}
     * @memberof EntBillEdges
     */
    payments?: EntPayment;
    /**
     * 
     * @type {EntUser}
     * @memberof EntBillEdges
     */
    resident?: EntUser;
}

export function EntBillEdgesFromJSON(json: any): EntBillEdges {
    return EntBillEdgesFromJSONTyped(json, false);
}

export function EntBillEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBillEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'billstatus': !exists(json, 'Billstatus') ? undefined : EntBillStatusFromJSON(json['Billstatus']),
        'payments': !exists(json, 'Payments') ? undefined : EntPaymentFromJSON(json['Payments']),
        'resident': !exists(json, 'Resident') ? undefined : EntUserFromJSON(json['Resident']),
    };
}

export function EntBillEdgesToJSON(value?: EntBillEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'billstatus': EntBillStatusToJSON(value.billstatus),
        'payments': EntPaymentToJSON(value.payments),
        'resident': EntUserToJSON(value.resident),
    };
}


