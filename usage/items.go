package usage

import cache "genericCache"

type Item struct {
	ID    string
	Value string
}

type ItemsService struct {
	cache *cache.Cache[Item]
}

func NewItemsService() *ItemsService {
	return &ItemsService{
		cache: cache.NewCache[Item](),
	}
}

func (s *ItemsService) AddItem(item Item) {
	s.cache.Set(item.ID, item)
}

func (s *ItemsService) GetItem(id string) Item {
	return s.cache.Get(id)
}
