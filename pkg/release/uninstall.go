package release

import (
	"fmt"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

func (rel *config) Uninstall() (*release.UninstallReleaseResponse, error) {
	client := action.NewUninstall(rel.Cfg())
	client.Timeout = rel.Timeout

	resp, err := client.Run(rel.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to uninstall release %s: %w", rel.Uniq(), err)
	}

	return resp, nil
}
