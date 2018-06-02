// Package octicons provides GitHub Octicons.
//
// It's a Go package that statically embeds GitHub Octicons data, exposing it via an http.FileSystem.
//
// Deprecated: GitHub has deprecated the icon font version of Octicons in favor of SVG.
// See https://github.com/primer/octicons/issues/108 and https://github.com/blog/2112-delivering-octicons-with-svg.
// Use github.com/shurcooL/octicon instead, which provides Octicons in SVG format.
package octicons

//go:generate curl -L -o octicons.zip https://github.com/primer/octicons/archive/v4.3.0.zip
//go:generate unzip octicons.zip octicons-4.3.0/build/font/octicons.* -d octicons
//go:generate goexec -quiet "err := vfsgen.Generate(http.Dir(\"octicons/octicons-4.3.0/build/font\"), vfsgen.Options{PackageName: \"octicons\", VariableName: \"Assets\", VariableComment: \"Assets provides GitHub Octicons data.\"}); if err != nil { log.Fatalln(err) }"
//go:generate rm -rf octicons.zip octicons
