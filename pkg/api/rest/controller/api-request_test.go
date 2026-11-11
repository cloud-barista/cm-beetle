package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRestGetRequest_ProjectionAndTumblebugEnrichment(t *testing.T) {
	e := echo.New()
	reqID := "req-get-projection-001"
	common.RemoveAllRequests()
	t.Cleanup(common.RemoveAllRequests)

	require.NoError(t, common.SetRequest(reqID, common.RequestDetails{
		ReqId:           reqID,
		StartTime:       time.Now().Add(-5 * time.Second),
		Status:          common.RequestStatusHandling,
		TumblebugSource: common.FormatTumblebugSource(reqID),
		RequestInfo: common.RequestInfo{
			Method: http.MethodPost,
			URL:    "/beetle/migration/ns/mig01/infra",
			Body:   map[string]any{"large": "request-payload"},
		},
		ResponseData: map[string]any{"large": "response-payload"},
	}))

	subID, _ := common.StartSubRequest(reqID, http.MethodPost, "/ns/mig01/infra", "infra", "mig-infra-01")
	require.Equal(t, reqID+"-1", subID)

	t.Run("brief=true omits bodies while keeping subRequestSummary and subRequests", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/beetle/request/"+reqID+"?brief=true", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("reqId")
		c.SetParamValues(reqID)

		require.NoError(t, RestGetRequest(c))
		require.Equal(t, http.StatusOK, rec.Code)

		var resp model.ApiResponse[common.RequestDetails]
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
		assert.True(t, resp.Success)
		assert.Nil(t, resp.Data.RequestInfo.Body)
		assert.Nil(t, resp.Data.ResponseData)
		require.NotNil(t, resp.Data.SubRequestSummary)
		assert.Equal(t, 1, resp.Data.SubRequestSummary.Total)
		assert.Equal(t, 1, resp.Data.SubRequestSummary.Handling)
		require.Len(t, resp.Data.SubRequests, 1)
		assert.Equal(t, subID, resp.Data.SubRequests[0].SubReqId)
	})

	t.Run("withSubRequests=false omits subRequests list while retaining summary", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/beetle/request/"+reqID+"?withSubRequests=false", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("reqId")
		c.SetParamValues(reqID)

		require.NoError(t, RestGetRequest(c))
		require.Equal(t, http.StatusOK, rec.Code)

		var resp model.ApiResponse[common.RequestDetails]
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
		require.NotNil(t, resp.Data.SubRequestSummary)
		assert.Empty(t, resp.Data.SubRequests)
	})

	t.Run("withTumblebug=true enriches live upstreamStatus and upstreamProgress without new subRequests", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/request/"+subID, r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(tbclient.TumblebugRequestSummary{
				Status: "Handling",
				ResponseData: []map[string]any{
					{"title": "Provisioning node 1/2"},
				},
			})
		}))
		restore := tbclient.SetClient(tbclient.NewClient(tbclient.ApiConfig{
			RestUrl:  srv.URL,
			Username: "default",
			Password: "default",
		}))
		t.Cleanup(func() {
			restore()
			srv.Close()
		})

		req := httptest.NewRequest(http.MethodGet, "/beetle/request/"+reqID+"?brief=true&withTumblebug=true", nil)
		req = req.WithContext(common.WithRequestID(req.Context(), "req-poll-observer"))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("reqId")
		c.SetParamValues(reqID)

		require.NoError(t, RestGetRequest(c))
		require.Equal(t, http.StatusOK, rec.Code)

		var resp model.ApiResponse[common.RequestDetails]
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
		require.Len(t, resp.Data.SubRequests, 1)
		assert.Equal(t, "Handling", resp.Data.SubRequests[0].UpstreamStatus)
		assert.NotNil(t, resp.Data.SubRequests[0].UpstreamProgress)

		// Verify GET enrichment did not record any sub-request on either parent or observer.
		obsSubs, _ := common.SnapshotSubRequests("req-poll-observer")
		assert.Empty(t, obsSubs)
	})
}

func TestRestGetAllRequests_SortingFilteringAndValidation(t *testing.T) {
	e := echo.New()
	common.RemoveAllRequests()
	t.Cleanup(common.RemoveAllRequests)

	baseTime := time.Now().Add(-10 * time.Minute)
	require.NoError(t, common.SetRequest("req-older", common.RequestDetails{
		ReqId:           "req-older",
		StartTime:       baseTime,
		Status:          common.RequestStatusSuccess,
		TumblebugSource: common.FormatTumblebugSource("req-older"),
		RequestInfo: common.RequestInfo{
			Method: http.MethodPost,
			URL:    "/beetle/recommendation/infra",
			Header: map[string]string{"X-Request-Source": "beetle-ux-lab"},
			Body:   map[string]any{"heavy": true},
		},
		ResponseData: map[string]any{"candidates": 3},
	}))
	require.NoError(t, common.SetRequest("req-newer", common.RequestDetails{
		ReqId:           "req-newer",
		StartTime:       baseTime.Add(5 * time.Minute),
		Status:          common.RequestStatusHandling,
		TumblebugSource: common.FormatTumblebugSource("req-newer"),
		RequestInfo: common.RequestInfo{
			Method: http.MethodPost,
			URL:    "/beetle/migration/ns/mig01/infra",
			Header: map[string]string{"X-Request-Source": "cm-butterfly"},
			Body:   map[string]any{"heavy": true},
		},
	}))

	t.Run("Sorts by startTime descending and applies brief and limit", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/beetle/requests?brief=true&limit=1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		require.NoError(t, RestGetAllRequests(c))
		require.Equal(t, http.StatusOK, rec.Code)

		var resp model.ApiResponse[map[string][]common.RequestDetails]
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
		items := resp.Data["requests"]
		require.Len(t, items, 1)
		assert.Equal(t, "req-newer", items[0].ReqId)
		assert.Nil(t, items[0].RequestInfo.Body)
	})

	t.Run("Filters by source header or tumblebugSource", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/beetle/requests?source=cm-beetle:req-older", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		require.NoError(t, RestGetAllRequests(c))
		require.Equal(t, http.StatusOK, rec.Code)

		var resp model.ApiResponse[map[string][]common.RequestDetails]
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
		items := resp.Data["requests"]
		require.Len(t, items, 1)
		assert.Equal(t, "req-older", items[0].ReqId)
	})

	t.Run("Rejects negative or non-integer limit with 400", func(t *testing.T) {
		for _, badLimit := range []string{"-1", "abc"} {
			req := httptest.NewRequest(http.MethodGet, "/beetle/requests?limit="+badLimit, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			require.NoError(t, RestGetAllRequests(c))
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestRestDeleteRequest_ReturnsSuccessEnvelope(t *testing.T) {
	e := echo.New()
	common.RemoveAllRequests()
	t.Cleanup(common.RemoveAllRequests)

	reqID := "req-delete-001"
	require.NoError(t, common.SetRequest(reqID, common.RequestDetails{
		ReqId:     reqID,
		StartTime: time.Now(),
		Status:    common.RequestStatusSuccess,
	}))

	req := httptest.NewRequest(http.MethodDelete, "/beetle/request/"+reqID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("reqId")
	c.SetParamValues(reqID)

	require.NoError(t, RestDeleteRequest(c))
	require.Equal(t, http.StatusOK, rec.Code)

	var resp model.ApiResponse[any]
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.True(t, resp.Success)
	assert.False(t, common.HasRequest(reqID))
}
