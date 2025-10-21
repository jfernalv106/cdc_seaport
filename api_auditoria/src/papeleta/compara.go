package papeleta

import "api_auditoria/src/model"

func ComparaPapeleta(src model.PapeletaRecepcionTopic) string {
	var diferencia string = ""

	if src.Op == "c" || src.Op == "r" {
		return ""
	}
	if src.Before.Aga != src.After.Aga {
		diferencia += "Agente aduana:" + *src.After.Aga
	}

	return diferencia

}
