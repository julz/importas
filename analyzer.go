package importas

import (
	"flag"
	"go/ast"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "importas",
	Doc:  "Enforces consistent import aliases",
	Run:  run,

	Flags: flags(),

	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

type Config struct {
	RequiredAlias map[string]string // path -> alias.
}

var config = Config{
	RequiredAlias: make(map[string]string),
}

func flags() flag.FlagSet {
	fs := flag.FlagSet{}
	fs.Var(stringMap(config.RequiredAlias), "alias", "required import alias in form path:alias")
	return fs
}

func run(pass *analysis.Pass) (interface{}, error) {
	return runWithConfig(config, pass)
}

func runWithConfig(config Config, pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder([]ast.Node{(*ast.ImportSpec)(nil)}, func(n ast.Node) {
		node := n.(*ast.ImportSpec)

		if node.Name == nil {
			return // not aliased at all, ignore. (Maybe add strict mode for this?).
		}

		alias := node.Name.String()
		if alias == "." {
			return // Dot aliases are generally used in tests, so ignore.
		}

		path, err := strconv.Unquote(node.Path.Value)
		if err != nil {
			pass.Reportf(node.Pos(), "import not quoted")
		}

		if required, exists := config.RequiredAlias[path]; exists && required != alias {
			pass.Reportf(node.Pos(), "import %q imported as %q but must be %q according to config", path, alias, required)
		}
	})

	return nil, nil
}
