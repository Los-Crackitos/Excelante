# Excelante

![Build Status](https://github.com/Los-Crackitos/Excelante/workflows/.github/workflows/main.yml/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Los-Crackitos/Excelante)](https://goreportcard.com/report/github.com/Los-Crackitos/Excelante)
![Code Size](https://img.shields.io/github/languages/code-size/Los-Crackitos/Excelante)
[![Go Doc](https://godoc.org/github.com/Los-Crackitos/Excelante?status.svg)](https://godoc.org/github.com/Los-Crackitos/Excelante)
[![License](https://img.shields.io/github/license/Los-Crackitos/Excelante)](LICENSE)

Excelante is an API used to read and write Excel files through HTTP requests.

# Table of Contents
1. [Description](#Description)
2. [Write an Excel File](#Write-an-excel-file)
  2.1 [Json scheme](#Json-scheme)
  2.2 [Style references](#Style-references)
    2.2.1 [Border references](#Border-references)
    2.2.2 [Fill pattern references](#Fill-pattern-references)
    2.2.3 [Font references](#Font-references)
    2.2.4 [Alignment references](#Alignment-references)
3. [Read an excel file](#Read-an-excel-file)
4. [Swagger API Documentation](#Swagger-API-Documentation)
5. [Clients Examples](#Clients-Examples)
  5.1 [JavaScript](#JavaScript)
6. [Want to contribute ?](#Want-to-contribute-?)
7. [Contributors](#Contributors)
8. [Sponsors](#Sponsors)
9. [License](#License)

# Description

Excelante is a result of a simple problem : how is the best way to create or to read some excel file? Many times, the answer is to use a library in your project language. But, it is the best way ? Do you have all performances you need ?
Excelante provide an HTTP API, optimised system based on Golang. You can use it in any language you want.
Our API can write excel file with only a json input, or read all or only specified column in an input excel file.

Try it now : link excelante-front

## Benchmark

To do


# Write an excel file

To create an excel file with Excelante, you need to call the API with POST method and pass the following json scheme to the route.
Some fiedls are optionnal, feel free to use it or not.

## Json scheme

```
{
  "sheets": [{                                // Array of sheets
    "name": "Sheet1",                         // Name of the sheet
    "orientation": "OrientationLanscape",     // (Optionnal) Is the orientation of layout. (OrientationPortrait || OrientationLanscape) Default is set to OrientationPortrait.
    "items": [{                               // Array of items. An item is an object that can be write into excel file.
      "starting_cell_coordonates": "B3",      // (Optionnal) Represent the begining coordonates of the item.
      "tables": [{                            // Array of data. Represent all data of an excel table that can be write into file. An object of tables is write by column or by row
        "orientation": "column",              // Represent the mode of writing. (column || row)
        "cells": [{                           // Represent an array of cells to wwrite
          "value": "myValue",                 // Value of the cell
          "style": {                          // (Optionnal) Describe the style of the cell
            "border": [{                      // (Optionnal) Describe all the border of the cell.
              "type": "top",                  // (Optionnal) Type is the position of border. Can be (top || bottom || left || right)
              "color": "#000000",             // (Optionnal) Color is the border color.
              "style":1                       // (Optionnal) Style is the border format (See border style reference)
            }],
            "fill": {                          // (Optionnal) Represent cell fill
              "type": "pattern",              // (Optionnal) Type of fill
              "color": ["#87CEFA"],           // (Optionnal) Fill color
              "pattern": 1,                   // (Optionnal) Fill pattern (See fill style reference)
              "shading":1                     // (Optionnal) Shading fill (See fill style reference)
            },
            "font": {                         // (Optionnal) Represent cell text font
              "bold": true,                   // (Optionnal) Bold (true) or not (false)
              "italic": true,                 // (Optionnal) Italic (true) or not (false)
              "underline": "single",          // (Optionnal) Underline text (See font style reference)
              "family": "Times New Roman",    // (Optionnal) Family of font
              "size": 14.5,                   // (Optionnal) Size of font
              "strike": true,                 // (Optionnal) Strike (true) or not (false)
              "color": "#000000"              // (Optionnal) Color of font
            },
            "alignment": {                    // (Optionnal) Represent cell content aligment
              "horizontal":"center",          // (Optionnal) Represent horizontal alignment into the cell (See alignment style reference)
              "vertical": "top",              // (Optionnal) Represent vertical alignment into the cell (See alignment style reference)
              "shrink_to_fit": true,           // (Optionnal) Shrink to fit cell text
              "wrap_text": false              // (Optionnal) Wrap cell text
            },
            "protection": {                   // (Optionnal) Represent cell protection
              "hidden": 0,                    // (Optionnal) Cell is hidden (true) or not (false)
              "locked": 0                     // (Optionnal) Cell is locked (true) or not (false)
            }
          }
        }]
      }],
      "graph": {                              // (Optionnal) Represent a graph item
        "type": "bar",                        // Type of chart (See graph type reference)
        "series": [{                          // Array of data to draw the chart.
            "name": "Sheet1!$A$2",            // Name of the serie
            "categories": "Sheet1!$A$2",      // Category of the serie
            "values": "Sheet1!$B$2:$D$2"      // Value of the serie
        }],
        "format":{
            "x_scale": 1.0,                   // Set x scale for chart
            "y_scale": 1.0,                   // Set y scale for chart
            "x_offset": 15,                   // Set x offset for chart
            "y_offset": 10,                   // Set y offset for chart
            "print_obj": true,                //
            "lock_aspect_ratio": false,       // Lock x/y ratio
            "locked": false                   // Lock chart
        },
        "legend":{                            // (Optionnal) Legend of the graph
            "position": "left",               // Legend position (See legend graph reference)
            "show_legend_key": false          // Show legend but not overlap with chart
        },
        "title":{                             // Title of the graph
            "name": "Chart of my data"        // Name of chart
        },
        "dimension": {                        // Set graph dimension
            "height": 290,                    // Set height. Default 290
            "width": 480                      // Set width. Default 480
        }
        "plotarea":{                          // Set the position of the chart plot area
            "show_bubble_size": true,         // Specifies the bubble size shall be shown in a data label. Default false.
            "show_cat_name": false,           // Show category name. Default true.
            "show_leader_lines": false,       // Specifies that the category name shall be shown in the data label. Default false.
            "show_percent": true,             // Specifies that the percentage shall be shown in a data label. Default false.
            "show_series_name": true,         // Specifies that the series name shall be shown in a data label. Default false.
            "show_val": true                  // Specifies that the value shall be shown in a data label. Default false.
        }
      }
    }]
  }, {
    "name": "Sheet2",
    ...
  }]
}


```

## Style references

### Border references
```
"border": [{
    "type": "top"
          /**
            * top
            * bottom
            * left
            * right
          */
    "color": "#000000",
    "style": 0,
            /**
              * 0 : Nothing
              * 1 : Continuous
              * 2 : Continuous with double lines
              * 3 : Dash
              * 4 : Dot
              * 5 : Continuous with 3 lines
              * 6 : Double
              * 7 : Continuous with no line
              * 8 : Dash with 2 lines
              * 9 : Dash Dot
              * 10 : Dash Dot with 2 lines
              * 11 : Dash Dot Dot
              * 12 : Dash Dot Dot with 2 lines
              * 13 : SlantDash Dot
            */
  }]
```

### Fill Pattern reference

To do

### Font reference

```
"font": {
  "bold": true,
  "italic": true,
  "underline": "single",
              /**
              * single
              * double
              */
  "family": "Times New Roman",
  "size": 14.5,
  "strike": true,
            /** false */
  "color": ""
}
```

### Alignment reference

```
"alignment": {
  "horizontal":"center",
                /**
                * left
                * center
                * right
                * fill
                * justify
                * centerContinous
                * distributed
                */
  "vertical": "top",
                /**
                * top
                * center
                * justify
                * distributed
                */
  "shrink_to_fit": true,
                /** false */
  "wrap_text": false
                /** true */
}

```

# Graph references

## Graph type references

```

"graph": {
  "type": "bar",
          /**
          * areaStacked || areaPercentStacked || area3D || area3DStacked || area3DPercentStacked
          * bar || barStacked || barPercentStacked || bar3DClustered || bar3DStacked || bar3DPercentStacked || bar3DConeClustered || bar3DConeStacked || bar3DConePercentStacked || bar3DPyramidClustered || bar3DPyramidStacked ||
          * bar3DPyramidPercentStacked || bar3DCylinderClustered || bar3DCylinderStacked || bar3DCylinderPercentStacked
          * col || colStacked || colPercentStacked || col3DClustered || col3D || col3DStacked || col3DPercentStacked || col3DCone || col3DConeClustered || col3DConeStacked || col3DConePercentStacked || col3DPyramid ||
          * col3DPyramidClustered || col3DPyramidStacked || col3DPyramidPercentStacked || col3DCylinder || col3DCylinderClustered || col3DCylinderStacked || col3DCylinderPercentStacked
          * doughnut
          * line
          * pie || pie3D
          * radar
          * scatter
          * surface3D || wireframeSurface3D || wireframeContour
          * contour
          * bubble || bubble3D
          */
  "series": [{}]
}


```

## Legend references

```
"legend":{
  "position": "left",
              /**
              * left
              * right
              * top
              * bottom
              * top_right
              */
  "show_legend_key": false
              /** true */
}
```

# Read an excel file

# Swagger API Documentation

Swagger API documentation can be found at `/api/v1/swagger/`

# Clients Examples

To do

## JavaScript

# Want to contribute ?

As contributors and maintainers of this project, we pledge to respect all people who contribute through reporting issues, posting feature requests, updating documentation, submitting pull requests or patches, and other activities.

We are committed to making participation in this project a harassment-free experience for everyone, regardless of level of experience, gender, gender identity and expression, sexual orientation, disability, personal appearance, body size, race, ethnicity, age, or religion.

Examples of unacceptable behavior by participants include the use of sexual language or imagery, derogatory comments or personal attacks, trolling, public or private harassment, insults, or other unprofessional conduct.

Project maintainers have the right and responsibility to remove, edit, or reject comments, commits, code, wiki edits, issues, and other contributions that are not aligned to this Code of Conduct. Project maintainers who do not follow the Code of Conduct may be removed from the project team.

Instances of abusive, harassing, or otherwise unacceptable behavior may be reported by opening an issue or contacting one or more of the project maintainers.

This Code of Conduct is adapted from the Contributor Covenant, version 1.0.0, available at https://www.contributor-covenant.org/version/1/0/0/code-of-conduct.html

# Contributors


# Sponsors

If you want to support us in our open source work :)

# License

Excelante is under MIT liense. Feel free to use it as u want.
