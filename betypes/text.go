package betypes

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

//Start command data
const (
	StartCommand = "start"
	StartText    = "\"Slippy world\" is a game that reflects the real world." +
		" A world in which you have to fight for every minute of your life." +
		" Amazing adventures await you here. Here, you have to take part in big wars between clans," +
		" form alliances, become a real spy and that's not all ..." +
		" Try it yourself!"
	RegistrationIsSuccessfulText = "Your registration is successful!!"
	AlreadyRegisteredText        = "You already registered!!"
)

const (
	ReregisterCommand          = "reregister"
	ReReregisterSuccessfulText = "Re-register successful!!"
)

//Help command data
const (
	HelpCommand = "help"
	HelpText    = "\"Slippy world\" is a game that reflects the real world." +
		" A world in which you have to fight for every minute of your life." +
		" Amazing adventures await you here. Here, you have to take part in big wars between clans," +
		" form alliances, become a real spy and that's not all ..." +
		" Try it yourself!" +
		"List of main commands: " +
		"\n/start" +
		"\n/help"
)

const (
	StartANewGameCommand    = "startgame"
	StartANewGameText       = "The set for the game has started!!"
	EnrollmentInTheGame     = "Enrollment in the game is possible only in group chats!!"
	GameHasAlreadyBegun     = "Game has already begun!!"
	AlreadyInLine           = "You are already in line!!"
	UserIsNotRegisteredText = "User is not registered!!"
)

//Unclear command data
const (
	UnclearCommandText = "There is no such command!!"
)

const (
	RightArrow = "➡"
	UpArrow    = "⬆"
	LeftArrow  = "⬅"
	DownArrow  = "⬇"
)

var SelectMoveButton = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(RightArrow, RightArrow),
		tgbotapi.NewInlineKeyboardButtonData(UpArrow, UpArrow),
		tgbotapi.NewInlineKeyboardButtonData(LeftArrow, LeftArrow),
		tgbotapi.NewInlineKeyboardButtonData(DownArrow, DownArrow),
	),
)

const JoinToGameButtonData = "join_to_game"

var JoinButton = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Join to game! ", JoinToGameButtonData),
	),
)
