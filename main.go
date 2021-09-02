package main

import (
	"ctmSetup/config"
	"ctmSetup/utils"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"

	"github.com/briandowns/spinner"
	"github.com/gookit/color"
)

func main() {
	var softList []config.Soft
	selectSoft := []string{}
	promtSoft := []string{}
	defaultList := []string{}

	softList = config.GetConfig()
	versionURL := "http://ftp.ctm.ru/ctm/Scripts/Versions.ini"
	maindir := "./CTM_SETUP/"

	color.White.Print("Загрузка информации о номерах версий  ")
	response, err := http.Get(versionURL)
	if err != nil {
		color.Red.Println("CANCEL [" + err.Error() + "]")
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		bodyText, err := ioutil.ReadAll(response.Body)
		if err != nil {
			color.Red.Println("CANCEL [" + err.Error() + "]")
		}
		for _, item := range strings.Split(string(bodyText), "\n") {
			if strings.Contains(item, "=") {
				vers := strings.Split(item, "=")
				for i, item2 := range softList {
					if item2.Folder == vers[0] {
						item2.Version = vers[1]
						softList[i] = item2
					}
				}
			}
		}
	}
	color.Green.Println("OK")
	for _, item := range softList {
		promtSoft = append(promtSoft, item.Name+" "+item.Version)
		if item.Checked {
			defaultList = append(defaultList, item.Name+" "+item.Version)
		}
	}

	prompt := &survey.MultiSelect{
		Message: "Выберите программы для загрузки:",
		Options: promtSoft,
		Default: defaultList,
	}

	if err := os.MkdirAll(filepath.Dir(maindir), os.ModePerm); err != nil {
		color.Red.Println("Ошибка создания каталога" + maindir)
		return
	}

	downloader := utils.NewDownloader(maindir)

	survey.AskOne(prompt, &selectSoft, survey.WithPageSize(15))

	cnt := 0
	for _, item1 := range selectSoft {
		for i, item2 := range softList {
			if item1 == (item2.Name + " " + item2.Version) {
				downloader.AppendResource(item2.File, item2.URL)
				item2.Checked = true
				softList[i] = item2
				cnt++
			} else {
				item2.Checked = false
				softList[i] = item2
			}
		}
	}
	if cnt == 0 {
		color.Red.Println("Не были выбраны программы для скачивания. Операция отменена.")
		return
	}

	color.Yellow.Println("Старт загрузки файлов. Всего ", cnt)
	downloader.Concurrent = 3
	downloader.Start()

	color.Yellow.Println("Распаковка файлов ")

	for _, item := range softList {
		if item.Checked == true {
			color.White.Print(item.File)
			utils.Unzip(maindir+item.File, maindir+item.Folder)
			color.Green.Println(" OK ")
		}
	}

	result := false
	prompt2 := &survey.Confirm{
		Message: "Выполнить установку программы?",
	}
	survey.AskOne(prompt2, &result)
	if result == false {
		return
	}

	color.Yellow.Println("Начала установки")
	spinner := spinner.New(spinner.CharSets[8], 100*time.Millisecond)
	for _, item := range softList {
		if item.Checked == true {
			color.White.Print(item.Name)
			execuFile := filepath.Join(maindir, item.Folder, "setup.exe")

			cmd := exec.Command(execuFile, "/CD", "/AUTO", "/HIDE")
			spinner.Start()
			if err := cmd.Run(); err != nil {
				color.Red.Println("CANCEL [" + err.Error() + "]")
			} else {
				color.Green.Println("OK")
			}
			spinner.Stop()
		}
	}

}
