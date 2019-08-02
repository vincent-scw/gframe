workflow _Wkf_Start {
	parallel {
		Invoke-Expression -Command "go run .\oauth" 
		Invoke-Expression -Command "go run .\broker_svc" 
		Invoke-Expression -Command "go run .\notification_svc"
		Invoke-Expression -Command "go run .\reception_svc"
		Invoke-Expression -Command "go run .\admin_svc"
	}
}

_Wkf_Start