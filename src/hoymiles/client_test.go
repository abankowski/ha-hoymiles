package hoymiles

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestUmmarshall(t *testing.T) {

	input := `
{
  "status":"0",
  "message":"success",
  "data":[
    {
      "id":574820,
      "sn":"10F881835718",
      "dtu_sn":"10F881835718",
      "type":1,
      "version":3,
      "replace_num":0,
      "model_no":"DTU-Pro",
      "soft_ver":"V00.02.13",
      "hard_ver":"H09.02.02",
      "extend_data":{
        
      },
      "warn_data":{
        "_rw":"",
        "connect":true,
        "warn":false
      },
      "children":[
        {
          "id":2622510,
          "sn":"116181850007",
          "dtu_sn":"10F881835718",
          "type":3,
          "version":3,
          "repl  ace_num":0,
          "model_no":"HM-1500",
          "soft_ver":"V01.00.18",
          "hard_ver":"H00.04.10",
          "extend_data":{
            "grid_name":"",
            "port_array":[
              1,
              2,
              3,
              4
            ],
            "grid_id":0,
            "grid_version":""
          },
          "warn_data":{
            "_raw_warns":"148",
            "track_time":"2022-10-21 07:01",
            "_l_time":18,
            "warn":true,
            "conn  ect":true,
            "_dtu_link":true
          },
          "children":[
            
          ]
        },
        {
          "id":2622500,
          "sn":"116181849296",
          "dtu_sn":"10F881835718",
          "type":3,
          "version":3,
          "repla  ce_num":0,
          "model_no":"HM-1500",
          "soft_ver":"V01.00.18",
          "hard_ver":"H00.04.10",
          "extend_data":{
            "grid_name":"",
            "port_array":[
              1,
              2,
              3,
              4
            ],
            "grid_id":0,
            "grid_version":""
          },
          "warn_data":{
            "_raw_warns":"148",
            "track_time":"2022-10-21 07:01",
            "_l_time":18,
            "warn":true,
            "connect":true,
            "_dtu_link":true
          },
          "children":[
            
          ]
        },
        {
          "id":2622490,
          "sn":"116173414141",
          "dtu_sn":"10F881835718",
          "type":3,
          "version":3,
          "replac  e_num":0,
          "model_no":"HM-1500",
          "soft_ver":"V01.00.16",
          "hard_ver":"H00.04.00",
          "extend_data":{
            "grid_name":"",
            "port_array":[
              1,
              2,
              3,
              4
            ],
            "grid_id":0,
            "grid_version":""
          },
          "warn_data":{
            "_raw_warns":"148",
            "track_time":"2022-10-21 07:01",
            "_l_time":18,
            "warn":true,
            "connect":true,
            "_dtu_link":true
          },
          "children":[
            
          ]
        },
        {
          "id":2622520,
          "sn":"114174940119",
          "dtu_sn":"10F881835718",
          "type":3,
          "version":3,
          "replace  _num":0,
          "model_no":"HM-800",
          "soft_ver":"V01.00.10",
          "hard_ver":"H00.04.00",
          "extend_data":{
            "grid_name":"",
            "port_array":[
              1,
              2
            ],
            "grid_id":0,
            "grid_version":""
          },
          "warn_data":{
            "_raw_warns":"148",
            "track_time":"2022-10-21 07:01",
            "_l_time":18,
            "warn":true,
            "connect":true,
            "_dtu_link":true
          },
          "children":[
            
          ]
        }
      ]
    }
  ],
  "systemNotice":null
}`

	a := StationSelectDeviceOfTreeResponse{}
	err := json.Unmarshal([]byte(input), &a)

	if err != nil {
		t.Fatalf("failed to unmarshal json %s", err)
	}

	if a.Status != "0" {
		log.Fatal("status is not 0")
	}

	if a.Message != "success" {
		log.Fatal("status is not 0")
	}

	if len(a.Data) != 1 {
		log.Fatal("body len does not match")
	}

	fmt.Println("a", a)
}
