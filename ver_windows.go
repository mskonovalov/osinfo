package osinfo

import (
	"runtime"
	"strconv"

	"golang.org/x/sys/windows/registry"
)

// GetVersion Windows returns version info
// fetching os info for modern windows versions is fairly simple
// version info is easily fetched in the registry
// Returns:
//		- r.Runtime
//		- r.Arch
//		- r.Name
//		- r.Version
//		- r.win.Build
func GetVersion() Release {
	inf := Release{
		Runtime: runtime.GOOS,
		Arch:    runtime.GOARCH,
		Name:    "unknown",
		Version: "unknown",
		Windows: WindowsRelease{
			Build: "unknown",
		},
	}

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return inf
	}

	defer k.Close()

	// extract info
	pname, _, err := k.GetStringValue("ProductName")
	if err == nil {
		inf.Name = pname
	}

	ver, _, err := k.GetIntegerValue("CurrentMajorVersionNumber")
	if err == nil {
		inf.Version = strconv.Itoa(int(ver))
	}

	build, _, err := k.GetStringValue("CurrentBuild")
	if err == nil {
		inf.Windows.Build = build
	}

	return inf
}
