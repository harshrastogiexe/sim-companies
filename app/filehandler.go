package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/harshrastogiexe/simcompanies/types"
)

func ReadFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	log.Println(filename, "open for reading")
	return file, err
}

func UnMarshall[T any](reader io.Reader) (*T, error) {
	value := new(T)
	if err := json.NewDecoder(reader).Decode(value); err != nil {
		return nil, err
	}
	return value, nil
}

func ReadResourceFromFile(filename string) (*types.Resource, error) {
	file, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("failed to close file", filename)
		}
		log.Println(filename, "closed successfully")
	}()

	resource, err := UnMarshall[types.Resource](file)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func ReadResourcesFromDisk() (<-chan []types.Resource, error) {
	var (
		start           = time.Now()
		currentDir, _   = os.Getwd()
		ch, resourcesCh = make(chan *types.Resource, 15), make(chan []types.Resource)
	)

	for i := 1; i <= int(SIM_COMPANIES_MAX_RESOURCE_COUNT); i++ {
		go func(id int) {
			filename := filepath.Join(currentDir, "temp/encyclopedia", strconv.Itoa(id)+".json")
			resource, err := ReadResourceFromFile(filename)

			if err != nil {
				log.Println(fmt.Errorf("failed to read resource id [%d]: [%w]", id, err))
				ch <- nil
				return
			}

			ch <- resource
		}(i)
	}

	go func() {
		resources := []types.Resource{}

		for i := 1; i <= int(SIM_COMPANIES_MAX_RESOURCE_COUNT); i++ {
			if r := <-ch; r != nil {
				resources = append(resources, *r)
			}
		}
		resourcesCh <- resources
		close(resourcesCh)
	}()

	log.Println("Time elapsed", time.Since(start).Milliseconds(), "milliseconds")
	return resourcesCh, nil
}
