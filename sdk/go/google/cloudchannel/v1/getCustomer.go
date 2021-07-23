// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the requested Customer resource. Possible error codes: * PERMISSION_DENIED: The reseller account making the request is different from the reseller account in the API request. * INVALID_ARGUMENT: Required request parameters are missing or invalid. * NOT_FOUND: The customer resource doesn't exist. Usually the result of an invalid name parameter. Return value: The Customer resource.
func LookupCustomer(ctx *pulumi.Context, args *LookupCustomerArgs, opts ...pulumi.InvokeOption) (*LookupCustomerResult, error) {
	var rv LookupCustomerResult
	err := ctx.Invoke("google-native:cloudchannel/v1:getCustomer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomerArgs struct {
	AccountId            string `pulumi:"accountId"`
	ChannelPartnerLinkId string `pulumi:"channelPartnerLinkId"`
	CustomerId           string `pulumi:"customerId"`
}

type LookupCustomerResult struct {
	// Secondary contact email. You need to provide an alternate email to create different domains if a primary contact email already exists. Users will receive a notification with credentials when you create an admin.google.com account. Secondary emails are also recovery email addresses. Alternate emails are optional when you create Team customers.
	AlternateEmail string `pulumi:"alternateEmail"`
	// Cloud Identity ID of the customer's channel partner. Populated only if a channel partner exists for this customer.
	ChannelPartnerId string `pulumi:"channelPartnerId"`
	// The customer's Cloud Identity ID if the customer has a Cloud Identity resource.
	CloudIdentityId string `pulumi:"cloudIdentityId"`
	// Cloud Identity information for the customer. Populated only if a Cloud Identity account exists for this customer.
	CloudIdentityInfo GoogleCloudChannelV1CloudIdentityInfoResponse `pulumi:"cloudIdentityInfo"`
	// Time when the customer was created.
	CreateTime string `pulumi:"createTime"`
	// The customer's primary domain. Must match the primary contact email's domain.
	Domain string `pulumi:"domain"`
	// Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more information, see https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode string `pulumi:"languageCode"`
	// Resource name of the customer. Format: accounts/{account_id}/customers/{customer_id}
	Name string `pulumi:"name"`
	// Name of the organization that the customer entity represents.
	OrgDisplayName string `pulumi:"orgDisplayName"`
	// The organization address for the customer. To enforce US laws and embargoes, we require a region and zip code. You must provide valid addresses for every customer. To set the customer's language, use the Customer-level language code.
	OrgPostalAddress GoogleTypePostalAddressResponse `pulumi:"orgPostalAddress"`
	// Primary contact info.
	PrimaryContactInfo GoogleCloudChannelV1ContactInfoResponse `pulumi:"primaryContactInfo"`
	// Time when the customer was updated.
	UpdateTime string `pulumi:"updateTime"`
}