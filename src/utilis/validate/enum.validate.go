package validate

import (
	"errors"

	"github.com/devkishor8007/email-sender/src/utilis/enums"
)

func ValidateStatus(status enums.Status) error {
	validStatuses := []enums.Status{enums.StatusPending, enums.StatusActive, enums.StatusInactive}
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return nil
		}
	}
	return errors.New("Invalid status")
}
