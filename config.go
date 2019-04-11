package requestgateway

import (
	"context"
	"log"
	"os"
	"strconv"

	lbcf "github.com/lidstromberg/config"
)

var (
	//EnvDebugOn controls verbose logging
	EnvDebugOn bool
)

//preflight config checks
func preflight(ctx context.Context, bc lbcf.ConfigSetting) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	log.Println("Started Gateway preflight..")

	//get the session config and apply it to the config
	bc.LoadConfigMap(ctx, preflightConfigLoader())

	//then check that we have everything we need
	if bc.GetConfigValue(ctx, "EnvDebugOn") == "" {
		log.Fatal("Could not parse environment variable EnvDebugOn")
	}

	if bc.GetConfigValue(ctx, "EnvGtwaySqlType") == "" {
		log.Fatal("Could not parse environment variable EnvGtwaySqlType")
	}

	if bc.GetConfigValue(ctx, "EnvGtwaySqlConnection") == "" {
		log.Fatal("Could not parse environment variable EnvGtwaySqlConnection")
	}

	//set the debug value
	constlog, err := strconv.ParseBool(bc.GetConfigValue(ctx, "EnvDebugOn"))

	if err != nil {
		log.Fatal("Could not parse environment variable EnvDebugOn")
	}

	EnvDebugOn = constlog

	log.Println("..Finished Gateway preflight.")
}

//preflightConfigLoader loads the config vars
func preflightConfigLoader() map[string]string {
	cfm := make(map[string]string)

	//EnvDebugOn controls verbose logging
	cfm["EnvDebugOn"] = os.Getenv("GTWAYPG_DEBUGON")
	//EnvGtwaySqlType is the driver type
	cfm["EnvGtwaySqlType"] = os.Getenv("GTWAYPG_SQLDST")
	//EnvGtwaySqlConnection is the connection string
	cfm["EnvGtwaySqlConnection"] = os.Getenv("GTWAYPG_SQLCNX")

	if cfm["EnvDebugOn"] == "" {
		log.Fatal("Could not parse environment variable EnvDebugOn")
	}

	if cfm["EnvGtwaySqlType"] == "" {
		log.Fatal("Could not parse environment variable EnvGtwaySqlType")
	}

	if cfm["EnvGtwaySqlConnection"] == "" {
		log.Fatal("Could not parse environment variable EnvGtwaySqlConnection")
	}

	return cfm
}
