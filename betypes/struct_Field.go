package betypes

//Field - The location of each clan and mine is preserved in the field
type Field struct {
	Mines []Mine
	Clans []Clan
}

func (f *Field) AddMine(mine Mine) {
	f.Mines = append(f.Mines, mine)
}

func (f *Field) AddClan(clan Clan) {
	f.Clans = append(f.Clans, clan)
}

/*func (f *Field) ShowField(bot *tgbotapi.BotAPI) {
	field := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Join to game!", JoinToGameButtonData),
		),
	)
}*/
