package text2speech

// Config represents a configuration set for Text2Speech service
type Config struct {
	IsAsync bool // the document seems not being up to date, i can't implement
	AsyncCallbackURL string
	Voice string // male or female
	Catalog string // don't know what is this? Me too!
	SubCatalog string
	Speed int
	Format string
}