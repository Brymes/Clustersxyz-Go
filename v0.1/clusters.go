package clusters

import (
	"fmt"
	"net/http"
)

const BaseURL = "https://api.clusters.xyz/v0.1/"

type ClustersSDK struct {
	URL string
	body interface{}
	err  error
}

func NewCluster() *ClustersSDK {
	return &ClustersSDK{
		URL: BaseURL,
	}
}

// GetName  GET /v0.1/name/:address
func (cl *ClustersSDK) GetName(address string) (string, error) {
	name := ""
	cl.URL = BaseURL + "name/" + address
	cl.makeRequest(http.MethodGet, &name)
	return name, cl.err
}

// GetAddress GET /v0.1/address/:name
func (cl *ClustersSDK) GetAddress(name string) (target GetAddressResponse, err error) {
	cl.URL = BaseURL + fmt.Sprintf("address/%s", name)
	cl.makeRequest(http.MethodGet, &target)
	return target, cl.err
}

// GetAddressWithName GET /v0.1/address/:name/:addressName
func (cl *ClustersSDK) GetAddressWithName(name, addressName string) (target GetAddressResponse, err error) {
	cl.URL = BaseURL + fmt.Sprintf("address/%s/%s", name, addressName)
	cl.makeRequest(http.MethodGet, &target)
	return target, cl.err
}

// BulkGetNames  POST /v0.1/name/addresses
func (cl *ClustersSDK) BulkGetNames(addresses []string) (target []GetClusterNameResponse, err error) {
	cl.body = addresses
	cl.URL = BaseURL + "name/addresses"
	cl.makeRequest(http.MethodPost, &target)
	return target, cl.err
}

// BulkGetAddresses POST /v0.1/address/names
func (cl *ClustersSDK) BulkGetAddresses(addresses []string) (target []GetAddressResponse, err error) {
	cl.body = addresses
	cl.URL = BaseURL + "address/names"

	cl.makeRequest(http.MethodPost, &target)

	return target, cl.err
}

// GetCluster GET /v0.1/cluster/:name
func (cl *ClustersSDK) GetCluster(cluster string) (target GetClusterResponse, err error) {
	cl.URL = fmt.Sprintf("%scluster/%s", BaseURL, cluster)
	cl.makeRequest(http.MethodGet, &target)
	return target, cl.err
}

// BulkGetClusters  POST /v0.1/cluster/names
func (cl *ClustersSDK) BulkGetClusters(names []string) (target []GetClusterResponse, err error) {
	cl.body = names
	cl.URL = BaseURL + "cluster/names"

	cl.makeRequest(http.MethodPost, &target)

	return target, cl.err
}
