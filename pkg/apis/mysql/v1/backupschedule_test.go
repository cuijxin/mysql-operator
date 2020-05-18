package v1

import (
	"strings"
	"testing"

	"github.com/cuijxin/mysql-operator/pkg/version"
	corev1 "k8s.io/api/core/v1"
)

func TestEmptyBackupScheduleIsInvalid(t *testing.T) {
	bs := MySQLBackupSchedule{}
	err := bs.Validate()
	if err == nil {
		t.Error("An empty backup schedule should be invalid")
	}
}

func TestValidateValidBackupSchedule(t *testing.T) {
	bs := MySQLBackupSchedule{
		Spec: BackupScheduleSpec{
			Schedule: "* * * * * *",
			BackupTemplate: BackupSpec{
				Executor: &Executor{
					Provider:  "mysqldump",
					Databases: []string{"db1", "db2"},
				},
				Storage: &Storage{
					Provider: "s3",
					SecretRef: &corev1.LocalObjectReference{
						Name: "backup-storage-creds",
					},
					Config: map[string]string{
						"endpoint": "endpoint",
						"region":   "region",
						"bucket":   "bucket",
					},
				},
				ClusterRef: &corev1.LocalObjectReference{
					Name: "test-cluster",
				},
			},
		},
	}
	bs.Labels = make(map[string]string)
	SetOperatorVersionLabel(bs.Labels, "v1.0.0")
	err := bs.Validate()
	if err != nil {
		t.Errorf("Expected no validation errors but got %s", err)
	}
}

func TestBackupScheduleEnsureDefaultVersionSet(t *testing.T) {
	expected := version.GetBuildVersion()
	bs := &MySQLBackupSchedule{}
	bs = bs.EnsureDefaults()

	actual := GetOperatorVersionLabel(bs.Labels)
	if actual != expected {
		t.Errorf("Expected version '%s' but got '%s'", expected, actual)
	}
}

func TestBackupScheduleEnsureDefaultVersionNotSetIfExists(t *testing.T) {
	version := "v1.0.0"
	bs := &MySQLBackupSchedule{}
	bs.Labels = make(map[string]string)
	SetOperatorVersionLabel(bs.Labels, version)
	bs = bs.EnsureDefaults()

	actual := GetOperatorVersionLabel(bs.Labels)

	if actual != version {
		t.Errorf("Expected version '%s' but got '%s'", version, actual)
	}
}

func TestValidateBackupScheduleMissingCluster(t *testing.T) {
	bs := MySQLBackupSchedule{
		Spec: BackupScheduleSpec{
			Schedule: "* * * * * *",
			BackupTemplate: BackupSpec{
				Executor: &Executor{
					Provider:  "mysqldump",
					Databases: []string{"db1", "db2"},
				},
				Storage: &Storage{
					Provider: "s3",
					SecretRef: &corev1.LocalObjectReference{
						Name: "backup-storage-creds",
					},
					Config: map[string]string{
						"endpoint": "endpoint",
						"region":   "region",
						"bucket":   "bucket",
					},
				},
				AgentScheduled: "hostname-1",
			},
		},
	}

	err := bs.Validate()
	if !strings.Contains(err.Error(), "missing cluster") {
		t.Errorf("Expected backup schedule with missing Cluster to show 'missing cluster' error. Error is: %s", err)
	}
}

func TestValidateBackupScheduleMissingSecretRef(t *testing.T) {
	bs := MySQLBackupSchedule{
		Spec: BackupScheduleSpec{
			Schedule: "* * * * * *",
			BackupTemplate: BackupSpec{
				Executor: &Executor{
					Provider:  "mysqldump",
					Databases: []string{"db1", "db2"},
				},
				Storage: &Storage{
					Provider: "s3",
					Config: map[string]string{
						"endpoint": "endpoint",
						"region":   "region",
						"bucket":   "bucket",
					},
				},
				ClusterRef: &corev1.LocalObjectReference{
					Name: "test-cluster",
				},
				AgentScheduled: "hostname-1",
			},
		},
	}

	err := bs.Validate()
	if !strings.Contains(err.Error(), "storage.secretRef: Required value") {
		t.Errorf("Expected backup schedule with missing SecretRef to show 'storage.secretRef: Required value' error. Error is: %s", err)
	}
}
