package flags

import (
	"flag"

	"github.com/spf13/pflag"
	"k8s.io/klog"
)

const (
	allowDownloadFlagShort = "d"
	allowDownloadFlagLong  = "allow-download"
	allowDownloadFlagUsage = "allow downloading Kubernetes binaries"
)

// RegisterAllowDownloadFlag registers allow download flag.
func RegisterAllowDownloadFlag(local *pflag.FlagSet) {
	local.Bool(allowDownloadFlagLong, true, "Allow downloading Kubernetes binaries")

	if f := flag.CommandLine.Lookup(allowDownloadFlagShort); f != nil {
		pflagFlag := pflag.PFlagFromGoFlag(f)
		pflagFlag.Name = allowDownloadFlagLong
		pflagFlag.Usage = allowDownloadFlagUsage
		local.AddFlag(pflagFlag)
	} else {
		klog.Fatalf("failed to find flag in global flagset (flag): %s", allowDownloadFlagShort)
	}
}
