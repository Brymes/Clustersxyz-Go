package clusters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetName(t *testing.T) {
	var (
		address string
		client  = NewCluster()
	)
	tests := []struct {
		testcase string
		setup    func(*testing.T)
		expect   func(*testing.T, string, error)
	}{
		{
			testcase: "Test Get Name Not Found",
			setup: func(t *testing.T) {
				address = "asfas"
			},
			expect: func(t *testing.T, name string, err error) {
				assert.Equal(t, name, "")
				assert.NotNil(t, err)
				assert.ErrorContains(t, err, ErrorNotFound.Error())
			},
		},
		{
			testcase: "Test Get Name Successful",
			setup: func(t *testing.T) {
				address = "0x5755d1dcea21caa687339c305d143e6e78f96adf"
			},
			expect: func(t *testing.T, name string, err error) {
				assert.Nil(t, err)
				assert.Equal(t, name, "clusters/main")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testcase, func(t *testing.T) {
			test.setup(t)
			res, err := client.GetName(address)
			test.expect(t, res, err)
		})
	}
}

func TestGetAddress(t *testing.T) {
	var (
		clustersName, walletName string
		client                   = NewCluster()
	)
	tests := []struct {
		testcase string
		setup    func(*testing.T) (GetAddressResponse, error)
		expect   func(*testing.T, *GetAddressResponse, error)
	}{
		{
			testcase: "Test Get Wallet by ClusterName",
			setup: func(t *testing.T) (GetAddressResponse, error) {
				clustersName = "clusters"
				return client.GetAddress(clustersName)
			},
			expect: func(t *testing.T, target *GetAddressResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, target)
				assert.Equal(t, target.Address, "0x5755d1dcea21caa687339c305d143e6e78f96adf")
			},
		},
		{
			testcase: "Test Get Wallet by ClusterName and AddressName",
			setup: func(t *testing.T) (GetAddressResponse, error) {
				clustersName, walletName = "clusters", "main"
				return client.GetAddressWithName(clustersName, walletName)
			},
			expect: func(t *testing.T, target *GetAddressResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, target.Address, "0x5755d1dcea21caa687339c305d143e6e78f96adf")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testcase, func(t *testing.T) {
			res, err := test.setup(t)
			test.expect(t, &res, err)
		})
	}
}

func TestBulkGetNames(t *testing.T) {
	var (
		addresses []string
		client    = NewCluster()
	)
	tests := []struct {
		testcase string
		setup    func(*testing.T)
		expect   func(*testing.T, []GetClusterNameResponse, error)
	}{
		{
			testcase: "Test Bulk  Get Name Not Found",
			setup: func(t *testing.T) {
				addresses = []string{"add"}
			},
			expect: func(t *testing.T, resp []GetClusterNameResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, resp[0].Name, "")
			},
		},
		{
			testcase: "Test Bulk Get Names Successful",
			setup: func(t *testing.T) {
				addresses = []string{"0x5755d1dcea21caa687339c305d143e6e78f96adf", "0xccdead94e8cf17de32044d9701c4f5668ad0bef9"}

			},
			expect: func(t *testing.T, resp []GetClusterNameResponse, Error error) {
				assert.Nil(t, Error)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testcase, func(t *testing.T) {
			test.setup(t)
			res, Error := client.BulkGetNames(addresses)
			test.expect(t, res, Error)
		})
	}
}

func TestBulkGetAddresses(t *testing.T) {
	var (
		addresses []string
		client    = NewCluster()
	)
	tests := []struct {
		testcase string
		setup    func(*testing.T)
		expect   func(*testing.T, []GetAddressResponse, error)
	}{
		{
			testcase: "Test Bulk Get Addresses Successful",
			setup: func(t *testing.T) {
				addresses = []string{"0x5755d1dcea21caa687339c305d143e6e78f96adf"}

			},
			expect: func(t *testing.T, resp []GetAddressResponse, Error error) {
				assert.Nil(t, Error)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testcase, func(t *testing.T) {
			test.setup(t)
			res, Error := client.BulkGetAddresses(addresses)
			test.expect(t, res, Error)
		})
	}
}

func TestGetCluster(t *testing.T) {
	var (
		cluster string
		client  = NewCluster()
	)
	tests := []struct {
		testcase string
		setup    func(*testing.T)
		expect   func(*testing.T, *GetClusterResponse, error)
	}{
		{
			testcase: "Test Get Cluster Successful",
			setup: func(t *testing.T) {
				cluster = "clusters/"

			},
			expect: func(t *testing.T, target *GetClusterResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, target.Name, cluster)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testcase, func(t *testing.T) {
			test.setup(t)
			res, err := client.GetCluster(cluster)
			test.expect(t, &res, err)
		})
	}
}

func TestBulkGetClusters(t *testing.T) {
	var (
		names  []string
		client = NewCluster()
	)
	tests := []struct {
		testcase string
		setup    func(*testing.T)
		expect   func(*testing.T, []GetClusterResponse, error)
	}{
		{
			testcase: "Test Bulk Get Clusters Successful",
			setup: func(t *testing.T) {
				names = []string{"clusters"}

			},
			expect: func(t *testing.T, resp []GetClusterResponse, Error error) {
				assert.Nil(t, Error)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testcase, func(t *testing.T) {
			test.setup(t)
			res, Error := client.BulkGetClusters(names)
			test.expect(t, res, Error)
		})
	}
}
