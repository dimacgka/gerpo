package sql

import (
	"context"
	"testing"

	"github.com/insei/fmap/v3"

	"github.com/insei/gerpo/types"
)

// Column implementation for testing
type testColumn struct {
	field          fmap.Field
	allowedActions map[types.SQLAction]bool
	sql            string
	types.Column
}

func (c *testColumn) GetPtr(model any) any {
	fields, _ := fmap.GetFrom(model)
	return fields.MustFind(c.sql).GetPtr(model)
}

func (c *testColumn) GetField() fmap.Field {
	return c.field
}

func (c *testColumn) Name() (string, bool) {
	if c.sql == "" {
		return "", false
	}
	return c.sql, true
}

func (c *testColumn) IsAllowedAction(action types.SQLAction) bool {
	return c.allowedActions[action]
}

func (c *testColumn) ToSQL(ctx context.Context) string {
	return c.sql
}

func TestStringGroupBuilder(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name       string
		initialSQL string
		columns    []types.Column
	}{
		{
			name:       "Initial SQL, no columns",
			initialSQL: "SELECT * FROM table ",
			columns:    nil,
		},
		{
			name:       "Columns with allowed actions",
			initialSQL: "SELECT * FROM table ",
			columns: []types.Column{
				&testColumn{
					allowedActions: map[types.SQLAction]bool{
						types.SQLActionGroup: true,
					},
					sql: "GROUP BY col1 ",
				},
				&testColumn{
					allowedActions: map[types.SQLAction]bool{
						types.SQLActionGroup: true,
					},
					sql: "GROUP BY col2 ",
				},
			},
		},
		{
			name:       "Columns with mixed allowed and disallowed actions",
			initialSQL: "SELECT * FROM table ",
			columns: []types.Column{
				&testColumn{
					allowedActions: map[types.SQLAction]bool{
						types.SQLActionGroup: true,
					},
					sql: "GROUP BY col1 ",
				},
				&testColumn{
					allowedActions: map[types.SQLAction]bool{
						types.SQLActionGroup: false,
					},
					sql: "DISALLOWED_GROUP_BY col2 ",
				},
				&testColumn{
					allowedActions: map[types.SQLAction]bool{
						types.SQLActionGroup: true,
					},
					sql: "GROUP BY col3 ",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			builder := &StringGroupBuilder{
				ctx: ctx,
				sql: tc.initialSQL,
			}

			builder.GroupBy(tc.columns...)
		})
	}
}