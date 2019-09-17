package assos

import (
	"github.com/genesixx/coalibot/utils"
	"github.com/nlopes/slack"
)

func Event(option string, event *utils.Message) bool {
	params := utils.SlackParams
	params.IconURL = "https://bde.student42.fr/img/bde42-logo-1538664197.jpg"
	params.Username = "Undefined Bot"
	// attachment := slack.Attachment{
	// 	Title:     "AfterWork (Ouverture du 95) Happy Hour de 16h à 20h !",
	// 	TitleLink: "https://www.facebook.com/events/409610906300164",
	// 	Text:      "Venez Networker autour d'une petite mousse 😎",
	// 	Footer:    "Powered by Coalibot",
	// 	Color:     "#87CEEB",
	// 	Fields: []slack.AttachmentField{
	// 		slack.AttachmentField{
	// 			Title: "Date",
	// 			Value: "Ce soir! (25/06)",
	// 			Short: false,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Monaco / 1664 / Grimbergen :beer:",
	// 			Value: "Pinte 5 €",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Gallia / Chouffe :beer:",
	// 			Value: "Pinte 6€",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Verre de vin :wine_glass:",
	// 			Value: "Pinte 3,50€",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Cocktails :cocktail:",
	// 			Value: "Avec alcool(sauf long island): 6€/Sans alcool: 5€",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Planche de charcuterie ou fromage",
	// 			Value: "12 €",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Guacamole et chips",
	// 			Value: "6 €",
	// 			Short: true,
	// 		},
	// 		slack.AttachmentField{
	// 			Title: "Planche et bouteille de vin",
	// 			Value: "27 €",
	// 			Short: true,
	// 		},
	// 	},
	// }
	// utils.PostMsg(event, slack.MsgOptionAttachments(attachment), slack.MsgOptionPostMessageParameters(params))
	utils.PostMsg(event, slack.MsgOptionText("Coming soon!", false), slack.MsgOptionPostMessageParameters(params))

	return true
}
