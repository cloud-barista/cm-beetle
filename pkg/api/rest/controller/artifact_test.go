package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/artifact"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArtifactController_CRUDAndLineage(t *testing.T) {
	_, err := artifact.InitStore(t.TempDir())
	require.NoError(t, err)

	e := echo.New()

	// 1. POST /beetle/artifacts (Source stage)
	createBody := `{
		"id": "art-http-src",
		"name": "http-source-infra",
		"category": "infra",
		"stage": "source",
		"sourceGroupId": "sg-http-test",
		"payload": {
			"servers": [
				{
					"hostname": "node-01",
					"machineId": "m-01",
					"cpu": {"cpus": 1, "cores": 2, "threads": 4},
					"memory": {"totalSize": 8}
				}
			]
		}
	}`
	req := httptest.NewRequest(http.MethodPost, "/beetle/artifacts", bytes.NewBufferString(createBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	require.NoError(t, CreateArtifact(c))
	assert.Equal(t, http.StatusCreated, rec.Code)

	var createResp model.ApiResponse[artifact.RawArtifact]
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &createResp))
	assert.True(t, createResp.Success)
	assert.Equal(t, "OnpremInfra", createResp.Data.SchemaType)
	assert.Equal(t, "art-http-src", createResp.Data.Lineage.SourceModelId)

	// 2. POST /beetle/artifacts (Planning stage derived from Source)
	planBody := `{
		"id": "art-http-plan",
		"name": "http-plan-infra",
		"category": "infra",
		"stage": "planning",
		"parentId": "art-http-src",
		"payload": {
			"nodes": [
				{
					"hostname": "node-01-plan",
					"machineId": "m-01",
					"cpu": {"cpus": 1, "cores": 4, "threads": 8},
					"memory": {"totalSize": 16}
				}
			]
		}
	}`
	reqPlan := httptest.NewRequest(http.MethodPost, "/beetle/artifacts", bytes.NewBufferString(planBody))
	reqPlan.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recPlan := httptest.NewRecorder()
	cPlan := e.NewContext(reqPlan, recPlan)
	require.NoError(t, CreateArtifact(cPlan))
	assert.Equal(t, http.StatusCreated, recPlan.Code)

	// 3. GET /beetle/artifacts?category=infra&stage=planning
	reqList := httptest.NewRequest(http.MethodGet, "/beetle/artifacts?category=infra&stage=planning", nil)
	recList := httptest.NewRecorder()
	cList := e.NewContext(reqList, recList)
	require.NoError(t, ListArtifacts(cList))
	assert.Equal(t, http.StatusOK, recList.Code)

	var listResp model.ApiResponse[[]artifact.RawArtifact]
	require.NoError(t, json.Unmarshal(recList.Body.Bytes(), &listResp))
	require.Len(t, listResp.Data, 1)
	assert.Equal(t, "art-http-plan", listResp.Data[0].Id)

	// 4. GET /beetle/artifacts/:id/lineage
	reqLin := httptest.NewRequest(http.MethodGet, "/beetle/artifacts/art-http-plan/lineage", nil)
	recLin := httptest.NewRecorder()
	cLin := e.NewContext(reqLin, recLin)
	cLin.SetParamNames("id")
	cLin.SetParamValues("art-http-plan")
	require.NoError(t, GetArtifactLineage(cLin))
	assert.Equal(t, http.StatusOK, recLin.Code)

	var linResp model.ApiResponse[artifact.LineageGraph]
	require.NoError(t, json.Unmarshal(recLin.Body.Bytes(), &linResp))
	assert.Equal(t, "art-http-src", linResp.Data.Lineage.SourceModelId)
	assert.Equal(t, "art-http-plan", linResp.Data.Lineage.PlanningModelId)

	// 5. PUT /beetle/artifacts/:id
	updateBody := `{"name": "http-plan-infra-updated", "version": "1.2.0"}`
	reqPut := httptest.NewRequest(http.MethodPut, "/beetle/artifacts/art-http-plan", bytes.NewBufferString(updateBody))
	reqPut.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recPut := httptest.NewRecorder()
	cPut := e.NewContext(reqPut, recPut)
	cPut.SetParamNames("id")
	cPut.SetParamValues("art-http-plan")
	require.NoError(t, UpdateArtifact(cPut))
	assert.Equal(t, http.StatusOK, recPut.Code)

	// 6. DELETE /beetle/artifacts/:id
	reqDel := httptest.NewRequest(http.MethodDelete, "/beetle/artifacts/art-http-plan", nil)
	recDel := httptest.NewRecorder()
	cDel := e.NewContext(reqDel, recDel)
	cDel.SetParamNames("id")
	cDel.SetParamValues("art-http-plan")
	require.NoError(t, DeleteArtifact(cDel))
	assert.Equal(t, http.StatusOK, recDel.Code)
}

