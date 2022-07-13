package dorado

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func (d *Device) UtcTime(ctx context.Context) (time.Time, error) {
	spath := "/system_utc_time"

	req, err := d.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return time.Time{}, fmt.Errorf(ErrCreateRequest+": %w", err)
	}

	var data struct {
		CMO_SYS_UTC_TIME string
	}
	if err = d.requestWithRetry(req, &data, DefaultHTTPRetryCount); err != nil {
		return time.Time{}, fmt.Errorf(ErrRequestWithRetry+": %w", err)
	}

	i64, err := strconv.ParseInt(data.CMO_SYS_UTC_TIME, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("CMO_SYS_UTC_TIME '%s' is invalid", data.CMO_SYS_UTC_TIME)
	}
	return time.Unix(i64, 0), nil
}

func (d *Device) PowerOff(ctx context.Context, superAdminpassword string) error {
	spath := "/SYSTEM/POWEROFF"

	jb, err := json.Marshal(map[string]interface{}{
		"IMPORTANTPSW": superAdminpassword,
	})
	if err != nil {
		return fmt.Errorf(ErrCreatePostValue+": %w", err)
	}
	req, err := d.newRequest(ctx, "PUT", spath, bytes.NewBuffer(jb))
	if err != nil {
		return fmt.Errorf(ErrCreateRequest+": %w", err)
	}

	data := map[string]interface{}{}
	if err = d.requestWithRetry(req, &data, DefaultHTTPRetryCount); err != nil {
		return fmt.Errorf(ErrRequestWithRetry+": %w", err)
	}

	return nil
}

func (d *Device) PowerReboot(ctx context.Context, superAdminpassword string) error {
	spath := "/SYSTEM/REBOOT"

	jb, err := json.Marshal(map[string]interface{}{
		"IMPORTANTPSW": superAdminpassword,
	})
	if err != nil {
		return fmt.Errorf(ErrCreatePostValue+": %w", err)
	}
	req, err := d.newRequest(ctx, "PUT", spath, bytes.NewBuffer(jb))
	if err != nil {
		return fmt.Errorf(ErrCreateRequest+": %w", err)
	}

	data := map[string]interface{}{}
	if err = d.requestWithRetry(req, &data, DefaultHTTPRetryCount); err != nil {
		return fmt.Errorf(ErrRequestWithRetry+": %w", err)
	}

	return nil
}
