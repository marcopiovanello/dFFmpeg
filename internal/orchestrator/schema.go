package orchestrator

import "github.com/hashicorp/go-memdb"

type FFmpegJob struct {
	Id               string
	Node             string
	OriginalFilePath string
}

var schema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"jobs": {
			Name: "jobs",
			Indexes: map[string]*memdb.IndexSchema{
				"id": {
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "Id"},
				},
				"node": {
					Name:    "node",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Node"},
				},
				"original_file_path": {
					Name:    "original_file_path",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "OriginalFilePath"},
				},
			},
		},
	},
}
