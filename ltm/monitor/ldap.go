package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type MonitorLDAPList struct {
	Items    []MonitorLDAP `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}
type MonitorLDAP struct {
	AppService          string `json:"appService,omitempty"`
	Base                string `json:"base,omitempty"`
	ChaseReferrals      string `json:"chaseReferrals,omitempty"`
	Debug               string `json:"debug,omitempty"`
	DefaultsFrom        string `json:"defaultsFrom,omitempty"`
	Description         string `json:"description,omitempty"`
	Destination         string `json:"destination,omitempty"`
	Filter              string `json:"filter,omitempty"`
	FullPath            string `json:"fullPath,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	Interval            int    `json:"interval,omitempty"`
	Kind                string `json:"kind,omitempty"`
	MandatoryAttributes string `json:"mandatoryAttributes,omitempty"`
	ManualResume        string `json:"manualResume,omitempty"`
	Name                string `json:"name,omitempty"`
	Partition           string `json:"partition,omitempty"`
	Security            string `json:"security,omitempty"`
	SelfLink            string `json:"selfLink,omitempty"`
	TimeUntilUp         int    `json:"timeUntilUp,omitempty"`
	Timeout             int    `json:"timeout,omitempty"`
	UpInterval          int    `json:"upInterval,omitempty"`
}

const MonitorLDAPEndpoint = "/monitor/ldap"

type MonitorLDAPResource struct {
	b *bigip.BigIP
}

func (mlr *MonitorLDAPResource) List() (*MonitorLDAPList, error) {
	var mlcl MonitorLDAPList
	res, err := mlr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorLDAPEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &mlcl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mlcl, nil
}

func (mlr *MonitorLDAPResource) Get(fullPathName string) (*MonitorLDAP, error) {
	var mlc MonitorLDAP
	res, err := mlr.b.RestClient.Get().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorLDAPEndpoint).ResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &mlc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &mlc, nil
}

func (mlr *MonitorLDAPResource) Create(item MonitorLDAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mlr.b.RestClient.Post().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorLDAPEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mlr *MonitorLDAPResource) Update(name string, item MonitorLDAP) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = mlr.b.RestClient.Put().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorLDAPEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mlr *MonitorLDAPResource) Delete(name string) error {
	_, err := mlr.b.RestClient.Delete().Prefix(BasePath).ResourceCategory(TMResource).ManagerName(LtmManager).
		Resource(MonitorLDAPEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
