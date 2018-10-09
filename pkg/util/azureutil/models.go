package azureutil

import (
	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2018-03-31/containerservice"
)

// ClusterClientInput is used to create a new cluster client
type ClusterClientInput struct {
	TenantID       string
	ClientID       string
	ClientSecret   string
	SubscriptionID string
}

// ClusterClientOutput shows results of creating a new cluster client
type ClusterClientOutput struct {
	Client containerservice.ManagedClustersClient
}

// CreateClusterInput is used to create a new cluster
type CreateClusterInput struct {
	Name         string
	Location     string
	K8sVersion   string
	ClientID     string
	ClientSecret string
	AgentPools   []Agent
	Tags         map[string]*string
}

// Agent is used to define properties for each instance group
type Agent struct {
	Name    *string
	Count   *int32
	MaxPods *int32
	Type    string
}

// CreateClusterOutput shows results of creating a new cluster
type CreateClusterOutput struct {
	Status string
}

// GetClusterInput is used to retrieve a cluster
type GetClusterInput struct {
	Name string
}

// GetClusterOutput shows the results of retrieving a cluster
type GetClusterOutput struct {
	Cluster ClusterDetailItem
}

// UpgradeClusterInput is used to upgrade an existing cluster
type UpgradeClusterInput struct {
	Name       string
	K8sVersion string
}

// UpgradeClusterOutput shows the results of upgrading an existing cluster
type UpgradeClusterOutput struct {
	Status string
}

// ClusterDetailItem shows the details of a specific cluster
type ClusterDetailItem struct {
	ID                string
	Name              string
	Status            string
	Kubeconfig        string
	AgentPoolProfiles *[]containerservice.ManagedClusterAgentPoolProfile
	NodeResourceGroup *string
}

// DeleteClusterInput is used to delete an existing cluster
type DeleteClusterInput struct {
	Name string
}

// DeleteClusterOutput shows the results of deleting a cluster
type DeleteClusterOutput struct {
	Status string
}

// ListClusterInput is used to retrieve a list of clusters
type ListClusterInput struct{}

// ListClusterOutput shows the list of available clusters
type ListClusterOutput struct {
	Clusters []containerservice.ManagedCluster
}

// GetClusterUpgradeInput is used to retrieve the available k8s upgrades for a cluster
type GetClusterUpgradeInput struct {
	Name string
}

// GetClusterUpgradeOutput shows the available upgrades of a cluster
type GetClusterUpgradeOutput struct {
	Upgrades []string
}

// ClusterNodeCountInput is used to retrieve the current number of worker nodes in the cluster
type ClusterNodeCountInput struct {
	Name string
}

// ClusterNodeCountOutput shows the number of worker nodes available in the cluster
type ClusterNodeCountOutput struct {
	Agent Agent
}

// ScaleClusterInput is used to scale the number of worker nodes in a cluster up or down
type ScaleClusterInput struct {
	Name     string
	NodePool string
	Count    int32
}

// ScaleClusterOutput shows the results of scaling a cluster
type ScaleClusterOutput struct {
	Status string
}
