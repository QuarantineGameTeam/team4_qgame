package methodsAndStructs

func (f *Field) AddMine(mine Mine) {
	f.Mines = append(f.Mines, mine)
}

func (f *Field) AddClan(clan Clan) {
	f.Clans = append(f.Clans, clan)
}
