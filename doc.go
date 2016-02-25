// Package octicons provides Octicons.
//
// It's a Go package that statically embeds Octicons data, making it available via an http.FileSystem.
package octicons

//go:generate curl -L -O https://github.com/github/octicons/releases/download/v3.5.0/octicons.zip
//go:generate unzip octicons.zip -d octicons
//go:generate rm -f octicons.zip octicons/README.md octicons/octicons-local.ttf octicons/sprockets-octicons.scss
//go:generate goexec -quiet "err := vfsgen.Generate(http.Dir(\"octicons\"), vfsgen.Options{PackageName: \"octicons\", VariableName: \"Assets\", VariableComment: \"Assets provides Octicons data.\"}); if err != nil { log.Fatalln(err) }"
//go:generate rm -rf octicons
