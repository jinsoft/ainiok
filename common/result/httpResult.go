package result

import (
	"fmt"
	"github.com/jinsoft/ainiok/common/xerr"
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// Success
		r := Success(resp)
		httpx.OkJson(w, r)
	} else {
		// Error
		errCode := xerr.SERVER_COMMON_ERROR
		errMsg := xerr.MapErrMsg(errCode)

		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*xerr.CodeError); ok {
			// 自定义错误类型
			fmt.Println("自定义错误")
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			fmt.Println("系统错误")
			if gstatus, ok := status.FromError(causeErr); ok {
				// grpc error
				grpcCode := uint32(gstatus.Code())
				// 区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
				if xerr.IsCodeErr(grpcCode) {
					errCode = grpcCode
					errMsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("[API ERROR] : %v", err)
		httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
	}
}

func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.REQUES_PARAM_ERROR), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.REQUES_PARAM_ERROR, errMsg))
}
