package e2e

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	v1 "github.com/cuijxin/mysql-operator/pkg/apis/mysql/v1"
	"github.com/cuijxin/mysql-operator/test/e2e/framework"
)

var _ = Describe("MySQLCluster creation", func() {
	f := framework.NewDefaultFramework("cluster-creation")

	It("should be possible to create a basic 3 member cluster with a 28 character name", func() {
		clusterName := "basic-twenty-eight-char-name"
		Expect(clusterName).To(HaveLen(28))

		jig := framework.NewMySQLClusterTestJig(f.MySQLClientSet, f.ClientSet, clusterName)

		cluster := jig.CreateAndAwaitMySQLClusterOrFail(f.Namespace.Name, 3, nil, framework.DefaultTimeout)

		expected, err := framework.WriteSQLTest(cluster, cluster.Name+"-0")
		Expect(err).NotTo(HaveOccurred())

		actual, err := framework.ReadSQLTest(cluster, cluster.Name+"-0")
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(Equal(expected))
	})

	It("should be possible to create a multi-master cluster", func() {
		clusterName := "multi-master"
		replicas := int32(3)

		jig := framework.NewMySQLClusterTestJig(f.MySQLClientSet, f.ClientSet, clusterName)

		cluster := jig.CreateAndAwaitMySQLClusterOrFail(f.Namespace.Name, replicas, func(cluster *v1.MySQLCluster) {
			cluster.Spec.MultiMaster = true
		}, framework.DefaultTimeout)

		By("Checking we can write to and read from to all members")
		for i := int32(0); i < replicas; i++ {
			member := fmt.Sprintf("%s-%d", cluster.Name, i)
			By(fmt.Sprintf("Checking that we can write to and read from %q", member))

			expected, err := framework.WriteSQLTest(cluster, member)
			Expect(err).NotTo(HaveOccurred())

			actual, err := framework.ReadSQLTest(cluster, member)
			Expect(err).NotTo(HaveOccurred())

			Expect(actual).To(Equal(expected))
		}
	})
})
