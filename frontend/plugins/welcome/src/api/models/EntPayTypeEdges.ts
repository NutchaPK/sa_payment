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
    EntPayment,
    EntPaymentFromJSON,
    EntPaymentFromJSONTyped,
    EntPaymentToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPayTypeEdges
 */
export interface EntPayTypeEdges {
    /**
     * Payments holds the value of the payments edge.
     * @type {Array<EntPayment>}
     * @memberof EntPayTypeEdges
     */
    payments?: Array<EntPayment>;
}

export function EntPayTypeEdgesFromJSON(json: any): EntPayTypeEdges {
    return EntPayTypeEdgesFromJSONTyped(json, false);
}

export function EntPayTypeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPayTypeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'payments': !exists(json, 'Payments') ? undefined : ((json['Payments'] as Array<any>).map(EntPaymentFromJSON)),
    };
}

export function EntPayTypeEdgesToJSON(value?: EntPayTypeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'payments': value.payments === undefined ? undefined : ((value.payments as Array<any>).map(EntPaymentToJSON)),
    };
}


