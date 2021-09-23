package hahamut

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

func doFileUploadRequest(uri, path string, file io.Reader) ([]byte, error) {
	// prepare request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("filedata", filepath.Base(path))
	if err != nil {
		return []byte{}, err
	}
	if _, err = io.Copy(part, file); err != nil {
		return []byte{}, err
	}
	if err = writer.Close(); err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(http.MethodPost, uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	// read response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	return bodyBytes, err
}

func doPostRequest(url string, body []byte) (string, error) {
	// prepare request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	header := http.Header{}
	header.Add("Content-Type", "application/json")
	req.Header = header

	// do request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read response body
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}
