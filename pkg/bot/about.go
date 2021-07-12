package bot

import (
	"fmt"

	"github.com/mikeletux/EMT-Go-Telegram-Bot/pkg/emt"
)

func GetAbout(emt emt.Emt, args []string) (string, error) {
	return fmt.Sprintf("%s This bot has been developed by MikeletuX.\n"+
		"%s API provided by mobilitylabs (https://mobilitylabs.emtmadrid.es)\n"+
		"%s Code available at https://github.com/mikeletux/EMT-Go-Telegram-Bot\n"+
		"2021", phone, key, paper), nil
}
