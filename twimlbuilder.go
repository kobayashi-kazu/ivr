package main

import (
	"bytes"
	"fmt"
)

type TwimlBuilder interface {
	Text(string) TwimlBuilder
	Say(string) TwimlBuilder
	Pause(int) TwimlBuilder
	Gather(string, string, string, int, string) TwimlBuilder
	String() string
	Response() string
	Hangup() TwimlBuilder
}

type twimlBuilder struct {
	buf bytes.Buffer
}

func (r *twimlBuilder) Text(text string) TwimlBuilder {
	r.buf.WriteString(text)
	return r
}

func (r *twimlBuilder) Gather(text string, action string, method string, timeout int, finishOnKey string) TwimlBuilder {
	r.Text("<Gather")
	if len(action) != 0 {
		r.Text(fmt.Sprintf(" action=\"%s\"", action))
	}
	if len(method) != 0 {
		r.Text(fmt.Sprintf(" method=\"%s\"", method))
	}
	if timeout != 0 {
		r.Text(fmt.Sprintf(" timeout=\"%d\"", timeout))
	}
	if len(finishOnKey) != 0 {
		r.Text(fmt.Sprintf(" finishOnKey=\"%s\"", finishOnKey))
	}

	r.Text(" >")

	return r.Text(text).Text("</Gather>")
}

func (r *twimlBuilder) Say(text string) TwimlBuilder {
	return r.Text(fmt.Sprintf("<Say>%s</Say>", text))
}

func (r *twimlBuilder) Pause(length int) TwimlBuilder {
	return r.Text(fmt.Sprintf("<Pause length=\"%d\" />", length))
}

func (r twimlBuilder) String() string {
	return r.buf.String()
}

func (r twimlBuilder) Response() string {
	return fmt.Sprintf("<Response>%s</Response>", r.String())
}

func (r *twimlBuilder) Hangup() TwimlBuilder {
	return r.Text("<Hangup/>")
}

func NewTwimlBuilder() TwimlBuilder {
	return &twimlBuilder{}
}