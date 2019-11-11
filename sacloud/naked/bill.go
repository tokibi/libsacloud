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

package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Bill 請求情報
type Bill struct {
	ID             types.ID   `json:"BillID,omitempty" yaml:"bill_id,omitempty" structs:",omitempty"`    // 請求ID
	Amount         int64      `json:",omitempty" yaml:"amount,omitempty" structs:",omitempty"`           // 金額
	Date           *time.Time `json:",omitempty" yaml:"date,omitempty" structs:",omitempty"`             // 請求日
	MemberID       string     `json:",omitempty" yaml:"member_id,omitempty" structs:",omitempty"`        // 会員ID
	Paid           bool       `json:",omitempty" yaml:"paid,omitempty" structs:",omitempty"`             // 支払済フラグ
	PayLimit       *time.Time `json:",omitempty" yaml:"pay_limit,omitempty" structs:",omitempty"`        // 支払い期限
	PaymentClassID types.ID   `json:",omitempty" yaml:"payment_class_id,omitempty" structs:",omitempty"` // 支払いクラスID
}

// BillDetail 支払い明細情報
type BillDetail struct {
	ID             types.ID   `json:"ContractID,omitempty" yaml:"contract_id,omitempty" structs:",omitempty"` // 契約ID
	Amount         int64      `json:",omitempty" yaml:"amount,omitempty" structs:",omitempty"`                // 金額
	Description    string     `yaml:"description"`                                                            // 説明
	Index          int        `json:",omitempty" yaml:"index,omitempty" structs:",omitempty"`                 // インデックス
	ServiceClassID types.ID   `json:",omitempty" yaml:"service_class_id,omitempty" structs:",omitempty"`      // サービスクラスID
	Usage          int64      `json:",omitempty" yaml:"usage,omitempty" structs:",omitempty"`                 // 秒数
	Zone           string     `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`                  // ゾーン
	ContractEndAt  *time.Time `json:",omitempty" yaml:"contract_end_at,omitempty" structs:",omitempty"`       // 契約終了日時
}

// IsContractEnded 支払済か判定
func (d *BillDetail) IsContractEnded(t time.Time) bool {
	return d.ContractEndAt != nil && d.ContractEndAt.Before(t)
}

// BillDetailCSV 請求明細CSVレスポンス
type BillDetailCSV struct {
	// Count 件数
	Count int `json:",omitempty"`
	// ResponsedAt 応答日時
	ResponsedAt *time.Time `json:",omitempty"`
	// Filename ファイル名
	Filename string `json:",omitempty"`
	// RawBody ボディ(未加工)
	RawBody string `json:"Body,omitempty"`
	// HeaderRow ヘッダ行
	HeaderRow []string
	// BodyRows ボディ(各行/各列での配列)
	BodyRows [][]string
}
