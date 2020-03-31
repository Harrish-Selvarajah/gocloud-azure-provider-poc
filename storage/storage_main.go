package storage

import (
	"../internal"
	"context"
	"fmt"
	"time"
)

func Create(accountName, accountGroupName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()
	internal.LoadConfig()
	fmt.Print(CreateStorageAccount(ctx, accountName, accountGroupName))
}

func Update(accountName, accountGroupName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()
	internal.LoadConfig()
	fmt.Print(UpdateAccount(ctx, accountName, accountGroupName))
}

func Delete(accountName, accountGroupName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()
	internal.LoadConfig()
	fmt.Print(DeleteStorageAccount(ctx, accountName, accountGroupName))
}

func CheckAccountNameAvailable(acaccountName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()
	internal.LoadConfig()
	fmt.Print(CheckAccountNameAvailability(ctx, acaccountName))
}

func ListAccountsByResourceGroupName(accountGroupName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()
	internal.LoadConfig()
	fmt.Print(ListAccountsByResourceGroup(ctx, accountGroupName))
}
