package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/schollz/progressbar/v3"
)

type ProgressBarListener struct {
	bar *progressbar.ProgressBar
}

func NewDefaultProgressListener() *ProgressBarListener {
	return &ProgressBarListener{}
}

func (pl *ProgressBarListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		pl.bar = progressbar.DefaultBytes(
			event.TotalBytes, "file start upload",
		)
	case oss.TransferDataEvent:
		pl.bar.Add64(event.RwBytes)
	case oss.TransferCompletedEvent:
		fmt.Println("file upload completed")
	case oss.TransferFailedEvent:
		fmt.Println("file upload is failed")
	}
}
