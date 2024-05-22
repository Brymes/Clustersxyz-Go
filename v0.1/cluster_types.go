package clusters

type GetAddressResponse struct {
	IsVerified bool   `json:"isVerified"`
	Type       string `json:"type"`
	Address    string `json:"address"`
	Name       string `json:"name"`
}

type GetClusterResponse struct {
	Name       string          `json:"name"`
	ImageUrl   string          `json:"imageUrl"`
	ProfileUrl string          `json:"profileUrl"`
	Wallets    []ClusterWallet `json:"wallets"`
}

type CoreWallet struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type ClusterWallet struct {
	CoreWallet
	Type       string `json:"type"`
	IsVerified bool   `json:"isVerified"`
}

type SearchClusterBidsResponse struct {
	Name                 string      `json:"name"`
	HighestBidWei        string      `json:"highestBidWei"`
	HighestBoostedBidWei string      `json:"highestBoostedBidWei"`
	HighestBidAddress    interface{} `json:"highestBidAddress"`
	IsAvailable          bool        `json:"isAvailable"`
}

type ClusterAddresses struct {
	NextPage   string           `json:"nextPage"`
	LastUpdate int              `json:"lastUpdate"`
	Items      []ClusterAddress `json:"items"`
}

type ClusterAddress struct {
	Bytes32Address string `json:"bytes32Address"`
	Address        string `json:"address"`
	Type           string `json:"type"`
	ClusterName    string `json:"clusterName"`
	Name           string `json:"name"`
	IsVerified     bool   `json:"isVerified"`
	UpdatedAt      int    `json:"updatedAt"`
}

type ManageWalletsPayload struct {
	Add    []CoreWallet `json:"add"`
	Update []CoreWallet `json:"update"`
	Remove []string     `json:"remove"`
}

type GetClusterNameResponse struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}
