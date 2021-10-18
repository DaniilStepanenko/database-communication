package jet

import "github.com/go-jet/jet/v2/postgres"

func asIs(col postgres.Column) postgres.Projection {
	return col.AS(col.Name())
}
