package driver_ego

import (
	"github.com/dtm-labs/dtmdriver"
	_ "github.com/dtm-labs/dtmdriver-ego"
)

func init() {
	err := dtmdriver.Use("dtm-driver-ego")
	if err != nil {
		panic(err)
	}
}
