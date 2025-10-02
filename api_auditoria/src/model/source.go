package model

type Source struct {
	Version       *string     `json:"version,omitempty"`
	Connector     *string     `json:"connector,omitempty"`
	Name          *string     `json:"name,omitempty"`
	TsMS          *int64      `json:"ts_ms,omitempty"`
	Snapshot      *string     `json:"snapshot,omitempty"`
	DB            *string     `json:"db,omitempty"`
	Sequence      interface{} `json:"sequence"`
	Schema        *string     `json:"schema,omitempty"`
	Table         *string     `json:"table,omitempty"`
	ChangeLsn     *string     `json:"change_lsn,omitempty"`
	CommitLsn     *string     `json:"commit_lsn,omitempty"`
	EventSerialNo *int64      `json:"event_serial_no,omitempty"`
}
