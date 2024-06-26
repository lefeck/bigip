package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type ICAPList struct {
	Items    []ICAP `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type ICAP struct {
	Kind          string `json:"kind,omitempty"`
	Name          string `json:"name,omitempty"`
	Partition     string `json:"partition,omitempty"`
	FullPath      string `json:"fullPath,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	SelfLink      string `json:"selfLink,omitempty"`
	AppService    string `json:"appService,omitempty"`
	DefaultsFrom  string `json:"defaultsFrom,omitempty"`
	HeaderFrom    string `json:"headerFrom,omitempty"`
	Host          string `json:"host,omitempty"`
	PreviewLength int    `json:"previewLength,omitempty"`
	Referer       string `json:"referer,omitempty"`
	URI           string `json:"uri,omitempty"`
	UserAgent     string `json:"userAgent,omitempty"`
}

const ICAPEndpoint = "icap"

type ICAPResource struct {
	b *bigip.BigIP
}

// List retrieves a list of ICAP resources.
func (cr *ICAPResource) List() (*ICAPList, error) {
	var items ICAPList
	// Perform a GET request to retrieve a list of ICAP resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ICAPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into ICAPList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves an ICAP resource by its full path name.
func (cr *ICAPResource) Get(fullPathName string) (*ICAP, error) {
	var item ICAP
	// Perform a GET request to retrieve a specific ICAP resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ICAPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into ICAP struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new ICAP resource using the provided ICAP item.
func (cr *ICAPResource) Create(item ICAP) error {
	// Marshal the ICAP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new ICAP resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ICAPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an ICAP resource identified by its full path name using the provided ICAP item.
func (cr *ICAPResource) Update(fullPathName string, item ICAP) error {
	// Marshal the ICAP struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified ICAP resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ICAPEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes an ICAP resource by its full path name.
func (cr *ICAPResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified ICAP resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(ICAPEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
