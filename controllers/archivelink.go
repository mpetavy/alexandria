package controllers

import (
	"strings"

	"github.com/mpetavy/common"

	"github.com/astaxie/beego"
	"github.com/mpetavy/alexandria/models"
)

const (
	SERVERINFO = "serverInfo"
	PUTCERT    = "putCert"

	P_VERSION_INFO = "045"

	ACCESS_READ   = "r"
	ACCESS_CREATE = "c"
	ACCESS_UPDATE = "u"
	ACCESS_DELETE = "d"

	ACCESSMODE                 = "accessMode"
	ASCII                      = "ascii"
	AUTHID                     = "authId"
	CASE_SENSITIVE             = "caseSensitive"
	CHARSET                    = "charSet"
	COMP_DATE_C                = "compDateC"
	COMP_DATE_M                = "compDateM"
	COMP_ID                    = "compId"
	COMP_TIME_C                = "compTimeC"
	COMP_TIME_M                = "compTimeM"
	COMP_STATUS                = "compStatus"
	CONTENT_DISPOSITION        = "Content-Disposition"
	CONTENT_LENGTH             = "Content-Length"
	CONTENT_TYPE               = "Content-Type"
	CONTREP                    = "contRep"
	CONTREP_DATE               = "contRepDate"
	CONTREP_DESCRIPTION        = "contRepDescription"
	CONTREP_ERROR_DESCRIPTION  = "contRepErrorDescription"
	CONTREP_STATUS             = "contRepStatus"
	CONTREP_STATUS_DESCRIPTION = "contRepStatusDescription"
	CONTREP_TIME               = "contRepTime"
	CONTREP_VENDOR_ID          = "contRepVendorId"
	DATE_C                     = "dateC"
	DATE_M                     = "dateM"
	DOC_ID                     = "docId"
	DOC_PROT                   = "docProt"
	DOC_STATUS                 = "docStatus"
	ERROR                      = "error"
	ERROR_DESCRIPTION          = "errorDescription"
	EXPIRATION                 = "expiration"
	FROM_OFFSET                = "fromOffset"
	HTML                       = "html"
	NUM_COMPS                  = "numComps"
	NUM_RESULTS                = "numResults"
	ONLINE                     = "online"
	OFFLINE                    = "offline"
	PATTERN                    = "pattern"
	P_VERSION                  = "pVersion"
	RESULT_AS                  = "resultAs"
	RET_CODE                   = "retCode"
	RUNNING                    = "running"
	SEC_KEY                    = "secKey"
	SERVER_BUILD               = "serverBuild"
	SERVER_DATE                = "serverDate"
	SERVER_ERROR_DESCRIPTION   = "serverErrorDescription"
	SERVER_STATUS              = "serverStatus"
	SERVER_STATUS_DESCRIPTION  = "serverStatusDescription"
	SERVER_TIME                = "serverTime"
	SERVER_VENDOR_ID           = "serverVendorId"
	SERVER_VERSION             = "serverVersion"
	OPERATION                  = "operation"
	STOPPED                    = "stopped"
	TIME_C                     = "timeC"
	TIME_M                     = "timeM"
	TO_OFFSET                  = "toOffset"
	VERSION                    = "version"

	X_ERROR_DESCRIPTION = "X-errorDescription"
	X_DATE_C            = "X-dateC"
	X_TIME_C            = "X-timeC"
	X_DATE_M            = "X-dateM"
	X_TIME_M            = "X-timeM"
	X_NUMBER_COMPS      = "X-numberComps"
	X_CONTREP           = "X-contentRep"
	X_DOC_ID            = "X-docId"
	X_DOC_STATUS        = "X-docStatus"
	X_P_VERSION         = "X-pVersion"
	X_CONTENT_LENGTH    = "X-Content-Length"
	X_COMP_ID           = "X-compId"
	X_COMP_DATE_C       = "X-compDateC"
	X_COMP_TIME_C       = "X-compTimeC"
	X_COMP_DATE_M       = "X-compDateM"
	X_COMP_TIME_M       = "X-compTimeM"
	X_COMP_STATUS       = "X-compStatus"

	OPTIONAL    = "optional"
	MANDATORY   = "mandatory"
	S_MANDATORY = "s_mandatory"
)

type Archivelink struct {
	beego.Controller

	contreps map[string]models.Contrep
}

type ParameterInfo struct {
	optmand string
	def     string
	sign    bool
}

type ParameterInfos map[string]ParameterInfo

// New instantiates new object of type
func NewArchivelink() *Archivelink {
	result := new(Archivelink)

	result.contreps = make(map[string]models.Contrep)

	return result
}

func (c *Archivelink) BadRequest(txt string) {
	c.CustomAbort(400, txt)
}

func (c *Archivelink) Param(name string) string {
	values, ok := c.Ctx.Request.URL.Query()[name]

	if ok {
		return values[0]
	} else {
		return ""
	}
}

func (c *Archivelink) check(command string, pis *ParameterInfos) bool {
	result := false

	isSigned := false
	params := c.Ctx.Request.URL.Query()

	for param := range params {
		value := params[param][0]

		if param != command {
			if param == SEC_KEY {
				isSigned = true
			} else {
				info, ok := (*pis)[param]

				if !ok {
					c.BadRequest("invalid parameter " + param)
					break
				} else {
					if (value == "") && (info.def != "") {
						params.Set(param, info.def)
						value = info.def
					}

					common.Debug("parameter %s : %s'", param, value)
				}
			}
		}
	}

	result = !isSigned

	if isSigned {
		//TODO isSigned handling
	}

	return result
}

func (c *Archivelink) Put() {
	c.Get()
}
func (c *Archivelink) Get() {
	uri := c.Ctx.Request.RequestURI

	common.Debug("uri : %s", uri)

	var cmd string

	p := strings.Index(uri, "?")
	if p != -1 {
		cmd = uri[p+1:]
	}

	p = strings.IndexAny(cmd, "&#")
	if p != -1 {
		cmd = cmd[:p]
	}

	common.Debug("cmd : %s", cmd)

	if cmd != "" {
		switch cmd {
		case SERVERINFO:
			c.serverInfo()
		case PUTCERT:
			c.putCert()
		default:
			c.BadRequest(common.Translate("Unknown command : %s", cmd))
		}
	} else {
		c.BadRequest(common.Translate("Undefined SAP Archivelink command"))
	}
}
