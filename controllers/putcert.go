package controllers

import (
	"io/ioutil"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mpetavy/alexandria/models"
	"github.com/mpetavy/alexandria/openssl"
	"github.com/mpetavy/common"
)

func (c *Archivelink) putCert() {
	var pis = ParameterInfos{
		CONTREP:   ParameterInfo{optmand: MANDATORY},
		P_VERSION: ParameterInfo{optmand: MANDATORY},
		AUTHID:    ParameterInfo{optmand: MANDATORY},
	}

	if c.check(PUTCERT, &pis) {
		contrep := models.Contrep{Name: c.Param(CONTREP)}

		err := models.Orm.Read(&contrep, "NAme")
		if err != orm.ErrNoRows {
			defer c.Ctx.Request.Body.Close()

			body, err := ioutil.ReadAll(c.Ctx.Request.Body)
			common.Fatal(err)

			derFile, err := common.CreateTempFile()
			common.Fatal(err)

			derFile.Write(body)
			derFile.Close()

			pemFilename, err := common.CreateTempFile()
			common.Fatal(err)

			err = openssl.ConvertDER2PEM(derFile.Name(), pemFilename.Name())
			common.Fatal(err)

			beego.Debug("read ConvertDER2PEM PEM file")

			pem, err := ioutil.ReadFile(pemFilename.Name())
			common.Fatal(err)

			if contrep.IsCertProtected {
				contrep.ReceivedCert = string(pem)
				contrep.ReceivedCertDatetime = time.Now()
			} else {
				contrep.Cert = string(pem)
				contrep.CertDatetime = time.Now()
			}

			_, err = models.Orm.Update(&contrep)
			common.Fatal(err)
		} else {
			c.CustomAbort(406, common.Translate("Unknown Content Repository %s", c.Param(CONTREP)))
		}
	}
}
