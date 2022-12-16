package main

//------------------------------HTTP request and response structure-----------------------------------------
    
type responseGetIdentity struct {
    Server_name   string         `json:"server_name"`   
}

type requestPostConvert struct {
    Value   int                  `json:"value"`
}

type responsePostConvert struct {
    Value            int         `json:"value"`
    Value_in_words   string      `json:"value_in_words"`
}
//------------------------------HTTP request and response structure-----------------------------------------
