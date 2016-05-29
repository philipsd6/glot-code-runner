package language

import (
	"./assembly"
	"./ats"
	"./bash"
	"./c"
	"./clojure"
	"./coffeescript"
	"./cpp"
	"./csharp"
	"./d"
	"./elixir"
	"./elm"
	"./erlang"
	"./fsharp"
	"./golang"
	"./groovy"
	"./haskell"
	"./idris"
	"./java"
	"./javascript"
	"./julia"
	"./lua"
	"./nim"
	"./ocaml"
	"./perl"
	"./perl6"
	"./php"
	"./python"
	"./ruby"
	"./rust"
	"./scala"
	"./swift"
	"./zsh"
)

type runFn func([]string, string) (string, string, error)

var languages = map[string]runFn{
	"assembly":     assembly.Run,
	"ats":          ats.Run,
	"bash":         bash.Run,
	"c":            c.Run,
	"clojure":      clojure.Run,
	"coffeescript": coffeescript.Run,
	"csharp":       csharp.Run,
	"d":            d.Run,
	"elixir":       elixir.Run,
	"elm":          elm.Run,
	"cpp":          cpp.Run,
	"erlang":       erlang.Run,
	"fsharp":       fsharp.Run,
	"haskell":      haskell.Run,
	"idris":        idris.Run,
	"go":           golang.Run,
	"groovy":       groovy.Run,
	"java":         java.Run,
	"javascript":   javascript.Run,
	"julia":        julia.Run,
	"lua":          lua.Run,
	"nim":          nim.Run,
	"ocaml":        ocaml.Run,
	"perl":         perl.Run,
	"perl6":        perl6.Run,
	"php":          php.Run,
	"python":       python.Run,
	"ruby":         ruby.Run,
	"rust":         rust.Run,
	"scala":        scala.Run,
	"swift":        swift.Run,
	"zsh":          zsh.Run,
}

func IsSupported(lang string) bool {
	_, supported := languages[lang]
	return supported
}

func Run(lang string, files []string, stdin string) (string, string, error) {
	return languages[lang](files, stdin)
}
