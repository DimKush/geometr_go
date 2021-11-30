package unit

type Unit struct {
	Id   int64           `json:"id"`
	Name string          `json:"name"`
	Poly []byte `json:"geometry" sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}
