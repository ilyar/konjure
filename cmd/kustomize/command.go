/*
Copyright 2019 GramLabs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kustomize

import (
	"github.com/carbonrelay/konjure/cmd/berglas"
	"github.com/carbonrelay/konjure/cmd/helm"
	"github.com/carbonrelay/konjure/cmd/jsonnet"
	"github.com/carbonrelay/konjure/cmd/kustomize/edit"
	"github.com/carbonrelay/konjure/cmd/label"
	"github.com/spf13/cobra"
)

// The Kustomize command really just aggregates all the exec plugin commands in one place

const example = `
# Edit a kustomization to include a generator configuration file
# NOTE: This functionality will be removed when it makes it into Kustomize proper
konjure kustomize edit add generator my-konjure-plugin-config.yaml

# Install Konjure as a series of Kustomize plugins
konjure kustomize init
`

// TODO Add a helper for creating configurations

func NewKustomizeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "kustomize",
		Short:   "Extensions for Kustomize",
		Example: example,
	}

	cmd.AddCommand(newInitializeCommand())
	cmd.AddCommand(edit.NewEditCommand())

	cmd.AddCommand(berglas.NewBerglasGenerator())
	cmd.AddCommand(berglas.NewBerglasTransformer())
	cmd.AddCommand(helm.NewHelmGenerator())
	cmd.AddCommand(jsonnet.NewJsonnetGenerator())
	cmd.AddCommand(label.NewLabelTransformer())

	return cmd
}

func newInitializeCommand() *cobra.Command {
	opts := NewInitializeOptions()

	cmd := &cobra.Command{
		Use:          "init [PLUGIN...]",
		Short:        "Configure Kustomize plugins",
		Long:         "Manages your '~/.config/kustomize/plugin' directory to include symlinks back to the Konjure executable",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.Kinds = args
			return opts.Complete()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.OutOrStdout())
		},
	}

	cmd.Flags().StringVar(&opts.PluginDir, "plugins", "", "override the `path` to the plugin directory")
	cmd.Flags().StringVar(&opts.Source, "source", "", "override the `path` to the source executable")
	cmd.Flags().BoolVar(&opts.Prune, "prune", false, "remove old versions")
	cmd.Flags().BoolVarP(&opts.Verbose, "verbose", "v", false, "be more verbose")
	cmd.Flags().BoolVar(&opts.DryRun, "dry-run", false, "check existing plugins")

	return cmd
}
