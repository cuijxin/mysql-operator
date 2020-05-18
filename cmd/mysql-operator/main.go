package main

import (
	goflag "flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
	utilflag "k8s.io/apiserver/pkg/util/flag"

	"k8s.io/apiserver/pkg/util/logs"

	"github.com/cuijxin/mysql-operator/cmd/mysql-operator/app"
	"github.com/cuijxin/mysql-operator/cmd/mysql-operator/app/options"
	"github.com/cuijxin/mysql-operator/pkg/version"
)

const (
	configPath      = "/etc/mysql-operator/mysql-operator-config.yaml"
	metricsEndpoint = "0.0.0.0:8080"
)

func main() {
	fmt.Fprintf(os.Stderr, "Starting mysql-operator version '%s'\n", version.GetBuildVersion())

	opts, err := options.NewMySQLOperatorServer(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading config: %v\n", err)
		os.Exit(1)
	}

	opts.AddFlags(pflag.CommandLine)
	pflag.CommandLine.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()
	goflag.CommandLine.Parse([]string{})

	logs.InitLogs()
	defer logs.FlushLogs()

	pflag.VisitAll(func(flag *pflag.Flag) {
		glog.V(2).Infof("FLAG: --%s=%q", flag.Name, flag.Value)
	})

	if err := app.Run(opts); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
