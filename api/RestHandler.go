package api

import (
	"encoding/json"
	"github.com/devtron-labs/telemetry/pkg/telemetry"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type RestHandler interface {
	GetApiKey(w http.ResponseWriter, r *http.Request)
	CheckForOptOut(w http.ResponseWriter, r *http.Request)
	GetPostHogInfo(w http.ResponseWriter, r *http.Request)
}

func NewRestHandlerImpl(logger *zap.SugaredLogger,
	telemetryEventService telemetry.TelemetryService) *RestHandlerImpl {
	return &RestHandlerImpl{
		logger:                logger,
		telemetryEventService: telemetryEventService,
	}
}

type RestHandlerImpl struct {
	logger                *zap.SugaredLogger
	telemetryEventService telemetry.TelemetryService
}
type Response struct {
	Code   int         `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
	Result interface{} `json:"result,omitempty"`
	Errors []*ApiError `json:"errors,omitempty"`
}
type ApiError struct {
	HttpStatusCode    int         `json:"-"`
	Code              string      `json:"code,omitempty"`
	InternalMessage   string      `json:"internalMessage,omitempty"`
	UserMessage       interface{} `json:"userMessage,omitempty"`
	UserDetailMessage string      `json:"userDetailMessage,omitempty"`
}

func (impl RestHandlerImpl) writeJsonResp(w http.ResponseWriter, err error, respBody interface{}, status int) {
	response := Response{}
	response.Code = status
	response.Status = http.StatusText(status)
	if err == nil {
		response.Result = respBody
	} else {
		apiErr := &ApiError{}
		apiErr.Code = "000" // 000=unknown
		apiErr.InternalMessage = err.Error()
		apiErr.UserMessage = respBody
		response.Errors = []*ApiError{apiErr}

	}
	b, err := json.Marshal(response)
	if err != nil {
		impl.logger.Error("error in marshaling err object", err)
		status = 500
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(b)
}

type ResetRequest struct {
	AppId         int `json:"appId"`
	EnvironmentId int `json:"environmentId"`
}

func (impl *RestHandlerImpl) GetApiKey(w http.ResponseWriter, r *http.Request) {
	result, err := impl.telemetryEventService.GetByAPIKey()
	if err != nil {
		impl.logger.Errorw("error on getting telemetry api key", "err", err)
		impl.writeJsonResp(w, err, nil, http.StatusInternalServerError)
		return
	}
	impl.writeJsonResp(w, err, result, 200)
}

func (impl *RestHandlerImpl) CheckForOptOut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ucid := vars["ucid"]
	isOptOut, err := impl.telemetryEventService.CheckForOptOut(ucid)
	if err != nil {
		impl.logger.Errorw("error on checking ucid opt-out or not", "err", err)
		impl.writeJsonResp(w, err, nil, http.StatusInternalServerError)
		return
	}
	impl.writeJsonResp(w, err, isOptOut, 200)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (impl *RestHandlerImpl) GetPostHogInfo(w http.ResponseWriter, r *http.Request) {
	result, err := impl.telemetryEventService.GetByPosthogInfo()
	if err != nil {
		impl.logger.Errorw("error on getting telemetry api key", "err", err)
		impl.writeJsonResp(w, err, nil, http.StatusInternalServerError)
		return
	}
	impl.writeJsonResp(w, err, result, 200)
}
