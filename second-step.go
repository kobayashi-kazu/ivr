package main

import (
	"github.com/aws/aws-lambda-go/events"
)

func secondStep(request events.APIGatewayProxyRequest) string {

	ssmlBuilder := NewBuilder()
	twimlBuilder := NewTwimlBuilder()

	ssmlBuilder.SayAs("????" + "0321", INTERPRET_AS_DATE).
		Paragraph("のご予約を承りました。").
		Paragraph("ご予約ありがとうございました。")

	twimlBuilder.Say(ssmlBuilder.String())
	twimlBuilder.Pause(3).Hangup()

	return twimlBuilder.Response()
}
