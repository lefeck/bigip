package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type QOEList struct {
	Items    []QOE  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

type QOE struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	AppService   string `json:"appService"`
	DefaultsFrom string `json:"defaultsFrom"`
	Description  string `json:"description"`
	Video        string `json:"video"`
}

const QOEEndpoint = "qoe"

type QOEResource struct {
	b *bigip.BigIP
}

// List retrieves a list of QOE resources.
func (qr *QOEResource) List() (*QOEList, error) {
	var items QOEList
	res, err := qr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(QOEEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a QOE resource by its full path name.
func (qr *QOEResource) Get(fullPathName string) (*QOE, error) {
	var item QOE
	res, err := qr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(QOEEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new QOE resource using the provided QOE item.
func (qr *QOEResource) Create(item QOE) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	_, err = qr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(QOEEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a QOE resource identified by its full path name using the provided QOE item.
func (qr *QOEResource) Update(fullPathName string, item QOE) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	_, err = qr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(QOEEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a QOE resource by its full path name.
func (qr *QOEResource) Delete(fullPathName string) error {
	_, err := qr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(QOEEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
