package entity

import (
	"github.com/photoprism/photoprism/pkg/s2"
)

type CellMap map[string]Cell

func (m CellMap) Get(name string) Cell {
	if result, ok := m[name]; ok {
		return result
	}

	return UnknownLocation
}

func (m CellMap) Pointer(name string) *Cell {
	if result, ok := m[name]; ok {
		return &result
	}

	return &UnknownLocation
}

var CellFixtures = CellMap{
	"mexico": {
		ID:           s2.TokenPrefix + "85d1ea7d382c",
		PlaceID:      PlaceFixtures.Get("mexico").ID,
		CellName:     "Adosada Platform",
		CellCategory: "botanical garden",
		Place:        PlaceFixtures.Pointer("mexico"),
		CreatedAt:    Timestamp(),
		UpdatedAt:    Timestamp(),
	},
	"caravan park": {
		ID:      s2.TokenPrefix + "1ef75a71a36c",
		PlaceID: s2.TokenPrefix + "1ef75a71a36c",
		Place: &Place{
			ID:           s2.TokenPrefix + "1ef75a71a36",
			PlaceLabel:   "Mandeni, KwaZulu-Natal, South Africa",
			PlaceCity:    "Mandeni",
			PlaceState:   "KwaZulu-Natal",
			PlaceCountry: "za",
			CreatedAt:    Timestamp(),
			UpdatedAt:    Timestamp(),
		},
		CellName:     "Lobotes Caravan Park",
		CellCategory: "camping",
		CreatedAt:    Timestamp(),
		UpdatedAt:    Timestamp(),
	},
	"zinkwazi": {
		ID:           s2.TokenPrefix + "1ef744d1e28c",
		PlaceID:      PlaceFixtures.Get("zinkwazi").ID,
		Place:        PlaceFixtures.Pointer("zinkwazi"),
		CellName:     "Zinkwazi Beach",
		CellCategory: "beach",
		CreatedAt:    Timestamp(),
		UpdatedAt:    Timestamp(),
	},
	"hassloch": {
		ID:           s2.TokenPrefix + "1ef744d1e280",
		PlaceID:      PlaceFixtures.Get("holidaypark").ID,
		Place:        PlaceFixtures.Pointer("holidaypark"),
		CellName:     "Holiday Park",
		CellCategory: "park",
		CreatedAt:    Timestamp(),
		UpdatedAt:    Timestamp(),
	},
	"emptyNameLongCity": {
		ID:           s2.TokenPrefix + "1ef744d1e281",
		PlaceID:      PlaceFixtures.Get("emptyNameLongCity").ID,
		Place:        PlaceFixtures.Pointer("emptyNameLongCity"),
		CellName:     "",
		CellCategory: "botanical garden",
		CreatedAt:    Timestamp(),
		UpdatedAt:    Timestamp(),
	},
	"emptyNameShortCity": {
		ID:           s2.TokenPrefix + "1ef744d1e282",
		PlaceID:      PlaceFixtures.Get("emptyNameShortCity").ID,
		Place:        PlaceFixtures.Pointer("emptyNameShortCity"),
		CellName:     "",
		CellCategory: "botanical garden",
		CreatedAt:    Timestamp(),
		UpdatedAt:    Timestamp(),
	},
	"veryLongLocName": {
		ID:           s2.TokenPrefix + "1ef744d1e283",
		PlaceID:      PlaceFixtures.Get("veryLongLocName").ID,
		Place:        PlaceFixtures.Pointer("veryLongLocName"),
		CellName:     "longlonglonglonglonglonglonglonglonglonglonglonglongName",
		CellCategory: "cape",
		CreatedAt:    Timestamp(),
		UpdatedAt:    Timestamp(),
	},
	"mediumLongLocName": {
		ID:           s2.TokenPrefix + "1ef744d1e283",
		PlaceID:      PlaceFixtures.Get("mediumLongLocName").ID,
		Place:        PlaceFixtures.Pointer("mediumLongLocName"),
		CellName:     "longlonglonglonglonglongName",
		CellCategory: "botanical garden",
		CreatedAt:    Timestamp(),
		UpdatedAt:    Timestamp(),
	},
	"Neckarbrücke": {
		ID:           s2.TokenPrefix + "1ef744d1e284",
		PlaceID:      PlaceFixtures.Get("Germany").ID,
		Place:        PlaceFixtures.Pointer("Germany"),
		CellName:     "Neckarbrücke",
		CellCategory: "",
		CreatedAt:    Timestamp(),
		UpdatedAt:    Timestamp(),
	},
}

// CreateCellFixtures inserts known entities into the database for testing.
func CreateCellFixtures() {
	for _, entity := range CellFixtures {
		Db().Create(&entity)
	}
}
