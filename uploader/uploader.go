package uploader

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

const (
	OPT_HEADLESS_BROWSER = "headless"
	OPT_USER_AGENT       = "Mozilla/5.0 (Macintosh; Intel Mac OS X 8_9_7; en-US) Gecko/20130401 Firefox/59.3"
	OPT_WIDTH            = 1920
	OPT_HEIGHT           = 1080
	OPT_NO_SANDBOX       = "no-sandbox"
	OPT_NO_DEV_SHM       = "disable-dev-shm-usage"
)

var (
	cancels []context.CancelFunc
)

type Uploader struct {
	debug    bool
	jobs     []Job
	settings context.Context
}

func New(debug bool, jobs []Job) *Uploader {
	ctx := getContext(debug, len(jobs))

	return &Uploader{
		debug:    debug,
		jobs:     jobs,
		settings: ctx,
	}
}

func (u *Uploader) Run() error {
	dur, _ := time.ParseDuration("2m")
	username, password := credentials(!u.debug)

	err := chromedp.Run(u.settings,

		chromedp.Navigate(`https://www.teepublic.com/`),

		login(username, password),

		uploadDesigns(u.jobs),

		chromedp.Sleep(dur),
	)

	defer cancelAll()

	return err
}

func cancelAll() {
	for _, fn := range cancels {
		fn()
	}
}

func getContext(debug bool, size int) context.Context {
	opts := chromedp.DefaultExecAllocatorOptions[:]
	opts = append(opts,
		chromedp.UserAgent(OPT_USER_AGENT),
		chromedp.Flag(OPT_NO_SANDBOX, true),
		chromedp.Flag(OPT_NO_DEV_SHM, true),
		chromedp.WindowSize(OPT_WIDTH, OPT_HEIGHT),
	)

	if debug == true {
		opts = append(opts,
			chromedp.Flag(OPT_HEADLESS_BROWSER, true),
		)
	}
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)

	if debug == true {
		ctx, cancel = chromedp.NewContext(ctx,
			chromedp.WithDebugf(log.Printf),
		)
	} else {
		ctx, cancel = chromedp.NewContext(ctx, 
			chromedp.WithDebugf(log.Printf),
		)
	}
	cancels = append(cancels, cancel)

	/*
		This timeout is set for the whole run of chromedriver
		It needs to be dynamic in order to scale timeout duration
		with the size of job list
	*/
	ctx, cancel = context.WithTimeout(ctx, 45*time.Duration(size)*time.Second)
	cancels = append(cancels, cancel)

	return ctx
}
