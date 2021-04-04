package AutoSendAction

import (
	"errors"
)

func App_autosend_verify(sep, count int, Type string) error {
	if sep < 0 {
		return errors.New("时间间隔需要大于0分钟")
	}
	if count < 0 {
		return errors.New("剩余重复次数应该大于0")
	}
	switch Type {
	case "fix", "sep":
		break

	default:
		return errors.New("只支持fix和sep模式")
	}

	return nil
}
