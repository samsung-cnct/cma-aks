package azureutil

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
)

// GetGroupsClient get new groups client
func GetGroupsClient(tenantID string, clientID string, clientSecret string, subscriptionID string) resources.GroupsClient {

	// used to authenticate with azure
	os.Setenv("AZURE_TENANT_ID", tenantID)
	os.Setenv("AZURE_CLIENT_ID", clientID)
	os.Setenv("AZURE_CLIENT_SECRET", clientSecret)

	groupsClient := resources.NewGroupsClient(subscriptionID)

	// create an authorizer from env vars or Azure Managed Service Idenity
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	groupsClient.Authorizer = authorizer
	return groupsClient
}

// TODO: add GetGroup function

// CheckForGroup verify group exists
func CheckForGroup(ctx context.Context, groupsClient resources.GroupsClient, resourceName string) bool {
	// generate resource group name based on resource name
	resourceGroupName := resourceName + "-group"

	result, err := groupsClient.CheckExistence(ctx, resourceGroupName)
	if err != nil {
		log.Fatalf("error checking for group: %v", err)
	}
	if result.StatusCode == 404 {
		return false
	}

	return true
}

// CreateGroup will create a new group with the provided name
func CreateGroup(ctx context.Context, groupsClient resources.GroupsClient, resourceName string, location string) (resources.Group, error) {
	// generate resource group name based on resource name
	resourceGroupName := resourceName + "-group"

	log.Println(fmt.Sprintf("creating resource group '%s' on location: %v", resourceGroupName, location))
	grp, err := groupsClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resources.Group{
			Location: to.StringPtr(location),
		},
	)
	if err != nil {
		return grp, fmt.Errorf("cannot create resource group: %v", err)
	}

	return grp, nil
}

// DeleteGroup will delete an azure resource group
func DeleteGroup(ctx context.Context, groupsClient resources.GroupsClient, resourceName string) error {
	// generate resource group name based on resource name
	resourceGroupName := resourceName + "-group"

	_, err := groupsClient.Delete(
		ctx,
		resourceGroupName,
	)
	if err != nil {
		return fmt.Errorf("error while deleting resource group %s, error message: %v", resourceGroupName, err)
	}
	return nil
}
