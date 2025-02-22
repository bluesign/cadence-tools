/*
 * Cadence lint - The Cadence linter
 *
 * Copyright Flow Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lint_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence/ast"
	"github.com/onflow/cadence/tools/analysis"

	"github.com/onflow/cadence-tools/lint"
)

func TestForceOperatorAnalyzer(t *testing.T) {

	t.Parallel()

	t.Run("unnecessary", func(t *testing.T) {

		t.Parallel()

		diagnostics := testAnalyzers(t,
			`
			access(all) contract Test {
				access(all) fun test() {
					let x = 3
					let y = x!
				}
			}
			`,
			lint.UnnecessaryForceAnalyzer,
		)

		require.Equal(
			t,
			[]analysis.Diagnostic{
				{
					Range: ast.Range{
						StartPos: ast.Position{Offset: 89, Line: 5, Column: 13},
						EndPos:   ast.Position{Offset: 90, Line: 5, Column: 14},
					},
					Location: testLocation,
					Category: lint.RemovalCategory,
					Message:  "unnecessary force operator",
				},
			},
			diagnostics,
		)
	})

	t.Run("valid", func(t *testing.T) {

		t.Parallel()

		diagnostics := testAnalyzers(t,
			`
			access(all) contract Test {
				access(all) fun test() {
					let x: Int? = 3
					let y = x!
				}
			}
			`,
			lint.UnnecessaryForceAnalyzer,
		)

		require.Equal(
			t,
			[]analysis.Diagnostic(nil),
			diagnostics,
		)
	})
}
