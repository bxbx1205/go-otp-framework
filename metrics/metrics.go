package metrics

import "sync/atomic"

var OTPSent uint64
var OTPVerified uint64
var SMSSuccess uint64
var SMSFailed uint64

func IncrementOTPSent() {
	atomic.AddUint64(&OTPSent, 1)
}

func IncrementOTPVerified() {
	atomic.AddUint64(&OTPVerified, 1)
}

func IncrementSMSSuccess() {
	atomic.AddUint64(&SMSSuccess, 1)
}

func IncrementSMSFailed() {
	atomic.AddUint64(&SMSFailed, 1)
}
