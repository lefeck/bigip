package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"github.com/lefeck/go-bigip/ltm"
	"strings"
)

type MonitorGatewayICMPConfigList struct {
	Items    []MonitorGatewayICMPConfig `json:"items,omitempty"`
	Kind     string                     `json:"kind,omitempty"`
	SelfLink string                     `json:"selflink,omitempty"`
}

type MonitorGatewayICMPConfig struct {
	Adaptive                 string `json:"adaptive,omitempty"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType,omitempty"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue,omitempty"`
	AdaptiveLimit            int    `json:"adaptiveLimit,omitempty"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan,omitempty"`
	AppService               string `json:"appService,omitempty"`
	DefaultsFrom             string `json:"defaultsFrom,omitempty"`
	Description              string `json:"description,omitempty"`
	Destination              string `json:"destination,omitempty"`
	FullPath                 string `json:"fullPath,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	Interval                 int    `json:"interval,omitempty"`
	Kind                     string `json:"kind,omitempty"`
	ManualResume             string `json:"manualResume,omitempty"`
	Name                     string `json:"name,omitempty"`
	Partition                string `json:"partition,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty"`
	TimeUntilUp              int    `json:"timeUntilUp,omitempty"`
	Timeout                  int    `json:"timeout,omitempty"`
	Transparent              string `json:"transparent,omitempty"`
	UpInterval               int    `json:"upInterval,omitempty"`
}

const MonitorGatewayICMPEndpoint = "/monitor/gateway-icmp"

type MonitorGatewayICMPResource struct {
	b *bigip.BigIP
}

func (mgir *MonitorGatewayICMPResource) List() (*MonitorGatewayICMPConfigList, error) {
	var mgicl MonitorGatewayICMPConfigList
	res, err := mgir.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorGatewayICMPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mgicl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mgicl, nil
}

func (mgir *MonitorGatewayICMPResource) Get(fullPathName string) (*MonitorGatewayICMPConfig, error) {
	var mgic MonitorGatewayICMPConfig
	res, err := mgir.b.RestClient.Get().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorGatewayICMPEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mgic); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mgic, nil
}

func (mgir *MonitorGatewayICMPResource) Create(item MonitorGatewayICMPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mgir.b.RestClient.Post().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorGatewayICMPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mgir *MonitorGatewayICMPResource) Update(name string, item MonitorGatewayICMPConfig) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mgir.b.RestClient.Put().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorGatewayICMPEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mgir *MonitorGatewayICMPResource) Delete(name string) error {
	_, err := mgir.b.RestClient.Delete().Prefix(ltm.BasePath).ResourceCategory(ltm.TMResource).ManagerName(ltm.LtmManager).
		Resource(MonitorGatewayICMPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}