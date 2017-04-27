package openfpt

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
	"strings"
	"strconv"
	"time"

	"github.com/nguyendangminh/openfpt/text2speech"

	"github.com/pkg/errors"
)

type Application struct {
	token string
	text2speechConfig text2speech.Config
}

func (app *Application) ConfigText2Speech(config text2speech.Config) {
	app.text2speechConfig = config
}

func (app *Application) Text2Speech(text string) (*text2speech.Output, error) {
	return app.Text2SpeechWithConfig(text, app.text2speechConfig)
}

func (app *Application) Text2SpeechWithConfig(text string, config text2speech.Config) (*text2speech.Output, error) {
	body := strings.NewReader(text)
	req, err := http.NewRequest("POST", text2speech.Endpoint, body)
	if err != nil {
		return nil, errors.Wrap(err, "NewRequest failed")
	}
	req.Header.Set("api_key", app.token)
	req.Header.Set("voice", config.Voice)
	req.Header.Set("speed", strconv.Itoa(config.Speed))
	req.Header.Set("format", config.Format)
	if config.Catalog != "" {
		req.Header.Set("catalog_name", config.Catalog) 
	}
	if config.SubCatalog != "" {
		req.Header.Set("sub_catalog_name", config.SubCatalog)
	}

	var client = &http.Client{
		Timeout: time.Duration(TIMEOUT) * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Client.Do failed")
	}
	defer resp.Body.Close()
	
	// b, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, errors.New("failed to read body")
	// }
	// if resp.StatusCode != 200 {
	// 	return nil, errors.New(fmt.Sprintf("request to text2speech failed, status code = %d, body = %s", resp.StatusCode, string(b)))
	// }

	var output text2speech.Output
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Decode body failed, body = %+v\n", resp.Body))
	}

	return &output, nil
}