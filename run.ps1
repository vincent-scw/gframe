param (
	[bool] $i = 1
)

workflow _Wkf_Start {
	parallel {
		Invoke-Expression -Command "go run .\oauth" 
		Invoke-Expression -Command "go run .\broker_svc" 
		Invoke-Expression -Command "go run .\notification_svc"
		Invoke-Expression -Command "go run .\reception_svc"
		Invoke-Expression -Command "go run .\admin_svc"
		if (isInt($i) -eq 1) {
			Invoke-Expression -Command "cd .\playercli; npm start"
			Invoke-Expression -Command "cd .\admincli; npm start"
		}
	}
}

_Wkf_Start