package modrinth_test

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/bananapizzuh/miced/modrinth"
)

func TestGetProject(t *testing.T) {
	ModrinthClient := modrinth.NewClient("", "github.com/bananapizzuh/miced/pkg/modrinth")
	project, err := ModrinthClient.GetProject("cobblemon-fabric")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	project_json, err := json.Marshal(project)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	expected, err := os.ReadFile("testdata/cobblemon-fabric.json")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if string(expected) != string(project_json) {
		t.Log(reflect.DeepEqual(string(expected), string(project_json)))
	}

}
