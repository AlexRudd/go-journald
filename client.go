package gojournald

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const DefaultHost = "http://localhost:19531"

type Journaler interface {
	Machine() (*MachineOutput, error)
}

type MachineOutput struct {
	MachineID *string `json:"machine_id"`
}

type Journal struct {
	c   *http.Client
	url string
}

func NewJournal() *Journal {
	return &Journal{
		c:   http.DefaultClient,
		url: DefaultHost,
	}
}

func (j *Journal) Configure(opt func(*Journal) error) error {
	return opt(j)
}

func (j *Journal) Machine() (*MachineOutput, error) {
	req, _ := http.NewRequest("get", j.url+"/machine", nil)
	req.Header.Set("Accept", "application/json")
	resp, err := j.c.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	mo := &MachineOutput{}
	err = json.Unmarshal(body, mo)
	if err != nil {
		return nil, err
	}
	return mo, nil
}
