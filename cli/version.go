package cli

import (
	"context"
	"encoding/json"
	"github.com/lefeck/go-bigip"
)

type VersionStats struct {
	Kind     string  `json:"kind"`
	SelfLink string  `json:"selfLink"`
	Entries  Entries `json:"entries"`
}

type Entries struct {
	HTTPSLocalhostMgmtTmCliVersion0 HTTPSLocalhostMgmtTmCliVersion0 `json:"https://localhost/mgmt/tm/cli/version/0"`
}

type HTTPSLocalhostMgmtTmCliVersion0 struct {
	NestedStats NestedStats `json:"nestedStats"`
}

type NestedStats struct {
	EntriesMenu EntriesMenu `json:"entries"`
}

type EntriesMenu struct {
	Active    Active    `json:"active"`
	Latest    Latest    `json:"latest"`
	Supported Supported `json:"supported"`
}

type Active struct {
	Description string `json:"description"`
}

type Latest struct {
	Description string `json:"description"`
}

type Supported struct {
	Description string `json:"description"`
}

type VersionStatsResoure struct {
	b *bigip.BigIP
}

// VersionEndpoint is the base path of the TM API.
const VersionEndpoint = "version"

// Get bigip device version
func (vsr *VersionStatsResoure) Get() (*VersionStats, error) {
	var vs *VersionStats
	res, err := vsr.b.RestClient.Get().Prefix(bigip.BasePath).ResourceCategory(bigip.TMResource).ManagerName(CliManager).
		Resource(VersionEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &vs); err != nil {
		panic(err)
	}

	return vs, nil
}
