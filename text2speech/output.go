package text2speech

// type Output struct {
// 	RequestID string `json:"request_id"`
// 	AudioURL string `json:"audio_menv_url"`
// 	Catalog string `json:"catalog_name"`
// 	SubCatalog string `json:"sub_catalog_name"`
// 	EngineVersion string `json:"engin_version"`
// }

// {
//   "request_id": "-356438514",
//   "async": "https://s3-ap-southeast-1.amazonaws.com/fho-fti-tts-output/cache-356438514.male.0.mp3",
//   "message": "The content will be returned after a few minutes under the async link."
// }
type Output struct {
	RequestID string `json:"request_id"`
	AudioURL string `json:"async"`
	Message string `json:"message"`
}