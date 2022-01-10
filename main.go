package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/logicmonitor/terraform-provider-logicmonitor/logicmonitor"
)

/*
func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: logicmonitor.Provider,
	})
}
*/

func main() {

	opts := &plugin.ServeOpts{ProviderFunc: logicmonitor.Provider}

	// TODO: update this string with the full name of your provider as used in your configs
	/*
		err := plugin.Debug(context.Background(), "registry.terraform.io/logicmonitor/logicmonitor", opts)

		if err != nil {
			log.Fatal(err.Error())
		}
	*/
	plugin.Serve(opts)
}
