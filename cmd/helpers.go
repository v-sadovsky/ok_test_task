package main

import (
	"encoding/xml"
	"github.com/grokify/html-strip-tags-go"
	"github.com/v-sadovsky/ok_test_task/pkg/models"
	"os"
	"sync"
)

func (app *application) requestData(wp int, tm string) ([]byte, error) {
	shops, err := app.stores.Get(tm)
	if err != nil {
		return nil, err
	}

	data := app.handleData(shops, wp)

	return data, nil
}

func (app *application) handleData(shops map[int]*models.Shop, pool int) []byte {
	var wg sync.WaitGroup

	jobs := make(chan *models.Shop)
	chunks := make(chan []byte)

	// Start workers
	for i := 0; i < pool; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			for shop := range jobs {
				app.infoLog.Printf("Worker (id=%d) has been started...", i)

				for _, item := range shop.Offers {
					item.Description = strip.StripTags(item.Description)
				}

				chunk, err := xml.MarshalIndent(shop, "", "  ")
				if err != nil {
					app.errorLog.Printf("worker (id=%d, shop.ID=%d): %v", i, shop.ID, err)
					continue
				}

				chunks <- chunk
				app.infoLog.Printf("Worker (id=%d) has been finished.", i)
			}
		}(i)
	}

	// Provide jobs
	go func() {
		for _, shop := range shops {
			jobs <- shop
		}
		close(jobs)
	}()

	// Wait for finish all jobs
	go func() {
		wg.Wait()
		close(chunks)
	}()

	// Collect all the chunks together
	var data []byte
	for c := range chunks {
		c = append(c, '\n')
		data = append(data, c...)
	}

	return data
}

func provideFile(fName string, data []byte) error {
	f, err := os.OpenFile(fName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)

	return err
}
