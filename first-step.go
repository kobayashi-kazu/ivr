package main

import (
	"github.com/aws/aws-lambda-go/events"
)

func firstStep(request events.APIGatewayProxyRequest) string {

	ssmlBuilder := NewBuilder()
	twimlBuilder := NewTwimlBuilder()

	ssmlBuilder.Paragraph("お電話ありがとうございます。").
		Paragraph("こちらは居酒屋かずま").
		Paragraph("ミッドタウン日比谷店です。")

	twimlBuilder.Say(ssmlBuilder.String())
	twimlBuilder.Pause(3)

	ssmlBuilder = NewBuilder()
	ssmlBuilder.Paragraph("ご予約の希望日時を４桁の数字で押して")
	ssmlBuilder.Paragraph("最後にシャープを押してください。")
	gatherBuilder := NewTwimlBuilder()
	twimlBuilder.Gather(gatherBuilder.Say(ssmlBuilder.String()).String(), "https://j3mase6i3f.execute-api.us-east-2.amazonaws.com/Staging/step-second", "POST", 10, "#")

	return twimlBuilder.Response()
}
