package model

import (
	"github.com/spf13/viper"
	"github.com/zzzgydi/thanks/common"
	"github.com/zzzgydi/thanks/common/initializer"
)

func initModel() error {
	if viper.GetBool("DATABASE_AUTO_MIGRATE") {
		return common.MDB.AutoMigrate(
			&GitRepo{},
			&GitContributor{},
			&NodeRepo{},
		)
	}

	return nil
}

func init() {
	initializer.Register("migrate", initModel)
}
