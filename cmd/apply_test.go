package cmd

import (
	"io/ioutil"
	"net/http"
	"testing"

	models "github.com/semaphoreci/cli/api/models"
	assert "github.com/stretchr/testify/assert"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func Test__ApplySecret__FromYaml__Response200(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	yaml_file := `
apiVersion: v1beta
kind: Secret
metadata:
  name: Test
  id: "8f100520-5ab9-469f-854a-87bae95f19b9"
data:
  env_vars:
  - value: A
    name: B
  files:
  - path: "a.txt"
    content: "21313123"
`

	yaml_file_path := "/tmp/secret.yaml"

	ioutil.WriteFile(yaml_file_path, []byte(yaml_file), 0644)

	received := ""

	httpmock.RegisterResponder("PATCH", "https://org.semaphoretext.xyz/api/v1beta/secrets/8f100520-5ab9-469f-854a-87bae95f19b9",
		func(req *http.Request) (*http.Response, error) {
			body, _ := ioutil.ReadAll(req.Body)

			received = string(body)

			return httpmock.NewStringResponse(200, received), nil
		},
	)

	RootCmd.SetArgs([]string{"apply", "-f", yaml_file_path})
	RootCmd.Execute()

	expected := `{"apiVersion":"v1beta","kind":"Secret","metadata":{"name":"Test","id":"8f100520-5ab9-469f-854a-87bae95f19b9"},"data":{"env_vars":[{"name":"B","value":"A"}],"files":[{"path":"a.txt","content":"21313123"}]}}`

	if received != expected {
		t.Errorf("Expected the API to receive PATCH secret with: %s, got: %s", expected, received)
	}
}

func Test__ApplyNotification__Response200(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var received *models.NotificationV1Alpha

	endpoint := "https://org.semaphoretext.xyz/api/v1alpha/notifications/test"

	httpmock.RegisterResponder("PATCH", endpoint,
		func(req *http.Request) (*http.Response, error) {
			body, _ := ioutil.ReadAll(req.Body)
			received, _ = models.NewNotificationV1AlphaFromJson(body)

			return httpmock.NewStringResponse(200, string(body)), nil
		},
	)

	RootCmd.SetArgs([]string{"apply", "-f", "../fixtures/notification.yml"})
	RootCmd.Execute()

	assert.Equal(t, received.Metadata.Name, "test")

	rule := received.Spec.Rules[0]

	assert.Equal(t, rule.Name, "Rule #1")
	assert.Equal(t, rule.Filter.Projects, []string{"cli"})
	assert.Equal(t, rule.Filter.Branches, []string{"master"})
	assert.Equal(t, rule.Filter.Pipelines, []string{"semaphore.yml"})
}

func Test__ApplyDashboard__Response200(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	yaml_file := `
apiVersion: v1alpha
kind: Dashboard
metadata:
  name: Test
  title: "Test Something"
  id: "8f100520-5ab9-469f-854a-87bae95f19b9"
spec:
  widgets:
    - name: "Workflows"
      type: list
      filters:
         github_uid: "{{ github_uid }}"
`

	yaml_file_path := "/tmp/project.yaml"

	ioutil.WriteFile(yaml_file_path, []byte(yaml_file), 0644)

	received := ""

	httpmock.RegisterResponder("PATCH", "https://org.semaphoretext.xyz/api/v1alpha/dashboards/8f100520-5ab9-469f-854a-87bae95f19b9",
		func(req *http.Request) (*http.Response, error) {
			body, _ := ioutil.ReadAll(req.Body)

			received = string(body)

			return httpmock.NewStringResponse(200, received), nil
		},
	)

	RootCmd.SetArgs([]string{"apply", "-f", yaml_file_path})
	RootCmd.Execute()

	expected := `{"apiVersion":"v1alpha","kind":"Dashboard","metadata":{"name":"Test","title":"Test Something","id":"8f100520-5ab9-469f-854a-87bae95f19b9"},"spec":{"widgets":[{"name":"Workflows","type":"list","filters":{"github_uid":"{{ github_uid }}"}}]}}`

	if received != expected {
		t.Errorf("Expected the API to receive PATCH dashbord with: %s, got: %s", expected, received)
	}
}

func Test__ApplyProject__FromYaml_Response200(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	yaml_file := `
apiVersion: v1alpha
kind: Project
metadata:
  name: Test
  id: a13949b7-b2f6-4286-8f26-3962d7e97828
spec:
  visibility: public
  repository:
    url: "git@github.com:/semaphoreci/cli.git"
    integration_type: github_token
`

	yaml_file_path := "/tmp/project.yaml"
	ioutil.WriteFile(yaml_file_path, []byte(yaml_file), 0644)

	received := ""

	httpmock.RegisterResponder("PATCH", "https://org.semaphoretext.xyz/api/v1alpha/projects/a13949b7-b2f6-4286-8f26-3962d7e97828",
		func(req *http.Request) (*http.Response, error) {
			body, _ := ioutil.ReadAll(req.Body)

			received = string(body)

			return httpmock.NewStringResponse(200, received), nil
		},
	)

	RootCmd.SetArgs([]string{"apply", "-f", yaml_file_path})
	RootCmd.Execute()

	expected := `{"apiVersion":"v1alpha","kind":"Project","metadata":{"name":"Test","id":"a13949b7-b2f6-4286-8f26-3962d7e97828"},"spec":{"visibility":"public","repository":{"url":"git@github.com:/semaphoreci/cli.git","forked_pull_requests":{},"pipeline_file":"","whitelist":{},"integration_type":"github_token"}}}`

	if received != expected {
		t.Errorf("Expected the API to receive PATCH project with: %s, got: %s", expected, received)
	}
}
