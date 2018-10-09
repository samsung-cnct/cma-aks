# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api.proto](#api.proto)
    - [AzureClusterServiceAccount](#cmaaks.AzureClusterServiceAccount)
    - [AzureCredentials](#cmaaks.AzureCredentials)
    - [ClusterDetailItem](#cmaaks.ClusterDetailItem)
    - [ClusterItem](#cmaaks.ClusterItem)
    - [CreateClusterAKSSpec](#cmaaks.CreateClusterAKSSpec)
    - [CreateClusterAKSSpec.AKSInstanceGroup](#cmaaks.CreateClusterAKSSpec.AKSInstanceGroup)
    - [CreateClusterAKSSpec.Tags](#cmaaks.CreateClusterAKSSpec.Tags)
    - [CreateClusterMsg](#cmaaks.CreateClusterMsg)
    - [CreateClusterProviderSpec](#cmaaks.CreateClusterProviderSpec)
    - [CreateClusterReply](#cmaaks.CreateClusterReply)
    - [DeleteClusterMsg](#cmaaks.DeleteClusterMsg)
    - [DeleteClusterReply](#cmaaks.DeleteClusterReply)
    - [EnableClusterAutoscalingMsg](#cmaaks.EnableClusterAutoscalingMsg)
    - [EnableClusterAutoscalingReply](#cmaaks.EnableClusterAutoscalingReply)
    - [GetClusterListMsg](#cmaaks.GetClusterListMsg)
    - [GetClusterListReply](#cmaaks.GetClusterListReply)
    - [GetClusterMsg](#cmaaks.GetClusterMsg)
    - [GetClusterNodeCountMsg](#cmaaks.GetClusterNodeCountMsg)
    - [GetClusterNodeCountReply](#cmaaks.GetClusterNodeCountReply)
    - [GetClusterReply](#cmaaks.GetClusterReply)
    - [GetClusterUpgradesMsg](#cmaaks.GetClusterUpgradesMsg)
    - [GetClusterUpgradesReply](#cmaaks.GetClusterUpgradesReply)
    - [GetVersionMsg](#cmaaks.GetVersionMsg)
    - [GetVersionReply](#cmaaks.GetVersionReply)
    - [GetVersionReply.VersionInformation](#cmaaks.GetVersionReply.VersionInformation)
    - [ScaleClusterMsg](#cmaaks.ScaleClusterMsg)
    - [ScaleClusterReply](#cmaaks.ScaleClusterReply)
    - [Upgrade](#cmaaks.Upgrade)
    - [UpgradeClusterAKSSpec](#cmaaks.UpgradeClusterAKSSpec)
    - [UpgradeClusterMsg](#cmaaks.UpgradeClusterMsg)
    - [UpgradeClusterProviderSpec](#cmaaks.UpgradeClusterProviderSpec)
    - [UpgradeClusterReply](#cmaaks.UpgradeClusterReply)
  
    - [ClusterStatus](#cmaaks.ClusterStatus)
  
  
    - [Cluster](#cmaaks.Cluster)
  

- [api.proto](#api.proto)
    - [AzureClusterServiceAccount](#cmaaks.AzureClusterServiceAccount)
    - [AzureCredentials](#cmaaks.AzureCredentials)
    - [ClusterDetailItem](#cmaaks.ClusterDetailItem)
    - [ClusterItem](#cmaaks.ClusterItem)
    - [CreateClusterAKSSpec](#cmaaks.CreateClusterAKSSpec)
    - [CreateClusterAKSSpec.AKSInstanceGroup](#cmaaks.CreateClusterAKSSpec.AKSInstanceGroup)
    - [CreateClusterAKSSpec.Tags](#cmaaks.CreateClusterAKSSpec.Tags)
    - [CreateClusterMsg](#cmaaks.CreateClusterMsg)
    - [CreateClusterProviderSpec](#cmaaks.CreateClusterProviderSpec)
    - [CreateClusterReply](#cmaaks.CreateClusterReply)
    - [DeleteClusterMsg](#cmaaks.DeleteClusterMsg)
    - [DeleteClusterReply](#cmaaks.DeleteClusterReply)
    - [EnableClusterAutoscalingMsg](#cmaaks.EnableClusterAutoscalingMsg)
    - [EnableClusterAutoscalingReply](#cmaaks.EnableClusterAutoscalingReply)
    - [GetClusterListMsg](#cmaaks.GetClusterListMsg)
    - [GetClusterListReply](#cmaaks.GetClusterListReply)
    - [GetClusterMsg](#cmaaks.GetClusterMsg)
    - [GetClusterNodeCountMsg](#cmaaks.GetClusterNodeCountMsg)
    - [GetClusterNodeCountReply](#cmaaks.GetClusterNodeCountReply)
    - [GetClusterReply](#cmaaks.GetClusterReply)
    - [GetClusterUpgradesMsg](#cmaaks.GetClusterUpgradesMsg)
    - [GetClusterUpgradesReply](#cmaaks.GetClusterUpgradesReply)
    - [GetVersionMsg](#cmaaks.GetVersionMsg)
    - [GetVersionReply](#cmaaks.GetVersionReply)
    - [GetVersionReply.VersionInformation](#cmaaks.GetVersionReply.VersionInformation)
    - [ScaleClusterMsg](#cmaaks.ScaleClusterMsg)
    - [ScaleClusterReply](#cmaaks.ScaleClusterReply)
    - [Upgrade](#cmaaks.Upgrade)
    - [UpgradeClusterAKSSpec](#cmaaks.UpgradeClusterAKSSpec)
    - [UpgradeClusterMsg](#cmaaks.UpgradeClusterMsg)
    - [UpgradeClusterProviderSpec](#cmaaks.UpgradeClusterProviderSpec)
    - [UpgradeClusterReply](#cmaaks.UpgradeClusterReply)
  
    - [ClusterStatus](#cmaaks.ClusterStatus)
  
  
    - [Cluster](#cmaaks.Cluster)
  

- [Scalar Value Types](#scalar-value-types)



<a name="api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="cmaaks.AzureClusterServiceAccount"></a>

### AzureClusterServiceAccount
the account used by the cluster to create azure resources (ex: load balancer)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_id | [string](#string) |  | The ClientId (aka: AppID) |
| client_secret | [string](#string) |  | The ClientSecret (aka: password) |






<a name="cmaaks.AzureCredentials"></a>

### AzureCredentials
The credentials to use for creating the cluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| app_id | [string](#string) |  | The AppId for API Access |
| tenant | [string](#string) |  | The Tenant for API access |
| password | [string](#string) |  | The Password for API access |
| subscription_id | [string](#string) |  | The Subscription for API access |






<a name="cmaaks.ClusterDetailItem"></a>

### ClusterDetailItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| kubeconfig | [string](#string) |  | What is the kubeconfig to connect to the cluster |
| status | [ClusterStatus](#cmaaks.ClusterStatus) |  | The status of the cluster |






<a name="cmaaks.ClusterItem"></a>

### ClusterItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| status | [ClusterStatus](#cmaaks.ClusterStatus) |  | The status of the cluster |






<a name="cmaaks.CreateClusterAKSSpec"></a>

### CreateClusterAKSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| location | [string](#string) |  | The Azure Data Center |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to build the cluster |
| clusterAccount | [AzureClusterServiceAccount](#cmaaks.AzureClusterServiceAccount) |  | Cluster service account used to talk to azure (ex: creating load balancer) |
| instance_groups | [CreateClusterAKSSpec.AKSInstanceGroup](#cmaaks.CreateClusterAKSSpec.AKSInstanceGroup) | repeated | Instance groups |
| tags | [CreateClusterAKSSpec.Tags](#cmaaks.CreateClusterAKSSpec.Tags) | repeated | Tags |






<a name="cmaaks.CreateClusterAKSSpec.AKSInstanceGroup"></a>

### CreateClusterAKSSpec.AKSInstanceGroup
Instance groups define a type and number of instances


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the group |
| type | [string](#string) |  | Instance type (Standard_D2_v2, etc.) |
| min_quantity | [int32](#int32) |  | Minimum number of instances (defaults to zero) |






<a name="cmaaks.CreateClusterAKSSpec.Tags"></a>

### CreateClusterAKSSpec.Tags
Tags are name/value pairs that enable you to categorize resources and view consolidated billing


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | Tag key |
| value | [string](#string) |  | Tag value |






<a name="cmaaks.CreateClusterMsg"></a>

### CreateClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be provisioned |
| provider | [CreateClusterProviderSpec](#cmaaks.CreateClusterProviderSpec) |  | The provider specification |






<a name="cmaaks.CreateClusterProviderSpec"></a>

### CreateClusterProviderSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the provider - like aks |
| k8s_version | [string](#string) |  | The version of Kubernetes |
| azure | [CreateClusterAKSSpec](#cmaaks.CreateClusterAKSSpec) |  | The AKS specification |
| high_availability | [bool](#bool) |  | Whether or not the cluster is HA |
| network_fabric | [string](#string) |  | The fabric to be used |






<a name="cmaaks.CreateClusterReply"></a>

### CreateClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was provisioned by this request |
| cluster | [ClusterItem](#cmaaks.ClusterItem) |  | The details of the cluster request response |






<a name="cmaaks.DeleteClusterMsg"></a>

### DeleteClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster&#39;s name to destroy |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to delete the cluster |






<a name="cmaaks.DeleteClusterReply"></a>

### DeleteClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Could the cluster be destroyed |
| status | [string](#string) |  | Status of the request |






<a name="cmaaks.EnableClusterAutoscalingMsg"></a>

### EnableClusterAutoscalingMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster |
| min_quantity | [int32](#int32) |  | Minimum number of nodes for autoscaling |
| max_quantity | [int32](#int32) |  | Maximum number of nodes for autoscaling |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to access the cluster |






<a name="cmaaks.EnableClusterAutoscalingReply"></a>

### EnableClusterAutoscalingReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not cluster autoscaling was enabled by this request |






<a name="cmaaks.GetClusterListMsg"></a>

### GetClusterListMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to search subscription for clusters |






<a name="cmaaks.GetClusterListReply"></a>

### GetClusterListReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| clusters | [ClusterItem](#cmaaks.ClusterItem) | repeated | List of clusters |






<a name="cmaaks.GetClusterMsg"></a>

### GetClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be looked up |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to query for the cluster |






<a name="cmaaks.GetClusterNodeCountMsg"></a>

### GetClusterNodeCountMsg
Get available node count


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be looked up |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to query for the cluster |






<a name="cmaaks.GetClusterNodeCountReply"></a>

### GetClusterNodeCountReply
Reply with available node count


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| name | [string](#string) |  | the available node pool name |
| count | [int32](#int32) |  | The available node count |






<a name="cmaaks.GetClusterReply"></a>

### GetClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| cluster | [ClusterDetailItem](#cmaaks.ClusterDetailItem) |  |  |






<a name="cmaaks.GetClusterUpgradesMsg"></a>

### GetClusterUpgradesMsg
Get available cluster upgrades


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be looked up |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to query for the cluster |






<a name="cmaaks.GetClusterUpgradesReply"></a>

### GetClusterUpgradesReply
Reply with available upgrades


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| upgrades | [Upgrade](#cmaaks.Upgrade) | repeated | The available upgrades |






<a name="cmaaks.GetVersionMsg"></a>

### GetVersionMsg
Get version of API Server






<a name="cmaaks.GetVersionReply"></a>

### GetVersionReply
Reply for version request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | If operation was OK |
| version_information | [GetVersionReply.VersionInformation](#cmaaks.GetVersionReply.VersionInformation) |  | Version Information |






<a name="cmaaks.GetVersionReply.VersionInformation"></a>

### GetVersionReply.VersionInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| git_version | [string](#string) |  | The tag on the git repository |
| git_commit | [string](#string) |  | The hash of the git commit |
| git_tree_state | [string](#string) |  | Whether or not the tree was clean when built |
| build_date | [string](#string) |  | Date of build |
| go_version | [string](#string) |  | Version of go used to compile |
| compiler | [string](#string) |  | Compiler used |
| platform | [string](#string) |  | Platform it was compiled for / running on |






<a name="cmaaks.ScaleClusterMsg"></a>

### ScaleClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of cluster |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to search subscription for clusters |
| node_pool | [string](#string) |  | name of node pool |
| count | [int32](#int32) |  | total desired nodes |






<a name="cmaaks.ScaleClusterReply"></a>

### ScaleClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | wheather or not the cluster was scaled |
| status | [string](#string) |  | Status of the request |






<a name="cmaaks.Upgrade"></a>

### Upgrade
available upgrade version


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  |  |






<a name="cmaaks.UpgradeClusterAKSSpec"></a>

### UpgradeClusterAKSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to build the cluster |






<a name="cmaaks.UpgradeClusterMsg"></a>

### UpgradeClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be upgraded |
| provider | [UpgradeClusterProviderSpec](#cmaaks.UpgradeClusterProviderSpec) |  | The provider specification |






<a name="cmaaks.UpgradeClusterProviderSpec"></a>

### UpgradeClusterProviderSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the provider - like aks |
| k8s_version | [string](#string) |  | The version of Kubernetes |
| azure | [UpgradeClusterAKSSpec](#cmaaks.UpgradeClusterAKSSpec) |  | The AKS specification |






<a name="cmaaks.UpgradeClusterReply"></a>

### UpgradeClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was upgrade by this request |
| cluster | [ClusterItem](#cmaaks.ClusterItem) |  | The details of the cluster request response |





 


<a name="cmaaks.ClusterStatus"></a>

### ClusterStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 | Not set |
| PROVISIONING | 1 | The PROVISIONING state indicates the cluster is being created. |
| RUNNING | 2 | The RUNNING state indicates the cluster has been created and is fully usable. |
| RECONCILING | 3 | The RECONCILING state indicates that some work is actively being done on the cluster, such as upgrading the master or node software. |
| STOPPING | 4 | The STOPPING state indicates the cluster is being deleted |
| ERROR | 5 | The ERROR state indicates the cluster may be unusable |
| DEGRADED | 6 | The DEGRADED state indicates the cluster requires user action to restore full functionality |


 

 


<a name="cmaaks.Cluster"></a>

### Cluster


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCluster | [CreateClusterMsg](#cmaaks.CreateClusterMsg) | [CreateClusterReply](#cmaaks.CreateClusterReply) | Will provision a cluster |
| GetCluster | [GetClusterMsg](#cmaaks.GetClusterMsg) | [GetClusterReply](#cmaaks.GetClusterReply) | Will retrieve the status of a cluster and its kubeconfig for connectivity |
| DeleteCluster | [DeleteClusterMsg](#cmaaks.DeleteClusterMsg) | [DeleteClusterReply](#cmaaks.DeleteClusterReply) | Will delete a cluster |
| GetClusterList | [GetClusterListMsg](#cmaaks.GetClusterListMsg) | [GetClusterListReply](#cmaaks.GetClusterListReply) | Will retrieve a list of clusters |
| GetVersionInformation | [GetVersionMsg](#cmaaks.GetVersionMsg) | [GetVersionReply](#cmaaks.GetVersionReply) | Will return version information about api server |
| GetClusterUpgrades | [GetClusterUpgradesMsg](#cmaaks.GetClusterUpgradesMsg) | [GetClusterUpgradesReply](#cmaaks.GetClusterUpgradesReply) | Will retrieve available upgrades of a cluster |
| UpgradeCluster | [UpgradeClusterMsg](#cmaaks.UpgradeClusterMsg) | [UpgradeClusterReply](#cmaaks.UpgradeClusterReply) | Will upgrade a cluster |
| GetClusterNodeCount | [GetClusterNodeCountMsg](#cmaaks.GetClusterNodeCountMsg) | [GetClusterNodeCountReply](#cmaaks.GetClusterNodeCountReply) | Will retrieve node count |
| ScaleCluster | [ScaleClusterMsg](#cmaaks.ScaleClusterMsg) | [ScaleClusterReply](#cmaaks.ScaleClusterReply) | Will scale a clusters node count |
| EnableClusterAutoscaling | [EnableClusterAutoscalingMsg](#cmaaks.EnableClusterAutoscalingMsg) | [EnableClusterAutoscalingReply](#cmaaks.EnableClusterAutoscalingReply) | Will enable cluster autoscaling |

 



<a name="api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="cmaaks.AzureClusterServiceAccount"></a>

### AzureClusterServiceAccount
the account used by the cluster to create azure resources (ex: load balancer)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_id | [string](#string) |  | The ClientId (aka: AppID) |
| client_secret | [string](#string) |  | The ClientSecret (aka: password) |






<a name="cmaaks.AzureCredentials"></a>

### AzureCredentials
The credentials to use for creating the cluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| app_id | [string](#string) |  | The AppId for API Access |
| tenant | [string](#string) |  | The Tenant for API access |
| password | [string](#string) |  | The Password for API access |
| subscription_id | [string](#string) |  | The Subscription for API access |






<a name="cmaaks.ClusterDetailItem"></a>

### ClusterDetailItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| kubeconfig | [string](#string) |  | What is the kubeconfig to connect to the cluster |
| status | [ClusterStatus](#cmaaks.ClusterStatus) |  | The status of the cluster |






<a name="cmaaks.ClusterItem"></a>

### ClusterItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| status | [ClusterStatus](#cmaaks.ClusterStatus) |  | The status of the cluster |






<a name="cmaaks.CreateClusterAKSSpec"></a>

### CreateClusterAKSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| location | [string](#string) |  | The Azure Data Center |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to build the cluster |
| clusterAccount | [AzureClusterServiceAccount](#cmaaks.AzureClusterServiceAccount) |  | Cluster service account used to talk to azure (ex: creating load balancer) |
| instance_groups | [CreateClusterAKSSpec.AKSInstanceGroup](#cmaaks.CreateClusterAKSSpec.AKSInstanceGroup) | repeated | Instance groups |
| tags | [CreateClusterAKSSpec.Tags](#cmaaks.CreateClusterAKSSpec.Tags) | repeated | Tags |






<a name="cmaaks.CreateClusterAKSSpec.AKSInstanceGroup"></a>

### CreateClusterAKSSpec.AKSInstanceGroup
Instance groups define a type and number of instances


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the group |
| type | [string](#string) |  | Instance type (Standard_D2_v2, etc.) |
| min_quantity | [int32](#int32) |  | Minimum number of instances (defaults to zero) |






<a name="cmaaks.CreateClusterAKSSpec.Tags"></a>

### CreateClusterAKSSpec.Tags
Tags are name/value pairs that enable you to categorize resources and view consolidated billing


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | Tag key |
| value | [string](#string) |  | Tag value |






<a name="cmaaks.CreateClusterMsg"></a>

### CreateClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be provisioned |
| provider | [CreateClusterProviderSpec](#cmaaks.CreateClusterProviderSpec) |  | The provider specification |






<a name="cmaaks.CreateClusterProviderSpec"></a>

### CreateClusterProviderSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the provider - like aks |
| k8s_version | [string](#string) |  | The version of Kubernetes |
| azure | [CreateClusterAKSSpec](#cmaaks.CreateClusterAKSSpec) |  | The AKS specification |
| high_availability | [bool](#bool) |  | Whether or not the cluster is HA |
| network_fabric | [string](#string) |  | The fabric to be used |






<a name="cmaaks.CreateClusterReply"></a>

### CreateClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was provisioned by this request |
| cluster | [ClusterItem](#cmaaks.ClusterItem) |  | The details of the cluster request response |






<a name="cmaaks.DeleteClusterMsg"></a>

### DeleteClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster&#39;s name to destroy |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to delete the cluster |






<a name="cmaaks.DeleteClusterReply"></a>

### DeleteClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Could the cluster be destroyed |
| status | [string](#string) |  | Status of the request |






<a name="cmaaks.EnableClusterAutoscalingMsg"></a>

### EnableClusterAutoscalingMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster |
| min_quantity | [int32](#int32) |  | Minimum number of nodes for autoscaling |
| max_quantity | [int32](#int32) |  | Maximum number of nodes for autoscaling |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to access the cluster |






<a name="cmaaks.EnableClusterAutoscalingReply"></a>

### EnableClusterAutoscalingReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not cluster autoscaling was enabled by this request |






<a name="cmaaks.GetClusterListMsg"></a>

### GetClusterListMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to search subscription for clusters |






<a name="cmaaks.GetClusterListReply"></a>

### GetClusterListReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| clusters | [ClusterItem](#cmaaks.ClusterItem) | repeated | List of clusters |






<a name="cmaaks.GetClusterMsg"></a>

### GetClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be looked up |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to query for the cluster |






<a name="cmaaks.GetClusterNodeCountMsg"></a>

### GetClusterNodeCountMsg
Get available node count


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be looked up |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to query for the cluster |






<a name="cmaaks.GetClusterNodeCountReply"></a>

### GetClusterNodeCountReply
Reply with available node count


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| name | [string](#string) |  | the available node pool name |
| count | [int32](#int32) |  | The available node count |






<a name="cmaaks.GetClusterReply"></a>

### GetClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| cluster | [ClusterDetailItem](#cmaaks.ClusterDetailItem) |  |  |






<a name="cmaaks.GetClusterUpgradesMsg"></a>

### GetClusterUpgradesMsg
Get available cluster upgrades


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be looked up |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to query for the cluster |






<a name="cmaaks.GetClusterUpgradesReply"></a>

### GetClusterUpgradesReply
Reply with available upgrades


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| upgrades | [Upgrade](#cmaaks.Upgrade) | repeated | The available upgrades |






<a name="cmaaks.GetVersionMsg"></a>

### GetVersionMsg
Get version of API Server






<a name="cmaaks.GetVersionReply"></a>

### GetVersionReply
Reply for version request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | If operation was OK |
| version_information | [GetVersionReply.VersionInformation](#cmaaks.GetVersionReply.VersionInformation) |  | Version Information |






<a name="cmaaks.GetVersionReply.VersionInformation"></a>

### GetVersionReply.VersionInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| git_version | [string](#string) |  | The tag on the git repository |
| git_commit | [string](#string) |  | The hash of the git commit |
| git_tree_state | [string](#string) |  | Whether or not the tree was clean when built |
| build_date | [string](#string) |  | Date of build |
| go_version | [string](#string) |  | Version of go used to compile |
| compiler | [string](#string) |  | Compiler used |
| platform | [string](#string) |  | Platform it was compiled for / running on |






<a name="cmaaks.ScaleClusterMsg"></a>

### ScaleClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of cluster |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to search subscription for clusters |
| node_pool | [string](#string) |  | name of node pool |
| count | [int32](#int32) |  | total desired nodes |






<a name="cmaaks.ScaleClusterReply"></a>

### ScaleClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | wheather or not the cluster was scaled |
| status | [string](#string) |  | Status of the request |






<a name="cmaaks.Upgrade"></a>

### Upgrade
available upgrade version


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  |  |






<a name="cmaaks.UpgradeClusterAKSSpec"></a>

### UpgradeClusterAKSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credentials | [AzureCredentials](#cmaaks.AzureCredentials) |  | Credentials to build the cluster |






<a name="cmaaks.UpgradeClusterMsg"></a>

### UpgradeClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be upgraded |
| provider | [UpgradeClusterProviderSpec](#cmaaks.UpgradeClusterProviderSpec) |  | The provider specification |






<a name="cmaaks.UpgradeClusterProviderSpec"></a>

### UpgradeClusterProviderSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the provider - like aks |
| k8s_version | [string](#string) |  | The version of Kubernetes |
| azure | [UpgradeClusterAKSSpec](#cmaaks.UpgradeClusterAKSSpec) |  | The AKS specification |






<a name="cmaaks.UpgradeClusterReply"></a>

### UpgradeClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was upgrade by this request |
| cluster | [ClusterItem](#cmaaks.ClusterItem) |  | The details of the cluster request response |





 


<a name="cmaaks.ClusterStatus"></a>

### ClusterStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 | Not set |
| PROVISIONING | 1 | The PROVISIONING state indicates the cluster is being created. |
| RUNNING | 2 | The RUNNING state indicates the cluster has been created and is fully usable. |
| RECONCILING | 3 | The RECONCILING state indicates that some work is actively being done on the cluster, such as upgrading the master or node software. |
| STOPPING | 4 | The STOPPING state indicates the cluster is being deleted |
| ERROR | 5 | The ERROR state indicates the cluster may be unusable |
| DEGRADED | 6 | The DEGRADED state indicates the cluster requires user action to restore full functionality |


 

 


<a name="cmaaks.Cluster"></a>

### Cluster


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCluster | [CreateClusterMsg](#cmaaks.CreateClusterMsg) | [CreateClusterReply](#cmaaks.CreateClusterReply) | Will provision a cluster |
| GetCluster | [GetClusterMsg](#cmaaks.GetClusterMsg) | [GetClusterReply](#cmaaks.GetClusterReply) | Will retrieve the status of a cluster and its kubeconfig for connectivity |
| DeleteCluster | [DeleteClusterMsg](#cmaaks.DeleteClusterMsg) | [DeleteClusterReply](#cmaaks.DeleteClusterReply) | Will delete a cluster |
| GetClusterList | [GetClusterListMsg](#cmaaks.GetClusterListMsg) | [GetClusterListReply](#cmaaks.GetClusterListReply) | Will retrieve a list of clusters |
| GetVersionInformation | [GetVersionMsg](#cmaaks.GetVersionMsg) | [GetVersionReply](#cmaaks.GetVersionReply) | Will return version information about api server |
| GetClusterUpgrades | [GetClusterUpgradesMsg](#cmaaks.GetClusterUpgradesMsg) | [GetClusterUpgradesReply](#cmaaks.GetClusterUpgradesReply) | Will retrieve available upgrades of a cluster |
| UpgradeCluster | [UpgradeClusterMsg](#cmaaks.UpgradeClusterMsg) | [UpgradeClusterReply](#cmaaks.UpgradeClusterReply) | Will upgrade a cluster |
| GetClusterNodeCount | [GetClusterNodeCountMsg](#cmaaks.GetClusterNodeCountMsg) | [GetClusterNodeCountReply](#cmaaks.GetClusterNodeCountReply) | Will retrieve node count |
| ScaleCluster | [ScaleClusterMsg](#cmaaks.ScaleClusterMsg) | [ScaleClusterReply](#cmaaks.ScaleClusterReply) | Will scale a clusters node count |
| EnableClusterAutoscaling | [EnableClusterAutoscalingMsg](#cmaaks.EnableClusterAutoscalingMsg) | [EnableClusterAutoscalingReply](#cmaaks.EnableClusterAutoscalingReply) | Will enable cluster autoscaling |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

