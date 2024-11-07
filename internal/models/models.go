package models

type DoorLock struct {
	PartNumber           string   `json:"part_number"`
	Title                string   `json:"title"`
	Price                float32  `json:"price"`
	SalePrice            float32  `json:"sale_price"`
	Equipment            string   `json:"equipment"`
	Colors               []string `json:"colors"`
	Description          string   `json:"description"`
	Category             string   `json:"category"`
	CardMemory           int      `json:"card_memory"`
	Material             string   `json:"material"`
	HasMobileApplication bool     `json:"has_mobile_application"`
	PowerSupply          string   `json:"power_supply"`
	Size                 string   `json:"size"`
	Weight               int      `json:"weight"`
	DoorType             []string `json:"door_type"`
	DoorThicknessMin     int      `json:"door_thickness_min"`
	DoorThicknessMax     int      `json:"door_thickness_max"`
	Rating               float32  `json:"rating"`
	Quantity             int      `json:"quantity"`
}
