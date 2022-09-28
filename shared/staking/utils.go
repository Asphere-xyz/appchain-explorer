package staking

import (
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"github.com/sirupsen/logrus"
)

func logValidatorFields(val *types.Validator) logrus.Fields {
	return logrus.Fields{
		"address":    val.Address,
		"owner":      val.Owner,
		"status":     val.Status.String(),
		"commission": val.Commission,
	}
}

func logSlashingFields(val *types.ValidatorSlashing) logrus.Fields {
	return logrus.Fields{
		"validator": val.Validator,
		"epoch":     val.Epoch,
		"slashings": val.Slashings,
	}
}

func logDepositFields(val *types.ValidatorDeposit) logrus.Fields {
	return logrus.Fields{
		"validator": val.Validator,
		"epoch":     val.Epoch,
		"amount":    val.Amount,
	}
}
