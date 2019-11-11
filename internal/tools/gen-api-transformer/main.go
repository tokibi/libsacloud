// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud/v2/internal/define"
	"github.com/sacloud/libsacloud/v2/internal/tools"
)

const destination = "sacloud/zz_api_transformers.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-transformer: ")
}

func main() {
	outputPath := destination
	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), outputPath),
		Template:   tmpl,
		Parameter:  define.APIs,
	})
	log.Printf("generated: %s\n", outputPath)
}

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-transformer'; DO NOT EDIT

package sacloud

import (
{{- range .ImportStatements "encoding/json" "github.com/sacloud/libsacloud/v2/pkg/mapconv" }}
	{{ . }}
{{- end }}
)

{{ range . }}{{ $typeName := .TypeName }}{{$resource := .}}
{{ range .Operations }}{{$returnErrStatement := .ReturnErrorStatement}}{{ $operationName := .MethodName }}
{{ if .HasRequestEnvelope }}
func (o *{{ $typeName }}Op) transform{{.MethodName}}Args({{ range .Arguments }}{{ .ArgName }} {{ .TypeName }},{{ end }}) (*{{.RequestEnvelopeStructName}}, error) {
	{{- range $i, $v := .Arguments }}
	if {{.ArgName}} == {{.ZeroValueOnSource}} {
		{{.ArgName}} = {{.ZeroInitializer}}	
	}
	var arg{{$i}} interface{} = {{.ArgName}}
	if v , ok := arg{{$i}}.(argumentDefaulter); ok {
		arg{{$i}} = v.setDefaults()
	}
	{{- end }}
	args := &struct {
		{{- range $i, $v := .Arguments }}
		Arg{{ $i }} interface{} {{.MapConvTagSrc}}
		{{- end }}
	}{
		{{- range $i, $v := .Arguments }}
		Arg{{ $i }}: arg{{ $i }},
		{{- end }}
	}

	v := &{{.RequestEnvelopeStructName}}{}
	if err := mapconv.ConvertTo(args, v); err != nil {
		return nil, err
	}
	return v, nil
}
{{ end -}}

{{ if .HasResponseEnvelope }}
func (o *{{ $typeName }}Op) transform{{.MethodName}}Results(data []byte) (*{{.ResultTypeName}}, error) {
	nakedResponse := &{{.ResponseEnvelopeStructName}}{}
	if err := json.Unmarshal(data, nakedResponse); err != nil {
		return nil, err
	}
	
	results := &{{.ResultTypeName}}{}
	if err :=  mapconv.ConvertFrom(nakedResponse, results); err != nil {
		return  nil, err
	}	
	return results, nil
}
{{ end -}}
{{ end -}}
{{ end -}}
`
