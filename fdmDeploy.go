package goftd
import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/golang/glog"
)

func (f *FTD) postDeploy(limit int) error {
	var err error
	endpoint := apiDeploy

	filter := make(map[string]string)
	filter["limit"] = strconv.Itoa(limit)

	_, err = f.Post(endpoint, limit)
}