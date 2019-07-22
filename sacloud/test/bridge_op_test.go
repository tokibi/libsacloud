package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func TestBridgeOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &CRUDTestFunc{
			Func: testBridgeCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createBridgeExpected,
				IgnoreFields: ignoreBridgeFields,
			}),
		},

		Read: &CRUDTestFunc{
			Func: testBridgeRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createBridgeExpected,
				IgnoreFields: ignoreBridgeFields,
			}),
		},

		Updates: []*CRUDTestFunc{
			{
				Func: testBridgeUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateBridgeExpected,
					IgnoreFields: ignoreBridgeFields,
				}),
			},
			{
				Func: testBridgeUpdateToMin,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateBridgeToMinExpected,
					IgnoreFields: ignoreBridgeFields,
				}),
			},
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testBridgeDelete,
		},
	})
}

var (
	ignoreBridgeFields = []string{
		"ID",
		"CreatedAt",
		"Region",
		"SwitchInZone",
		"BridgeInfo",
	}

	createBridgeParam = &sacloud.BridgeCreateRequest{
		Name:        "libsacloud-bridge",
		Description: "desc",
	}
	createBridgeExpected = &sacloud.Bridge{
		Name:        createBridgeParam.Name,
		Description: createBridgeParam.Description,
	}
	updateBridgeParam = &sacloud.BridgeUpdateRequest{
		Name:        "libsacloud-bridge-upd",
		Description: "desc-upd",
	}
	updateBridgeExpected = &sacloud.Bridge{
		Name:        updateBridgeParam.Name,
		Description: updateBridgeParam.Description,
	}
	updateBridgeToMinParam = &sacloud.BridgeUpdateRequest{
		Name: "libsacloud-bridge-to-min",
	}
	updateBridgeToMinExpected = &sacloud.Bridge{
		Name: updateBridgeToMinParam.Name,
	}
)

func testBridgeCreate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewBridgeOp(caller)
	return client.Create(ctx, testZone, createBridgeParam)
}

func testBridgeRead(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewBridgeOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testBridgeUpdate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewBridgeOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateBridgeParam)
}

func testBridgeUpdateToMin(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewBridgeOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateBridgeToMinParam)
}

func testBridgeDelete(ctx *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewBridgeOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}
