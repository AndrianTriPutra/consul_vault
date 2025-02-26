package env

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"managenv/pkg/logger"
	"net/http"
	"time"
)

func ReadEnv_Consul(c Setting) (Config, error) {
	var conf Config

	// Request to Consul
	url := c.Host + c.Path
	resp, err := http.Get(url)
	if err != nil {
		errN := errors.New("[Get]" + fmt.Sprintf("fetching data:%s", err.Error()))
		return conf, errN
	}
	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errN := errors.New("[ReadAll]" + fmt.Sprintf("reading response:%s", err.Error()))
		return conf, errN
	}

	// Parsing JSON response
	var kvData []map[string]interface{}
	if err := json.Unmarshal(body, &kvData); err != nil {
		errN := errors.New("[Unmarshal]" + fmt.Sprintf("parsing JSON:%s", err.Error()))
		return conf, errN
	}

	// Base64-encoded value from Consul response
	encoded, ok := kvData[0]["Value"].(string)
	if !ok {
		errN := errors.New("[encoded]" + "Value not found in response")
		return conf, errN
	}

	// Decode Base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		errN := errors.New("[decoded]" + fmt.Sprintf("decoding base64:%s", err.Error()))
		return conf, errN
	}

	// Unmarshal JSON to struct
	if err := json.Unmarshal(decoded, &conf); err != nil {
		fmt.Println("Error parsing decoded JSON:", err)
		errN := errors.New("[decoded]" + fmt.Sprintf("parsing decoded JSON:%s", err.Error()))
		return conf, errN
	}

	//logger access
	logger.Load(conf.Internal.Logger)

	logger.Trace(" ============= [Secret from Consul] ============= ", "")
	logger.Trace(" ======= [internal] ======= ", "")
	logger.Trace("Logger       :", conf.Internal.Logger)
	logger.Trace("Core         :", conf.Internal.Core)

	logger.Trace(" ======= [interval] ======= ", "")
	conf.Interval.Environment *= time.Second
	conf.Interval.Schedulle *= time.Second
	logger.Trace("Environment  :", conf.Interval.Environment)
	logger.Trace("Schedulle    :", conf.Interval.Schedulle)

	logger.Trace(" =======  [database] ======= ", "")
	logger.Trace("Host         :", conf.Database.Host)
	logger.Trace("Port         :", conf.Database.Port)
	logger.Trace("Name         :", conf.Database.Name)
	logger.Trace("User         :", conf.Database.User)
	logger.Trace("Pass         :", conf.Database.Pass)

	logger.Trace(" =======   [rabbit]  ======= ", "")
	logger.Trace("Host         :", conf.Rabbit.Host)
	logger.Trace("Tag          :", conf.Rabbit.Tag)
	logger.Trace("Queue        :", conf.Rabbit.Que)
	logger.Trace("Key          :", conf.Rabbit.Key)

	return conf, nil
}
