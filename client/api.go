package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Client http request configuration
type Client struct {
	baseURL     string
	accessToken string
	httpClient  *http.Client
}

// NewClient new instance
func NewClient(accessToken, baseURL string) *Client {
	return &Client{
		baseURL:     baseURL,
		accessToken: accessToken,
		httpClient:  &http.Client{},
	}
}

// Folder represents a single Item
type Folder struct {
	Name   string `json:"name"`
	Parent Parent `json:"parent"`
}

// Parent json
type Parent struct {
	ID string `json:"id"`
}

func main() {

	url := "https://api.box.com/2.0/folders/0/items/"
	bearer := "Bearer " + "NvocCaINeboZfFLGXzzJGC5A3IURu8S5"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string([]byte(body)))
}

// CreateFolder cretaes new folder with the given name
func (c *Client) CreateFolder(name, id string) error {
	buf := bytes.Buffer{}
	folder := &Folder{
		Name: name,
		Parent: Parent{
			ID: id,
		},
	}
	err := json.NewEncoder(&buf).Encode(folder)

	if err != nil {
		return err
	}

	_, err = c.httpRequest("folders", "POST", "application/json", buf)

	if err != nil {
		return err
	}

	return nil
}

// GetFolderItems Gets all of the files, folders, or web links contained within a folder.
func (c *Client) GetFolderItems(id string) error {

	path := fmt.Sprintf("folders/%s/items", id)
	fmt.Println(path)
	// body, err := c.httpRequest(path, "GET", "", bytes.Buffer{})

	// if err != nil {
	// 	return err
	// }
	return nil
}

func (c *Client) httpRequest(endPoint, method, contentType string, body bytes.Buffer) (closer io.ReadCloser, err error) {

	req, err := http.NewRequest(method, c.baseURL+endPoint, &body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)

	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", contentType)
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(res.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", res.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", res.StatusCode, respBody.String())
	}

	return res.Body, nil

}
