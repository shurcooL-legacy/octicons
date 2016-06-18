// Package octicons provides Octicons.
//
// It's a Go package that statically embeds Octicons data, making it available via an http.FileSystem.
package octicons

//go:generate curl -L -o octicons.zip https://github.com/primer/octicons/archive/v4.2.0.zip
//go:generate unzip octicons.zip octicons-4.2.0/build/font/* -d octicons
//go:generate goexec -quiet "err := vfsgen.Generate(http.Dir(\"octicons/octicons-4.2.0/build/font\"), vfsgen.Options{PackageName: \"octicons\", VariableName: \"Assets\", VariableComment: \"Assets provides Octicons data.\"}); if err != nil { log.Fatalln(err) }"
//go:generate rm -rf octicons.zip octicons
