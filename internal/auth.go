package internal

import (
	"log"
	"os"
	"strconv"
)

type AzureCOnfiguration struct {
	tenantID       string
	subscriptionID string
	clientID       string
	clientSecret   string
}

func LoadConfig() error {

	groupName = os.Getenv("AZURE_GROUP_NAME")
	baseGroupName = "az-samples-go"
	locationDefault = "westus2"

	var err error
	useDeviceFlow, err = strconv.ParseBool(os.Getenv("AZURE_USE_DEVICEFLOW"))
	if err != nil {
		log.Printf("invalid value specified for AZURE_USE_DEVICEFLOW, disabling\n")
		useDeviceFlow = false
	}
	keepResources, err = strconv.ParseBool(os.Getenv("AZURE_SAMPLES_KEEP_RESOURCES"))
	if err != nil {
		log.Printf("invalid value specified for AZURE_SAMPLES_KEEP_RESOURCES, discarding\n")
		keepResources = false
	}

	clientID = os.Getenv("AZURE_CLIENT_ID")

	clientSecret = os.Getenv("AZURE_CLIENT_SECRET")

	tenantID = os.Getenv("AZURE_TENANT_ID")

	subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")

	return nil

}
