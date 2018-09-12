package apiserver

import (
	"fmt"
	pb "github.com/samsung-cnct/cma-aks/pkg/generated/api"
	az "github.com/samsung-cnct/cma-aks/pkg/util/azureutil"
	"golang.org/x/net/context"
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

	// create cluster
	clusterClient, err := az.GetClusterClient(in.Provider.Azure.Credentials.Tenant, in.Provider.Azure.Credentials.AppId, in.Provider.Azure.Credentials.Password, in.Provider.Azure.Credentials.SubscriptionId)
	if err != nil {
		return nil, fmt.Errorf("cannot get aks client: %v", err)
	}

	// set parameters for new cluster
	var parameters az.ClusterParameters

	parameters.Name = in.Name
	parameters.Location = in.Provider.Azure.Location
	parameters.KubernetesVersion = in.Provider.K8SVersion
	parameters.ClientID = in.Provider.Azure.ClusterAccount.ClientId
	parameters.ClientSecret = in.Provider.Azure.ClusterAccount.ClientSecret

	// sets up each instance group agent
	parameters.AgentPools = make([]az.Agent, len(in.Provider.Azure.InstanceGroups))
	for i := range in.Provider.Azure.InstanceGroups {
		parameters.AgentPools[i].Name = &in.Provider.Azure.InstanceGroups[i].Name
		parameters.AgentPools[i].Count = &in.Provider.Azure.InstanceGroups[i].MinQuantity
		parameters.AgentPools[i].MaxPods = &in.Provider.Azure.InstanceGroups[i].MaxQuantity
		parameters.AgentPools[i].Type = in.Provider.Azure.InstanceGroups[i].Type
	}

	// Tags
	parameters.Tags = make(map[string]*string)
	for _, tag := range in.Provider.Azure.Tags {
		parameters.Tags[tag.Key] = &tag.Value
	}

	// create cluster
	status, err := az.CreateCluster(ctx, clusterClient, parameters)
	if err != nil {
		return nil, fmt.Errorf("error creating cluster: %v", err)
	}

	clusterID := "/subscriptions/" + in.Provider.Azure.Credentials.SubscriptionId + "/resourcegroups/" + parameters.Name + "-group/providers/Microsoft.ContainerService/managedClusters/" + parameters.Name

	return &pb.CreateClusterReply{
		Ok: true,
		Cluster: &pb.ClusterItem{
			Id:     clusterID,
			Name:   parameters.Name,
			Status: status,
		},
	}, nil
}

func (s *Server) GetCluster(ctx context.Context, in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {

	clusterClient, err := az.GetClusterClient(in.Credentials.Tenant, in.Credentials.AppId, in.Credentials.Password, in.Credentials.SubscriptionId)
	if err != nil {
		return nil, fmt.Errorf("cannot get aks client: %v", err)
	}

	c, err := az.GetCluster(ctx, clusterClient, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.GetClusterReply{
		Ok: true,
		Cluster: &pb.ClusterDetailItem{
			Id:     *c.ID,
			Name:   *c.Name,
			Status: c.Status,
			// TODO: get kubeconfig?
			Kubeconfig: "xyz",
		},
	}, nil
}

func (s *Server) DeleteCluster(ctx context.Context, in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {

	clusterClient, err := az.GetClusterClient(in.Credentials.Tenant, in.Credentials.AppId, in.Credentials.Password, in.Credentials.SubscriptionId)
	if err != nil {
		return nil, fmt.Errorf("cannot get aks client: %v", err)
	}

	result, err := az.DeleteCluster(ctx, clusterClient, in.Name)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteClusterReply{Ok: true, Status: result}, nil
}

func (s *Server) GetClusterList(ctx context.Context, in *pb.GetClusterListMsg) (reply *pb.GetClusterListReply, err error) {

	clusterClient, err := az.GetClusterClient(in.Credentials.Tenant, in.Credentials.AppId, in.Credentials.Password, in.Credentials.SubscriptionId)
	if err != nil {
		return nil, fmt.Errorf("cannot get aks client: %v", err)
	}

	result, err := az.ListClusters(ctx, clusterClient)
	if err != nil {
		return nil, err
	}

	var clusters []*pb.ClusterItem

	for i := range result {
		cluster := pb.ClusterItem{
			Id:   *result[i].ID,
			Name: *result[i].Name,
		}
		clusters = append(clusters, &cluster)
	}

	reply = &pb.GetClusterListReply{
		Ok:       true,
		Clusters: clusters,
	}
	return reply, nil
}
