package make

import (
	"embed"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"gohub/pkg/console"
	"gohub/pkg/file"
	"gohub/pkg/str"
	"strings"
)

// Model 参数解释
// 单个词，用户命令传参，以 User 模型为例：
//     - user
//  - User
//  - users
//  - Users
// 整理好的数据：
// {
//     "TableName": "users",
//     "StructName": "User",
//     "StructNamePlural": "Users"
//     "VariableName": "user",
//     "VariableNamePlural": "users",
//     "PackageName": "user"
// }
// -
// 两个词或者以上，用户命令传参，以 TopicComment 模型为例：
//     - topic_comment
//  - topic_comments
//  - TopicComment
//  - TopicComments
// 整理好的数据：
// {
//     "TableName": "topic_comments",
//     "StructName": "TopicComment",
//     "StructNamePlural": "TopicComments"
//     "VariableName": "topicComment",
//     "VariableNamePlural": "topicComments",
//     "PackageName": "topic_comment"
// }

type Model struct {
	TableName          string
	StructName         string
	StructNamePlural   string
	VariableName       string
	VariableNamePlural string
	PackageName        string
}

// stubsFS package .stub files

//go:embed stubs
var stubsFS embed.FS

var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	// Registration Make sub commands
	CmdMake.AddCommand(
		CmdMakeCMD,
		CmdMakeModel,
		CmdMakeAPIController,
	)
}

// makeModelFromString format user input
func makeModelFromString(name string) Model {
	model := Model{}
	model.StructName = str.Singular(strcase.ToCamel(name))
	model.StructNamePlural = str.Plural(model.StructName)
	model.TableName = str.Snake(model.StructNamePlural)
	model.VariableName = str.LowerCamel(model.StructName)
	model.PackageName = str.Snake(model.StructName)
	model.VariableNamePlural = str.LowerCamel(model.StructNamePlural)
	return model
}

func createFileFromStub(filePath string, stubName string, model Model, variables ...interface{}) {
	// Implement last parameter
	replaces := make(map[string]string)
	if len(variables) > 0 {
		replaces = variables[0].(map[string]string)
	}

	// Check target file path exists
	if file.Exists(filePath) {
		console.Exit(filePath + " already exists!")
	}

	// Read stub template file
	modelData, err := stubsFS.ReadFile("stubs/" + stubName + ".stub")
	if err != nil {
		console.Exit(err.Error())
	}

	modelStub := string(modelData)

	// Add default replace variables
	replaces["{{VariableName}}"] = model.VariableName
	replaces["{{VariableNamePlural}}"] = model.VariableNamePlural
	replaces["{{StructName}}"] = model.StructName
	replaces["{{StructNamePlural}}"] = model.StructNamePlural
	replaces["{{PackageName}}"] = model.PackageName
	replaces["{{TableName}}"] = model.TableName

	// Replace template variables
	for search, replace := range replaces {
		modelStub = strings.ReplaceAll(modelStub, search, replace)
	}

	// Save to target file
	err = file.Put([]byte(modelStub), filePath)
	if err != nil {
		console.Exit(err.Error())
	}

	// Output success
	console.Success(fmt.Sprintf("[%s] created.", filePath))
}
