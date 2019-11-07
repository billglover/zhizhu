package feed

import (
	"fmt"
	"net/http"

	"github.com/billglover/zhizhu/model"
	"github.com/billglover/zhizhu/storage"
	"github.com/mmcdole/gofeed"
)

// Manager provides the ability to create, update and retrieve feeds
// from the underlying database.
type Manager struct {
	store *storage.Storage
}

// NewManager returns a Feed Manager configured to use the provided
// underlying data store.
func NewManager(store *storage.Storage) *Manager {
	return &Manager{store: store}
}

// Create takes a URL and fetches the feed, parses the information and
// stores it in the database.
func (m *Manager) Create(url string) (model.FeedInfo, error) {

	fi := model.FeedInfo{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fi, fmt.Errorf("feed.Manager.Create: %w", err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)

	fp := gofeed.NewParser()
	f, err := fp.Parse(resp.Body)
	if err != nil {
		return fi, fmt.Errorf("feed.Manager.Create: %w", err)
	}
	resp.Body.Close()

	fi = model.FeedInfo{
		Title:       f.Title,
		Link:        f.Link,
		Description: f.Description,
		Language:    f.Language,
		Copyright:   f.Copyright,
	}

	err = m.store.CreateFeed(&fi)
	if err != nil {
		return fi, fmt.Errorf("feed.Manager.Create: %w", err)
	}

	for _, i := range f.Items {
		item := model.Item{
			Title:       i.Title,
			Link:        i.Link,
			Description: i.Description,
			GUID:        i.GUID,
			Content:     i.Content,
		}

		err = m.store.CreateItem(fi.ID, &item)
		if err != nil {
			return fi, fmt.Errorf("feed.Manager.Create: %w", err)
		}
	}

	return fi, nil
}

// Refresh takes a feed ID, fetches the feed, parses the information
// and updates the database.
func (m *Manager) Refresh(id string) (model.FeedInfo, error) {
	fi := model.FeedInfo{}
	return fi, nil
}

// Delete takes a feed ID and removes it from the database.
func (m *Manager) Delete(id string) error {
	return nil
}

// List returns a slice of all feeds from the database.
func (m *Manager) List() ([]model.FeedInfo, error) {
	return nil, nil
}

// Get takes a feed UUID and returns the feed from the database.
func (m *Manager) Get() (model.FeedInfo, error) {
	fi := model.FeedInfo{}
	return fi, nil
}
