package dorado

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// PowerOff create lun from source lun by LUN Clone.
func (d *Device) PowerOff(ctx context.Context, superAdminpassword string)  error {
	spath := "/SYSTEM/POWEROFF"

	jb, err := json.Marshal(map[string]interface{}{
		"IMPORTANTPSW": superAdminpassword,
	})
	if err != nil {
		return fmt.Errorf(ErrCreatePostValue+": %w", err)
	}
	req, err := d.newRequest(ctx, "POST", spath, bytes.NewBuffer(jb))
	if err != nil {
		return fmt.Errorf(ErrCreateRequest+": %w", err)
	}

	data := map[string]interface{}{}
	if err = d.requestWithRetry(req, &data, DefaultHTTPRetryCount); err != nil {
		return fmt.Errorf(ErrRequestWithRetry+": %w", err)
	}

	return nil
}
