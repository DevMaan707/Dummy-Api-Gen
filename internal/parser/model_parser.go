package parser

import (
	"go/ast"
	"go/parser"
	"go/token"

	validator "github.com/DevMaan707/dummy-api-gen/internal/generator"
	"github.com/DevMaan707/dummy-api-gen/internal/shared"
)

func ParseModels(path string) ([]shared.ModelData, error) {
	var models []shared.ModelData
	fs := token.NewFileSet()
	pkgs, err := parser.ParseDir(fs, path, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				genDecl, ok := decl.(*ast.GenDecl)
				if !ok {
					continue
				}
				for _, spec := range genDecl.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok || !validator.ValidateResponseModelName(typeSpec.Name.Name) {
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

					models = append(models, shared.ModelData{
						Name:   typeSpec.Name.Name,
						Fields: fields,
					})
				}
			}
		}
	}
	return models, nil
}
