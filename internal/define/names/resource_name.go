package names

import (
	"strings"

	"github.com/sacloud/libsacloud/v2/internal/dsl"
)

// ResourceFieldName リソース名がペイロードなどで利用される場合のフィールド名、コード生成時に利用される
func ResourceFieldName(resourceName string, form dsl.PayloadForm) string {
	switch {
	case form.IsSingular():
		return resourceName
	case form.IsPlural():
		// TODO とりあえずワードで例外指定
		switch {
		case resourceName == "NFS":
			return resourceName
		case strings.HasSuffix(resourceName, "ch"):
			return resourceName + "es"
		default:
			return resourceName + "s"
		}
	default:
		return ""
	}
}

// CreateParameterName Create操作に渡すパラメータの名称
func CreateParameterName(resourceName string) string {
	return resourceName + "CreateRequest"
}

// UpdateParameterName Create操作に渡すパラメータの名称
func UpdateParameterName(resourceName string) string {
	return resourceName + "UpdateRequest"
}