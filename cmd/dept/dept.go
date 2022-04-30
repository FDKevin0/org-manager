package dept

import (
	"fmt"

	orgmanager "github.com/hduhelp/org-manager"
	"github.com/manifoldco/promptui"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "dept",
	Short: "dept management",
	Run: func(cmd *cobra.Command, args []string) {
		targets := lo.Keys(orgmanager.Targets)
		prompt := promptui.Select{
			Label: "Select Target",
			Items: targets,
		}

		_, target, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		nowDepartment := orgmanager.Targets[target].RootDepartment()
		for {
			depts := nowDepartment.SubDepartments()
			if len(depts) == 0 {
				return
			}
			deptsName := make([]string, 0)
			for _, v := range depts {
				deptsName = append(deptsName, v.Name())
			}
			prompt = promptui.Select{
				Label: "Select Department",
				Items: deptsName,
			}
			_, target, err = prompt.Run()
			for _, v := range depts {
				if v.Name() == target {
					nowDepartment = v
				}
			}
		}
	},
}