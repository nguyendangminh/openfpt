package text2speech

const Endpoint string = "http://api.openfpt.vn/text2speech/v3"

// voice type
var Male string = "male"
var Female string = "female"

// speed of speech
var VeryVerySlow int = -3
var VerySlow int = -2
var Slow int = -1
var Normal int = 0
var Fast int = 1
var VeryFast int = 2
var VeryVeryFast int = 3

// audio file format
var WAV string = "wav"
var MP3 string = "mp3"

var DefaultConfig Config = Config{
	Voice: Male,
	Catalog: "",
	SubCatalog: "",
	Speed: Normal,
	Format: MP3,
}