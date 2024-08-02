package common

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

//{
//  "timestamp": "Example text",
//  "process_name": "Example text",
//  "error_msg": "Example text",
//  "severity": "Example text"
//}

type SequencerAlert struct {
	ProcessName string `json:"process_name"`
	ErrorMsg    string `json:"error_msg"`
	Timestamp   string `json:"timestamp"`
	Severity    string `json:"severity"`
}

func (s SequencerAlert) String() string {
	return fmt.Sprintf("ProcessName: %s, ErrorMsg: %s, Timestamp: %s, Severity: %s",
		s.ProcessName, s.ErrorMsg, s.Timestamp, s.Severity)
}

// These issues should not be common but best to allow high volumes
var SequencerAlertsChannel = make(chan SequencerAlert, 10000)

func SendFailureNotification(processName, errorMsg, timestamp, severity string) {
	var issue SequencerAlert
	issue = SequencerAlert{
		processName,
		errorMsg,
		timestamp,
		severity,
	}

	select {
	case SequencerAlertsChannel <- issue:
		log.Debugln("Issue sent to channel: ", issue)
	default:
		log.Errorln("Issue channel is full, dropping issue: ", issue)
	}
}
