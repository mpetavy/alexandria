package controllers

// "time"
//
// "github.com/astaxie/beego"
// "github.com/astaxie/beego/orm"
// "github.com/mpetavy/alexandria/app"
// "github.com/mpetavy/alexandria/models"
// "github.com/mpetavy/alexandria/openssl"

func (c *Archivelink) create() {
	// 	var pis = ParameterInfos{
	// 		CONTREP:        ParameterInfo{optmand: MANDATORY},
	// 		COMP_ID:        ParameterInfo{optmand: MANDATORY},
	// 		DOC_ID:         ParameterInfo{optmand: MANDATORY},
	// 		P_VERSION:      ParameterInfo{optmand: MANDATORY},
	// 		CONTENT_TYPE:   ParameterInfo{optmand: OPTIONAL},
	// 		CHARSET:        ParameterInfo{optmand: OPTIONAL},
	// 		VERSION:        ParameterInfo{optmand: OPTIONAL},
	// 		CONTENT_LENGTH: ParameterInfo{optmand: MANDATORY},
	// 		DOC_PROT:       ParameterInfo{optmand: OPTIONAL},
	// 		ACCESSMODE:     ParameterInfo{optmand: S_MANDATORY},
	// 		AUTHID:         ParameterInfo{optmand: S_MANDATORY},
	// 		EXPIRATION:     ParameterInfo{optmand: S_MANDATORY},
	// 		SEC_KEY:        ParameterInfo{optmand: OPTIONAL},
	// 	}

	// if c.check(PUTCERT, &pis) {
	// 	contrep := models.Contrep{Name: c.Param(CONTREP)}
	//
	// 	err := models.Orm.Read(&contrep, "NAme")
	// 	if err != orm.ErrNoRows {
	// 		defer c.Ctx.Request.Body.Close()
	//
	// 		body, err := io.ReadAll(c.Ctx.Request.Body)
	// 		common.Error(err)
	//
	// 		derFile, err := app.CreateTempFile()
	// 		common.Error(err)
	//
	// 		derFile.Write(body)
	// 		derFile.Close()
	//
	// 		pemFilename, err := app.CreateTempFilename()
	// 		common.Error(err)
	//
	// 		err = openssl.ConvertDER2PEM(derFile.Name(), pemFilename)
	// 		common.Error(err)
	//
	// 		beego.Debug("read ConvertDER2PEM PEM file")
	//
	// 		pem, err := os.ReadFile(pemFilename)
	// 		common.Error(err)
	//
	// 		if contrep.IsCertProtected {
	// 			contrep.ReceivedCert = string(pem)
	// 			contrep.ReceivedCertDatetime = time.Now()
	// 		} else {
	// 			contrep.Cert = string(pem)
	// 			contrep.CertDatetime = time.Now()
	// 		}
	//
	// 		_, err = models.Orm.Update(&contrep)
	//
	// 		common.Error(err)
	// 	} else {
	// 		c.CustomAbort(406, app.Translate("Unknown Content Repository %s", c.Param(CONTREP)))
	// 	}
	// }
}
