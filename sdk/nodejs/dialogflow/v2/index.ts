// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./agentEntityType";
export * from "./agentEnvironment";
export * from "./agentEnvironmentUserSessionContext";
export * from "./agentEnvironmentUserSessionEntityType";
export * from "./agentIntent";
export * from "./agentKnowledgeBase";
export * from "./agentKnowledgeBaseDocument";
export * from "./agentSessionContext";
export * from "./agentSessionEntityType";
export * from "./agentVersion";
export * from "./conversation";
export * from "./conversationParticipant";
export * from "./conversationProfile";
export * from "./knowledgeBase";
export * from "./knowledgeBaseDocument";

// Import resources to register:
import { AgentEntityType } from "./agentEntityType";
import { AgentEnvironment } from "./agentEnvironment";
import { AgentEnvironmentUserSessionContext } from "./agentEnvironmentUserSessionContext";
import { AgentEnvironmentUserSessionEntityType } from "./agentEnvironmentUserSessionEntityType";
import { AgentIntent } from "./agentIntent";
import { AgentKnowledgeBase } from "./agentKnowledgeBase";
import { AgentKnowledgeBaseDocument } from "./agentKnowledgeBaseDocument";
import { AgentSessionContext } from "./agentSessionContext";
import { AgentSessionEntityType } from "./agentSessionEntityType";
import { AgentVersion } from "./agentVersion";
import { Conversation } from "./conversation";
import { ConversationParticipant } from "./conversationParticipant";
import { ConversationProfile } from "./conversationProfile";
import { KnowledgeBase } from "./knowledgeBase";
import { KnowledgeBaseDocument } from "./knowledgeBaseDocument";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:dialogflow/v2:AgentEntityType":
                return new AgentEntityType(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:AgentEnvironment":
                return new AgentEnvironment(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:AgentEnvironmentUserSessionContext":
                return new AgentEnvironmentUserSessionContext(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:AgentEnvironmentUserSessionEntityType":
                return new AgentEnvironmentUserSessionEntityType(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:AgentIntent":
                return new AgentIntent(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:AgentKnowledgeBase":
                return new AgentKnowledgeBase(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:AgentKnowledgeBaseDocument":
                return new AgentKnowledgeBaseDocument(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:AgentSessionContext":
                return new AgentSessionContext(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:AgentSessionEntityType":
                return new AgentSessionEntityType(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:AgentVersion":
                return new AgentVersion(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:Conversation":
                return new Conversation(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:ConversationParticipant":
                return new ConversationParticipant(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:ConversationProfile":
                return new ConversationProfile(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:KnowledgeBase":
                return new KnowledgeBase(name, <any>undefined, { urn })
            case "google-native:dialogflow/v2:KnowledgeBaseDocument":
                return new KnowledgeBaseDocument(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "dialogflow/v2", _module)
