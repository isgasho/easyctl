package cmd

import (
	"easyctl/sys"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	infoCmd.AddCommand(osInfoCmd)
	infoCmd.AddCommand(sysInfoCmd)
	rootCmd.AddCommand(infoCmd)
}

// set 命令合法参数
var infoValidArgs = []string{"os", "system"}

// 查询命令search指令集
var infoCmd = &cobra.Command{
	Use:     "info [OPTIONS]",
	Short:   "print information about current system through easyctl",
	Example: "info os",
	RunE: func(cmd *cobra.Command, args []string) error {
		return parseCommand(cmd, args, infoValidArgs)
	},
	Args: cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return infoValidArgs, cobra.ShellCompDirectiveNoFileComp
	},
}

// 操作系统版本信息
var osInfoCmd = &cobra.Command{
	Use:     "os",
	Short:   "print system version",
	Long:    `print system version...`,
	Example: "info os",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("当前系统版本：%+v\n", sys.SystemInfoObject.OSVersion.ReleaseContent)
	},
}

// 操作系统整体信息
var sysInfoCmd = &cobra.Command{
	Use:     "system",
	Short:   "print system version",
	Long:    `print system version...`,
	Example: "info os",
	Aliases: []string{"sys"},
	Run: func(cmd *cobra.Command, args []string) {
		sys.PrintSystemInfo()
		sys.PrintkernelInfo()
		sys.PrintMemoryInfo()
	},
}
