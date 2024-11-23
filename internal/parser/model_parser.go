package parser

import (
    "go/ast"
    "go/parser"
    "go/token"
    "strings"
)


func ParseModels(path string) ([]string, error) {
    var models []string
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
					if !ok || !ValidateResponseModelName(typeSpec.Name.Name) {
						continue
					}
					models = append(models, typeSpec.Name.Name)
				}
				
            }
        }
    }
    return models, nil
}
