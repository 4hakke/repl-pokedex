package networkclient

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/4hakke/repl-pokedex/internal/cache"
)

func NewClient(cache *cache.Cache) *NetworkClient {
	return &NetworkClient{cache: cache}
}

type NetworkClient struct {
	cache *cache.Cache
}

func (client *NetworkClient) Get(url string, resultedObject any) error {
	cachedResult, ok := client.cache.Get(url)
	if ok {
		return parse(cachedResult, &resultedObject)
	}

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = parse(body, &resultedObject)
	if err == nil {
		client.cache.Add(url, body)
	}
	return err
}

func parse(payload []byte, result any) error {
	err := json.Unmarshal(payload, &result)
	if err != nil {
		return err
	}

	return nil
}
