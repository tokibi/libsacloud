package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud/v2/internal/define"
	"github.com/sacloud/libsacloud/v2/internal/tools"
)

const destination = "sacloud/zz_envelopes.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-envelope: ")
}

func main() {
	outputPath := destination
	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), outputPath),
		Template:   tmpl,
		Parameter:  define.Resources,
	})
	log.Printf("generated: %s\n", outputPath)

}

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-envelope'; DO NOT EDIT

package sacloud

import (
{{- range .ImportStatements "github.com/sacloud/libsacloud/v2/sacloud/types" }}
	{{ . }}
{{- end }}
)

{{- range . }}
{{- range .Operations -}}

{{ if .HasRequestEnvelope }}
// {{ .RequestEnvelopeStructName }} is envelop of API request
type {{ .RequestEnvelopeStructName }} struct {
{{ if .IsRequestSingular }}
	{{- range .RequestPayloads}}
	{{.Name}} {{.TypeName}} {{.TagString}}
	{{- end }}
{{- else if .IsRequestPlural -}}
	{{- range .RequestPayloads}}
	{{.Name}} []{{.TypeName}} {{.TagString}}
	{{- end }}
{{ end }}
}
{{ end }}

{{ if .HasResponseEnvelope }}
// {{ .ResponseEnvelopeStructName }} is envelop of API response
type {{ .ResponseEnvelopeStructName }} struct {
{{- if .IsResponsePlural -}}
	Total       int        ` + "`" + `json:",omitempty"` + "`" + ` // トータル件数
	From        int        ` + "`" + `json:",omitempty"` + "`" + ` // ページング開始ページ
	Count       int        ` + "`" + `json:",omitempty"` + "`" + ` // 件数
{{ else }}
	IsOk    bool  ` + "`" + `json:"is_ok,omitempty"` + "`" + ` // is_ok項目
	Success types.APIResult  ` + "`" + `json:",omitempty"` + "`" + `      // success項目
{{ end }}
{{ if .IsResponseSingular }}
	{{- range .ResponsePayloads}}
	{{.Name}} {{.TypeName}} {{.TagString}}
	{{- end }}
{{- else if .IsResponsePlural -}}
	{{- range .ResponsePayloads}}
	{{.Name}} []{{.TypeName}} {{.TagString}}
	{{- end }}
{{ end }}
}
{{ end }}

{{- end -}}
{{- end -}}
`
