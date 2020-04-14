package dorado

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

// NOTE(whywaita): implement only GET.
// storagePool is a few under our usage.

type StoragePool struct {
	ASSOCIATEOBJID         string `json:"ASSOCIATEOBJID"`
	DESCRIPTION            string `json:"DESCRIPTION"`
	ENCRYPTDISKTYPE        string `json:"ENCRYPTDISKTYPE"`
	ENGINEIDLIST           string `json:"ENGINEIDLIST"`
	ENGINEINFO             string `json:"ENGINEINFO"`
	FREECAPACITY           string `json:"FREECAPACITY"`
	FREECAPACITYLIST       string `json:"FREECAPACITYLIST"`
	HEALTHSTATUS           string `json:"HEALTHSTATUS"`
	ID                     string `json:"ID"`
	NAME                   string `json:"NAME"`
	OWNERCONTROLLERLIST    string `json:"OWNERCONTROLLERLIST"`
	RUNNINGSTATUS          string `json:"RUNNINGSTATUS"`
	SPARECAPACITY          string `json:"SPARECAPACITY"`
	SSDDISKNUM             string `json:"SSDDISKNUM"`
	SSDFREECAPACITY        string `json:"SSDFREECAPACITY"`
	SSDHOTSPARESTRATEGY    string `json:"SSDHOTSPARESTRATEGY"`
	SSDSPARECAPACITY       string `json:"SSDSPARECAPACITY"`
	SSDTOTALCAPACITY       string `json:"SSDTOTALCAPACITY"`
	SSDUSEDCAPACITY        string `json:"SSDUSEDCAPACITY"`
	SSDUSEDSPARECAPACITY   string `json:"SSDUSEDSPARECAPACITY"`
	TIER0DISKTYPE          string `json:"TIER0DISKTYPE"`
	TOTALCAPACITY          string `json:"TOTALCAPACITY"`
	TYPE                   int    `json:"TYPE"`
	USEDCAPACITY           string `json:"USEDCAPACITY"`
	USEDSPARECAPACITY      string `json:"USEDSPARECAPACITY"`
	AbrasionRate           string `json:"abrasionRate"`
	EnduranceBalanceStatus string `json:"enduranceBalanceStatus"`
	EngineCapacityDetail   string `json:"engineCapacityDetail"`
	RemainLife             string `json:"remainLife"`
	UnbalanceDiskIDList    string `json:"unbalanceDiskIdList"`
}

type StoragePools struct {
	COMPRESSEDCAPACITY              string `json:"COMPRESSEDCAPACITY"`
	COMPRESSINVOLVEDCAPACITY        string `json:"COMPRESSINVOLVEDCAPACITY"`
	COMPRESSIONRATE                 string `json:"COMPRESSIONRATE"`
	DATASPACE                       string `json:"DATASPACE"`
	DEDUPEDCAPACITY                 string `json:"DEDUPEDCAPACITY"`
	DEDUPINVOLVEDCAPACITY           string `json:"DEDUPINVOLVEDCAPACITY"`
	DEDUPLICATIONRATE               string `json:"DEDUPLICATIONRATE"`
	DESCRIPTION                     string `json:"DESCRIPTION"`
	ENDINGUPTHRESHOLD               string `json:"ENDINGUPTHRESHOLD"`
	HEALTHSTATUS                    string `json:"HEALTHSTATUS"`
	ID                              string `json:"ID"`
	LUNCONFIGEDCAPACITY             string `json:"LUNCONFIGEDCAPACITY"`
	NAME                            string `json:"NAME"`
	PARENTID                        string `json:"PARENTID"`
	PARENTNAME                      string `json:"PARENTNAME"`
	PARENTTYPE                      int    `json:"PARENTTYPE"`
	PROVISIONINGLIMIT               string `json:"PROVISIONINGLIMIT"`
	PROVISIONINGLIMITSWITCH         string `json:"PROVISIONINGLIMITSWITCH"`
	REDUCTIONINVOLVEDCAPACITY       string `json:"REDUCTIONINVOLVEDCAPACITY"`
	REPLICATIONCAPACITY             string `json:"REPLICATIONCAPACITY"`
	RUNNINGSTATUS                   string `json:"RUNNINGSTATUS"`
	SAVECAPACITYRATE                string `json:"SAVECAPACITYRATE"`
	SPACEREDUCTIONRATE              string `json:"SPACEREDUCTIONRATE"`
	THINPROVISIONSAVEPERCENTAGE     string `json:"THINPROVISIONSAVEPERCENTAGE"`
	TIER0CAPACITY                   string `json:"TIER0CAPACITY"`
	TIER0DISKTYPE                   string `json:"TIER0DISKTYPE"`
	TIER0RAIDLV                     string `json:"TIER0RAIDLV"`
	TOTALLUNWRITECAPACITY           string `json:"TOTALLUNWRITECAPACITY"`
	TYPE                            int    `json:"TYPE"`
	USAGETYPE                       string `json:"USAGETYPE"`
	USERCONSUMEDCAPACITY            string `json:"USERCONSUMEDCAPACITY"`
	USERCONSUMEDCAPACITYPERCENTAGE  string `json:"USERCONSUMEDCAPACITYPERCENTAGE"`
	USERCONSUMEDCAPACITYTHRESHOLD   string `json:"USERCONSUMEDCAPACITYTHRESHOLD"`
	USERCONSUMEDCAPACITYWITHOUTMETA string `json:"USERCONSUMEDCAPACITYWITHOUTMETA"`
	USERFREECAPACITY                string `json:"USERFREECAPACITY"`
	USERTOTALCAPACITY               string `json:"USERTOTALCAPACITY"`
	USERWRITEALLOCCAPACITY          string `json:"USERWRITEALLOCCAPACITY"`
	AutoDeleteSwitch                string `json:"autoDeleteSwitch"`
	PoolProtectHighThreshold        string `json:"poolProtectHighThreshold"`
	PoolProtectLowThreshold         string `json:"poolProtectLowThreshold"`
	ProtectSize                     string `json:"protectSize"`
	TotalSizeWithoutSnap            string `json:"totalSizeWithoutSnap"`
}

func (d *Device) GetStoragePools(ctx context.Context) ([]StoragePools, error) {
	spath := "/storagepool"
	req, err := d.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, errors.Wrap(err, ErrCreateRequest)
	}
	resp, err := d.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, ErrHTTPRequestDo)
	}

	storagePools := []StoragePools{}
	err = decodeBody(resp, &storagePools)
	if err != nil {
		return nil, errors.Wrap(err, ErrDecodeBody)
	}

	return storagePools, nil
}

func (d *Device) GetStoragePool(ctx context.Context, storagePoolId int) (*StoragePool, error) {
	var storagePool *StoragePool

	spath := fmt.Sprintf("/storagepool/%d", storagePoolId)
	req, err := d.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, errors.Wrap(err, ErrCreateRequest)
	}

	resp, err := d.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, ErrHTTPRequestDo)
	}

	err = decodeBody(resp, storagePool)
	if err != nil {
		return nil, errors.Wrap(err, ErrDecodeBody)
	}

	return storagePool, nil
}