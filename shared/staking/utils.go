package staking

import (
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"github.com/sirupsen/logrus"
	"log"
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

func logValidatorEventFields(val *types.ValidatorEvent) logrus.Fields {
	res := logrus.Fields{
		"transaction_hash": val.TransactionHash,
		"block_number":     val.BlockNumber,
	}
	var event *types.ValidatorEvent_EventFields
	if val.GetDelegated() != nil {
		event = val.GetDelegated()
	} else if val.GetUndelegated() != nil {
		event = val.GetUndelegated()
	} else if val.GetClaimed() != nil {
		event = val.GetClaimed()
	} else if val.GetRedelegated() != nil {
		event = val.GetRedelegated()
	} else {
		log.Panicf("")
	}
	res["validator"] = event.Validator
	res["staker"] = event.Staker
	res["amount"] = event.Amount
	res["epoch"] = event.Epoch
	return res
}
