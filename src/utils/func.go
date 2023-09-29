package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func (p *Personnage) DisplayInfo() {
	fmt.Println("--------------------INFORMATIONS DU PERSONNAGE--------------------")
	fmt.Println("")
	fmt.Println("Name :", p.Name)
	fmt.Println("Classe :", p.Class)
	fmt.Println("Niveau :", p.Level)
	fmt.Println("Points de vie maximum :", p.MaxLife)
	fmt.Println("Points de vie actuels :", p.Life)
}

func (p *Personnage) AccessInventory() {
	fmt.Println("--------------------CONTENU DE L'Inventory--------------------")
	fmt.Println("")
	fmt.Println("Inventory :")
	for item := range p.Inventory {
		fmt.Println("  ", item, " : ", p.Inventory[item])
	}
}

func Exit() {
	fmt.Println("--------------------QUITTER--------------------")
	fmt.Println("")
	fmt.Println("Au revoir ...")
	os.Exit(1)
}

func ClearConsole() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported platform")
	}
}

func Inputint() (int, error) {
	fmt.Print(">> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	chiffre, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return chiffre, nil
}

func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func (p *Personnage) Menu() {
	ClearConsole()
	fmt.Println("--------------------MENU PRINCIPAL--------------------")
	fmt.Println("")
	fmt.Println("[1] - Afficher les informations du personnage")
	fmt.Println("[2] - Accéder au contenu de l'Inventory")
	if !p.Character {
		fmt.Println("[3] - Initialiser mon personnage")
		fmt.Println("[4] - Marchand")
		fmt.Println("[5] - Forgeron")
		fmt.Println("[6] - Équipementier")
		fmt.Println("[7] - Quitter")

		choice, err := Inputint()
		if err != nil {
			fmt.Println("Erreur de saisie")
			os.Exit(1)
		}

		switch choice {
		case 1:
			ClearConsole()
			fmt.Println("--------------------INFORMATIONS DU PERSONNAGE--------------------")
			fmt.Println("")
			fmt.Println("Personnage non-créé")
			p.ReturnMenu()
		case 2:
			ClearConsole()
			fmt.Println("--------------------Inventory DU PERSONNAGE--------------------")
			fmt.Println("")
			fmt.Println("Personnage non-créé")
			p.ReturnMenu()
		case 3:
			ClearConsole()
			fmt.Println("--------------------INITIALISATION DU PERSONNAGE--------------------")
			fmt.Println("")

			p.charCreation()
			fmt.Println("Initialisation du personnage ...")
			time.Sleep(3 * time.Second)
			fmt.Println("Personnage initialisé")
			fmt.Println("")
			p.ReturnMenu()
		case 4:
			ClearConsole()
			fmt.Println("--------------------MARCHAND--------------------")
			fmt.Println("")
			fmt.Println("Commerce fermé (personnage non créé)")
			p.ReturnMenu()
		case 5:
			ClearConsole()
			fmt.Println("--------------------FORGERON--------------------")
			fmt.Println("")
			fmt.Println("Fourneau fermé (personnage non créé)")
			p.ReturnMenu()
		case 6:
			ClearConsole()
			fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
			fmt.Println("")
			fmt.Println("Magasin fermé (personnage non créé)")
			p.ReturnMenu()
		case 7:
			ClearConsole()
			Exit()
		}
	} else {
		fmt.Println("[3] - Marchand")
		fmt.Println("[4] - Forgeron")
		fmt.Println("[5] - Équipementier")
		fmt.Println("[6] - Entraînement")
		fmt.Println("[7] - Quitter")

		choice, err := Inputint()
		if err != nil {
			fmt.Println("Erreur de saisie")
			os.Exit(1)
		}

		switch choice {
		case 1:
			ClearConsole()
			fmt.Println("--------------------INFORMATIONS DU PERSONNAGE--------------------")
			fmt.Println("")
			p.DisplayInfo()
			p.ReturnMenu()
		case 2:
			ClearConsole()
			if !p.InventoryExists {
				fmt.Println("--------------------Inventory DU PERSONNAGE--------------------")
				fmt.Println("")
				fmt.Println("Inventory vide")
			} else if p.InventoryExists {
				p.AccessInventory()
			}
			p.ReturnMenu()
		case 3:
			ClearConsole()
			fmt.Println("--------------------MARCHAND--------------------")
			fmt.Println("")
			p.Marchand()
			p.ReturnMenu()
		case 4:
			ClearConsole()
			fmt.Println("--------------------FORGERON--------------------")
			fmt.Println("")
			p.Forgeron()
			p.ReturnMenu()
		case 5:
			ClearConsole()
			fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
			fmt.Println("")
			p.EquipmentManufacturer()
			p.ReturnMenu()
		case 6:
			ClearConsole()
			fmt.Println("--------------------ENTRAÎNEMENT--------------------")
			fmt.Println("")
			p.trainingFight(Monster{})
			p.ReturnMenu()
		case 7:
			ClearConsole()
			Exit()
		}
	}
}

func (p *Personnage) ReturnMenu() {
	fmt.Println("")
	fmt.Println("[1] - Retourner au menu")
	fmt.Println("[2] - Quitter")

	choice, err := Inputint()
	if err != nil {
		fmt.Println("Erreur de saisie")
		os.Exit(1)
	}

	switch choice {
	case 1:
		p.Menu()
	case 2:
		fmt.Println("Au revoir ...")
		os.Exit(1)
	}
}

func (p *Personnage) Init(name string, class string, level int, maxLife int, life int, inventory map[string]int, character bool, inventoryExists bool, skill []string, money int, maximumInventoryCapacity int, upgradeInventorySlotCount int) {
	p.Name = name
	p.Class = class
	p.Level = level
	p.MaxLife = maxLife
	p.Life = life
	p.Inventory = inventory
	p.Character = character
	p.InventoryExists = inventoryExists
	p.Skill = skill
	p.Money = money
	p.MaximumInventoryCapacity = maximumInventoryCapacity
	p.UpgradeInventorySlotCount = upgradeInventorySlotCount
}

func (p *Personnage) charCreation() {
	fmt.Println("Choisissez un Name :")
	name := Input()

	var class string
	var pvMax int

	fmt.Println("Choisissez votre classe :")
	fmt.Println("[1] - Humain")
	fmt.Println("[2] - Elfe")
	fmt.Println("[3] - Nain")

	choice, err := Inputint()
	if err != nil {
		fmt.Println("Erreur de saisie")
		os.Exit(1)
	}

	switch choice {
	case 1:
		class = "Humain"
		pvMax = 100
	case 2:
		class = "Elfe"
		pvMax = 80
	case 3:
		class = "Nain"
		pvMax = 120
	}

	p.Init(name, class, 1, pvMax, pvMax/2, map[string]int{"Potions de vie :": 3}, true, true, []string{"Coup de poing"}, 100, 10, 0)
}

func (p *Personnage) TakePot() {
	ClearConsole()
	if p.Inventory["Potions de vie"] < 1 {
		fmt.Println("--------------------POTION DE VIE--------------------")
		fmt.Println("")
		fmt.Println("Vous n'avez pas assez de potion dans votre Inventory")
	} else {
		fmt.Println("--------------------POTION DE VIE--------------------")
		fmt.Println("")

		p.removeInventory("Potions de vie", 1)
		if p.Inventory["Potions de vie"] == 0 {
			delete(p.Inventory, "Potions de vie")
		}

		p.addLife(50)
		if p.Life > p.MaxLife {
			p.MaxLife = p.Life
		}

		fmt.Println("Vous avez bu une potion de vie")
		fmt.Println("")
		fmt.Println("Points de vie actuels :", p.Life)
		fmt.Println("Points de vie maximum :", p.MaxLife)
	}
}

func (p *Personnage) Marchand() {
	fmt.Println("--------------------MARCHAND--------------------")
	fmt.Println("")
	fmt.Println("[1] - Potion de vie (x1)")
	fmt.Println("[2] - Potion de poison (x1)")
	fmt.Println("[3] - Livre de sort : Boule de feu (x1)")
	fmt.Println("[4] - Fourrure de Loup (x1)")
	fmt.Println("[5] - Peau de Troll (x1)")
	fmt.Println("[6] - Cuir de Sanglier (x1)")
	fmt.Println("[7] - Plume de Corbeau (x1)")
	fmt.Println("[8] - Augmentation d'Inventory (x10)")

	choice, err := Inputint()
	if err != nil {
		fmt.Println("Erreur de saisie")
		os.Exit(1)
	}

	switch choice {
	case 1:
		ClearConsole()
		fmt.Println("--------------------POTION DE VIE--------------------")
		fmt.Println("")

		p.addInventory("Potions de vie", 1)
		potions := p.Inventory["Potions de vie"]

		p.removeMoney(3)

		fmt.Println("Vous avez acheté une potion de vie")
		fmt.Println("")
		fmt.Println("Vous avez maintenant", potions, "potions de vie")
	case 2:
		ClearConsole()
		fmt.Println("--------------------POTION DE POISON--------------------")
		fmt.Println("")

		p.addInventory("Potions de poison", 1)
		potions := p.Inventory["Potions de poison"]

		p.removeMoney(6)

		fmt.Println("Vous avez acheté une potion de poison")
		fmt.Println("")
		fmt.Println("Vous avez maintenant", potions, "potions de poison")
	case 3:
		ClearConsole()
		fmt.Println("--------------------LIVRE DE SORT--------------------")
		fmt.Println("")

		p.SpellBlock("Boule de feu")

		p.removeMoney(25)

		fmt.Println("Vous avez maintenant le sort Boule de feu")
	case 4:
		ClearConsole()
		fmt.Println("--------------------FOURRURE DE LOUP--------------------")
		fmt.Println("")

		p.addInventory("Fourrure de Loup", 1)
		fourrures := p.Inventory["Fourrure de Loup"]

		p.removeMoney(4)

		fmt.Println("Vous avez acheté une fourrure de loup")
		fmt.Println("")
		fmt.Println("Vous avez maintenant", fourrures, "fourrure de Loup")
	case 5:
		ClearConsole()
		fmt.Println("--------------------PEAU DE TROLL--------------------")
		fmt.Println("")

		p.addInventory("Peau de Troll", 1)
		peaux := p.Inventory["Peau de Troll"]

		p.removeMoney(7)

		fmt.Println("Vous avez acheté une peau de troll")
		fmt.Println("")
		fmt.Println("Vous avez maintenant", peaux, "peau de troll")
	case 6:
		ClearConsole()
		fmt.Println("--------------------CUIR DE SANGLIER--------------------")
		fmt.Println("")

		p.addInventory("Cuir de Sanglier", 1)
		cuirs := p.Inventory["Cuir de Sanglier"]

		p.removeMoney(3)

		fmt.Println("Vous avez acheté un cuir de sanglier")
		fmt.Println("")
		fmt.Println("Vous avez maintenant", cuirs, "cuir de sanglier")
	case 7:
		ClearConsole()
		fmt.Println("--------------------PLUME DE CORBEAU--------------------")
		fmt.Println("")

		p.addInventory("Plume de Corbeau", 1)
		plumes := p.Inventory["Plume de Corbeau"]

		p.removeMoney(1)

		fmt.Println("Vous avez acheté une plume de corbeau")
		fmt.Println("")
		fmt.Println("Vous avez maintenant", plumes, "plume de corbeau")
	case 8:
		ClearConsole()
		fmt.Println("--------------------AUGMENTATION DE L'Inventory--------------------")
		fmt.Println("")
		fmt.Println("[1] - Augmentation de la capacité maximale de l'Inventory (30 pièces d'or)")

		C, ERR := Inputint()
		if ERR != nil {
			fmt.Println("Erreur de saisie")
			os.Exit(1)
		}

		switch C {
		case 1:
			ClearConsole()
			fmt.Println("--------------------AUGMENTATION DE L'Inventory--------------------")
			fmt.Println("")

			p.UpgradeInventorySlot()
			fmt.Println("La capacité maximale de votre Inventory a été augmentée de 10")
		}
	}
	p.ReturnMenu()
}

func (p *Personnage) Forgeron() {
	fmt.Println("--------------------FORGERON--------------------")
	fmt.Println("")
	fmt.Println("[1] - Chapeau de l'aventurier (x1)")
	fmt.Println("[2] - Tunique de l'aventurier (x1)")
	fmt.Println("[3] - Bottes de l'aventurier (x1)")

	choice, err := Inputint()
	if err != nil {
		fmt.Println("Erreur de saisie")
		os.Exit(1)
	}

	switch choice {
	case 1:
		if p.Inventory["Plume de Corbeau"] <= 0 || p.Inventory["Cuir de Sanglier"] <= 0 {
			fmt.Println("Vous n'avez pas les ressources suffisantes. Rappel : 1 Plume de Corbeau & 1 Cuir de Saglier nécessaires")
		} else if p.Inventory["Plume de Corbeau"] >= 1 && p.Inventory["Cuir de Sanglier"] >= 1 {
			AnimateText("Fabrication en cours ...")
			time.Sleep(2 * time.Second)

			p.removeMoney(5)
			p.addInventory("Chapeau de l'aventurier", 1)
			p.removeInventory("Plume de Corbeau", 1)
			p.removeInventory("Cuir de Sanglier", 1)

			fmt.Println("Vous avez fabriqué un chapeau de l'aventurier")
		}
		p.ReturnMenu()
	case 2:
		if p.Inventory["Fourrure de Loup"] <= 1 || p.Inventory["Peau de Troll"] <= 0 {
			fmt.Println("Vous n'avez pas les ressources suffisantes. Rappel : 2 Fourrure de Loup & 1 Peau de Troll nécessaires")
		} else if p.Inventory["Fourrure de Loup"] >= 2 && p.Inventory["Peau de Troll"] >= 1 {
			AnimateText("Fabrication en cours ...")
			time.Sleep(2 * time.Second)

			p.removeMoney(5)
			p.addInventory("Tunique de l'aventurier", 1)
			p.removeInventory("Fourrure de Loup", 2)
			p.removeInventory("Peau de Troll", 1)

			fmt.Println("Vous avez fabriqué une tunique de l'aventurier")
		}
		p.ReturnMenu()
	case 3:
		if p.Inventory["Fourrure de Loup"] <= 0 || p.Inventory["Cuir de Sanglier"] <= 0 {
			fmt.Println("Vous n'avez pas les ressources suffisantes. Rappel : 1 Fourrure de Loup & 1 Cuir de Saglier nécessaires")
		} else if p.Inventory["Fourrure de Loup"] >= 1 && p.Inventory["Cuir de Sanglier"] >= 1 {
			AnimateText("Fabrication en cours ...")
			time.Sleep(2 * time.Second)

			p.removeMoney(5)
			p.addInventory("Bottes de l'aventurier", 1)
			p.removeInventory("Fourrure de Loup", 1)
			p.removeInventory("Cuir de Sanglier", 1)

			fmt.Println("Vous avez fabriqué des bottes de l'aventurier")
		}
		p.ReturnMenu()
	}
}

func (p *Personnage) EquipmentManufacturer() {
	fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
	fmt.Println("")

	if p.Inventory["Chapeau de l'aventurier"] >= 1 {
		fmt.Println("[1] - Chapeau de l'aventurier (x1)")
	}
	if p.Inventory["Tunique de l'aventurier"] >= 1 {
		fmt.Println("[2] - Tunique de l'aventurier (x1)")
	}
	if p.Inventory["Bottes de l'aventurier"] >= 1 {
		fmt.Println("[3] - Bottes de l'aventurier (x1)")
	}

	choice, err := Inputint()
	if err != nil {
		fmt.Println("Erreur de saisie")
		os.Exit(1)
	}

	switch choice {
	case 1:
		ClearConsole()
		fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
		fmt.Println("")

		fmt.Println("[1] - S'équiper d'un chapeau de l'aventurier")
		fmt.Println("[2] - Se déséquiper d'un chapeau de l'aventurier")

		c, Err := Inputint()
		if Err != nil {
			fmt.Println("Erreur de saisie")
			os.Exit(1)
		}

		switch c {
		case 1:
			if p.Equipment.Head != "" {
				ClearConsole()
				fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
				fmt.Println("")

				p.Equipment.Head = "Chapeau de l'aventurier"
				p.addMaxLife(10)

				fmt.Println("Vous avez été équipé d'un chapeau de l'aventurier")
			} else if p.Equipment.Head == "Chapeau de l'aventurier" {
				ClearConsole()
				fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
				fmt.Println("")
				fmt.Println("Vous êtes déjà équipé d'un chapeau de l'aventurier")
			}
			p.ReturnMenu()
		case 2:
			if p.Equipment.Head == "Chapeau de l'aventurier" {
				ClearConsole()
				fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
				fmt.Println("")

				p.Equipment.Head = ""
				p.removeMaxLife(10)

				fmt.Println("Vous avez été déséquipé d'un chapeau de l'aventurier")
			} else if p.Equipment.Head == "" {
				ClearConsole()
				fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
				fmt.Println("")
				fmt.Println("Vous n'êtes équipé de rien")
			}
			p.ReturnMenu()
		}
	case 2:
		ClearConsole()
		fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
		fmt.Println("")

		fmt.Println("[1] - S'équiper d'une tunique de l'aventurier")
		fmt.Println("[2] - Se déséquiper d'une tunique de l'aventurier")

		c, Err := Inputint()
		if Err != nil {
			fmt.Println("Erreur de saisie")
			os.Exit(1)
		}

		switch c {
		case 1:
			if p.Equipment.Body != "" {
				ClearConsole()
				fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
				fmt.Println("")

				p.Equipment.Body = "Tunique de l'aventurier"
				p.addMaxLife(25)

				fmt.Println("Vous avez été équipé d'une tunique de l'aventurier")
			} else if p.Equipment.Body == "Tunique de l'aventurier" {
				ClearConsole()
				fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
				fmt.Println("")
				fmt.Println("Vous êtes déjà équipé d'une tunique de l'aventurier")
			}
			p.ReturnMenu()
		case 2:
			if p.Equipment.Body == "Tunique de l'aventurier" {
				ClearConsole()
				fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
				fmt.Println("")

				p.Equipment.Body = ""
				p.removeMaxLife(10)

				fmt.Println("Vous avez été déséquipé d'une tunique de l'aventurier")
			} else if p.Equipment.Body == "" {
				ClearConsole()
				fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
				fmt.Println("")
				fmt.Println("Vous n'êtes équipé de rien")
			}
			p.ReturnMenu()
		}
	case 3:
		ClearConsole()
		fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
		fmt.Println("")

		fmt.Println("[1] - S'équiper des bottes de l'aventurier")
		fmt.Println("[2] - Se déséquiper des bottes de l'aventurier")

		c, Err := Inputint()
		if Err != nil {
			fmt.Println("Erreur de saisie")
			os.Exit(1)
		}

		switch c {
		case 1:
			if p.Equipment.Foot != "" {
				ClearConsole()
				fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
				fmt.Println("")

				p.Equipment.Foot = "Bottes de l'aventurier"
				p.addMaxLife(25)

				fmt.Println("Vous avez été équipé des bottes de l'aventurier")
			} else if p.Equipment.Foot == "Bottes de l'aventurier" {
				ClearConsole()
				fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
				fmt.Println("")
				fmt.Println("Vous êtes déjà équipé des bottes de l'aventurier")
			}
			p.ReturnMenu()
		case 2:
			if p.Equipment.Foot == "Bottes de l'aventurier" {
				ClearConsole()
				fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
				fmt.Println("")

				p.Equipment.Foot = ""
				p.removeMaxLife(10)

				fmt.Println("Vous avez été déséquipé des bottes de l'aventurier")
			} else if p.Equipment.Foot == "" {
				ClearConsole()
				fmt.Println("--------------------ÉQUIPEMENTIER--------------------")
				fmt.Println("")
				fmt.Println("Vous n'êtes équipé de rien")
			}
			p.ReturnMenu()
		}
	}
}

func (p *Personnage) addInventory(key string, quantity int) {
	p.MaxInv()
	p.Inventory[key] += quantity
}

func (p *Personnage) addLife(quantity int) {
	p.Life += quantity
}

func (p *Personnage) addMaxLife(quantity int) {
	p.MaxLife += quantity
}

func (p *Personnage) addMoney(quantity int) {
	p.Money += quantity
}

func (p *Personnage) removeInventory(key string, quantity int) {
	p.Inventory[key] -= quantity
}

func (p *Personnage) removeLife(quantity int) {
	p.Life -= quantity
}

func (p *Personnage) removeMaxLife(quantity int) {
	p.MaxLife -= quantity
}

func (p *Personnage) removeMoney(quantity int) {
	p.Money -= quantity
}

func (p *Personnage) dead() {
	ClearConsole()
	if p.Life <= 0 {
		fmt.Println("--------------------DEATH CHECK--------------------")
		fmt.Println("")
		fmt.Println(p.Name, "est mort")
		fmt.Println("En attente ...")
		time.Sleep(3 * time.Second)
		p.Life = p.MaxLife / 2
		fmt.Println("")
		fmt.Println(p.Name, "est ressuscité avec", p.Life, "points de vie")
	}
}

func (p *Personnage) PoisonPot() {
	time.Sleep(1 * time.Second) // 1 seconde
	p.removeLife(10)
	time.Sleep(1 * time.Second) // 2 secondes
	p.removeLife(10)
	time.Sleep(1 * time.Second) // 3 secondes
	p.removeLife(10)
}

func AnimateText(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(150 * time.Millisecond)
	}
	fmt.Println("")
}

func (p *Personnage) SpellBlock(sort string) {
	ClearConsole()
	fmt.Println("--------------------SORT--------------------")
	fmt.Println("")
	for _, s := range p.Skill {
		if s == sort {
			fmt.Println("Vous avez déjà le sort", sort)
		}
	}

	p.Skill = append(p.Skill, sort)
	fmt.Println("Vous avez appris le sort", sort)
}

func (p *Personnage) MaxInv() bool {
	if len(p.Inventory) > 10 {
		fmt.Println("Votre Inventory est plein")
		return false
	} else {
		return true
	}
}

func (p *Personnage) UpgradeInventorySlot() {
	if p.UpgradeInventorySlotCount >= 3 {
		fmt.Println("--------------------AUGMENTATION D'Inventory--------------------")
		fmt.Println("")
		fmt.Println("Vous avez atteint votre quota maximum d'augmentation de la capacité de l'Inventory (3)")
	} else {
		p.MaximumInventoryCapacity += 10
		p.UpgradeInventorySlotCount += 1
	}
}

func (m *Monster) InitGoblin() {
	m.Name = "Gobelin d'entraînement"
	m.MaxLife = 40
	m.Life = m.MaxLife
	m.AttackPoints = 5
}

func (p *Personnage) trainingFight(monster Monster) {
	// Initialisation du tour de combat
	playerTurn := true

	fmt.Println("--------------------COMBAT D'ENTRAÎNEMENT--------------------")
	fmt.Println("")

	for {
		charTurn(p, &monster)
		goblinPattern(p, &monster)
		// Vérifier si le combat est terminé
		if p.Life <= 0 {
			fmt.Println("Vous avez été vaincu par le", monster.Name)
			break
		} else if monster.Life <= 0 {
			fmt.Printf("Vous avez vaincu le %s !\n", monster.Name)
			break
		}

		// Changer le tour
		playerTurn = !playerTurn
	}
}

func goblinPattern(player *Personnage, goblin *Monster) {
	// Initialisation des tours et des dégâts du Gobelin
	tour := 1
	goblinDommages := goblin.AttackPoints

	for {
		// Tour du Gobelin
		if tour%3 == 0 {
			// Le Gobelin inflige 200% de son attaque en dégâts tous les 3 tours
			dommagesInfliges := goblinDommages * 2
			player.Life -= dommagesInfliges
			fmt.Printf("%s inflige à %s %d de dégâts (Tour %d).\n", goblin.Name, player.Name, dommagesInfliges, tour)
		} else {
			// Le Gobelin inflige 100% de son attaque en dégâts
			player.Life -= goblinDommages
			fmt.Printf("%s inflige à %s %d de dégâts (Tour %d).\n", goblin.Name, player.Name, goblinDommages, tour)
		}

		// Affichage des points de vie actuels sur les points de vie max du joueur
		fmt.Printf("%s : %d / %d\n", player.Name, player.Life, player.MaxLife)

		// Vérification si le combat est terminé
		if player.Life <= 0 {
			fmt.Printf("%s a été vaincu par %s.\n", player.Name, goblin.Name)
			break
		}

		// Passage au tour suivant
		tour++
	}
}

func charTurn(player *Personnage, monster *Monster) {
	fmt.Println("--------------------TOUR DU PERSONNAGE--------------------")
	fmt.Println("")

	fmt.Println("[1] - Attaquer")
	fmt.Println("[2] - Inventaire")

	choice, err := Inputint()
	if err != nil {
		fmt.Println("Erreur de saisie")
		os.Exit(1)
	}

	switch choice {
	case 1:
		// Attaque de base
		damage := 5 // Dégâts de l'attaque de base (à adapter selon vos règles)
		monster.Life -= damage
		fmt.Printf("%s inflige %d dégâts à %s.\n", player.Name, damage, monster.Name)
	case 2:
		// Inventaire
		fmt.Println("--------------------INVENTAIRE--------------------")
		fmt.Println("")

		for item, quantity := range player.Inventory {
			fmt.Printf("[%s] - x%d\n", item, quantity)
		}

		fmt.Println("Choisissez un objet à utiliser :")
		itemChoice := Input()

		if quantity, ok := player.Inventory[itemChoice]; ok && quantity > 0 {
			// Utiliser l'objet et appliquer son effet (à implémenter)
			// Exemple : player.UseItem(itemChoice)

			// Mettez ici le code pour gérer l'effet de l'objet sur le personnage

			// Réduire la quantité de l'objet de l'inventaire
			player.Inventory[itemChoice]--

			fmt.Printf("%s utilise %s.\n", player.Name, itemChoice)
		} else {
			fmt.Println("Objet non valide ou quantité insuffisante.")
		}
	default:
		fmt.Println("Choix invalide. Réessayez.")
	}
}
