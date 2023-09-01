package uploader

const (
	ATTR_VALUE                = `value`
	URL_NEW_DESIGN            = `https://www.teepublic.com/design/quick_create`
	XPATH_ACC_NAV_A_TAG       = `/html/body/header/div/nav/div[3]/div/div[1]/a`
	XPATH_SNG_UPL_DIV_TAG     = `/html/body/div[4]/div/div[2]/div[3]/div/form/section/div/div[1]`
	XPATH_FILE_UPLOAD_INP_TAG = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[2]/input`
	XPATH_IMG_PREVIEW_IMG_TAG = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[2]/div[3]/img`
	XPATH_TITLE_INP_TAG       = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[1]/div[1]/div[1]/div[1]/input`
	XPATH_DESC_TXT_TAG        = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[1]/div[1]/div[1]/div[2]/textarea`
	XPATH_MAIN_TAG_INP_TAG    = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[1]/div[1]/div[2]/div[1]/input`
	XPATH_SUP_TAG_INP_TAG     = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[1]/div[1]/div[2]/div[2]/div[2]/ul[1]/li/input`
	XPATH_MATURE_NO_INP_TAG   = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[1]/div[1]/div[3]/div/div[2]/div[2]/input`
	XPATH_MATURE_NO_DIV_TAG   = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[1]/div[1]/div[3]/div/div[2]/div[2]`
	XPATH_MATURE_YES_INP_TAG  = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[1]/div[1]/div[3]/div/div[2]/div[1]/input`
	XPATH_TERMS_INP_TAG       = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[4]/label/input`
	XPATH_SUBMIT_BUTT_TAG     = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[4]/div[2]/button[1]`
)

/*
Default colors to specific apparel
*/
const (
	XPATH_TSHIRT_COLOR      = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[3]/div[2]/div/table/tbody/tr[2]/td/div[1]/div[1]/input`
	XPATH_HOODIE_COLOR      = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[3]/div[2]/div/table/tbody/tr[3]/td/div/div[1]/input`
	XPATH_TANK_COLOR        = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[3]/div[2]/div/table/tbody/tr[4]/td/div/div[1]/input`
	XPATH_CREWNECK_COLOR    = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[3]/div[2]/div/table/tbody/tr[5]/td/div/div[1]/input`
	XPATH_LONG_COLOR        = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[3]/div[2]/div/table/tbody/tr[6]/td/div/div[1]/input`
	XPATH_BASEBALL_COLOR    = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[3]/div[2]/div/table/tbody/tr[7]/td/div/div[1]/input`
	XPATH_KIDS_COLOR        = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[3]/div[2]/div/table/tbody/tr[8]/td/div/div[1]/input`
	XPATH_KIDS_HOODIE_COLOR = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[3]/div[2]/div/table/tbody/tr[9]/td/div/div[1]/input`
	XPATH_KIDS_LONG_COLOR   = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[3]/div[2]/div/table/tbody/tr[10]/td/div/div[1]/input`
	XPATH_BABY_BODY_COLOR   = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[3]/div[2]/div/table/tbody/tr[11]/td/div/div[1]/input`
)

/*
Additional products toggles
*/
const (
	ATTR_FALSE                 = `false`
	ATTR_TRUE                  = `true`
	XPATH_STICKERS_OFF_INP_TAG = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[11]/div[2]/div[2]/div[3]/div/input`
	XPATH_CASES_OFF_INP_TAG    = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[11]/div[2]/div[3]/div[3]/div/input`
	XPATH_MUGS_OFF_INP_TAG     = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[11]/div[2]/div[4]/div[3]/div/input`
	XPATH_WALLART_OFF_INP_TAG  = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[11]/div[2]/div[5]/div[4]/div/input`
	XPATH_PILLOWS_OFF_INP_TAG  = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[11]/div[2]/div[6]/div[3]/div/input`
	XPATH_TOTES_OFF_INP_TAG    = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[11]/div[2]/div[7]/div[3]/div/input`
	XPATH_PINS_OFF_INP_TAG     = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[11]/div[2]/div[9]/div[3]/div/input`
	XPATH_MAGNETS_OFF_INP_TAG  = `/html/body/div[4]/div/div[2]/div[3]/div/form/div/div[3]/div[2]/div[11]/div[2]/div[10]/div[3]/div/input`
)
