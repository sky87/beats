package status

import (
	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/metricbeat/schema"
	c "github.com/elastic/beats/metricbeat/schema/mapstrstr"
)

var (
	schema = s.Schema{
		"aborted": s.Object{
			"clients":  c.Int("Aborted_clients"),
			"connects": c.Int("Aborted_connects"),
		},
		"binlog": s.Object{
			"cache": s.Object{
				"disk_use": c.Int("Binlog_cache_disk_use"),
				"use":      c.Int("Binlog_cache_use"),
			},
		},
		"bytes": s.Object{
			"received": c.Int("Bytes_received"),
			"sent":     c.Int("Bytes_sent"),
		},
		"connections": c.Int("Connections"),
		"created": s.Object{
			"tmp": s.Object{
				"disk_tables": c.Int("Created_tmp_disk_tables"),
				"files":       c.Int("Created_tmp_files"),
				"tables":      c.Int("Created_tmp_tables"),
			},
		},
		"delayed": s.Object{
			"errors":         c.Int("Delayed_errors"),
			"insert_threads": c.Int("Delayed_insert_threads"),
			"writes":         c.Int("Delayed_writes"),
		},
		"flush_commands":       c.Int("Flush_commands"),
		"max_used_connections": c.Int("Max_used_connections"),
		"open": s.Object{
			"files":   c.Int("Open_files"),
			"streams": c.Int("Open_streams"),
			"tables":  c.Int("Open_tables"),
		},
		"opened_tables": c.Int("Opened_tables"),
	}
)

// Map data to MapStr of server stats variables: http://dev.mysql.com/doc/refman/5.7/en/server-status-variables.html
// This is only a subset of the available values
func eventMapping(status map[string]string) common.MapStr {
	source := map[string]interface{}{}
	for key, val := range status {
		source[key] = val
	}
	return schema.Apply(source)
}
