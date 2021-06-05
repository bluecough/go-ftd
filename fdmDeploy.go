package goftd

import (
	"fmt"
)

type DeployObject struct {
	ReferenceObject
	Description     string `json:"description,omitempty"`
	SubType         string `json:"subType"`
	Value           string `json:"value"`
	IsSystemDefined bool   `json:"isSystemDefined,omitempty"`
	Links           *Links `json:"links,omitempty"`
}

// Reference Returns a reference object
func (n *DeployObject) Reference() *ReferenceObject {
	r := ReferenceObject{
		ID:      n.ID,
		Name:    n.Name,
		Version: n.Version,
		Type:    n.Type,
	}

	return &r
}


func (f *FTD) PostDeploy(n *DeployObject) error {
	var err error

	// adding a comment line to see

	_, err = f.Post(apiDeploy, nil)
	if err != nil {
		fmt.Errorf("error: %s\n", err)
	}

	return nil
}
