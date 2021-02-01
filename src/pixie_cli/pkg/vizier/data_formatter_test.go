package vizier_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	public_vizierapipb "pixielabs.ai/pixielabs/src/api/public/vizierapipb"
	"pixielabs.ai/pixielabs/src/pixie_cli/pkg/vizier"
)

func TestBasic(t *testing.T) {
	relation := &public_vizierapipb.Relation{
		Columns: []*public_vizierapipb.Relation_ColumnInfo{
			&public_vizierapipb.Relation_ColumnInfo{
				ColumnName:         "bool",
				ColumnType:         public_vizierapipb.BOOLEAN,
				ColumnSemanticType: public_vizierapipb.ST_NONE,
			},
			&public_vizierapipb.Relation_ColumnInfo{
				ColumnName:         "int",
				ColumnType:         public_vizierapipb.INT64,
				ColumnSemanticType: public_vizierapipb.ST_NONE,
			},
			&public_vizierapipb.Relation_ColumnInfo{
				ColumnName:         "float",
				ColumnType:         public_vizierapipb.FLOAT64,
				ColumnSemanticType: public_vizierapipb.ST_NONE,
			},
			&public_vizierapipb.Relation_ColumnInfo{
				ColumnName:         "string",
				ColumnType:         public_vizierapipb.STRING,
				ColumnSemanticType: public_vizierapipb.ST_NONE,
			},
			&public_vizierapipb.Relation_ColumnInfo{
				ColumnName:         "time",
				ColumnType:         public_vizierapipb.TIME64NS,
				ColumnSemanticType: public_vizierapipb.ST_NONE,
			},
		},
	}

	formatter := vizier.NewDataFormatterForTable(relation)

	// bool
	assert.Equal(t, "true", formatter.FormatValue(0, true))
	assert.Equal(t, "false", formatter.FormatValue(0, false))

	// int
	assert.Equal(t, "-12", formatter.FormatValue(1, -12))
	assert.Equal(t, "1204", formatter.FormatValue(1, 1204))

	// float
	assert.Equal(t, "-0.1", formatter.FormatValue(2, -0.1))
	assert.Equal(t, "121214", formatter.FormatValue(2, 121214.04234))

	// string
	assert.Equal(t, "abc", formatter.FormatValue(3, "abc"))

	// time
	nowTime := time.Unix(0, 1601694759495000000)
	nowTimeStr := nowTime.String()
	assert.Equal(t, nowTimeStr, formatter.FormatValue(4, nowTime))
}

func TestDuration(t *testing.T) {
	relation := &public_vizierapipb.Relation{
		Columns: []*public_vizierapipb.Relation_ColumnInfo{
			&public_vizierapipb.Relation_ColumnInfo{
				ColumnName:         "bool",
				ColumnType:         public_vizierapipb.BOOLEAN,
				ColumnSemanticType: public_vizierapipb.ST_NONE,
			},
			&public_vizierapipb.Relation_ColumnInfo{
				ColumnName:         "duration",
				ColumnType:         public_vizierapipb.INT64,
				ColumnSemanticType: public_vizierapipb.ST_DURATION_NS,
			},
		},
	}

	formatter := vizier.NewDataFormatterForTable(relation)

	assert.Equal(t, "144 ns", formatter.FormatValue(1, int64(144)))
	assert.Equal(t, "5.1 µs", formatter.FormatValue(1, int64(5144)))
	assert.Equal(t, "5.0 ms", formatter.FormatValue(1, int64(5*1000*1000)))
	assert.Equal(t, "13.0 s", formatter.FormatValue(1, int64(13*1000*1000*1000+1242)))
	assert.Equal(t, "300 s", formatter.FormatValue(1, int64(5*60*1000*1000*1000+1334)))
	assert.Equal(t, "43200 s", formatter.FormatValue(1, int64(12*60*60*1000*1000*1000+1335144)))
	assert.Equal(t, "2160000 s", formatter.FormatValue(1, int64(25*24*60*60*1000*1000*1000+133514124)))
}

func TestBytes(t *testing.T) {
	relation := &public_vizierapipb.Relation{
		Columns: []*public_vizierapipb.Relation_ColumnInfo{
			&public_vizierapipb.Relation_ColumnInfo{
				ColumnName:         "bytes",
				ColumnType:         public_vizierapipb.INT64,
				ColumnSemanticType: public_vizierapipb.ST_BYTES,
			},
		},
	}

	formatter := vizier.NewDataFormatterForTable(relation)

	assert.Equal(t, "144 B", formatter.FormatValue(0, int64(144)))
	assert.Equal(t, "5.0 KiB", formatter.FormatValue(0, int64(5144)))
	assert.Equal(t, "4.8 MiB", formatter.FormatValue(0, int64(5*1000*1000)))
	assert.Equal(t, "12.1 GiB", formatter.FormatValue(0, int64(13*1000*1000*1000+1242)))
	assert.Equal(t, "39.3 TiB", formatter.FormatValue(0, int64(12*60*60*1000*1000*1000+1335144)))
}

func TestThroughput(t *testing.T) {
	relation := &public_vizierapipb.Relation{
		Columns: []*public_vizierapipb.Relation_ColumnInfo{
			&public_vizierapipb.Relation_ColumnInfo{
				ColumnName:         "throughput",
				ColumnType:         public_vizierapipb.FLOAT64,
				ColumnSemanticType: public_vizierapipb.ST_THROUGHPUT_PER_NS,
			},
			&public_vizierapipb.Relation_ColumnInfo{
				ColumnName:         "throughput_bytes",
				ColumnType:         public_vizierapipb.FLOAT64,
				ColumnSemanticType: public_vizierapipb.ST_THROUGHPUT_BYTES_PER_NS,
			},
		},
	}

	formatter := vizier.NewDataFormatterForTable(relation)

	assert.Equal(t, "-0.002342 /sec", formatter.FormatValue(0, float64(-0.0000000000023423424)))
	assert.Equal(t, "1243924924.000000 /sec", formatter.FormatValue(0, float64(1.243924924)))

	assert.Equal(t, "-229 KiB/sec", formatter.FormatValue(1, float64(-0.00023423424)))
	assert.Equal(t, "1.2 GiB/sec", formatter.FormatValue(1, float64(1.243924924)))
}

func TestPercent(t *testing.T) {
	relation := &public_vizierapipb.Relation{
		Columns: []*public_vizierapipb.Relation_ColumnInfo{
			&public_vizierapipb.Relation_ColumnInfo{
				ColumnName:         "percent",
				ColumnType:         public_vizierapipb.FLOAT64,
				ColumnSemanticType: public_vizierapipb.ST_PERCENT,
			},
		},
	}

	formatter := vizier.NewDataFormatterForTable(relation)

	assert.Equal(t, "-23.44 %", formatter.FormatValue(0, float64(-0.2344)))
	assert.Equal(t, "1223.40 %", formatter.FormatValue(0, float64(12.234)))
	assert.Equal(t, "14.00 %", formatter.FormatValue(0, float64(0.140000)))
}

func TestQuantiles(t *testing.T) {
	relation := &public_vizierapipb.Relation{
		Columns: []*public_vizierapipb.Relation_ColumnInfo{
			&public_vizierapipb.Relation_ColumnInfo{
				ColumnName:         "quantiles",
				ColumnType:         public_vizierapipb.STRING,
				ColumnSemanticType: public_vizierapipb.ST_QUANTILES,
			},
			&public_vizierapipb.Relation_ColumnInfo{
				ColumnName:         "quantiles_duration_ns",
				ColumnType:         public_vizierapipb.STRING,
				ColumnSemanticType: public_vizierapipb.ST_DURATION_NS_QUANTILES,
			},
		},
	}

	formatter := vizier.NewDataFormatterForTable(relation)

	rate1, err := json.Marshal(map[string]interface{}{
		"p50": float64(-1231),
		"p99": float64(0.000023423),
	})
	assert.Nil(t, err)
	rate2, err := json.Marshal(map[string]interface{}{
		"p50": float64(9234234.3),
		"p99": float64(42398243.001),
	})
	assert.Nil(t, err)

	assert.Equal(t, "p50: -1231, p99: 2.3423e-05", formatter.FormatValue(0, string(rate1)))
	assert.Equal(t, "p50: 9.23423e+06, p99: 4.23982e+07", formatter.FormatValue(0, string(rate2)))

	duration1, err := json.Marshal(map[string]interface{}{
		"p50": float64(-12313),
		"p99": float64(0.0001),
	})
	assert.Nil(t, err)
	duration2, err := json.Marshal(map[string]interface{}{
		"p50": float64(23409234),
		"p99": float64(234092234234),
	})
	assert.Nil(t, err)

	assert.Equal(t, "p50: -12.3 µs, p99: 0 s", formatter.FormatValue(1, string(duration1)))
	assert.Equal(t, "p50: 23.4 ms, p99: 234 s", formatter.FormatValue(1, string(duration2)))
}