package app

import (
	"log"
	"showcase/app/parser"
	"showcase/models"
	"showcase/repository"
)

const (
	toysCategoryID    = 687
	apparelCategoryID = 3515
	sportCategoryID   = 16
)

func (s *Server) Import() {
	currentTrademarks, err := repository.TrademarkMap(s.DB)
	if err != nil {
		log.Fatal(err)
	}

	categoryNames := map[int]string{
		toysCategoryID:    "Игрушки",
		apparelCategoryID: "Одежда",
		sportCategoryID:   "Спорт",
	}

	for id, name := range categoryNames {
		if err := repository.UpsertCategory(s.DB, models.Category{ID: sportCategoryID, Name: name}); err != nil {
			log.Printf("can not upsert category: %v\n", err)
			continue
		}

		products, newTrademarks, err := parser.ParseCategory(id)
		if err != nil {
			log.Printf("can not read api response: %v\n", err)

			continue
		}

		if err := repository.UpsertNewTrademarks(s.DB, newTrademarks, currentTrademarks); err != nil {
			log.Printf("can not insert products: %v\n", err)

			continue
		}

		if err = repository.UpsertProducts(s.DB, products); err != nil {
			log.Printf("can not insert products: %v\n", err)

			continue
		}

		log.Printf("Category \"%v\" saved", name)
	}
}
