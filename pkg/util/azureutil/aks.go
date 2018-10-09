// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/master/resources/groups.go
// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/master/compute/container_cluster.go

package azureutil

import (
	"context"
	"fmt"

	"github.com/Azure/go-autorest/autorest/azure/auth"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2018-03-31/containerservice"
	"github.com/Azure/go-autorest/autorest/to"
)

type AKS struct {
	Client  containerservice.ManagedClustersClient
	Context context.Context
}

// GetClusterClient creates a new aks cluster client and authorizes access
func (a *AKS) GetClusterClient(input ClusterClientInput) (ClusterClientOutput, error) {
	// create new cluster client in subscription
	client := containerservice.NewManagedClustersClient(input.SubscriptionID)

	// authorize request
	clientConfig := auth.NewClientCredentialsConfig(input.ClientID, input.ClientSecret, input.TenantID)
	authorizer, err := clientConfig.Authorizer()
	if err != nil {
		return ClusterClientOutput{}, fmt.Errorf("failed to initialize authorizer: %v", err)
	}

	client.Authorizer = authorizer

	output := ClusterClientOutput{
		Client: client,
	}
	return output, nil
}

// SetClient sets the client for an AKS
func (a *AKS) SetClient(client containerservice.ManagedClustersClient) *AKS {
	a.Client = client
	return a
}

// GetCluster retrieves a specific aks cluster
func (a *AKS) GetCluster(input GetClusterInput) (GetClusterOutput, error) {
	var kubeConfig string
	resourceGroupName := input.Name + "-group"

	// get kubeconfig for environment
	credentialResults, err := a.Client.ListClusterAdminCredentials(a.Context, resourceGroupName, input.Name)
	if err != nil {
		return GetClusterOutput{}, fmt.Errorf("err getting cluster credentails: %v", err)
	}
	if credentialResults.Kubeconfigs != nil {
		for _, v := range *credentialResults.Kubeconfigs {
			if *v.Name == "clusterAdmin" {
				kubeConfig = string(*v.Value)
			}
		}
	}

	c, err := a.Client.Get(a.Context, resourceGroupName, input.Name)
	if err != nil {
		return GetClusterOutput{}, fmt.Errorf("Error getting cluster %v: %v\n", input.Name, err)
	}

	output := GetClusterOutput{
		Cluster: ClusterDetailItem{
			ID:                *c.ID,
			Name:              *c.Name,
			Status:            *c.ProvisioningState,
			Kubeconfig:        kubeConfig,
			AgentPoolProfiles: c.AgentPoolProfiles,
			NodeResourceGroup: c.NodeResourceGroup,
		},
	}

	return output, err
}

// CreateCluster creates a new managed Kubernetes cluster
func (a *AKS) CreateCluster(input CreateClusterInput) (CreateClusterOutput, error) {
	resourceGroupName := input.Name + "-group"

	// create map of containerservice.ManagedClusterAgentPoolProfile from parameters.AgentPools
	agentPoolProfiles := make([]containerservice.ManagedClusterAgentPoolProfile, len(input.AgentPools))
	for i := range input.AgentPools {
		var vmSize containerservice.VMSizeTypes

		// get list of available VM size types
		vmSizeTypes := containerservice.PossibleVMSizeTypesValues()
		for j := range vmSizeTypes {
			// convert the vmSizeTypes to a string
			typeAsStr := string(vmSizeTypes[j])
			// compare input type against available vm size types
			if input.AgentPools[i].Type == typeAsStr {
				vmSize = vmSizeTypes[j]
			}
		}
		if vmSize == "" {
			return CreateClusterOutput{}, fmt.Errorf("invalid VM Size selected")
		}

		agentPoolProfiles[i] = containerservice.ManagedClusterAgentPoolProfile{
			Name:   input.AgentPools[i].Name,
			Count:  input.AgentPools[i].Count,
			VMSize: vmSize,
		}
	}

	future, err := a.Client.CreateOrUpdate(
		a.Context,
		resourceGroupName,
		input.Name,
		containerservice.ManagedCluster{
			Name:     &input.Name,
			Location: &input.Location,
			ManagedClusterProperties: &containerservice.ManagedClusterProperties{
				DNSPrefix:         &input.Name,
				KubernetesVersion: &input.K8sVersion,
				AgentPoolProfiles: &agentPoolProfiles,
				ServicePrincipalProfile: &containerservice.ManagedClusterServicePrincipalProfile{
					ClientID: to.StringPtr(input.ClientID),
					Secret:   to.StringPtr(input.ClientSecret),
				},
			},
			Tags: input.Tags,
		},
	)
	if err != nil {
		return CreateClusterOutput{}, fmt.Errorf("cannot create aks cluster: %v", err)
	}

	status := future.Status()
	if status != "Creating" {
		return CreateClusterOutput{}, fmt.Errorf("cannot create cluster: %v", status)
	}

	output := CreateClusterOutput{
		Status: status,
	}
	return output, nil
}

// DeleteCluster deletes an existing aks cluster
func (a *AKS) DeleteCluster(input DeleteClusterInput) (DeleteClusterOutput, error) {
	resourceGroupName := input.Name + "-group"

	future, err := a.Client.Delete(a.Context, resourceGroupName, input.Name)
	if err != nil {
		return DeleteClusterOutput{}, fmt.Errorf("error deleting cluster: %v", err)
	}

	status := future.Status()
	if status != "InProgress" {
		return DeleteClusterOutput{}, fmt.Errorf("current status of delete: %v", status)
	}

	output := DeleteClusterOutput{
		Status: "Deleting " + input.Name + " cluster",
	}

	return output, err

	// TODO: delete resource group also, if nothing else is in it
}

// ListClusters will list all clusters in the subscription
func (a *AKS) ListClusters(input ListClusterInput) (ListClusterOutput, error) {
	result, err := a.Client.List(a.Context)
	if err != nil {
		return ListClusterOutput{}, fmt.Errorf("error listing clusters: %v", err)
	}

	output := ListClusterOutput{
		Clusters: result.Values(),
	}
	return output, nil
}

// GetClusterUpgrades lists the kubernetes upgrades available on the cluster
func (a *AKS) GetClusterUpgrades(input GetClusterUpgradeInput) (GetClusterUpgradeOutput, error) {
	resourceGroupName := input.Name + "-group"

	result, err := a.Client.GetUpgradeProfile(a.Context, resourceGroupName, input.Name)
	if err != nil {
		return GetClusterUpgradeOutput{}, fmt.Errorf("error getting available upgrades: %v", err)
	}

	for _, v := range *result.AgentPoolProfiles {
		if v.Upgrades != nil {
			return GetClusterUpgradeOutput{
				Upgrades: *v.Upgrades,
			}, nil
		}
	}

	return GetClusterUpgradeOutput{}, nil
}

// UpgradeCluster upgrades the cluster to the provided kubernetes version
func (a *AKS) UpgradeCluster(input UpgradeClusterInput) (UpgradeClusterOutput, error) {
	resourceGroupName := input.Name + "-group"

	// Get the location from cluster properties
	c, err := a.Client.Get(a.Context, resourceGroupName, input.Name)
	if err != nil {
		fmt.Printf("Error getting location for cluster %v: %v\n", input.Name, err)
	}

	// check cluster status before upgrading
	if *c.ProvisioningState != "Succeeded" {
		return UpgradeClusterOutput{}, fmt.Errorf("Unable to upgrade cluster while it is currently %v", *c.ProvisioningState)
	}

	future, err := a.Client.CreateOrUpdate(
		a.Context,
		resourceGroupName,
		input.Name,
		containerservice.ManagedCluster{
			Location: c.Location,
			ManagedClusterProperties: &containerservice.ManagedClusterProperties{
				KubernetesVersion: &input.K8sVersion,
			},
		},
	)
	if err != nil {
		return UpgradeClusterOutput{}, fmt.Errorf("cannot upgrade cluster: %v", err)
	}

	status := future.Status()
	if status != "Upgrading" {
		return UpgradeClusterOutput{}, fmt.Errorf("cannot upgrade cluster: %v", status)
	}

	output := UpgradeClusterOutput{
		Status: status,
	}
	return output, nil
}

// GetClusterNodeCount returns the current number of nodes in the agent pool
func (a *AKS) GetClusterNodeCount(input ClusterNodeCountInput) (ClusterNodeCountOutput, error) {
	resourceGroupName := input.Name + "-group"

	c, err := a.Client.Get(a.Context, resourceGroupName, input.Name)
	if err != nil {
		return ClusterNodeCountOutput{}, fmt.Errorf("error getting cluster %v: %v", input.Name, err)
	}

	agentPool := c.ManagedClusterProperties.AgentPoolProfiles
	var agent Agent
	for _, v := range *agentPool {
		agent.Name = v.Name
		agent.Count = v.Count
	}
	return ClusterNodeCountOutput{
		Agent: agent,
	}, nil
}

// ScaleClusterNodeCount sets the total number of nodes based on the count input
func (a *AKS) ScaleClusterNodeCount(input ScaleClusterInput) (ScaleClusterOutput, error) {
	resourceGroupName := input.Name + "-group"

	// get current cluster
	c, err := a.Client.Get(a.Context, resourceGroupName, input.Name)
	if err != nil {
		return ScaleClusterOutput{}, fmt.Errorf("error getting cluster %v: %v", input.Name, err)
	}

	// check cluster status before scaling
	if *c.ProvisioningState != "Succeeded" {
		return ScaleClusterOutput{}, fmt.Errorf("Unable to update cluster while it is currently %v", *c.ProvisioningState)
	}

	// get the current VMSize from the cluster
	var vmSize containerservice.VMSizeTypes
	for _, v := range *c.ManagedClusterProperties.AgentPoolProfiles {
		if *v.Name == input.NodePool {
			vmSize = v.VMSize
		}
	}

	// scale cluster
	future, err := a.Client.CreateOrUpdate(
		a.Context,
		resourceGroupName,
		input.Name,
		containerservice.ManagedCluster{
			Location: c.Location,
			ManagedClusterProperties: &containerservice.ManagedClusterProperties{
				AgentPoolProfiles: &[]containerservice.ManagedClusterAgentPoolProfile{
					{
						Name:   to.StringPtr(input.NodePool),
						Count:  to.Int32Ptr(input.Count),
						VMSize: vmSize,
					},
				},
			},
		},
	)
	if err != nil {
		return ScaleClusterOutput{}, fmt.Errorf("cannot scale cluster: %v", err)
	}

	status := future.Status()
	if status != "Updating" {
		return ScaleClusterOutput{}, fmt.Errorf("unable to scale: %v", err)
	}

	return ScaleClusterOutput{
		Status: status,
	}, nil
}
