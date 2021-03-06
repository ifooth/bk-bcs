/*
Tencent is pleased to support the open source community by making Blueking Container Service available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package zone

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/viper"

	"bk-bscp/cmd/bscp-businessserver/modules/audit"
	"bk-bscp/internal/database"
	pb "bk-bscp/internal/protocol/businessserver"
	pbcommon "bk-bscp/internal/protocol/common"
	pbdatamanager "bk-bscp/internal/protocol/datamanager"
	"bk-bscp/pkg/logger"
)

// DeleteAction deletes target zone object.
type DeleteAction struct {
	viper      *viper.Viper
	dataMgrCli pbdatamanager.DataManagerClient

	req  *pb.DeleteZoneReq
	resp *pb.DeleteZoneResp
}

// NewDeleteAction creates new DeleteAction.
func NewDeleteAction(viper *viper.Viper, dataMgrCli pbdatamanager.DataManagerClient,
	req *pb.DeleteZoneReq, resp *pb.DeleteZoneResp) *DeleteAction {
	action := &DeleteAction{viper: viper, dataMgrCli: dataMgrCli, req: req, resp: resp}

	action.resp.Seq = req.Seq
	action.resp.ErrCode = pbcommon.ErrCode_E_OK
	action.resp.ErrMsg = "OK"

	return action
}

// Err setup error code message in response and return the error.
func (act *DeleteAction) Err(errCode pbcommon.ErrCode, errMsg string) error {
	act.resp.ErrCode = errCode
	act.resp.ErrMsg = errMsg
	return errors.New(errMsg)
}

// Input handles the input messages.
func (act *DeleteAction) Input() error {
	if err := act.verify(); err != nil {
		return act.Err(pbcommon.ErrCode_E_BS_PARAMS_INVALID, err.Error())
	}
	return nil
}

// Output handles the output messages.
func (act *DeleteAction) Output() error {
	// do nothing.
	return nil
}

func (act *DeleteAction) verify() error {
	length := len(act.req.Bid)
	if length == 0 {
		return errors.New("invalid params, bid missing")
	}
	if length > database.BSCPIDLENLIMIT {
		return errors.New("invalid params, bid too long")
	}

	length = len(act.req.Zoneid)
	if length == 0 {
		return errors.New("invalid params, zoneid missing")
	}
	if length > database.BSCPIDLENLIMIT {
		return errors.New("invalid params, zoneid too long")
	}

	length = len(act.req.Operator)
	if length == 0 {
		return errors.New("invalid params, operator missing")
	}
	if length > database.BSCPNAMELENLIMIT {
		return errors.New("invalid params, operator too long")
	}
	return nil
}

func (act *DeleteAction) delete() (pbcommon.ErrCode, string) {
	r := &pbdatamanager.DeleteZoneReq{
		Seq:      act.req.Seq,
		Bid:      act.req.Bid,
		Zoneid:   act.req.Zoneid,
		Operator: act.req.Operator,
	}

	ctx, cancel := context.WithTimeout(context.Background(), act.viper.GetDuration("datamanager.calltimeout"))
	defer cancel()

	logger.V(2).Infof("DeleteZone[%d]| request to datamanager DeleteZone, %+v", act.req.Seq, r)

	resp, err := act.dataMgrCli.DeleteZone(ctx, r)
	if err != nil {
		return pbcommon.ErrCode_E_BS_SYSTEM_UNKONW, fmt.Sprintf("request to datamanager DeleteZone, %+v", err)
	}
	if resp.ErrCode != pbcommon.ErrCode_E_OK {
		return resp.ErrCode, resp.ErrMsg
	}

	// audit here on zone deleted.
	audit.Audit(int32(pbcommon.SourceType_ST_ZONE), int32(pbcommon.SourceOpType_SOT_DELETE),
		act.req.Bid, act.req.Zoneid, act.req.Operator, "")

	return pbcommon.ErrCode_E_OK, ""
}

// Do makes the workflows of this action base on input messages.
func (act *DeleteAction) Do() error {
	// delete zone.
	if errCode, errMsg := act.delete(); errCode != pbcommon.ErrCode_E_OK {
		return act.Err(errCode, errMsg)
	}
	return nil
}
