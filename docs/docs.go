// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-07-21 18:19:03.3502139 +0200 CEST m=+0.407741701

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/read/columns": {
            "post": {
                "description": "transform given Excel file to JSON",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Readers"
                ],
                "summary": "Read Excel file column by column",
                "parameters": [
                    {
                        "description": "The Excel file to convert",
                        "name": "file",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.Output"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/read/lines": {
            "post": {
                "description": "transform given Excel file to JSON",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Readers"
                ],
                "summary": "Read Excel file line by line",
                "parameters": [
                    {
                        "description": "The Excel file to convert",
                        "name": "file",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.Output"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/write": {
            "post": {
                "description": "transform JSON to Excel file",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "Writers"
                ],
                "summary": "Write Excel file",
                "parameters": [
                    {
                        "description": "Items you want to convert into an Excel file",
                        "name": "file",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.File"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Alignment": {
            "type": "object",
            "properties": {
                "horizontal": {
                    "type": "string"
                },
                "shrink_to_fit": {
                    "type": "boolean"
                },
                "vertical": {
                    "type": "string"
                },
                "wrap_text": {
                    "type": "boolean"
                }
            }
        },
        "models.Border": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "style": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.Cell": {
            "type": "object",
            "properties": {
                "style": {
                    "type": "object",
                    "$ref": "#/definitions/models.Style"
                },
                "value": {
                    "type": "object"
                }
            }
        },
        "models.Chart": {
            "type": "object",
            "properties": {
                "dimension": {
                    "type": "object",
                    "$ref": "#/definitions/models.Dimension"
                },
                "format": {
                    "type": "object",
                    "$ref": "#/definitions/models.Format"
                },
                "legend": {
                    "type": "object",
                    "$ref": "#/definitions/models.Legend"
                },
                "plot_area": {
                    "type": "object",
                    "$ref": "#/definitions/models.PlotArea"
                },
                "series": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Series"
                    }
                },
                "title": {
                    "type": "object",
                    "$ref": "#/definitions/models.Title"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.ColOrRowValues": {
            "type": "object",
            "properties": {
                "cells": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Cell"
                    }
                },
                "orientation": {
                    "type": "string"
                }
            }
        },
        "models.Dimension": {
            "type": "object",
            "properties": {
                "height": {
                    "type": "integer"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "models.File": {
            "type": "object",
            "properties": {
                "sheets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Sheet"
                    }
                }
            }
        },
        "models.Fill": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "pattern": {
                    "type": "integer"
                },
                "shading": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.Font": {
            "type": "object",
            "properties": {
                "bold": {
                    "type": "boolean"
                },
                "color": {
                    "type": "string"
                },
                "family": {
                    "type": "string"
                },
                "italic": {
                    "type": "boolean"
                },
                "size": {
                    "type": "number"
                },
                "strike": {
                    "type": "boolean"
                },
                "underline": {
                    "type": "string"
                }
            }
        },
        "models.Format": {
            "type": "object",
            "properties": {
                "lock_aspect_ratio": {
                    "type": "boolean"
                },
                "locked": {
                    "type": "boolean"
                },
                "print_obj": {
                    "type": "boolean"
                },
                "x_offset": {
                    "type": "number"
                },
                "x_scale": {
                    "type": "number"
                },
                "y_offset": {
                    "type": "number"
                },
                "y_scale": {
                    "type": "number"
                }
            }
        },
        "models.Item": {
            "type": "object",
            "properties": {
                "chart": {
                    "type": "object",
                    "$ref": "#/definitions/models.Chart"
                },
                "starting_cell_coordinates": {
                    "type": "string"
                },
                "table": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ColOrRowValues"
                    }
                }
            }
        },
        "models.Legend": {
            "type": "object",
            "properties": {
                "position": {
                    "type": "string"
                },
                "show_legend_key": {
                    "type": "boolean"
                }
            }
        },
        "models.PlotArea": {
            "type": "object",
            "properties": {
                "show_bubble_size": {
                    "type": "boolean"
                },
                "show_cat_name": {
                    "type": "boolean"
                },
                "show_leader_lines": {
                    "type": "boolean"
                },
                "show_percent": {
                    "type": "boolean"
                },
                "show_series_name": {
                    "type": "boolean"
                },
                "show_val": {
                    "type": "boolean"
                }
            }
        },
        "models.Protection": {
            "type": "object",
            "properties": {
                "hidden": {
                    "type": "boolean"
                },
                "locked": {
                    "type": "boolean"
                }
            }
        },
        "models.Series": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "values": {
                    "type": "string"
                }
            }
        },
        "models.Sheet": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Item"
                    }
                },
                "name": {
                    "type": "string"
                },
                "orientation": {
                    "description": "Default : landscape",
                    "type": "string"
                }
            }
        },
        "models.Style": {
            "type": "object",
            "properties": {
                "alignment": {
                    "type": "object",
                    "$ref": "#/definitions/models.Alignment"
                },
                "border": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Border"
                    }
                },
                "fill": {
                    "type": "object",
                    "$ref": "#/definitions/models.Fill"
                },
                "font": {
                    "type": "object",
                    "$ref": "#/definitions/models.Font"
                },
                "protection": {
                    "type": "object",
                    "$ref": "#/definitions/models.Protection"
                }
            }
        },
        "models.Title": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "services.Output": {
            "type": "object",
            "additionalProperties": {
                "type": "object",
                "additionalProperties": true
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Excelante API",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
