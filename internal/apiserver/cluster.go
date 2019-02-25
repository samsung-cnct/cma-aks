package apiserver

import (
	"fmt"
	pb "github.com/samsung-cnct/cma-aks/pkg/generated/api"
	az "github.com/samsung-cnct/cma-aks/pkg/util/azureutil"
	"golang.org/x/net/context"

	k8s "github.com/samsung-cnct/cma-aks/pkg/util/k8s"
)

// match azure cluster status to api status enum
func matchStatus(status string) pb.ClusterStatus {
	switch s := status; s {
	case "Creating":
		return pb.ClusterStatus_PROVISIONING
	case "Updating":
		return pb.ClusterStatus_RECONCILING
	case "Upgrading":
		return pb.ClusterStatus_RECONCILING
	case "Succeeded":
		return pb.ClusterStatus_RUNNING
	case "Deleting":
		return pb.ClusterStatus_STOPPING
	case "Failed":
		return pb.ClusterStatus_ERROR
	default:
		return pb.ClusterStatus_STATUS_UNSPECIFIED
	}
}

func (s *Server) CreateCluster(ctx context.Context, in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {

	// check if resource group exists
	groupsClient := az.GetGroupsClient(in.Provider.Azure.Credentials.Tenant, in.Provider.Azure.Credentials.AppId, in.Provider.Azure.Credentials.Password, in.Provider.Azure.Credentials.SubscriptionId)
	exists := az.CheckForGroup(ctx, groupsClient, in.Name)
	// create group if it does not exist
	if !exists {
		_, err := az.CreateGroup(ctx, groupsClient, in.Name, in.Provider.Azure.Location)
		if err != nil {
			return nil, fmt.Errorf("error creating resource group: %v", err)
		}
	}

	// create cluster client
	aks := az.AKS{
		Context: ctx,
	}
	newClient, err := aks.GetClusterClient(az.ClusterClientInput{
		TenantID:       in.Provider.Azure.Credentials.Tenant,
		ClientID:       in.Provider.Azure.Credentials.AppId,
		ClientSecret:   in.Provider.Azure.Credentials.Password,
		SubscriptionID: in.Provider.Azure.Credentials.SubscriptionId,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot get aks client: %v", err)
	}
	aks.SetClient(newClient.Client)

	// sets up each instance group agent
	agentPools := make([]az.Agent, len(in.Provider.Azure.InstanceGroups))
	for i := range in.Provider.Azure.InstanceGroups {
		agentPools[i].Name = &in.Provider.Azure.InstanceGroups[i].Name
		agentPools[i].Count = &in.Provider.Azure.InstanceGroups[i].MinQuantity
		agentPools[i].Type = in.Provider.Azure.InstanceGroups[i].Type
	}

	// setup cluster tags
	tags := make(map[string]*string)
	for _, tag := range in.Provider.Azure.Tags {
		tags[tag.Key] = &tag.Value
	}

	// create cluster
	output, err := aks.CreateCluster(az.CreateClusterInput{
		Name:         in.Name,
		Location:     in.Provider.Azure.Location,
		K8sVersion:   in.Provider.K8SVersion,
		ClientID:     in.Provider.Azure.ClusterAccount.ClientId,
		ClientSecret: in.Provider.Azure.ClusterAccount.ClientSecret,
		AgentPools:   agentPools,
		Tags:         tags,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating cluster: %v", err)
	}

	clusterID := "/subscriptions/" + in.Provider.Azure.Credentials.SubscriptionId + "/resourcegroups/" + in.Name + "-group/providers/Microsoft.ContainerService/managedClusters/" + in.Name

	enumeratedStatus := matchStatus(output.Status)

	if enumeratedStatus != pb.ClusterStatus_PROVISIONING {
		logger.Errorf("expected status -->%s<-- on provision but instead received -->%s<-- on cluster -->%s<--! ... ", pb.ClusterStatus_PROVISIONING, enumeratedStatus, in.Name)
	}

	return &pb.CreateClusterReply{
		Ok: true,
		Cluster: &pb.ClusterItem{
			Id:     clusterID,
			Name:   in.Name,
			Status: pb.ClusterStatus(enumeratedStatus),
		},
	}, nil
}

func (s *Server) GetCluster(ctx context.Context, in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {

	aks := az.AKS{
		Context: ctx,
	}
	newClient, err := aks.GetClusterClient(az.ClusterClientInput{
		TenantID:       in.Credentials.Tenant,
		ClientID:       in.Credentials.AppId,
		ClientSecret:   in.Credentials.Password,
		SubscriptionID: in.Credentials.SubscriptionId,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot get aks client: %v", err)
	}
	aks.SetClient(newClient.Client)

	output, err := aks.GetCluster(az.GetClusterInput{
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}
	enumeratedStatus := matchStatus(output.Cluster.Status)

	return &pb.GetClusterReply{
		Ok: true,
		Cluster: &pb.ClusterDetailItem{
			Id:         output.Cluster.ID,
			Name:       output.Cluster.Name,
			Status:     pb.ClusterStatus(enumeratedStatus),
			Kubeconfig: output.Cluster.Kubeconfig,
		},
	}, nil
}

func (s *Server) DeleteCluster(ctx context.Context, in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {

	aks := az.AKS{
		Context: ctx,
	}
	newClient, err := aks.GetClusterClient(az.ClusterClientInput{
		TenantID:       in.Credentials.Tenant,
		ClientID:       in.Credentials.AppId,
		ClientSecret:   in.Credentials.Password,
		SubscriptionID: in.Credentials.SubscriptionId,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot get aks client: %v", err)
	}
	aks.SetClient(newClient.Client)

	output, err := aks.DeleteCluster(az.DeleteClusterInput{
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}
	enumeratedStatus := matchStatus(output.Status)

	if enumeratedStatus != pb.ClusterStatus_STOPPING {
		logger.Errorf("expected status -->%s<-- on stopping but instead received -->%s<-- on cluster -->%s<--! ... ", pb.ClusterStatus_PROVISIONING, enumeratedStatus, in.Name)
	}
	// delete resource group
	groupsClient := az.GetGroupsClient(in.Credentials.Tenant, in.Credentials.AppId, in.Credentials.Password, in.Credentials.SubscriptionId)
	deleteGroupErr := az.DeleteGroup(ctx, groupsClient, in.Name)
	if err != nil {
		return nil, fmt.Errorf("error deleting resource group: %v", deleteGroupErr)
	}

	return &pb.DeleteClusterReply{
		Ok:     true,
		Status: pb.ClusterStatus(enumeratedStatus),
	}, nil
}

func (s *Server) GetClusterList(ctx context.Context, in *pb.GetClusterListMsg) (reply *pb.GetClusterListReply, err error) {

	aks := az.AKS{
		Context: ctx,
	}
	newClient, err := aks.GetClusterClient(az.ClusterClientInput{
		TenantID:       in.Credentials.Tenant,
		ClientID:       in.Credentials.AppId,
		ClientSecret:   in.Credentials.Password,
		SubscriptionID: in.Credentials.SubscriptionId,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot get aks client: %v", err)
	}
	aks.SetClient(newClient.Client)

	output, err := aks.ListClusters(az.ListClusterInput{})
	if err != nil {
		return nil, err
	}

	var clusters []*pb.ClusterItem

	for i := range output.Clusters {
		enumeratedStatus := matchStatus(*output.Clusters[i].ProvisioningState)

		cluster := pb.ClusterItem{
			Id:     *output.Clusters[i].ID,
			Name:   *output.Clusters[i].Name,
			Status: pb.ClusterStatus(enumeratedStatus),
		}
		clusters = append(clusters, &cluster)
	}

	reply = &pb.GetClusterListReply{
		Ok:       true,
		Clusters: clusters,
	}
	return reply, nil
}

func (s *Server) GetClusterUpgrades(ctx context.Context, in *pb.GetClusterUpgradesMsg) (reply *pb.GetClusterUpgradesReply, err error) {
	aks := az.AKS{
		Context: ctx,
	}

	newClient, err := aks.GetClusterClient(az.ClusterClientInput{
		TenantID:       in.Credentials.Tenant,
		ClientID:       in.Credentials.AppId,
		ClientSecret:   in.Credentials.Password,
		SubscriptionID: in.Credentials.SubscriptionId,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot get aks client: %v", err)
	}
	aks.SetClient(newClient.Client)

	output, err := aks.GetClusterUpgrades(az.GetClusterUpgradeInput{
		Name: in.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve available upgrades: %v", err)
	}

	// create slice of available upgrades
	var upgrades []*pb.Upgrade
	for i := range output.Upgrades {
		upgrade := pb.Upgrade{
			Version: output.Upgrades[i],
		}
		upgrades = append(upgrades, &upgrade)
	}

	return &pb.GetClusterUpgradesReply{
		Ok:       true,
		Upgrades: upgrades,
	}, nil
}

func (s *Server) UpgradeCluster(ctx context.Context, in *pb.UpgradeClusterMsg) (*pb.UpgradeClusterReply, error) {
	aks := az.AKS{
		Context: ctx,
	}

	// get cluster client
	newClient, err := aks.GetClusterClient(az.ClusterClientInput{
		TenantID:       in.Provider.Azure.Credentials.Tenant,
		ClientID:       in.Provider.Azure.Credentials.AppId,
		ClientSecret:   in.Provider.Azure.Credentials.Password,
		SubscriptionID: in.Provider.Azure.Credentials.SubscriptionId,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot get aks client: %v", err)
	}
	aks.SetClient(newClient.Client)

	// upgrade cluster
	output, err := aks.UpgradeCluster(az.UpgradeClusterInput{
		Name:       in.Name,
		K8sVersion: in.Provider.K8SVersion,
	})
	if err != nil {
		return nil, fmt.Errorf("error upgrading cluster: %v", err)
	}

	clusterID := "/subscriptions/" + in.Provider.Azure.Credentials.SubscriptionId + "/resourcegroups/" + in.Name + "-group/providers/Microsoft.ContainerService/managedClusters/" + in.Name

	enumeratedStatus := matchStatus(output.Status)

	return &pb.UpgradeClusterReply{
		Ok: true,
		Cluster: &pb.ClusterItem{
			Id:     clusterID,
			Name:   in.Name,
			Status: pb.ClusterStatus(enumeratedStatus),
		},
	}, nil
}

func (s *Server) GetClusterNodeCount(ctx context.Context, in *pb.GetClusterNodeCountMsg) (reply *pb.GetClusterNodeCountReply, err error) {
	aks := az.AKS{
		Context: ctx,
	}
	newClient, err := aks.GetClusterClient(az.ClusterClientInput{
		TenantID:       in.Credentials.Tenant,
		ClientID:       in.Credentials.AppId,
		ClientSecret:   in.Credentials.Password,
		SubscriptionID: in.Credentials.SubscriptionId,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot get aks client: %v", err)
	}
	aks.SetClient(newClient.Client)

	output, err := aks.GetClusterNodeCount(az.ClusterNodeCountInput{
		Name: in.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve cluster node count: %v", err)
	}

	return &pb.GetClusterNodeCountReply{
		Ok:    true,
		Name:  *output.Agent.Name,
		Count: *output.Agent.Count,
	}, nil
}

func (s *Server) ScaleCluster(ctx context.Context, in *pb.ScaleClusterMsg) (reply *pb.ScaleClusterReply, err error) {
	aks := az.AKS{
		Context: ctx,
	}
	// get cluster client
	newClient, err := aks.GetClusterClient(az.ClusterClientInput{
		TenantID:       in.Credentials.Tenant,
		ClientID:       in.Credentials.AppId,
		ClientSecret:   in.Credentials.Password,
		SubscriptionID: in.Credentials.SubscriptionId,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot get aks client: %v", err)
	}
	aks.SetClient(newClient.Client)

	output, err := aks.ScaleClusterNodeCount(az.ScaleClusterInput{
		Name:     in.Name,
		NodePool: in.NodePool,
		Count:    in.Count,
	})
	if err != nil {
		return nil, fmt.Errorf("error scaling cluster: %v", err)
	}

	enumeratedStatus := matchStatus(output.Status)

	return &pb.ScaleClusterReply{
		Ok:     true,
		Status: pb.ClusterStatus(enumeratedStatus),
	}, nil
}

func (s *Server) EnableClusterAutoscaling(ctx context.Context, in *pb.EnableClusterAutoscalingMsg) (reply *pb.EnableClusterAutoscalingReply, err error) {
	aks := az.AKS{
		Context: ctx,
	}
	// get cluster client
	newClient, err := aks.GetClusterClient(az.ClusterClientInput{
		TenantID:       in.Credentials.Tenant,
		ClientID:       in.Credentials.AppId,
		ClientSecret:   in.Credentials.Password,
		SubscriptionID: in.Credentials.SubscriptionId,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot get aks client: %v", err)
	}
	aks.SetClient(newClient.Client)

	// get agent pool name
	var agentPoolName string
	var minQuantity int32
	var maxQuantity int32
	cluster, err := aks.GetCluster(az.GetClusterInput{
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}
	agentPool := *cluster.Cluster.AgentPoolProfiles

	// find provided node group in existing cluster
	for i := range agentPool {
		for j := range in.Nodegroups {
			if in.Nodegroups[j].Name == *agentPool[i].Name {
				agentPoolName = in.Nodegroups[j].Name // AKS supports only 1 node pool at this time
				minQuantity = in.Nodegroups[j].MinQuantity
				maxQuantity = in.Nodegroups[j].MaxQuantity
			}
		}
	}
	if agentPoolName == "" {
		return nil, fmt.Errorf("Unable to find provided nodeGroup in cluster")
	}

	// create config
	autoscalingConfig := make(map[string][]byte)
	autoscalingConfig["ResourceGroup"] = []byte(in.Name + "-group")
	autoscalingConfig["NodeResourceGroup"] = []byte(*cluster.Cluster.NodeResourceGroup)
	autoscalingConfig["ClientID"] = []byte(in.Credentials.AppId)
	autoscalingConfig["ClientSecret"] = []byte(in.Credentials.Password)
	autoscalingConfig["TenantID"] = []byte(in.Credentials.Tenant)
	autoscalingConfig["VMType"] = []byte("AKS")
	autoscalingConfig["ClusterName"] = []byte(in.Name)
	autoscalingConfig["SubscriptionID"] = []byte(in.Credentials.SubscriptionId)

	// setup config to call remote cluster
	config, err := k8s.SetKubeConfig(cluster.Cluster.Name, cluster.Cluster.Kubeconfig)

	// generate the secret for cluster autoscaling
	secretName := "cluster-autoscaler-azure"
	secretNamespace := "kube-system"
	k8s.CreateAutoScaleSecret(secretName, secretNamespace, autoscalingConfig, config)

	// deploy cluster autoscaling
	err = k8s.CreateAutoScaleDeployment(agentPoolName, minQuantity, maxQuantity, config)

	//err = k8s.CreateAutoScaleDeployment(agentPoolName, in.MinQuantity, in.MaxQuantity, config)
	if err != nil {
		return nil, fmt.Errorf("error while enabling cluster autoscaling: %v", err)
	}

	return &pb.EnableClusterAutoscalingReply{
		Ok: true,
	}, nil
}
