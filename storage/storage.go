package storage

import (
	"../internal"
	"../internal/iam"
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2017-06-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	// "log"
)

var (
	testAccountName      string
	testAccountGroupName string
	accountCreation      bool
)

func getStorageAccountsClient() storage.AccountsClient {
	storageAccountsClient := storage.NewAccountsClient(internal.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	storageAccountsClient.Authorizer = auth
	storageAccountsClient.AddToUserAgent(internal.UserAgent())
	return storageAccountsClient
}

func CreateStorageAccount(ctx context.Context, accountName, accountGroupName string) (storage.Account, error) {
	var s storage.Account
	storageAccountsClient := getStorageAccountsClient()

	result, err := storageAccountsClient.CheckNameAvailability(
		ctx,
		storage.AccountCheckNameAvailabilityParameters{
			Name: to.StringPtr(accountName),
			Type: to.StringPtr("Microsoft.Storage/storageAccounts"),
		})
	if err != nil {
		return s, fmt.Errorf("storage account check-name-availability failed: %+v", err)
	}

	if !*result.NameAvailable {
		return s, fmt.Errorf(
			"storage account name [%s] not available: %v\nserver message: %v",
			accountName, err, *result.Message)
	}

	future, err := storageAccountsClient.Create(
		ctx,
		accountGroupName,
		accountName,
		storage.AccountCreateParameters{
			Sku: &storage.Sku{
				Name: storage.StandardLRS},
			Kind:                              storage.Storage,
			Location:                          to.StringPtr(internal.DefaultLocation()),
			AccountPropertiesCreateParameters: &storage.AccountPropertiesCreateParameters{},
		})

	if err != nil {
		return s, fmt.Errorf("failed to start creating storage account: %+v", err)
	}

	err = future.WaitForCompletionRef(ctx, storageAccountsClient.Client)
	if err != nil {
		return s, fmt.Errorf("failed to finish creating storage account: %+v", err)
	}

	if err == nil {
		accountCreation = true
		fmt.Println("Storage Account has been created")
	}

	return future.Result(storageAccountsClient)
}

// GetStorageAccount gets details on the specified storage account
func GetStorageAccount(ctx context.Context, accountName, accountGroupName string) (storage.Account, error) {
	storageAccountsClient := getStorageAccountsClient()
	return storageAccountsClient.GetProperties(ctx, accountGroupName, accountName)
}

func DeleteStorageAccount(ctx context.Context, accountName, accountGroupName string) (autorest.Response, error) {
	storageAccountsClient := getStorageAccountsClient()
	return storageAccountsClient.Delete(ctx, accountGroupName, accountName)
}

func CheckAccountNameAvailability(ctx context.Context, accountName string) (storage.CheckNameAvailabilityResult, error) {
	storageAccountsClient := getStorageAccountsClient()
	result, err := storageAccountsClient.CheckNameAvailability(
		ctx,
		storage.AccountCheckNameAvailabilityParameters{
			Name: to.StringPtr(accountName),
			Type: to.StringPtr("Microsoft.Storage/storageAccounts"),
		})
	return result, err
}

func ListAccountsByResourceGroup(ctx context.Context, groupName string) (storage.AccountListResult, error) {
	storageAccountsClient := getStorageAccountsClient()
	return storageAccountsClient.ListByResourceGroup(ctx, groupName)
}

func UpdateAccount(ctx context.Context, accountName, accountGroupName string) (storage.Account, error) {
	accountsClient := getStorageAccountsClient()
	return accountsClient.Update(
		ctx,
		accountGroupName,
		accountName,
		storage.AccountUpdateParameters{
			Tags: map[string]*string{
				"who rocks": to.StringPtr("gocloud"),
				"where":     to.StringPtr("on azure")},
		})
}
