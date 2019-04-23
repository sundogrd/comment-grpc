package user

import "errors"

var TransactionServiceErr = struct {
	AckAlarmRealReplyEmptyErr                           error
	AckAlarmReplyConvErr                                error
}{
	AckAlarmRealReplyEmptyErr:                           errors.New("rpc empty real response: Psm=content.opadmin.feedback_info, method=AckAlarm"),
	AckAlarmReplyConvErr:                                errors.New("rpc response conv error: Psm=content.opadmin.feedback_info, method=AckAlarm"),
}