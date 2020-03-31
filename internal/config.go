package internal

import (
	"bytes"
	"fmt"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/marstr/randname"
)

var (
	clientID               string
	clientSecret           string
	tenantID               string
	subscriptionID         string
	locationDefault        string
	groupName              string
	baseGroupName          string
	userAgent              string
	environment            *azure.Environment
	useDeviceFlow          bool
	keepResources          bool
	authorizationServerURL string
	cloudName              string = "AzurePublicCloud"
)

func ClientID() string {
	return clientID
}

func ClientSecret() string {
	return clientSecret
}

func TenantID() string {
	return tenantID
}

func SubscriptionID() string {
	return subscriptionID
}

func Location() string {
	return locationDefault
}

func DefaultLocation() string {
	return locationDefault
}

func AuthorizationServerURL() string {
	return authorizationServerURL
}

func UseDeviceFlow() bool {
	return useDeviceFlow
}

func GroupName() string {
	return groupName
}

func SetGroupName(name string) {
	groupName = name
}

func BaseGroupName() string {
	return baseGroupName
}

func KeepResources() bool {
	return keepResources
}

func UserAgent() string {
	if len(userAgent) > 0 {
		return userAgent
	}
	return "sdk-samples"
}

func Environment() *azure.Environment {
	if environment != nil {
		return environment
	}
	env, err := azure.EnvironmentFromName(cloudName)
	if err != nil {
		panic(fmt.Sprintf(
			"invalid cloud name '%s' specified, cannot continue\n", cloudName))
	}
	environment = &env
	return environment
}

func GenerateGroupName(affixes ...string) string {
	b := bytes.NewBufferString(BaseGroupName())
	b.WriteRune('-')
	for _, affix := range affixes {
		b.WriteString(affix)
		b.WriteRune('-')
	}
	return randname.GenerateWithPrefix(b.String(), 5)
}

func AppendRandomSuffix(prefix string) string {
	return randname.GenerateWithPrefix(prefix, 5)
}
