package internal

import (
	_ "fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

/*
telegram bot + strpe api
*/

type Bot struct {
	api_token     string
	username      string
	payment_token string
	chat_id       int64
}

func init() {
}
func (bot *Bot) ReadConfig() {
	viper.AddConfigPath("./tgbotpay") // reading ./tgbotpay/config.env
	err := viper.ReadInConfig()
	checkErr(err)
	bot.api_token = viper.Get("API_TOKEN").(string)
	bot.username = viper.Get("USERNAME").(string)
	bot.payment_token = viper.Get("STRIPE_TEST_PAYMENT_TOKEN").(string)
	bot.chat_id, _ = strconv.ParseInt(viper.Get("CHAT_ID").(string), 10, 64)

}
func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error %v", err)
	}
}
func handlePreCheckoutQuery(tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	pca := tgbotapi.PreCheckoutConfig{
		OK:                 true,
		PreCheckoutQueryID: update.PreCheckoutQuery.ID,
	}
	_, err := tgbot.Request(pca)
	checkErr(err)
}

func handleShippingQuery(tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	ship := tgbotapi.ShippingConfig{
		OK:              true,
		ShippingQueryID: update.ShippingQuery.ID,
		ShippingOptions: []tgbotapi.ShippingOption{
			{ID: "1", Title: "test shipping 1", Prices: []tgbotapi.LabeledPrice{
				{Label: "test shipping 1 label", Amount: 20},
			}},
			{ID: "2", Title: "test shipping 2", Prices: []tgbotapi.LabeledPrice{
				{Label: "test shipping 2 label", Amount: 30},
			}},
		},
	}

	_, err := tgbot.Request(ship)
	checkErr(err)
}
func (bot *Bot) sendInvoice(tgbot *tgbotapi.BotAPI) {
	var invoice = tgbotapi.InvoiceConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: bot.chat_id,
		},
		Title:         "test title",
		Description:   "test desc",
		Payload:       "test payload",
		ProviderToken: bot.payment_token, // stripe test token
		Currency:      "USD",
		Prices: []tgbotapi.LabeledPrice{
			{Label: "test LabelA", Amount: 145},
			{Label: "test LabelB", Amount: 245},
		},
		StartParameter:      "multi-chat",
		SuggestedTipAmounts: []int{1, 2, 3, 4},
		MaxTipAmount:        4,
		IsFlexible:          true,
		//NeedShippingAddress: true,
	}
	_, err2 := tgbot.Send(invoice)
	checkErr(err2)
}

func PayStart() {
	bot := Bot{}
	bot.ReadConfig()
	tgbot, err := tgbotapi.NewBotAPI(bot.api_token)
	checkErr(err)
	updateConfig := tgbotapi.NewUpdate(0)
	//updateConfig.Timeout = 60
	updates := tgbot.GetUpdatesChan(updateConfig)
	//var message = tgbotapi.MessageConfig{}
	bot.sendInvoice(tgbot)

	for update := range updates {
		// if update.Message != nil {
		// 	fmt.Println("message chat Id", update.Message.Chat.ID)
		// 	return
		// }
		if update.Message != nil && update.Message.IsCommand() && update.Message.Command() == "invoice" {
			bot.sendInvoice(tgbot)
		}
		if update.PreCheckoutQuery != nil {
			handlePreCheckoutQuery(tgbot, update)
		}
		if update.ShippingQuery != nil {
			handleShippingQuery(tgbot, update)
		}
	}
}
