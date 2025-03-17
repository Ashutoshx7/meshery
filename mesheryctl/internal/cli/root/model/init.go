package model

import (
	"os"
	"strings"

	"github.com/layer5io/meshery/mesheryctl/internal/cli/root/config"
	"github.com/layer5io/meshery/mesheryctl/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// TODO must use schemas repo instead of local temp package
	// "github.com/meshery/schemas"
	schemas "github.com/layer5io/meshery/mesheryctl/internal/cli/root/model/temp_schemas"
)

var initModelCmd = &cobra.Command{
	Use:   "init [model-name]",
	Short: "generates scaffolding for convenient model creation",
	Long:  "generates a folder structure and guides user on model creation",
	Example: `
// generates a folder structure
mesheryctl model init [model-name]

// generates a folder structure and sets up model version
mesheryctl model init [model-name] --version [version] (default is 0.1.0)

// generates a folder structure under specified path
mesheryctl model init [model-name] --path [path-to-location] (default is current folder)

// generate a folder structure in json format
mesheryctl model init [model-name] --output-format [json|yaml|csv] (default is json)
    `,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := config.GetMesheryCtl(viper.GetViper())
		if err != nil {
			utils.Log.Error(err)
			// TODO use meshkit error format instead during implementation phase
			return err
		}

		modelName := args[0]
		path, _ := cmd.Flags().GetString("path")
		version, _ := cmd.Flags().GetString("version")
		outputFormat, _ := cmd.Flags().GetString("output-format")

		utils.Log.Info("init command will be here soon")
		utils.Log.Infof("model name = %s", modelName)
		utils.Log.Infof("path = %s", path)
		utils.Log.Infof("version = %s", version)
		utils.Log.Infof("output format = %s", outputFormat)

		for _, templatePath := range []string{
			templatePathModelJSON,
			templatePathDesignJSON,
			templatePathComponentJSON,
			templatePathConnectionJSON,
			templatePathRelathionshipJSON,
		} {
			// modelJSONContent, err := readTemplate(templatePath)
			_, err := readTemplate(templatePath)
			if err != nil {
				utils.Log.Error(err)
				// TODO use meshkit error format instead during implementation phase
				return err
			}

			//utils.Log.Debug(string(modelJSONContent))
		}

		folderPath := strings.Join([]string{path, modelName, version}, string(os.PathSeparator))
		err = os.MkdirAll(folderPath, 0755)
		if err != nil {
			// TODO use meshkit error format
			utils.Log.Error(err)
			return err
		}

		return nil
	},
}

func init() {
	initModelCmd.Flags().StringP("path", "p", ".", "(optional) target directory (default: current dir)")
	initModelCmd.Flags().StringP("version", "", "0.1.0", "(optional) model version (default: 0.1.0)")
	initModelCmd.Flags().StringP("output-format", "o", "json", "(optional) format to display in [json|yaml]")
}

const templatePathModelJSON = "json_models/constructs/v1beta1/model.json"
const templatePathDesignJSON = "json_models/constructs/v1beta1/design.json"
const templatePathComponentJSON = "json_models/constructs/v1beta1/component.json"
const templatePathConnectionJSON = "json_models/constructs/v1beta1/connection.json"
const templatePathRelathionshipJSON = "json_models/constructs/v1alpha3/relationship.json"

func readTemplate(templatePath string) ([]byte, error) {
	return schemas.Schemas.ReadFile(templatePath)
}
