package cmd

import (
	"fmt"
	"os"

	"github.com/60tonAngel/github_user_activity/activity"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ghmonitor",
		Short: "GitHub用户活动监控工具",
		Long:  "支持多用户配置的GitHub活动监控工具",
	}

	selectUserCmd = &cobra.Command{
		Use:   "select-user",
		Short: "选择要监控的已配置用户",
		Run: func(cmd *cobra.Command, args []string) {
			users := viper.GetStringMap("users")
			if len(users) == 0 {
				fmt.Println("错误: 没有配置任何用户")
				return
			}

			fmt.Println("已配置的用户:")
			i := 1
			var userNames []string
			for name := range users {
				fmt.Printf("%d. %s\n", i, name)
				userNames = append(userNames, name)
				i++
			}

			fmt.Print("请选择要监控的用户编号: ")
			var choice int
			_, err := fmt.Scanln(&choice)
			if err != nil || choice < 1 || choice > len(userNames) {
				fmt.Println("无效的选择")
				return
			}

			selectedUser := userNames[choice-1]
			fmt.Printf("正在监控用户: %s\n", selectedUser)

			// 执行监控命令
			monitorCmd.Run(cmd, []string{selectedUser})
		},
	}

	monitorCmd = &cobra.Command{
		Use:   "monitor [username]",
		Short: "监控指定用户的活动",
		Args:  cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			targetUser := args[0]
			if !viper.IsSet("users." + targetUser) {
				fmt.Printf("错误: 用户 %s 未配置\n", targetUser)
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			targetUser := args[0]
			token := viper.GetString("users." + targetUser + ".token")

			// 获取用户活动
			events, err := activity.GetUserActivity(targetUser, token)
			if err != nil {
				fmt.Printf("获取用户活动失败: %v\n", err)
				return
			}

			// 打印用户活动
			fmt.Printf("用户 %s 的最近活动:\n", targetUser)
			for _, event := range events {
				fmt.Printf("- 类型: %s, 仓库: %s, 时间: %s\n",
					event.Type,
					event.Repo.Name,
					event.Created)
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(monitorCmd)
	rootCmd.AddCommand(selectUserCmd)

	// 移除遗留的toggle参数
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到，创建一个新的
			err := viper.SafeWriteConfigAs("config/config.yaml")
			if err != nil {
				fmt.Printf("无法创建配置文件: %v\n", err)
			}
		} else {
			// 其他错误
			fmt.Printf("读取配置文件出错: %v\n", err)
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
