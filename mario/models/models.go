package models

type Player struct {
    Lives int
    stage int
    Inventory []string
}

// recv function to add a life
func (p *Player) Greenmushroom() {
    p.Lives++
}

// recv function to add an inventory item
func (p *Player) Pickup(powerup string) {
    p.Inventory = append(p.Inventory, powerup)
}

// rec function to check on the current stasge
func (p Player) CanWhistle() bool {
    return p.stage >= 5
    }
