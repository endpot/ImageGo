package corein

type WebHomeInp struct {
	Nsfw bool `in:"query" default:"false"`
}
