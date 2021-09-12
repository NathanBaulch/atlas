// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schemautil_test

import (
	"testing"

	"ariga.io/atlas/sql/internal/schemautil"
	"ariga.io/atlas/sql/schema/schemaspec"
	"github.com/stretchr/testify/require"
)

func TestOverride(t *testing.T) {
	spec := schemautil.ColSpec("name", "string")
	override := &schemaspec.Override{
		Dialect: "mysql",
		Resource: schemaspec.Resource{
			Attrs: []*schemaspec.Attr{
				// A string field
				schemautil.StrLitAttr("type", "varchar(123)"),

				// A boolean field
				schemautil.LitAttr("null", "true"),

				// A Literal
				schemautil.StrLitAttr("default", "howdy"),

				// A custom attribute
				schemautil.LitAttr("custom", "1234"),
			},
		},
	}

	err := schemautil.Override(spec, override)
	require.NoError(t, err)
	require.EqualValues(t, "varchar(123)", spec.Type)
	require.EqualValues(t, `"howdy"`, spec.Default.V)
	require.True(t, spec.Null)
	custom, ok := spec.Attr("custom")
	require.True(t, ok)
	i, err := custom.Int()
	require.NoError(t, err)
	require.EqualValues(t, 1234, i)
}
