package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type OneConnectList struct {
	Items    []OneConnect `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

type OneConnect struct {
	Kind                string `json:"kind,omitempty"`
	Name                string `json:"name,omitempty"`
	Partition           string `json:"partition,omitempty"`
	FullPath            string `json:"fullPath,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	SelfLink            string `json:"selfLink,omitempty"`
	AppService          string `json:"appService,omitempty"`
	DefaultsFrom        string `json:"defaultsFrom,omitempty"`
	Description         string `json:"description,omitempty"`
	IdleTimeoutOverride string `json:"idleTimeoutOverride,omitempty"`
	LimitType           string `json:"limitType,omitempty"`
	MaxAge              int    `json:"maxAge,omitempty"`
	MaxReuse            int    `json:"maxReuse,omitempty"`
	MaxSize             int    `json:"maxSize,omitempty"`
	SharePools          string `json:"sharePools,omitempty"`
	SourceMask          string `json:"sourceMask,omitempty"`
}

const OneConnectEndpoint = "oneconnect"

type OneConnectResource struct {
	b *bigip.BigIP
}

// List retrieves a list of OneConnect resources.
func (cr *OneConnectResource) List() (*OneConnectList, error) {
	var items OneConnectList
	// Perform a GET request to retrieve a list of OneConnect resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(OneConnectEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into OneConnectList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an OneConnect resource by its full path name.
func (cr *OneConnectResource) Get(fullPathName string) (*OneConnect, error) {
	var item OneConnect
	// Perform a GET request to retrieve a specific OneConnect resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(OneConnectEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into OneConnect struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new OneConnect resource using the provided OneConnect item.
func (cr *OneConnectResource) Create(item OneConnect) error {
	// Marshal the OneConnect struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new OneConnect resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(OneConnectEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an OneConnect resource identified by its full path name using the provided OneConnect item.
func (cr *OneConnectResource) Update(fullPathName string, item OneConnect) error {
	// Marshal the OneConnect struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified OneConnect resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(OneConnectEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an OneConnect resource by its full path name.
func (cr *OneConnectResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified OneConnect resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(OneConnectEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
