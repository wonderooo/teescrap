package uploader

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/chromedp/chromedp"
)

type Job struct {
	ImagePath   string       `json:"imagePath"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Tags        []string     `json:"tags"`
	Colors      ColorChoices `json:"colors"`
}

type JobAction chromedp.Action

func NewJob(imagePath string, title string, description string, colors ColorChoices, tags ...string) Job {
	colors.setDefaults()

	return Job{
		ImagePath:   imagePath,
		Title:       title,
		Description: description,
		Tags:        tags,
		Colors:      colors,
	}
}

func uploadDesigns(jobs *[]Job) JobAction {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		var err error

		for idx, job := range *jobs {
			err = toUploadDesign().Do(ctx)
			if err != nil {
				log.Fatalf("Could not go to upload new design page: %s", err)
			}

			/*
				Need to check if it is 1st iteration of cycle, because afterwards
				browser knows to redirect to single design upload page
			*/
			if idx == 0 {
				err = toSingleUpload().Do(ctx)
				if err != nil {
					log.Fatalf("Could not go to single design upload page: %s", err)
				}
			}

			err = uploadImage(job.ImagePath).Do(ctx)
			if err != nil {
				log.Fatalf("Could not upload image: %s", err)
			}

			err = fillInfo(job.Title, job.Description, job.Tags...).Do(ctx)
			if err != nil {
				log.Fatalf("Could not fill design info: %s", err)
			}

			err = acceptAgreements(false).Do(ctx)
			if err != nil {
				log.Fatalf("Could not accept design agreements: %s", err)
			}

			err = fillColors(job.Colors).Do(ctx)
			if err != nil {
				log.Fatalf("Could not fill design colors: %s", err)
			}

			err = offAdditionalProducts().Do(ctx)
			if err != nil {
				log.Fatalf("Could not turn off additional products: %s", err)
			}

			err = publish().Do(ctx)
			if err != nil {
				log.Fatalf("Could not publish design: %s", err)
			}
		}

		return err
	})
}

func toUploadDesign() chromedp.Action {
	return &chromedp.Tasks{
		chromedp.WaitVisible(XPATH_ACC_NAV_A_TAG),
		chromedp.Navigate(URL_NEW_DESIGN),
	}
}

func toSingleUpload() chromedp.Action {
	return chromedp.Click(XPATH_SNG_UPL_DIV_TAG, chromedp.NodeVisible)
}

func uploadImage(imagePath string) chromedp.Action {
	absPath, err := filepath.Abs(imagePath)
	if err != nil {
		ferr := fmt.Sprintf("Could not get absolute file path of %s: %s", imagePath, err)
		panic(ferr)
	}
	retPath := []string{absPath}

	return &chromedp.Tasks{
		chromedp.SetUploadFiles(
			XPATH_FILE_UPLOAD_INP_TAG,
			retPath,
			chromedp.NodeReady,
		),
		chromedp.WaitVisible(XPATH_IMG_PREVIEW_IMG_TAG),
	}
}

func fillInfo(title string, description string, tags ...string) chromedp.Action {
	return &chromedp.Tasks{
		chromedp.SendKeys(XPATH_TITLE_INP_TAG, title, chromedp.NodeVisible),
		chromedp.SendKeys(XPATH_DESC_TXT_TAG, description, chromedp.NodeVisible),
		chromedp.SendKeys(XPATH_MAIN_TAG_INP_TAG, tags[0], chromedp.NodeVisible),
		chromedp.SendKeys(XPATH_SUP_TAG_INP_TAG, yieldComma(tags[1:len(tags)]...), chromedp.NodeVisible),
	}
}

func acceptAgreements(matureContent bool) chromedp.Action {
	var xpathMature string
	if matureContent {
		xpathMature = XPATH_MATURE_YES_INP_TAG
	} else {
		xpathMature = XPATH_MATURE_NO_DIV_TAG
	}

	return &chromedp.Tasks{
		chromedp.Click(xpathMature, chromedp.NodeVisible),
		chromedp.Click(XPATH_TERMS_INP_TAG, chromedp.NodeVisible),
	}
}

func fillColors(colors ColorChoices) chromedp.Action {
	log.Printf("Got colors: %s", colors)
	return &chromedp.Tasks{
		chromedp.SetAttributeValue(XPATH_TSHIRT_COLOR, ATTR_VALUE, colorMapping[colors.Tshirt], chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_HOODIE_COLOR, ATTR_VALUE, colorMapping[colors.Hoodie], chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_TANK_COLOR, ATTR_VALUE, colorMapping[colors.Tank], chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_CREWNECK_COLOR, ATTR_VALUE, colorMapping[colors.Crewneck], chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_LONG_COLOR, ATTR_VALUE, colorMapping[colors.Longsleeve], chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_BASEBALL_COLOR, ATTR_VALUE, colorMapping[colors.Baseball], chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_KIDS_COLOR, ATTR_VALUE, colorMapping[colors.Kids], chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_KIDS_HOODIE_COLOR, ATTR_VALUE, colorMapping[colors.KidsHoodie], chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_KIDS_LONG_COLOR, ATTR_VALUE, colorMapping[colors.KidsLongsleeve], chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_BABY_BODY_COLOR, ATTR_VALUE, colorMapping[colors.BabyBody], chromedp.NodeReady),
	}
}

func offAdditionalProducts() chromedp.Action {
	return &chromedp.Tasks{
		chromedp.SetAttributeValue(XPATH_STICKERS_OFF_INP_TAG, ATTR_VALUE, ATTR_FALSE, chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_CASES_OFF_INP_TAG, ATTR_VALUE, ATTR_FALSE, chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_MUGS_OFF_INP_TAG, ATTR_VALUE, ATTR_FALSE, chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_WALLART_OFF_INP_TAG, ATTR_VALUE, ATTR_FALSE, chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_PILLOWS_OFF_INP_TAG, ATTR_VALUE, ATTR_FALSE, chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_TOTES_OFF_INP_TAG, ATTR_VALUE, ATTR_FALSE, chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_PINS_OFF_INP_TAG, ATTR_VALUE, ATTR_FALSE, chromedp.NodeReady),
		chromedp.SetAttributeValue(XPATH_MAGNETS_OFF_INP_TAG, ATTR_VALUE, ATTR_FALSE, chromedp.NodeReady),
	}
}

func publish() chromedp.Action {
	return &chromedp.Tasks{
		chromedp.Click(XPATH_SUBMIT_BUTT_TAG, chromedp.NodeVisible),
	}
}

func yieldComma(tags ...string) string {
	ret := ""
	for idx, tag := range tags {
		if idx == 0 {
			ret = fmt.Sprintf("%s,", tag)
		} else {
			ret = fmt.Sprintf("%s%s,", ret, tag)
		}
	}

	log.Printf("Got tags: %s", ret)
	return ret
}
