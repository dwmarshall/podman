package images

import (
	"fmt"
	"os"

	"github.com/containers/common/pkg/report"
	"github.com/containers/podman/v3/cmd/podman/common"
	"github.com/containers/podman/v3/cmd/podman/registry"
	"github.com/containers/podman/v3/pkg/domain/entities"
	"github.com/spf13/cobra"
)

var (
	showTrustDescription = "Display trust policy for the system"
	showTrustCommand     = &cobra.Command{
		Annotations:       map[string]string{registry.EngineMode: registry.ABIMode},
		Use:               "show [options] [REGISTRY]",
		Short:             "Display trust policy for the system",
		Long:              showTrustDescription,
		RunE:              showTrust,
		Args:              cobra.MaximumNArgs(1),
		ValidArgsFunction: common.AutocompleteRegistries,
		Example:           "",
	}
)

var (
	showTrustOptions entities.ShowTrustOptions
)

func init() {
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Command: showTrustCommand,
		Parent:  trustCmd,
	})
	showFlags := showTrustCommand.Flags()
	showFlags.BoolVarP(&showTrustOptions.JSON, "json", "j", false, "Output as json")
	showFlags.StringVar(&showTrustOptions.PolicyPath, "policypath", "", "")
	showFlags.BoolVar(&showTrustOptions.Raw, "raw", false, "Output raw policy file")
	_ = showFlags.MarkHidden("policypath")
	showFlags.StringVar(&showTrustOptions.RegistryPath, "registrypath", "", "")
	_ = showFlags.MarkHidden("registrypath")
}

func showTrust(cmd *cobra.Command, args []string) error {
	trust, err := registry.ImageEngine().ShowTrust(registry.Context(), args, showTrustOptions)
	if err != nil {
		return err
	}

	switch {
	case showTrustOptions.Raw:
		fmt.Println(string(trust.Raw))
		return nil
	case showTrustOptions.JSON:
		b, err := json.MarshalIndent(trust.Policies, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	}
	rpt := report.New(os.Stdout, cmd.Name())
	defer rpt.Flush()

	rpt, err = rpt.Parse(report.OriginPodman,
		"{{range . }}{{.RepoName}}\t{{.Type}}\t{{.GPGId}}\t{{.SignatureStore}}\n{{end -}}")
	if err != nil {
		return err
	}
	return rpt.Execute(trust.Policies)
}
