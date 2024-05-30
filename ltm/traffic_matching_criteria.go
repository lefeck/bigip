package ltm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

//	You can use this traffic-matching-criteria component to configure the
//	virtual server matching definitions on the Local Traffic Manager. This
//	defines the criteria used for matching destination and source
//	address/ports for a virtual server.
//
// A TrafficMatchingCriteriaList contains a list of TrafficMatchingCriteria.
type TrafficMatchingCriteriaList struct {
	Kind     string                    `json:"kind"`
	SelfLink string                    `json:"selfLink"`
	Items    []TrafficMatchingCriteria `json:"items"`
}

// TrafficMatchingCriteria represents an  BIG-IP LTM TrafficMatchingCriteria configuration.
type TrafficMatchingCriteria struct {
	Kind                            string `json:"kind"`
	Name                            string `json:"name"`
	Partition                       string `json:"partition"`
	FullPath                        string `json:"fullPath"`
	Generation                      int    `json:"generation"`
	SelfLink                        string `json:"selfLink"`
	DestinationAddressInline        string `json:"destinationAddressInline"`
	DestinationAddressList          string `json:"destinationAddressList"`
	DestinationAddressListReference struct {
		Link string `json:"link"`
	} `json:"destinationAddressListReference"`
	DestinationPortInline        string `json:"destinationPortInline"`
	DestinationPortList          string `json:"destinationPortList"`
	DestinationPortListReference struct {
		Link string `json:"link"`
	} `json:"destinationPortListReference"`
	Protocol                   string `json:"protocol"`
	RouteDomain                string `json:"routeDomain"`
	SourceAddressInline        string `json:"sourceAddressInline"`
	SourceAddressList          string `json:"sourceAddressList"`
	SourceAddressListReference struct {
		Link string `json:"link"`
	} `json:"sourceAddressListReference"`
	SourcePortInline int `json:"sourcePortInline"`
}

// TrafficMatchingCriteriaEndpoint represents the REST resource for managing TrafficMatchingCriteria.
const TrafficMatchingCriteriaEndpoint = "traffic-matching-criteria"

// TrafficMatchingCriteriaResource provides an API to manage virtual server of the address list object.
type TrafficMatchingCriteriaResource struct {
	b *bigip.BigIP
}

// List all TrafficMatchingCriteria details
func (nr *TrafficMatchingCriteriaResource) List() (*TrafficMatchingCriteriaList, error) {
	var nl TrafficMatchingCriteriaList
	res, err := nr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(TrafficMatchingCriteriaEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &nl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &nl, nil
}

// Get a single TrafficMatchingCriteria details by the TrafficMatchingCriteria name
func (nr *TrafficMatchingCriteriaResource) Get(name string) (*TrafficMatchingCriteria, error) {
	var node TrafficMatchingCriteria
	res, err := nr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(TrafficMatchingCriteriaEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &node); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &node, nil
}

// Create a new TrafficMatchingCriteria item
func (nr *TrafficMatchingCriteriaResource) Create(item Node) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = nr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(TrafficMatchingCriteriaEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update the TrafficMatchingCriteria item identified by the TrafficMatchingCriteria name, otherwise an error will be reported.
func (nr *TrafficMatchingCriteriaResource) Update(name string, item TrafficMatchingCriteria) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = nr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(TrafficMatchingCriteriaEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single TrafficMatchingCriteria identified by the TrafficMatchingCriteria name. if it is not exist return error
func (nr *TrafficMatchingCriteriaResource) Delete(name string) error {
	_, err := nr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(TrafficMatchingCriteriaEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
