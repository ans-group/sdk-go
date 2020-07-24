package helper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ukfast/sdk-go/pkg/gen/helper"
)

func TestGenerateImports(t *testing.T) {
	t.Run("NoImport_EmptyString", func(t *testing.T) {
		result := helper.GenerateImports([]string{})
		assert.Equal(t, "", result)
	})

	t.Run("SingleImport_ExpectedString", func(t *testing.T) {
		result := helper.GenerateImports([]string{"gitlab.com/ukfast/sdk-go/test1"})
		assert.Equal(t, "import \"gitlab.com/ukfast/sdk-go/test1\"", result)
	})

	t.Run("MultipleImports_ExpectedString", func(t *testing.T) {
		result := helper.GenerateImports([]string{"gitlab.com/ukfast/sdk-go/test1", "gitlab.com/ukfast/sdk-go/test2"})
		expected := `import (
	"gitlab.com/ukfast/sdk-go/test1"
	"gitlab.com/ukfast/sdk-go/test2"
)`

		assert.Equal(t, expected, result)
	})
}
