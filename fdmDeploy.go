package goftd

import (
	"fmt"
	"strconv"
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


func (f *FTD) postDeploy(n *DeployObject, limit int) error {
	var err error
	n.Type = "deployobject"
	endpoint := apiDeploy
	// adding a comment line to see
	filter := make(map[string]string)
	filter["limit"] = strconv.Itoa(limit)

	_, err = f.Post(endpoint, limit)
	if err != nil {
		fmt.Errorf("error: %s\n", err)
	}

	return nil
}
