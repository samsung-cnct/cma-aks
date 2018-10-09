package apiserver

import (
	"fmt"
	pb "github.com/samsung-cnct/cma-aks/pkg/generated/api"
	az "github.com/samsung-cnct/cma-aks/pkg/util/azureutil"
	"golang.org/x/net/context"

	k8s "github.com/samsung-cnct/cma-aks/pkg/util/k8s"
)

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
	_, err = aks.CreateCluster(az.CreateClusterInput{
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

	return &pb.CreateClusterReply{
		Ok: true,
		Cluster: &pb.ClusterItem{
			Id:     clusterID,
			Name:   in.Name,
			Status: pb.ClusterStatus_PROVISIONING,
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
	// TODO make this better
	enumeratedStatus, found := pb.ClusterStatus_value[output.Cluster.Status]
	if !found {
		enumeratedStatus = 0
	}

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

	return &pb.DeleteClusterReply{
		Ok:     true,
		Status: output.Status,
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
		cluster := pb.ClusterItem{
			Id:   *output.Clusters[i].ID,
			Name: *output.Clusters[i].Name,
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
	_, err = aks.UpgradeCluster(az.UpgradeClusterInput{
		Name:       in.Name,
		K8sVersion: in.Provider.K8SVersion,
	})
	if err != nil {
		return nil, fmt.Errorf("error upgrading cluster: %v", err)
	}

	clusterID := "/subscriptions/" + in.Provider.Azure.Credentials.SubscriptionId + "/resourcegroups/" + in.Name + "-group/providers/Microsoft.ContainerService/managedClusters/" + in.Name

	return &pb.UpgradeClusterReply{
		Ok: true,
		Cluster: &pb.ClusterItem{
			Id:     clusterID,
			Name:   in.Name,
			Status: pb.ClusterStatus_RECONCILING,
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

	return &pb.ScaleClusterReply{
		Ok:     true,
		Status: output.Status,
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
	cluster, err := aks.GetCluster(az.GetClusterInput{
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}
	agentPool := *cluster.Cluster.AgentPoolProfiles
	agentPoolName := *agentPool[0].Name // AKS supports only 1 node pool at this time

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

	// generate the secret for cluster autoscaling
	secretName := "cluster-autoscaler-azure"
	secretNamespace := "kube-system"
	k8s.CreateAutoScaleSecret(secretName, secretNamespace, autoscalingConfig)

	// deploy cluster autoscaling
	err = k8s.CreateAutoScaleDeployment(agentPoolName, in.MinQuantity, in.MaxQuantity)
	if err != nil {
		return nil, fmt.Errorf("error while enabling cluster autoscaling: %v", err)
	}

	return &pb.EnableClusterAutoscalingReply{
		Ok: true,
	}, nil
}
