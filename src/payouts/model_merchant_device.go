/*
 * Adyen Payout API
 *
 * A set of API endpoints that allow you to store payout details, confirm, or decline a payout.  For more information, refer to [Online payouts](https://docs.adyen.com/checkout/online-payouts).
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payouts

// MerchantDevice struct for MerchantDevice
type MerchantDevice struct {
	// Operating system running on the merchant device.
	Os string `json:"os,omitempty"`
	// Version of the operating system on the merchant device.
	OsVersion string `json:"osVersion,omitempty"`
	// Merchant device reference.
	Reference string `json:"reference,omitempty"`
}