package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type ConnectorConfigList struct {
	Items    []ConnectorConfig `json:"items,omitempty"`
	Kind     string            `json:"kind,omitempty"`
	SelfLink string            `json:"selflink,omitempty"`
}

type ConnectorConfig struct {
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	SelfLink           string `json:"selfLink"`
	AppService         string `json:"appService"`
	ConnectOnData      string `json:"connectOnData"`
	ConnectionTimeout  int    `json:"connectionTimeout"`
	EntryVirtualServer string `json:"entryVirtualServer"`
	ServiceDownAction  string `json:"serviceDownAction"`
}

const ConnectorEndpoint = "connector"

type ConnectorResource struct {
	B *bigip.BigIP
}

func (cr *ConnectorResource) List() (*ConnectorConfigList, error) {
	var items ConnectorConfigList
	res, err := cr.B.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(ProfileEndpoint).SubResource(ConnectorEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

func (cr *ConnectorResource) Get(fullPathName string) (*ConnectorConfig, error) {
	var item ConnectorConfig
	res, err := cr.B.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(ProfileEndpoint).SubResource(ConnectorEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

func (cr *ConnectorResource) Create(item ConnectorConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.B.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(ProfileEndpoint).SubResource(ConnectorEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *ConnectorResource) Update(name string, item ConnectorConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = cr.B.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(ProfileEndpoint).SubResource(ConnectorEndpoint).SubResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (cr *ConnectorResource) Delete(name string) error {
	_, err := cr.B.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(ProfileEndpoint).SubResource(ConnectorEndpoint).SubResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
