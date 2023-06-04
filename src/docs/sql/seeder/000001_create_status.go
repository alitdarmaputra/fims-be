package seeds

import (
	"fmt"

	"github.com/alitdarmaputra/fims-be/src/business/model"
)

func (s Seed) StatuSeed() {
	roles := []model.Status{
		{
			Name: model.StatusInProgress,
		},
		{
			Name: model.StatusProductReview,
		},
		{
			Name: model.StatusReadyForDevelopment,
		},
		{
			Name: model.StatusInDevelopment,
		},
		{
			Name: model.StatusDone,
		},
	}

	tx := s.db.CreateInBatches(roles, 3)
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
	}
}
