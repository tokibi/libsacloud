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

package types

import (
	"fmt"
	"strconv"
	"strings"
)

// EProxyLBPlan エンハンスドロードバランサのプラン
//
// エンハンスドロードバランサではプランはIDを受け取る形(Plan.ID)ではなく、
// ServiceClassに"cloud/proxylb/plain/100"のような形で文字列で指定する。
// このままでは扱いにくいためEProxyLBPlan型を設け、この型でjson.Marshaler/Unmarshalerを実装し
// プラン名とServiceClassでの文字列表現とで相互変換可能とする。
type EProxyLBPlan int

// ProxyLBPlans エンハンスドロードバランサのプラン
var ProxyLBPlans = struct {
	CPS100    EProxyLBPlan
	CPS500    EProxyLBPlan
	CPS1000   EProxyLBPlan
	CPS5000   EProxyLBPlan
	CPS10000  EProxyLBPlan
	CPS50000  EProxyLBPlan
	CPS100000 EProxyLBPlan
}{
	CPS100:    EProxyLBPlan(100),
	CPS500:    EProxyLBPlan(500),
	CPS1000:   EProxyLBPlan(1000),
	CPS5000:   EProxyLBPlan(5000),
	CPS10000:  EProxyLBPlan(10000),
	CPS50000:  EProxyLBPlan(50000),
	CPS100000: EProxyLBPlan(100000),
}

const (
	proxyLBServiceClassPrefix        = "cloud/proxylb/plain/"
	proxyLBServiceClassPrefixEscaped = "cloud\\/proxylb\\/plain\\/"
)

// MarshalJSON implements json.Marshaler
func (n *EProxyLBPlan) MarshalJSON() ([]byte, error) {
	if n == nil || int(*n) == 0 {
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf(`"%s%d"`, proxyLBServiceClassPrefix, int(*n))), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (n *EProxyLBPlan) UnmarshalJSON(b []byte) error {
	strPlan := string(b)
	if strPlan == `""` {
		*n = EProxyLBPlan(0)
		return nil
	}

	strPlan = strings.Replace(strPlan, `"`, "", -1)
	strPlan = strings.Replace(strPlan, proxyLBServiceClassPrefix, "", -1)
	strPlan = strings.Replace(strPlan, proxyLBServiceClassPrefixEscaped, "", -1)

	plan, err := strconv.Atoi(strPlan)
	if err != nil {
		return err
	}

	*n = EProxyLBPlan(plan)
	return nil
}
