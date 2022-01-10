package logicmonitor

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
	"github.com/logicmonitor/lm-sdk-go/client/lm"
)

func resourceLmotelCollector() *schema.Resource {
	return &schema.Resource{
		Create: addLmotelCollector,
		Read:   readLmotelCollector,
		Update: updateLmotelCollector,
		Delete: deleteLmotelCollector,

		Schema: map[string]*schema.Schema{

			"build": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"platform": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"collector_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_collector_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"collector_group_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_failback": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_collector_device_failover": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"escalation_chain_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"resend_interval": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"suppress_alert_clear": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"properties": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"downloadurl": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//resource function to add collector
func addLmotelCollector(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	lmotelcollector := makeDeviceLmotelCollectorObject(d) // need to add new function makeDeviceLmotelCollectorObject
	params := lm.NewAddLmotelCollectorParams()
	params.SetBody(&lmotelcollector)
	// calling API to add collector to portal
	log.Printf("Adding lmotel collector")
	restCollectorResponse, err := client.LM.AddLmotelCollector(params) //Changed restCollectorResponse to restLmotelCollectorResponse
	if err != nil {
		return err
	}

	// calling API to save deviceId
	d.SetId(restCollectorResponse.Payload.Data.AllIds[0].ID)
	d.Set("downloadurl", restCollectorResponse.Payload.Data.ByID.LmotelCollectors[restCollectorResponse.Payload.Data.AllIds[0].ID].DownloadURL)
	//d.SetId(getCo)
	return nil
}

// function to sync collector data
func readLmotelCollector(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)

	log.Printf("Starting reading lmotel collector...")
	// get collector id
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	params := lm.NewGetLmotelCollectorByIDParams()
	params.SetID(int32(id))

	restCollectorResponse, err := client.LM.GetLmotelCollectorByID(params)
	if err != nil {
		log.Printf("Failed to find collector %q", err)
		d.SetId("")
		return nil
	}
	log.Printf("ashwani: %+v", restCollectorResponse.Payload)

	// d.Set("backup_collector_id", restCollectorResponse.Payload.BackupAgentID)
	// d.Set("collector_group_id", restCollectorResponse.Payload.CollectorGroupID)
	d.Set("description", restCollectorResponse.Payload.Data.ByID.LmotelCollectors[restCollectorResponse.Payload.Data.AllIds[0].ID].Description)
	// d.Set("enable_failback", restCollectorResponse.Payload.EnableFailBack)
	// d.Set("enable_collector_device_failover", restCollectorResponse.Payload.EnableFailOverOnCollectorDevice)
	// d.Set("escalation_chain_id", restCollectorResponse.Payload.EscalatingChainID)
	// d.Set("resend_interval", restCollectorResponse.Payload.ResendIval)
	// d.Set("suppress_alert_clear", restCollectorResponse.Payload.SuppressAlertClear)

	// properties := make(map[*string]*string)
	// props := restCollectorResponse.Payload.CustomProperties
	// for _, v := range props {
	// 	properties[v.Name] = v.Value
	// }
	// d.Set("properties", properties)

	return nil
}

// function to update collector data
func updateLmotelCollector(d *schema.ResourceData, m interface{}) error {
	//d.Partial(true)
	// client := m.(*lmclient.LMSdkGo)
	log.Printf("Currently Update lmotel-collector is not supported. ")
	log.Printf("You may can delete it by **terrarorm destroy** and create it again... ")

	// collector := makeDeviceLmotelCollectorObject(d) // Replased makeDeviceCollectorObject with makeDeviceLmotelCollectorObject

	// // get collector id
	// id, err := strconv.Atoi(d.Id())
	// if err != nil {
	// 	return err
	// }

	// params := lm.NewUpdateLmotelCollectorByIDParams()
	// params.SetID(int32(id))
	// params.SetBody(&collector)

	// // list of available properties
	// s := []string{
	// 	"backup_collector_id",
	// 	"collector_group_id",
	// 	"description",
	// 	"enable_failback",
	// 	"enable_collector_device_failover",
	// 	"escalation_chain_id",
	// 	"properties",
	// 	"resend_interval",
	// 	"suppress_alert_clear",
	// }

	// // loops through array of properties to see which one has changed, the ones that did not change are removed from the list
	// for _, v := range s {
	// 	if d.HasChange(v) {
	// 	} else {
	// 		s = remove(s, v)
	// 	}
	// }

	// // makes a bulk update for all properties that were changed
	// restCollectorResponse, err := client.LM.UpdateLmotelCollectorByID(params)
	// if err != nil {
	// 	return err
	// }

	// for _, v := range s {
	// 	d.SetPartial(v)
	// }
	// log.Printf("updating collector response %v", restCollectorResponse.Payload)
	// d.Partial(false)
	return nil
}

// function to delete collector
func deleteLmotelCollector(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)

	// get collector id
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	params := lm.NewDeleteLmotelCollectorByIDParams()
	params.SetID(int32(id))
	log.Printf("ID %d", int32(id))
	restCollectorResponse, err := client.LM.DeleteLmotelCollectorByID(params)
	if err != nil {
		log.Printf("Error deleting collector %s", err)
		d.SetId("")
	}

	log.Printf("delete collector response %v", restCollectorResponse.Payload)
	d.SetId("")
	return nil
}
