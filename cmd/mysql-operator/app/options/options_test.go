package options

import (
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestEnsureDefaults(t *testing.T) {
	server := MySQLOperatorServer{}
	server.EnsureDefaults()
	assertRequiredDefaults(t, server)
}

func assertRequiredDefaults(t *testing.T, s MySQLOperatorServer) {
	if &s == nil {
		t.Error("MySQLOperatorServer: was nil, expected a valid configuration.")
	}
	if len(s.Hostname) <= 0 {
		t.Errorf("MySQLOperatorServer: expected a non-zero length hostname, got '%s'.", s.Hostname)
	}
	if &s.Images == nil {
		t.Error("MySQLOperatorServer.Images: was nil, expected a valid configuration.")
	}
	if s.Images.MySQLServerImage != mysqlServer {
		t.Errorf("MySQLOperatorServer.Images.MySQLServerImage: was '%s', expected '%s'.", s.Images.MySQLServerImage, mysqlServer)
	}
	if s.Images.MySQLAgentImage != mysqlAgent {
		t.Errorf("MySQLOperatorServer.Images.MySQLAgentImage: was '%s', expected '%s'.", s.Images.MySQLAgentImage, mysqlAgent)
	}
	expectedDuration := v1.Duration{Duration: 43200000000000}
	if &s.MinResyncPeriod == nil {
		t.Errorf("MySQLOperatorServer.MinResyncPeriod: was nil, expected '%s'.", expectedDuration)
	}
	if s.MinResyncPeriod != expectedDuration {
		t.Errorf("MySQLOperatorServer.MinResyncPeriod: was '%s', expected '%s'.", s.MinResyncPeriod, expectedDuration)
	}
}

func TestEnsureDefaultsOverrideSafety(t *testing.T) {
	expected := mockMySQLOperatorServer()
	ensured := mockMySQLOperatorServer()
	ensured.EnsureDefaults()
	if expected != ensured {
		t.Errorf("MySQLOperatorServer.EnsureDefaults() should not modify pre-configured values.")
	}
}

func mockMySQLOperatorServer() MySQLOperatorServer {
	return MySQLOperatorServer{
		KubeConfig: "some-kube-config",
		Master:     "some-master",
		Hostname:   "some-hostname",
		Images: Images{
			MySQLServerImage: "some-mysql-img",
			MySQLAgentImage:  "some-agent-img",
		},
		MinResyncPeriod: v1.Duration{Duration: 42},
	}
}
