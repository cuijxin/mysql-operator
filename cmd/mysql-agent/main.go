package main

import (
	goflag "flag"
	"fmt"
	"os"

	// "k8s.io/apiserver/pkg/util/logs"

	"github.com/golang/glog"
	"github.com/spf13/pflag"

	"github.com/cuijxin/mysql-operator/cmd/mysql-agent/app"
	"github.com/cuijxin/mysql-operator/cmd/mysql-agent/app/options"
	"github.com/cuijxin/mysql-operator/pkg/util/flag"
	"github.com/cuijxin/mysql-operator/pkg/util/logs"
	"github.com/cuijxin/mysql-operator/pkg/version"
)

func main() {
	fmt.Fprintf(os.Stderr, "Starting mysql-agent version %s\n", version.GetBuildVersion())

	opts := options.NewMySQLAgentOpts()

	opts.AddFlags(pflag.CommandLine)
	pflag.CommandLine.SetNormalizeFunc(flag.WordSepNormalizeFunc)
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
