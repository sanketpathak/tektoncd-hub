go test $(go list -f '{{ .ImportPath }} {{ .TestImports }}' ./... | grep gotest.tools/v3/golden | awk '{print $1}' | tr '\n' ' ') -test.update-golden=true