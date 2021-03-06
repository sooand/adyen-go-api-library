/*
 * Adyen Payment API
 *
 * A set of API endpoints that allow you to initiate, settle, and modify payments on the Adyen payments platform. You can use the API to accept card payments (including One-Click and 3D Secure), bank transfers, ewallets, and many other payment methods.  To learn more about the API, visit [Classic integration](https://docs.adyen.com/classic-integration).  ## Authentication To connect to the Payments API, you must use your basic authentication credentials. For this, create your web service user, as described in [How to get the WS user password](https://docs.adyen.com/user-management/how-to-get-the-web-service-ws-user-password). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@Company.YourCompany\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Payments API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://pal-test.adyen.com/pal/servlet/Payment/v52/authorise ```
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payments

// Split struct for Split
type Split struct {
	// The account to which this split applies.  >Required if the type is `MarketPlace`.
	Account string      `json:"account,omitempty"`
	Amount  SplitAmount `json:"amount"`
	// A description of this split.
	Description string `json:"description,omitempty"`
	// The reference of this split. Used to link other operations (e.g. captures and refunds) to this split.  >Required if the type is `MarketPlace`.
	Reference string `json:"reference,omitempty"`
	// The type of this split.  >Permitted values: `Default`, `PaymentFee`, `VAT`, `Commission`, `MarketPlace`, `BalanceAccount`.
	Type string `json:"type"`
}
