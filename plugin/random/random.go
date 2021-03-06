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

package random

import "github.com/carbonrelay/konjure/plugin/random/generator"

var (
	// NewRandomCommand creates a new command for running random from the CLI
	NewRandomCommand = generator.NewRandomGeneratorCommand
	// NewRandomGeneratorExecPlugin creates a new command for running random as an executable plugin
	NewRandomGeneratorExecPlugin = generator.NewRandomGeneratorExecPlugin
)
