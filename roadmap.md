# Excelante

Excelante is an Excel API that allow user to create or to read some excel file.

### Roadmap
- [V1.0]
  - You can create an excel file with multiple sheets
  - You can create an excel table with a json input file
    - The json input file can take some optionnal value, that correspond to style applied to excel table.
    - List of style must be present in readme
  - You can read an excel file and extract columns values
   - This service take an excel file and a list of fileds (or * ) to extract in input.
   - Json file with the values is return by the service.
  - Optimisation & best practices
   - Using goroutine to optimise the treatment (I can create a table with more than 50,000 lines in few seconds, or i can read an excel file with 50,000 lines in few seconds)
   - Using all the best practices to have a good & stable starter solution, that can be used in production environment
   - Set up deployment options
     - OpenFaas
     - Docker ready


- [V2.0]
  - Create excel graph with json input file
  - Create pivot table with json input file
  - Change the extract return format for excel reading
  - Add formula to cell when excel table is created
  - Online website to create/read excel file using Excelante API
