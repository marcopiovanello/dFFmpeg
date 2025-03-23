package processor

type QualityPreset struct {
	Quality int `json:"quality"` // -q paramenter
	CRF     int `json:"crc"`     // crf
	Preset  int `json:"preset"`
}
