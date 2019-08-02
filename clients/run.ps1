workflow _Wkf_Start {
	parallel {
		Invoke-Expression -Command "cd .\playercli; npm start"
		Invoke-Expression -Command "cd .\admincli; npm start"
	}
}

_Wkf_Start