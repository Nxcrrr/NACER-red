package utils

type Personnage struct {
	Name                      string
	Class                     string
	Level                     int
	MaxLife                   int
	Life                      int
	Inventory                 map[string]int
	Character                 bool
	InventoryExists           bool
	Skill                     []string
	Money                     int
	Equipment                 Equipment
	MaximumInventoryCapacity  int
	UpgradeInventorySlotCount int
}

type Equipment struct {
	Head string
	Body string
	Foot string
}

type Monster struct {
	Name         string
	MaxLife      int
	Life         int
	AttackPoints int
}
