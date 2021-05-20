// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new channel in the specified site.
type SiteChannel struct {
	pulumi.CustomResourceState

	// The time at which the channel was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Text labels used for extra metadata and/or filtering.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
	Name pulumi.StringOutput `pulumi:"name"`
	// The current release for the channel, if any.
	Release ReleaseResponseOutput `pulumi:"release"`
	// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount pulumi.IntOutput `pulumi:"retainedReleaseCount"`
	// Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
	Ttl pulumi.StringOutput `pulumi:"ttl"`
	// The time at which the channel was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The URL at which the content of this channel's current release can be viewed. This URL is a Firebase-provided subdomain of `web.app`. The content of this channel's current release can also be viewed at the Firebase-provided subdomain of `firebaseapp.com`. If this channel is the `live` channel for the Hosting site, then the content of this channel's current release can also be viewed at any connected custom domains.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewSiteChannel registers a new resource with the given unique name, arguments, and options.
func NewSiteChannel(ctx *pulumi.Context,
	name string, args *SiteChannelArgs, opts ...pulumi.ResourceOption) (*SiteChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ChannelId == nil {
		return nil, errors.New("invalid value for required argument 'ChannelId'")
	}
	if args.SiteId == nil {
		return nil, errors.New("invalid value for required argument 'SiteId'")
	}
	var resource SiteChannel
	err := ctx.RegisterResource("google-native:firebasehosting/v1beta1:SiteChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteChannel gets an existing SiteChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteChannelState, opts ...pulumi.ResourceOption) (*SiteChannel, error) {
	var resource SiteChannel
	err := ctx.ReadResource("google-native:firebasehosting/v1beta1:SiteChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteChannel resources.
type siteChannelState struct {
	// The time at which the channel was created.
	CreateTime *string `pulumi:"createTime"`
	// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
	ExpireTime *string `pulumi:"expireTime"`
	// Text labels used for extra metadata and/or filtering.
	Labels map[string]string `pulumi:"labels"`
	// The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
	Name *string `pulumi:"name"`
	// The current release for the channel, if any.
	Release *ReleaseResponse `pulumi:"release"`
	// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount *int `pulumi:"retainedReleaseCount"`
	// Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
	Ttl *string `pulumi:"ttl"`
	// The time at which the channel was last updated.
	UpdateTime *string `pulumi:"updateTime"`
	// The URL at which the content of this channel's current release can be viewed. This URL is a Firebase-provided subdomain of `web.app`. The content of this channel's current release can also be viewed at the Firebase-provided subdomain of `firebaseapp.com`. If this channel is the `live` channel for the Hosting site, then the content of this channel's current release can also be viewed at any connected custom domains.
	Url *string `pulumi:"url"`
}

type SiteChannelState struct {
	// The time at which the channel was created.
	CreateTime pulumi.StringPtrInput
	// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
	ExpireTime pulumi.StringPtrInput
	// Text labels used for extra metadata and/or filtering.
	Labels pulumi.StringMapInput
	// The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
	Name pulumi.StringPtrInput
	// The current release for the channel, if any.
	Release ReleaseResponsePtrInput
	// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount pulumi.IntPtrInput
	// Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
	Ttl pulumi.StringPtrInput
	// The time at which the channel was last updated.
	UpdateTime pulumi.StringPtrInput
	// The URL at which the content of this channel's current release can be viewed. This URL is a Firebase-provided subdomain of `web.app`. The content of this channel's current release can also be viewed at the Firebase-provided subdomain of `firebaseapp.com`. If this channel is the `live` channel for the Hosting site, then the content of this channel's current release can also be viewed at any connected custom domains.
	Url pulumi.StringPtrInput
}

func (SiteChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteChannelState)(nil)).Elem()
}

type siteChannelArgs struct {
	ChannelId string `pulumi:"channelId"`
	// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
	ExpireTime *string `pulumi:"expireTime"`
	// Text labels used for extra metadata and/or filtering.
	Labels map[string]string `pulumi:"labels"`
	// The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
	Name *string `pulumi:"name"`
	// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount *int   `pulumi:"retainedReleaseCount"`
	SiteId               string `pulumi:"siteId"`
	// Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a SiteChannel resource.
type SiteChannelArgs struct {
	ChannelId pulumi.StringInput
	// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
	ExpireTime pulumi.StringPtrInput
	// Text labels used for extra metadata and/or filtering.
	Labels pulumi.StringMapInput
	// The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
	Name pulumi.StringPtrInput
	// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount pulumi.IntPtrInput
	SiteId               pulumi.StringInput
	// Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
	Ttl pulumi.StringPtrInput
}

func (SiteChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteChannelArgs)(nil)).Elem()
}

type SiteChannelInput interface {
	pulumi.Input

	ToSiteChannelOutput() SiteChannelOutput
	ToSiteChannelOutputWithContext(ctx context.Context) SiteChannelOutput
}

func (*SiteChannel) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteChannel)(nil))
}

func (i *SiteChannel) ToSiteChannelOutput() SiteChannelOutput {
	return i.ToSiteChannelOutputWithContext(context.Background())
}

func (i *SiteChannel) ToSiteChannelOutputWithContext(ctx context.Context) SiteChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteChannelOutput)
}

type SiteChannelOutput struct {
	*pulumi.OutputState
}

func (SiteChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteChannel)(nil))
}

func (o SiteChannelOutput) ToSiteChannelOutput() SiteChannelOutput {
	return o
}

func (o SiteChannelOutput) ToSiteChannelOutputWithContext(ctx context.Context) SiteChannelOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteChannelOutput{})
}
