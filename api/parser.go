package api

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/DevMaan707/dummy-api-gen/shared"
)

func ParseModels(path string) ([]shared.ModelData, error) {
	fs := token.NewFileSet()
	pkgs, err := parser.ParseDir(fs, path, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	var models []shared.ModelData
	modelMap := make(map[string]*shared.ModelData)

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				genDecl, ok := decl.(*ast.GenDecl)
				if !ok {
					continue
				}
				for _, spec := range genDecl.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}

					modelName := typeSpec.Name.Name
					isRequestModel := strings.HasSuffix(modelName, "RequestModel")
					isResponseModel := strings.HasSuffix(modelName, "ResponseModel")
					if !isRequestModel && !isResponseModel {
						continue
					}

					structType, ok := typeSpec.Type.(*ast.StructType)
					if !ok {
						continue
					}

					fields := make(map[string]string)
					for _, field := range structType.Fields.List {
						fieldType := ""
						if ident, ok := field.Type.(*ast.Ident); ok {
							fieldType = ident.Name
						}
						for _, name := range field.Names {
							fields[name.Name] = fieldType
						}
					}

					baseName := strings.TrimSuffix(modelName, "RequestModel")
					baseName = strings.TrimSuffix(baseName, "ResponseModel")

					if modelMap[baseName] == nil {
						modelMap[baseName] = &shared.ModelData{Name: baseName}
					}

					if isRequestModel {
						modelMap[baseName].RequestFields = fields
					}
					if isResponseModel {
						modelMap[baseName].ResponseFields = fields
					}
				}
			}
		}
	}

	for _, model := range modelMap {
		models = append(models, *model)
	}

	return models, nil
}
