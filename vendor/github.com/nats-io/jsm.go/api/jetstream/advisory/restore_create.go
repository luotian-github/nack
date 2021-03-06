package advisory

import (
	"github.com/nats-io/jsm.go/api/event"
)

// JSRestoreCreateAdvisoryV1 is an advisory sent after a snapshot is successfully started
//
// NATS Schema io.nats.jetstream.advisory.v1.restore_create
type JSRestoreCreateAdvisoryV1 struct {
	event.NATSEvent

	Stream string           `json:"stream"`
	Client APIAuditClientV1 `json:"client"`
}

func init() {
	err := event.RegisterTextCompactTemplate("io.nats.jetstream.advisory.v1.restore_create", `{{ .Time | ShortTime }} [Restore Create] {{ .Stream }}`)
	if err != nil {
		panic(err)
	}

	err = event.RegisterTextExtendedTemplate("io.nats.jetstream.advisory.v1.restore_create", `
[{{ .Time | ShortTime }}] [{{ .ID }}] Stream Restore Created

        Stream: {{ .Stream }}
        Client:
{{- if .Client.User }}
                      User: {{ .Client.User }} Account: {{ .Client.Account }}
{{- end }}
                      Host: {{ HostPort .Client.Host .Client.Port }}
                       CID: {{ .Client.CID }}
{{- if .Client.Name }}
                      Name: {{ .Client.Name }}
{{- end }}
           Library Version: {{ .Client.Version }}  Language: {{ with .Client.Lang }}{{ . }}{{ else }}Unknown{{ end }}
`)
	if err != nil {
		panic(err)
	}
}
