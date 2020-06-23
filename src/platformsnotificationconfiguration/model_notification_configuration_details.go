/*
 * Adyen for Platforms: Notification Configuration API
 *
 * The Notification Configuration API provides endpoints for setting up and testing notifications that inform you of events on your platform, for example when a KYC check or a payout has been completed.  For more information, refer to our [documentation](https://docs.adyen.com/marketpay/marketpay-notifications). ## Authentication To connect to the Notification Configuration API, you must use basic authentication credentials of your web service user. If you don't have one, contact our [Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Notification Configuration API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Notification/v1/createNotificationConfiguration ```
 *
 * API version: 5
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformsnotificationconfiguration
// NotificationConfigurationDetails struct for NotificationConfigurationDetails
type NotificationConfigurationDetails struct {
	// Indicates whether the notification subscription is active.
	Active bool `json:"active,omitempty"`
	// The API version of the notification to send.
	ApiVersion int32 `json:"apiVersion,omitempty"`
	// A description of the notification subscription configuration. >Required when creating a configuration, forbidden when updating a configuration.
	Description string `json:"description,omitempty"`
	// The types of events whose notifications apply to this configuration.
	EventConfigs []NotificationEventConfiguration `json:"eventConfigs"`
	// A string with which to salt the notification(s) before hashing. If this field is provided, a hash value will be included under the notification header `HmacSignature` and the hash protocol will be included under the notification header `Protocol`. A notification body along with its `hmacSignatureKey` and `Protocol` can be used to calculate a hash value; matching this hash value with the `HmacSignature` will ensure that the notification body has not been tampered with or corrupted.  >Must be a 32-byte hex-encoded string (i.e. a string containing 64 hexadecimal characters; e.g. \"b0ea55c2fe60d4d1d605e9c385e0e7f7e6cafbb939ce07010f31a327a0871f27\").  The omission of this field will preclude the provision of the `HmacSignature` and `Protocol` headers in notification(s).
	HmacSignatureKey string `json:"hmacSignatureKey,omitempty"`
	// The ID of the configuration. >Required if updating an existing configuration, ignored during the creation of a configuration.
	NotificationId int64 `json:"notificationId,omitempty"`
	// The password to use when accessing the notifyURL with the specified username.
	NotifyPassword string `json:"notifyPassword,omitempty"`
	// The URL to which the notifications are to be sent.
	NotifyURL string `json:"notifyURL"`
	// The username to use when accessing the notifyURL.
	NotifyUsername string `json:"notifyUsername,omitempty"`
	// The SSL protocol employed by the endpoint. >Permitted values: `SSL`, `SSLInsecureCiphers`, `TLS`, `TLSv10`, `TLSv10InsecureCiphers`, `TLSv11`, `TLSv12`.
	SslProtocol string `json:"sslProtocol,omitempty"`
}