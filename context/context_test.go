package context

import (
	"testing"

	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/token"
)

func TestContext(t *testing.T) {
	t.Run("ScenarioFilepath", func(t *testing.T) {
		path := "test.yaml"
		ctx := FromT(t).WithScenarioFilepath(path)
		if got := ctx.ScenarioFilepath(); got != path {
			t.Errorf("expect %q but got %q", path, got)
		}
	})
	t.Run("node", func(t *testing.T) {
		ctx := FromT(t)
		node := ast.String(token.String("", "", nil))
		ctx = ctx.WithNode(node)
		if ctx.Node() != node {
			t.Fatal("failed to get node")
		}
	})
	t.Run("enabledColor", func(t *testing.T) {
		ctx := FromT(t)
		ctx = ctx.WithEnabledColor(true)
		if !ctx.EnabledColor() {
			t.Fatal("failed to get enabledColor")
		}
	})
}
