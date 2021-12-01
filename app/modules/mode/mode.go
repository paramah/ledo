package mode

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Mode struct {
	modeCfg        string
	CurrentMode    string
	availableModes map[string]string
}

func InitMode(modeFileName string, cfgFile string) Mode {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	modeFile := filepath.FromSlash(filepath.Join(wd, modeFileName))
	if _, err := os.Stat(modeFile); os.IsNotExist(err) {
		err := ioutil.WriteFile(modeFile, []byte("dev"), 0644)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
	}
	currentModeFromFile, err := ioutil.ReadFile(modeFile)
	currentMode := strings.TrimSpace(string(currentModeFromFile))
	// TODO: dirty implement mode read
	// Start dirty code ;]
	viper.AddConfigPath(wd)
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
	elements := viper.Get("modes")
	// End dirty code :)

	var el = make(map[string]string)
	for mode, cfg := range elements.(map[string]interface{}) {
		el[mode] = fmt.Sprintf("%v", cfg)
	}
	return Mode{
		CurrentMode:    string(currentMode),
		availableModes: el,
		modeCfg:        modeFileName,
	}

}

func (c *Mode) GetMode() string {
	return c.CurrentMode
}

func (c *Mode) PrintListModes() {
	fmt.Printf("Available modes:\n")
	for m := range c.availableModes {
		fmt.Printf("- %v\n", m)
	}
}

func (c *Mode) CheckMode(mode string) (bool, error) {
	for m := range c.availableModes {
		if m == mode {
			return true, nil
		}
	}
	return false, errors.New("Selected mode not exists in config file")
}

func (c *Mode) GetModes() map[string]string {
	return c.availableModes
}

func (c *Mode) SetMode(modeName string) (bool, error) {
	_, err := c.CheckMode(modeName)
	if err != nil {
		return false, err
	}

	wd, err := os.Getwd()
	if err != nil {
		return false, errors.New("Get working directory error!?")
	}

	modeFile := filepath.FromSlash(filepath.Join(wd, c.modeCfg))
	wrErr := ioutil.WriteFile(modeFile, []byte(modeName), 0644)
	if wrErr != nil {
		return false, errors.New("Write mode error")
	}
	return true, nil
}

func (c *Mode) GetModeConfig() ([]string, error) {
	_, err := c.CheckMode(c.CurrentMode)

	if err == nil {
		composes := strings.Fields(c.availableModes[c.CurrentMode])
		return composes, nil
	}

	return nil, errors.New("Selected mode not exists in config file")
}
