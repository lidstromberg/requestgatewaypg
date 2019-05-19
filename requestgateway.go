package requestgateway

import (
	"database/sql"

	lbcf "github.com/lidstromberg/config"
	lblog "github.com/lidstromberg/log"

	//required for cloudsql socket library
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"golang.org/x/net/context"
)

//GtwyMgr handles interactions with the datastore
type GtwyMgr struct {
	ds *sql.DB
}

//NewMgr creates a new gateway manager
func NewMgr(ctx context.Context, bc lbcf.ConfigSetting) (*GtwyMgr, error) {
	preflight(ctx, bc)

	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "NewMgr", "info", "start")
	}

	db, err := sql.Open(bc.GetConfigValue(ctx, "EnvGtwaySqlType"), bc.GetConfigValue(ctx, "EnvGtwaySqlConnection"))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	cm1 := &GtwyMgr{
		ds: db,
	}

	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "NewMgr", "info", "end")
	}

	return cm1, nil
}

//Set sets a gateway address
func (gt GtwyMgr) Set(ctx context.Context, appcontext, remoteAddress string) error {
	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "Set", "info", "start")
	}

	_, err := gt.ds.Exec("select public.set_gateway($1,$2)", appcontext, remoteAddress)
	if err != nil {
		return err
	}

	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "Set", "info", "end")
	}
	return nil
}

//IsPermitted indicates if the address is approved
func (gt GtwyMgr) IsPermitted(ctx context.Context, appcontext, remoteAddress string) (bool, error) {
	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "IsPermitted", "info", "start")
	}

	var result bool

	//check the approval list
	err := gt.ds.QueryRow("select public.get_gateway($1,$2)", appcontext, remoteAddress).Scan(&result)
	if err != nil {
		return false, err
	}

	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "IsPermitted", "info", remoteAddress)
		lblog.LogEvent("GtwyMgr", "IsPermitted", "info", "end")
	}

	//otherwise the address is valid
	return result, nil
}

//Delete removes a gateway address
func (gt GtwyMgr) Delete(ctx context.Context, appcontext, remoteAddress string) error {
	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "Delete", "info", "start")
	}

	_, err := gt.ds.Exec("select public.delete_gateway($1,$2)", appcontext, remoteAddress)
	if err != nil {
		return err
	}

	if EnvDebugOn {
		lblog.LogEvent("GtwyMgr", "Delete", "info", "end")
	}
	return nil
}
