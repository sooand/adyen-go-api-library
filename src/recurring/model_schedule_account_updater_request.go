/*
 * Adyen Recurring API
 *
 * The Recurring APIs allow you to manage and remove your tokens or saved payment details. Tokens should be created with validation during a payment request.  For more information, refer to our [Tokenization documentation](https://docs.adyen.com/checkout/tokenization). ## Authentication To connect to the Recurring API, you must use your basic authentication credentials. For this, create your web service user, as described in [How to get the WS user password](https://docs.adyen.com/user-management/how-to-get-the-web-service-ws-user-password). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@Company.YourCompany\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Recurring API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://pal-test.adyen.com/pal/servlet/Recurring/v49/disable ```
 *
 * API version: 49
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package recurring
// ScheduleAccountUpdaterRequest struct for ScheduleAccountUpdaterRequest
type ScheduleAccountUpdaterRequest struct {
	// This field contains additional data, which may be required for a particular request.
	AdditionalData *map[string]interface{} `json:"additionalData,omitempty"`
	Card *Card `json:"card,omitempty"`
	// Account of the merchant.
	MerchantAccount string `json:"merchantAccount"`
	// A reference that merchants can apply for the call.
	Reference string `json:"reference"`
	// The selected detail recurring reference.  Optional if `card` is provided.
	SelectedRecurringDetailReference string `json:"selectedRecurringDetailReference,omitempty"`
	// The reference of the shopper that owns the recurring contract.  Optional if `card` is provided.
	ShopperReference string `json:"shopperReference,omitempty"`
}
