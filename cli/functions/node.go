package functions

import (
	"fmt"
	"net/http"

	"github.com/gravitl/netmaker/models"
)

// GetNodes - fetch all nodes
func GetNodes(networkName ...string) *[]models.ApiNode {
	if len(networkName) == 1 {
		return request[[]models.ApiNode](http.MethodGet, "/api/nodes/"+networkName[0], nil)
	} else {
		return request[[]models.ApiNode](http.MethodGet, "/api/nodes", nil)
	}
}

// GetNodeByID - fetch a single node by ID
func GetNodeByID(networkName, nodeID string) *models.NodeGet {
	return request[models.NodeGet](http.MethodGet, fmt.Sprintf("/api/nodes/%s/%s", networkName, nodeID), nil)
}

// UpdateNode - update a single node
func UpdateNode(networkName, nodeID string, node *models.ApiNode) *models.ApiNode {
	return request[models.ApiNode](http.MethodPut, fmt.Sprintf("/api/nodes/%s/%s", networkName, nodeID), node)
}

// DeleteNode - delete a node
func DeleteNode(networkName, nodeID string) *models.SuccessResponse {
	return request[models.SuccessResponse](http.MethodDelete, fmt.Sprintf("/api/nodes/%s/%s", networkName, nodeID), nil)
}

// CreateEgress - turn a node into an egress
func CreateEgress(networkName, nodeID string, payload *models.EgressGatewayRequest) *models.ApiNode {
	return request[models.ApiNode](http.MethodPost, fmt.Sprintf("/api/nodes/%s/%s/creategateway", networkName, nodeID), payload)
}

// DeleteEgress - remove egress role from a node
func DeleteEgress(networkName, nodeID string) *models.ApiNode {
	return request[models.ApiNode](http.MethodDelete, fmt.Sprintf("/api/nodes/%s/%s/deletegateway", networkName, nodeID), nil)
}

// CreateIngress - turn a node into an ingress
func CreateIngress(networkName, nodeID string, failover bool) *models.ApiNode {
	return request[models.ApiNode](http.MethodPost, fmt.Sprintf("/api/nodes/%s/%s/createingress", networkName, nodeID), &struct {
		Failover bool `json:"failover"`
	}{Failover: failover})
}

// DeleteIngress - remove ingress role from a node
func DeleteIngress(networkName, nodeID string) *models.ApiNode {
	return request[models.ApiNode](http.MethodDelete, fmt.Sprintf("/api/nodes/%s/%s/deleteingress", networkName, nodeID), nil)
}

// UncordonNode - uncordon a node
func UncordonNode(networkName, nodeID string) *string {
	return request[string](http.MethodPost, fmt.Sprintf("/api/nodes/%s/%s/approve", networkName, nodeID), nil)
}
