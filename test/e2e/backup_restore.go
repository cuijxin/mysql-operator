package e2e

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	clientset "k8s.io/client-go/kubernetes"

	corev1 "k8s.io/api/core/v1"

	v1 "github.com/cuijxin/mysql-operator/pkg/apis/mysql/v1"

	mysqlclientset "github.com/cuijxin/mysql-operator/pkg/generated/clientset/versioned"
)

var _ = Describe("Backup/Restore", func() {
	f := framework.NewDefaultFramework("backup-restore")

	var cs clientset.Interface
	var mcs mysqlclientset.Interface
	BeforeEach(func() {
		cs = f.ClientSet
		mcs = f.MySQLClientSet
	})

	It("should be possible to backup a cluster and restore the created backup", func() {
		clusterName := "backup-restore"
		ns := f.Namespace.Name

		clusterJig := framework.NewMySQLClusterTestJig(mcs, cs, clusterName)
		backupJig := framework.NewMySQLBackupTestJig(mcs, cs, clusterName)
		restoreJig := framework.NewMySQLRestoreTestJig(mcs, cs, clusterName)

		By("Creating a cluster to backup")

		cluster := clusterJig.CreateAndAwaitMySQLClusterOrFail(ns, 3, nil, framework.DefaultTimeout)

		By("Creating testdb in the cluster to be backed up")

		member := cluster.Name + "-0"
		expected, err := framework.WriteSQLTest(cluster, member)
		Expect(err).NotTo(HaveOccurred())

		By("Checking testdb is present")

		actual, err := framework.ReadSQLTest(cluster, member)
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(Equal(expected))

		By("creating a secret containing the S3 (compat.) upload credentials")

		secret, err := backupJig.CreateS3AuthSecret(ns, "s3-upload-creds")
		Expect(err).NotTo(HaveOccurred())

		By("Backing up testdb")

		dbs := []string{framework.TestDBName}
		backup := backupJig.CreateAndAwaitMySQLDumpBackupOrFail(ns, clusterName, dbs, func(b *v1.MySQLBackup) {
			b.Spec.Storage = &v1.Storage{
				Provider: "s3",
				SecretRef: &corev1.LocalObjectReference{
					Name: secret.Name,
				},
				Config: map[string]string{
					"endpoint": "bristoldev.compat.objectstorage.us-phoenix-1.oraclecloud.com",
					"region":   "us-phoenix-1",
					"bucket":   "trjl-test",
				},
			}
		}, framework.DefaultTimeout)

		Expect(backup.Status.Outcome.Location).NotTo(BeEmpty())

		By("Dropping testdb")

		_, err = framework.ExecuteSQL(cluster, member,
			fmt.Sprintf("DROP DATABASE IF EXISTS %s", framework.TestDBName))
		Expect(err).NotTo(HaveOccurred())

		By("Checking that testdb has been dropped")

		_, err = framework.ReadSQLTest(cluster, member)
		Expect(err).To(HaveOccurred())

		By("Restoring the backup")

		restore := restoreJig.CreateAndAwaitMySQLRestoreOrFail(ns, clusterName, backup.Name, nil, framework.DefaultTimeout)
		Expect(restore.Status.TimeCompleted).ToNot(BeZero())

		By("Checking testdb is present and contains the correct uuid")

		actual, err = framework.ReadSQLTest(cluster, member)
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(Equal(expected))
	})
})