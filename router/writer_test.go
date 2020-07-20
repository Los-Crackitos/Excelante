package router

import (
	"net/http"
	"testing"
)

func TestPostFile(t *testing.T) {
	response := executeRequest(
		createTestRequest(
			t,
			"POST",
			"/api/v1/write",
			[]byte(`{
				"sheets": [{
					"name": "Sheet toto",
					"orientation": "OrientationPortrait",
					"items": [{
						"mode": "table",
						"startingCellCoordinates": "A1",
						"tables": [{
							"orientation": "column",
							"name": "Product",
							"style": {},
							"cells": [{
								"value": "Table",
								"style": {}
							}, {
								"value": "Chair"
							}, {
								"value": "Cut"
							}, {
								"value": "Glass"
							}]
						}, {
							"orientation": "column",
							"name": "Price",
							"style": {},
							"cells": [{
								"value": "150,00"
							}, {
								"value": "40,00"
							}, {
								"value": "1,50"
							}, {
								"value": "2"
							}]
						}]
					}]
				}]
			}`),
		),
	)

	checkResponseCode(t, http.StatusCreated, response.Code)
}
