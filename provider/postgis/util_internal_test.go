package postgis

import (
	"testing"

	"github.com/terranodo/tegola"
)

func TestReplaceTokens(t *testing.T) {
	testcases := []struct {
		layer    Layer
		tile     *tegola.Tile
		expected string
	}{
		{
			layer: Layer{
				sql:  "SELECT * FROM foo WHERE geom && !BBOX!",
				srid: tegola.WebMercator,
			},
			tile:     tegola.NewTile(2, 1, 1),
			expected: "SELECT * FROM foo WHERE geom && ST_MakeEnvelope(-1.0097025686953126e+07,1.0097025686953126e+07,78271.5169531256,-78271.5169531256,3857)",
		},
		{
			layer: Layer{
				sql:  "SELECT id, scalerank=!ZOOM! FROM foo WHERE geom && !BBOX!",
				srid: tegola.WebMercator,
			},
			tile:     tegola.NewTile(2, 1, 1),
			expected: "SELECT id, scalerank=2 FROM foo WHERE geom && ST_MakeEnvelope(-1.0097025686953126e+07,1.0097025686953126e+07,78271.5169531256,-78271.5169531256,3857)",
		},
	}

	for i, tc := range testcases {
		sql, err := replaceTokens(&tc.layer, tc.tile)
		if err != nil {
			t.Errorf("Failed test %v. err: %v", i, err)
			return
		}

		if sql != tc.expected {
			t.Errorf("Failed test %v. Expected (%v), got (%v)", i, tc.expected, sql)
			return
		}
	}
}
