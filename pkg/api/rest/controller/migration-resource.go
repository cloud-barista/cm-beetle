/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package controller has handlers and their request/response bodies for migration APIs
package controller

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	tbresource "github.com/cloud-barista/cb-tumblebug/src/interface/rest/server/resource"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/middlewares"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// JSONResult is a dummy struct for Swagger annotations.
type JSONResult struct {
}

// createTumblebugProxyHandler creates a handler function that proxies requests to Tumblebug API
// sourcePattern: the source URL pattern with wildcards (*)
// targetPattern: the target URL pattern with wildcards ($1, $2, etc.)
func createTumblebugProxyHandler(sourcePattern string, targetPattern string) echo.HandlerFunc {

	log.Trace().Msgf("Just keep importing tbmodel %s", tbmodel.DefaultNamespace)
	log.Trace().Msgf("Just keep importing tbresource %v", tbresource.RestGetAllVNetResponse{})

	// Parse Tumblebug endpoint URL
	tbURL, err := url.Parse(config.Tumblebug.RestUrl)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse Tumblebug endpoint URL")
		return func(c echo.Context) error {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Server configuration error",
			})
		}
	}

	// Return a handler that processes the request and forwards it to Tumblebug
	return func(c echo.Context) error {
		// Create the rewrite rule based on the source and target patterns
		rewriteRules := map[string]string{
			sourcePattern: targetPattern,
		}

		log.Debug().Msgf("Proxying with rewrite rule: %s -> %s", sourcePattern, targetPattern)

		// Use the existing Proxy middleware
		proxyMiddleware := middlewares.Proxy(middlewares.ProxyConfig{
			URL:     tbURL,
			Rewrite: rewriteRules,
			ModifyResponse: func(res *http.Response) error {
				resBytes, err := io.ReadAll(res.Body)
				if err != nil {
					return err
				}

				log.Debug().Msgf("[Proxy] response from %s", res.Request.URL)
				log.Trace().Msgf("[Proxy] response body: %s", string(resBytes))

				res.Body = io.NopCloser(bytes.NewReader(resBytes))
				return nil
			},
		})

		// Create a handler that will be wrapped by the proxy middleware
		handler := echo.HandlerFunc(func(c echo.Context) error {
			return nil // This will never be called because the proxy middleware will handle the request
		})

		// Apply the proxy middleware to the handler
		return proxyMiddleware(handler)(c)
	}
}

// ========== vNet Resource APIs ==========

// ListMigratedVNets godoc
// @ID ListMigratedVNets
// @Summary List all migrated virtual networks
// @Description Get the list of all migrated virtual networks in the namespace
// @Tags [Migration] Resources for VM infrastructure
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID"
// @Success 200 {object} JSONResult{[DEFAULT]=tbresource.RestGetAllVNetResponse,[ID]=tbmodel.IdList} "Different return structures by the given option param"
// @Failure 404 {object} tbmodel.SimpleMsg
// @Failure 500 {object} tbmodel.SimpleMsg
// @Router /migration/ns/{nsId}/resources/vNet [get]
func ListMigratedVNets(c echo.Context) error {
	// Source path pattern with * to capture nsId
	sourcePattern := "/migration/ns/*/resources/vNet"
	// Target path pattern using $1 for captured nsId
	targetPattern := "/ns/$1/resources/vNet"

	proxyHandler := createTumblebugProxyHandler(sourcePattern, targetPattern)
	return proxyHandler(c)
}

// GetMigratedVNet godoc
// @ID GetMigratedVNet
// @Summary Get a specific migrated virtual network
// @Description Get details of a specific virtual network in the namespace
// @Tags [Migration] Resources for VM infrastructure
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID"
// @Param vNetId path string true "Virtual Network ID"
// @Success 200 {object} tbmodel.TbVNetInfo
// @Failure 404 {object} tbmodel.SimpleMsg
// @Failure 500 {object} tbmodel.SimpleMsg
// @Router /migration/ns/{nsId}/resources/vNet/{vNetId} [get]
func GetMigratedVNet(c echo.Context) error {
	// Path pattern that captures two path parameters
	sourcePattern := "/migration/ns/*/resources/vNet/*"
	// First * is used as $1(nsId), second * as $2(vNetId)
	targetPattern := "/ns/$1/resources/vNet/$2"

	proxyHandler := createTumblebugProxyHandler(sourcePattern, targetPattern)
	return proxyHandler(c)
}

// CreateVNet godoc
// @ID CreateVNet
// @Summary Create a migrated virtual network
// @Description Create a new migrated virtual network in the namespace
// @Tags [Migration] Resources for VM infrastructure
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID"
// @Param vNetReq body tbmodel.TbVNetReq true "Virtual Network creation request"
// @Success 200 {object} tbmodel.TbVNetInfo
// @Failure 400 {object} tbmodel.SimpleMsg
// @Failure 500 {object} tbmodel.SimpleMsg
// @Router /migration/ns/{nsId}/resources/vNet [post]
func CreateVNet(c echo.Context) error {
	// Source path pattern with * to capture nsId
	sourcePattern := "/migration/ns/*/resources/vNet"
	// Target path pattern using $1 for captured nsId
	targetPattern := "/ns/$1/resources/vNet"

	proxyHandler := createTumblebugProxyHandler(sourcePattern, targetPattern)
	return proxyHandler(c)
}

// DeleteMigratedVNet godoc
// @ID DeleteMigratedVNet
// @Summary Delete a migrated virtual network
// @Description Delete a specific migrated virtual network in the namespace
// @Tags [Migration] Resources for VM infrastructure
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID"
// @Param vNetId path string true "Virtual Network ID"
// @Param action query string false "Action" Enums(withsubnets,refine,force)
// @Success 200 {object} tbmodel.SimpleMsg
// @Failure 404 {object} tbmodel.SimpleMsg
// @Router /migration/ns/{nsId}/resources/vNet/{vNetId} [delete]
func DeleteMigratedVNet(c echo.Context) error {
	// Path pattern that captures two path parameters
	sourcePattern := "/migration/ns/*/resources/vNet/*"
	// First * is used as $1(nsId), second * as $2(vNetId)
	targetPattern := "/ns/$1/resources/vNet/$2"

	proxyHandler := createTumblebugProxyHandler(sourcePattern, targetPattern)
	return proxyHandler(c)
}

// ========== SSH Key Resource APIs ==========
// ListMigratedSSHKeys godoc
// @ID ListMigratedSSHKeys
// @Summary List all migrated SSH keys
// @Description Get the list of all migrated SSH keys in the namespace
// @Tags [Migration] Resources for VM infrastructure
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(default)
// @Param option query string false "Option" Enums(id)
// @Param filterKey query string false "Field key for filtering (ex: systemLabel)"
// @Param filterVal query string false "Field value for filtering (ex: Registered from CSP resource)"
// @Success 200 {object} JSONResult{[DEFAULT]=tbresource.RestGetAllSshKeyResponse,[ID]=tbmodel.IdList} "Different return structures by the given option param"
// @Failure 404 {object} tbmodel.SimpleMsg
// @Failure 500 {object} tbmodel.SimpleMsg
// @Router /migration/ns/{nsId}/resources/sshKey [get]
func ListMigratedSSHKeys(c echo.Context) error {
	// Source path pattern with * to capture nsId
	sourcePattern := "/migration/ns/*/resources/sshKey"
	// Target path pattern using $1 for captured nsId
	targetPattern := "/ns/$1/resources/sshKey"

	proxyHandler := createTumblebugProxyHandler(sourcePattern, targetPattern)
	return proxyHandler(c)
}

// GetMigratedSSHKey godoc
// @ID GetMigratedSSHKey
// @Summary Get a specific migrated SSH key
// @Description Get details of a specific migrated SSH key in the namespace
// @Tags [Migration] Resources for VM infrastructure
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID"
// @Param sshKeyId path string true "SSH Key ID"
// @Success 200 {object} tbmodel.TbSshKeyInfo
// @Failure 404 {object} tbmodel.SimpleMsg
// @Failure 500 {object} tbmodel.SimpleMsg
// @Router /migration/ns/{nsId}/resources/sshKey/{sshKeyId} [get]
func GetMigratedSSHKey(c echo.Context) error {
	// Path pattern that captures two path parameters
	sourcePattern := "/migration/ns/*/resources/sshKey/*"
	// First * is used as $1(nsId), second * as $2(sshKeyId)
	targetPattern := "/ns/$1/resources/sshKey/$2"

	proxyHandler := createTumblebugProxyHandler(sourcePattern, targetPattern)
	return proxyHandler(c)
}

// CreateMigratedSSHKey godoc
// @ID CreateMigratedSSHKey
// @Summary Create a migrated SSH key
// @Description Create a new migrated SSH key in the namespace
// @Tags [Migration] Resources for VM infrastructure
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(default)
// @Param option query string false "Option: [required params for register] connectionName, name, cspKeyId" Enums(register)
// @Param sshKeyReq body tbmodel.TbSshKeyReq true "Details for an SSH key object"
// @Success 200 {object} tbmodel.TbSshKeyInfo
// @Failure 404 {object} tbmodel.SimpleMsg
// @Failure 500 {object} tbmodel.SimpleMsg
// @Router /migration/ns/{nsId}/resources/sshKey [post]
func CreateMigratedSSHKey(c echo.Context) error {
	// Source path pattern with * to capture nsId
	sourcePattern := "/migration/ns/*/resources/sshKey"
	// Target path pattern using $1 for captured nsId
	targetPattern := "/ns/$1/resources/sshKey"

	proxyHandler := createTumblebugProxyHandler(sourcePattern, targetPattern)
	return proxyHandler(c)
}

// DeleteMigratedSSHKey godoc
// @ID DeleteMigratedSSHKey
// @Summary Delete a migrated SSH key
// @Description Delete a specific migrated SSH key in the namespace
// @Tags [Migration] Resources for VM infrastructure
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(default)
// @Param sshKeyId path string true "SSH Key ID"
// @Success 200 {object} tbmodel.SimpleMsg
// @Failure 404 {object} tbmodel.SimpleMsg
// @Failure 500 {object} tbmodel.SimpleMsg
// @Router /migration/ns/{nsId}/resources/sshKey/{sshKeyId} [delete]
func DeleteMigratedSSHKey(c echo.Context) error {
	// Path pattern that captures two path parameters
	sourcePattern := "/migration/ns/*/resources/sshKey/*"
	// First * is used as $1(nsId), second * as $2(sshKeyId)
	targetPattern := "/ns/$1/resources/sshKey/$2"

	proxyHandler := createTumblebugProxyHandler(sourcePattern, targetPattern)
	return proxyHandler(c)
}

// ========== Security Group Resource APIs ==========
// ListMigratedSecurityGroups godoc
// @ID ListMigratedSecurityGroups
// @Summary List all migrated security groups
// @Description Get the list of all migrated security groups in the namespace
// @Tags [Migration] Resources for VM infrastructure
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(default)
// @Param option query string false "Option" Enums(id)
// @Param filterKey query string false "Field key for filtering (ex: systemLabel)"
// @Param filterVal query string false "Field value for filtering (ex: Registered from CSP resource)"
// @Success 200 {object} JSONResult{[DEFAULT]=tbresource.RestGetAllSecurityGroupResponse,[ID]=tbmodel.IdList} "Different return structures by the given option param"
// @Failure 404 {object} tbmodel.SimpleMsg
// @Failure 500 {object} tbmodel.SimpleMsg
// @Router /migration/ns/{nsId}/resources/securityGroup [get]
func ListMigratedSecurityGroups(c echo.Context) error {
	// Source path pattern with * to capture nsId
	sourcePattern := "/migration/ns/*/resources/securityGroup"
	// Target path pattern using $1 for captured nsId
	targetPattern := "/ns/$1/resources/securityGroup"

	proxyHandler := createTumblebugProxyHandler(sourcePattern, targetPattern)
	return proxyHandler(c)
}

// GetMigratedSecurityGroup godoc
// @ID GetMigratedSecurityGroup
// @Summary Get a specific migrated security group
// @Description Get details of a specific migrated security group in the namespace
// @Tags [Migration] Resources for VM infrastructure
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID"
// @Param sgId path string true "Security Group ID"
// @Success 200 {object} tbmodel.TbSecurityGroupInfo
// @Failure 404 {object} tbmodel.SimpleMsg
// @Failure 500 {object} tbmodel.SimpleMsg
// @Router /migration/ns/{nsId}/resources/securityGroup/{sgId} [get]
func GetMigratedSecurityGroup(c echo.Context) error {
	// Path pattern that captures two path parameters
	sourcePattern := "/migration/ns/*/resources/securityGroup/*"
	// First * is used as $1(nsId), second * as $2(sgId)
	targetPattern := "/ns/$1/resources/securityGroup/$2"

	proxyHandler := createTumblebugProxyHandler(sourcePattern, targetPattern)
	return proxyHandler(c)
}

// CreateMigratedSecurityGroup godoc
// @ID CreateMigratedSecurityGroup
// @Summary Create a migrated security group
// @Description Create a new migrated security group in the namespace
// @Tags [Migration] Resources for VM infrastructure
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(default)
// @Param option query string false "Option: [required params for register] connectionName, name, vNetId, cspResourceId" Enums(register)
// @Param securityGroupReq body tbmodel.TbSecurityGroupReq true "Details for an securityGroup object"
// @Success 200 {object} tbmodel.TbSecurityGroupInfo
// @Failure 404 {object} tbmodel.SimpleMsg
// @Failure 500 {object} tbmodel.SimpleMsg
// @Router /migration/ns/{nsId}/resources/securityGroup [post]
func CreateMigratedSecurityGroup(c echo.Context) error {
	// Source path pattern with * to capture nsId
	sourcePattern := "/migration/ns/*/resources/securityGroup"
	// Target path pattern using $1 for captured nsId
	targetPattern := "/ns/$1/resources/securityGroup"

	proxyHandler := createTumblebugProxyHandler(sourcePattern, targetPattern)
	return proxyHandler(c)
}

// DeleteMigratedSecurityGroup godoc
// @ID DeleteMigratedSecurityGroup
// @Summary Delete a migrated security group
// @Description Delete a specific migrated security group in the namespace
// @Tags [Migration] Resources for VM infrastructure
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(default)
// @Param securityGroupId path string true "Security Group ID"
// @Success 200 {object} tbmodel.SimpleMsg
// @Failure 404 {object} tbmodel.SimpleMsg
// @Failure 500 {object} tbmodel.SimpleMsg
// @Router /migration/ns/{nsId}/resources/securityGroup/{sgId} [delete]
func DeleteMigratedSecurityGroup(c echo.Context) error {
	// Path pattern that captures two path parameters
	sourcePattern := "/migration/ns/*/resources/securityGroup/*"
	// First * is used as $1(nsId), second * as $2(securityGroupId)
	targetPattern := "/ns/$1/resources/securityGroup/$2"

	proxyHandler := createTumblebugProxyHandler(sourcePattern, targetPattern)
	return proxyHandler(c)
}
