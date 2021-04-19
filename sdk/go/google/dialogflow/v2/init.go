// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-google-native/sdk/go/google"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "google-native:dialogflow/v2:AgentEntityType":
		r = &AgentEntityType{}
	case "google-native:dialogflow/v2:AgentEnvironment":
		r = &AgentEnvironment{}
	case "google-native:dialogflow/v2:AgentEnvironmentUserSessionContext":
		r = &AgentEnvironmentUserSessionContext{}
	case "google-native:dialogflow/v2:AgentEnvironmentUserSessionEntityType":
		r = &AgentEnvironmentUserSessionEntityType{}
	case "google-native:dialogflow/v2:AgentIntent":
		r = &AgentIntent{}
	case "google-native:dialogflow/v2:AgentKnowledgeBase":
		r = &AgentKnowledgeBase{}
	case "google-native:dialogflow/v2:AgentKnowledgeBaseDocument":
		r = &AgentKnowledgeBaseDocument{}
	case "google-native:dialogflow/v2:AgentSessionContext":
		r = &AgentSessionContext{}
	case "google-native:dialogflow/v2:AgentSessionEntityType":
		r = &AgentSessionEntityType{}
	case "google-native:dialogflow/v2:AgentVersion":
		r = &AgentVersion{}
	case "google-native:dialogflow/v2:Conversation":
		r = &Conversation{}
	case "google-native:dialogflow/v2:ConversationParticipant":
		r = &ConversationParticipant{}
	case "google-native:dialogflow/v2:ConversationProfile":
		r = &ConversationProfile{}
	case "google-native:dialogflow/v2:KnowledgeBase":
		r = &KnowledgeBase{}
	case "google-native:dialogflow/v2:KnowledgeBaseDocument":
		r = &KnowledgeBaseDocument{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := google.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"google-native",
		"dialogflow/v2",
		&module{version},
	)
}
