default:
	cd cmd/cm-beetle && $(MAKE)

cc:
	cd cmd/cm-beetle && $(MAKE)

run:
	cd cmd/cm-beetle && $(MAKE) run

runwithport:
	cd cmd/cm-beetle && $(MAKE) runwithport --port=$(PORT)

clean:
	cd cmd/cm-beetle && $(MAKE) clean

prod:
	cd cmd/cm-beetle && $(MAKE) prod

onprem-model:
	cd pkg/api/rest/model && $(MAKE) onprem-model

swag swagger:
	cd pkg/ && $(MAKE) swag
