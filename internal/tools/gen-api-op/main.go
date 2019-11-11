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

const destination = "sacloud/zz_api_ops.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-op: ")
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

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-op'; DO NOT EDIT

package sacloud

import (
{{- range .ImportStatements "context" "encoding/json" "github.com/sacloud/libsacloud/v2/sacloud/naked" "github.com/sacloud/libsacloud/v2/pkg/mapconv" }}
	{{ . }}
{{- end }}
)

func init() {
{{ range . }}
	SetClientFactoryFunc("{{.TypeName}}", func(caller APICaller) interface{} {
		return &{{ .TypeName }}Op {
			Client: caller,
			PathSuffix: "{{.GetPathSuffix}}",
			PathName: "{{.GetPathName}}",
		}
	})
{{ end -}}
}

{{ range . }}{{ $typeName := .TypeName }}{{$resource := .}}

/************************************************* 
* {{$typeName}}Op
*************************************************/

// {{ .TypeName }}Op implements {{ .TypeName }}API interface
type {{ .TypeName }}Op struct{
	// Client APICaller
    Client APICaller
	// PathSuffix is used when building URL
	PathSuffix string
	// PathName is used when building URL
	PathName string
}

// New{{ $typeName}}Op creates new {{ $typeName}}Op instance
func New{{ $typeName}}Op(caller APICaller) {{ $typeName}}API {
	return GetClientFactoryFunc("{{$typeName}}")(caller).({{$typeName}}API)
}

{{ range .Operations }}{{$returnErrStatement := .ReturnErrorStatement}}{{ $operationName := .MethodName }}
// {{ .MethodName }} is API call
func (o *{{ $typeName }}Op) {{ .MethodName }}(ctx context.Context{{if not $resource.IsGlobal}}, zone string{{end}}{{ range .Arguments }}, {{ .ArgName }} {{ .TypeName }}{{ end }}) {{.ResultsStatement}} {
	// build request URL
	url, err := buildURL("{{.GetPathFormat}}", map[string]interface{}{
		"rootURL": SakuraCloudAPIRoot,
		"pathSuffix": o.PathSuffix,
		"pathName": o.PathName,
		{{- if $resource.IsGlobal }}
		"zone": APIDefaultZone,
		{{- else }}
		"zone": zone,
		{{- end }}
		{{- range .Arguments }}
		"{{.PathFormatName}}": {{.Name}},
		{{- end }}
	})
	if err != nil {
		return {{ $returnErrStatement }}
	}

	{{if .IsPatch -}}
	original, err := o.Read(ctx{{if not $resource.IsGlobal}}, zone{{ end }}, id)
	if err != nil {
		return nil, err
	}
	patchParam := make(map[string]interface{})
	if err := mergo.Map(&patchParam, original); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	if err := mergo.Map(&patchParam, param); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	if err := mergo.Map(param, &patchParam); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}

	{{ range .PatchArgumentModel.Fields -}}
	{{ if .IsNeedPatchEmpty}}if param.PatchEmptyTo{{.Name}} {
		param.{{.Name}} = {{.Type.ZeroValueSourceCode}}	
	}
	{{ end }}{{ end -}}
	{{ end -}}

	// build request body
	var body interface{}
{{ if .HasRequestEnvelope -}}
	v, err := o.transform{{.MethodName}}Args({{ range .Arguments }}{{ .ArgName }},{{ end }})
	if err != nil {
		return {{ $returnErrStatement }}
	}
	body = v
{{ end }}

	// do request
	{{ if .HasResponseEnvelope -}}
	data, err := o.Client.Do(ctx, "{{.Method}}", url, body)
	{{ else -}}
	_, err = o.Client.Do(ctx, "{{.Method}}", url, body)
	{{ end -}}
	if err != nil {
		return {{ $returnErrStatement }}
	}

	// build results
	{{ if .HasResponseEnvelope -}}
	results, err := o.transform{{.MethodName}}Results(data)
	if err != nil {
		return {{ $returnErrStatement }}
	}	
	return {{ .ReturnStatement }}
	{{ else }}
	return nil
	{{ end -}}
}
{{ end -}}
{{ end -}}
`
