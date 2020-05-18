package constants

// MySQLClusterLabel is applied to all components of a MySQL cluster
const MySQLClusterLabel = "v1.mysql5.xintelligent.com/cluster"

// MySQLOperatorVersionLabel denotes the version of the MySQLOperator and
// MySQLOperatorAgent running in the cluster.
const MySQLOperatorVersionLabel = "v1.mysql5.xintelligent.com/version"

// LabelMySQLClusterRole specifies the role of a Pod within a MySQLCluster.
const LabelMySQLClusterRole = "v1.mysql5.xintelligent.com/role"

// MySQLClusterRolePrimary denotes a primary InnoDB cluster member.
const MySQLClusterRolePrimary = "primary"

// MySQLClusterRoleSecondary denotes a secondary InnoDB cluster member.
const MySQLClusterRoleSecondary = "secondary"
