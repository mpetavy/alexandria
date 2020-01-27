package controllers

import (
	"time"

	"github.com/mpetavy/common"

	"github.com/astaxie/beego"
	"github.com/mpetavy/alexandria/models"
)

func mapToString(m *map[string]string, separator string) string {
	s := ""

	for _, k := range *m {
		if len(s) > 0 {
			s = s + separator
		}

		s += k + "=" + (*m)[k]
	}

	return s
}

func (c *Archivelink) serverInfo() error {
	var pis = ParameterInfos{
		CONTREP:   ParameterInfo{optmand: OPTIONAL},
		P_VERSION: ParameterInfo{optmand: MANDATORY},
		RESULT_AS: ParameterInfo{def: ASCII},
	}

	if c.check(SERVERINFO, &pis) {
		t := time.Now().UTC()

		m := make(map[string]string)

		m[SERVER_STATUS] = RUNNING
		m[SERVER_VENDOR_ID] = common.Application().Developer
		m[SERVER_VERSION] = common.Application().Version
		m[SERVER_BUILD] = common.Application().Version
		m[SERVER_TIME] = t.Format(common.Hour + common.Minute + common.Second)
		m[SERVER_DATE] = t.Format(common.Year + common.Month + common.Day + common.Msec)
		m[SERVER_STATUS_DESCRIPTION] = RUNNING
		m[P_VERSION] = P_VERSION_INFO

		result := mapToString(&m, ";")

		filterName := c.Param(CONTREP)

		beego.Debug("filterName : " + filterName)

		qs := models.Orm.QueryTable(models.CONTREP).Limit(-1)

		if filterName != "" {
			qs.Filter("name", filterName)
		}

		var contreps []*models.Contrep

		_, err := qs.All(&contreps)

		if err != nil {
			return err
		}

		for _, contrep := range contreps {
			m = make(map[string]string)

			m[CONTREP] = contrep.Name
			m[CONTREP_DESCRIPTION] = contrep.Desc
			m[CONTREP_STATUS] = "??"
			m[CONTREP_STATUS_DESCRIPTION] = "??"

			result += mapToString(&m, ";")
		}

		if c.Param(RESULT_AS) == HTML {
			result = "<html><body>\n" + result + "</body></html>"
		}

		c.Ctx.Output.Body([]byte(result))
	}

	return nil
}
