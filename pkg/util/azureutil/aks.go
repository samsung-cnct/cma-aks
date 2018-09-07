// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/master/resources/groups.go
// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/master/compute/container_cluster.go

package azureutil

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/go-autorest/autorest/azure/auth"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2018-03-31/containerservice"
	"github.com/Azure/go-autorest/autorest/to"
)

// ClusterParameters are used to create a new aks cluster
type ClusterParameters struct {
	Name              string
	Location          string
	KubernetesVersion string
	AgentPoolName     string
	AgentPoolCount    int32
	AgentPoolMaxPods  int32
	ClientID          string
	ClientSecret      string
	// vmSize string
}

// GetClusterClient creates a new aks cluster client and authorizes access
func GetClusterClient(tenantID string, clientID string, clientSecret string, subscriptionID string) (containerservice.ManagedClustersClient, error) {

	// used to authenticate with azure
	// sets the environment variables expected by NewAuthorizerFromEnvironment()
	os.Setenv("AZURE_TENANT_ID", tenantID)
	os.Setenv("AZURE_CLIENT_ID", clientID)
	os.Setenv("AZURE_CLIENT_SECRET", clientSecret)

	// create new cluster client in subscription
	aksClient := containerservice.NewManagedClustersClient(subscriptionID)

	// create an authorizer from env vars
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	aksClient.Authorizer = authorizer
	return aksClient, nil
}

// GetCluster retrieves a specific aks cluster
func GetCluster(ctx context.Context, clusterClient containerservice.ManagedClustersClient, resourceName string) (containerservice.ManagedCluster, error) {
	resourceGroupName := resourceName + "-group"

	c, err := clusterClient.Get(ctx, resourceGroupName, resourceName)
	if err != nil {
		fmt.Printf("Error getting cluster %v: %v\n", resourceName, err)
	} else {
		fmt.Printf("Cluster %v status is %v\n", *c.Name, c.Status)
	}
	return c, err
}

// CreateCluster creates a new managed Kubernetes cluster
func CreateCluster(ctx context.Context, clusterClient containerservice.ManagedClustersClient, parameters ClusterParameters) (c containerservice.ManagedCluster, err error) {
	resourceGroupName := parameters.Name + "-group"

	future, err := clusterClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		parameters.Name,
		containerservice.ManagedCluster{
			Name:     &parameters.Name,
			Location: &parameters.Location,
			ManagedClusterProperties: &containerservice.ManagedClusterProperties{
				DNSPrefix:         &parameters.Name,
				KubernetesVersion: &parameters.KubernetesVersion,
				AgentPoolProfiles: &[]containerservice.ManagedClusterAgentPoolProfile{
					{
						Count:   to.Int32Ptr(parameters.AgentPoolCount),
						Name:    to.StringPtr(parameters.AgentPoolName),
						MaxPods: to.Int32Ptr(parameters.AgentPoolMaxPods),
						VMSize:  containerservice.StandardD2V2, // TODO: add to parameters
					},
				},
				ServicePrincipalProfile: &containerservice.ManagedClusterServicePrincipalProfile{
					ClientID: to.StringPtr(parameters.ClientID),
					Secret:   to.StringPtr(parameters.ClientSecret),
				},
			},
		},
	)
	if err != nil {
		return c, fmt.Errorf("cannot create aks cluster: %v", err)
	}

	err = future.WaitForCompletion(ctx, clusterClient.Client)
	if err != nil {
		return c, fmt.Errorf("cannot get the aks cluster create or update future response: %v", err)
	}

	return future.Result(clusterClient)
}

// DeleteCluster deletes an existing aks cluster
func DeleteCluster(ctx context.Context, clusterClient containerservice.ManagedClustersClient, resourceName string) (result string, err error) {
	resourceGroupName := resourceName + "-group"

	future, err := clusterClient.Delete(ctx, resourceGroupName, resourceName)
	if err != nil {
		return result, fmt.Errorf("error deleting cluster: %v", err)
	}

	// FIXME: handle future correctly
	// https://godoc.org/github.com/Azure/go-autorest/autorest/azure#Future
	fmt.Println("FUTURE", future)

	msg := "Deleting " + resourceName + " cluster"

	return msg, err

	// TODO: delete resource group also, if nothing else is in it
}

// ListClusters will list all clusters in the subscription
func ListClusters(ctx context.Context, clusterClient containerservice.ManagedClustersClient) ([]containerservice.ManagedCluster, error) {
	results, err := clusterClient.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing clusters: %v", err)
	}

	return results.Values(), nil
}
