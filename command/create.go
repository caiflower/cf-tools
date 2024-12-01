package command

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

type Project struct {
	Path      string
	URL       string
	Module    string
	Branch    string
	GitOrigin string
}

func NewCreateCommand() *cobra.Command {
	project := &Project{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "A command use for crate project base on caiflower-common-tools (github.com/caiflower/common-tools)",
		Long:  "path: ",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Flags().NFlag() == 0 {
				cmd.Help()
				return
			}

			splits := strings.Split(project.Module, "/")
			projectName := splits[len(splits)-1]

			path := project.Path
			// 优先使用path
			if project.Path == "" && project.URL != "" {
				fmt.Printf("git clone %s -b %s \n", project.URL, project.Branch)
				if err := RunCommand("git", []string{"clone", project.URL, "-b", project.Branch}, ""); err != nil {
					fmt.Println(err.Error())
					return
				}

				path = "./demo-api"
			}

			dataMap := make(map[string]interface{})
			dataMap["MODULE"] = project.Module
			if err := parse(path, dataMap); err != nil {
				fmt.Println(err.Error())
			}

			splits1 := strings.Split(path, "/")
			parentPath := strings.TrimSuffix(path, splits1[len(splits1)-1])
			fmt.Println("rename project")
			if err := os.Rename(path, parentPath+"/"+projectName); err != nil {
				fmt.Println(err.Error())
				return
			}

			if err := RunCommand("rm", []string{"-rf", ".git"}, parentPath+"/"+projectName); err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("git mod tidy")
			if err := RunCommand("go", []string{"mod", "tidy"}, parentPath+"/"+projectName); err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("git init")
			if err := RunCommand("git", []string{"init"}, parentPath+"/"+projectName); err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("git add .")
			if err := RunCommand("git", []string{"add", "."}, parentPath+"/"+projectName); err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("git commit -m 'init'")
			if err := RunCommand("git", []string{"commit", "-m", "init"}, parentPath+"/"+projectName); err != nil {
				fmt.Println(err.Error())
				return
			}

			if project.GitOrigin != "" {
				if err := RunCommand("git", []string{"remote", "add", "origin", project.GitOrigin}, parentPath+"/"+projectName); err != nil {
					fmt.Println(err.Error())
					return
				}
			}
		},
	}

	cmd.Flags().StringVar(&project.Module, "module", "", "module name")
	cmd.Flags().StringVar(&project.GitOrigin, "git-origin", "", "project git origin address")
	//cmd.Flags().StringVar(&project.Path, "path", "", "demo project path")
	cmd.Flags().StringVar(&project.URL, "url", "https://github.com/caiflower/demo-api.git", "demo project git address")
	cmd.Flags().StringVar(&project.Branch, "branch", "v1.0.0", "demo project git branch")
	return cmd
}

func parse(path string, dataMap map[string]interface{}) error {
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, v := range dir {
		if v.IsDir() {
			info, err := v.Info()
			// 忽略隐藏文件
			if strings.HasPrefix(info.Name(), ".") {
				continue
			}

			if err != nil {
				return err
			}
			if err := parse(path+"/"+info.Name(), dataMap); err != nil {
				return err
			}
		} else {
			if strings.HasSuffix(v.Name(), ".tpl") {
				fileName := path + "/" + v.Name()
				bytes, err := ExecTemplateGetBytes(fileName, dataMap)
				if err != nil {
					return err
				}

				file, err := os.OpenFile(strings.TrimSuffix(fileName, ".tpl"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
				file.Write(bytes)
				file.Close()

				// 删除tpl
				os.Remove(fileName)
			}
		}
	}

	return nil
}

func ExecTemplateGetBytes(path string, dataMap map[string]interface{}) ([]byte, error) {
	tt, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}
	objBuff := &bytes.Buffer{}
	if err = tt.Execute(objBuff, dataMap); err != nil {
		return nil, err
	}
	return objBuff.Bytes(), nil
}

func RunCommand(command string, args []string, dir string) (err error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return
	}

	if err = cmd.Start(); err != nil {
		return
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return
	}

	return
}
