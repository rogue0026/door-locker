package models

type Lock struct {
	PartNumber           string   `json:"part_number"`
	Title                string   `json:"title"`
	Image                []string `json:"image"`
	Price                int      `json:"price"`
	SalePrice            int      `json:"sale_price"`
	Equipment            string   `json:"equipment"`
	Colors               []string `json:"colors"`
	Description          string   `json:"description"`
	Category             string   `json:"category"`
	CardMemory           int      `json:"card_memory"`
	Material             []string `json:"material"`
	HasMobileApplication bool     `json:"has_mobile_application"`
	PowerSupply          int      `json:"power_supply"`
	Size                 string   `json:"size"`
	Weight               int      `json:"weight"`
	DoorType             []string `json:"door_type"`
	DoorThicknessMin     int      `json:"door_thickness_min"`
	DoorThicknessMax     int      `json:"door_thickness_max"`
	Rating               int      `json:"rating"`
	Quantity             int      `json:"quantity"`
}
