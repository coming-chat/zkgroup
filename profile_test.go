package zkgroup

import (
	"encoding/base64"
	"testing"
)

func TestCreateProfileKeyCredentialRequestContext(t *testing.T) {
	zkGroupServerPublicParams, _ := base64.StdEncoding.DecodeString("AMh1gu/ongPtTUjIejLX8fWKvJo5HkW6ajb5X5IGq0dABBjMz4KPsJYZ5BJEAavUMC7d8qHyAGUiRs4uIlwubQ4qVhlpEZtd8jIDDHgS0Bqi0RXr9B9fcw8wrNoEcdcnX7hnOuZV/8nZQ1WQdmNiP7LR7EvTeFk/iEj7/UV7bux1pjvSNq4E964apj+2Pux2Xo9kNctJ0oWehW/3vujoiy8=")
	type args struct {
		serverPublicParams []byte
		uuid               []byte
		profileKey         []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				serverPublicParams: zkGroupServerPublicParams,
				uuid: []byte{
					165,
					99,
					46,
					73,
					204,
					142,
					71,
					25,
					164,
					54,
					154,
					16,
					54,
					166,
					72,
					73,
				},
				profileKey: []byte{
					178,
					41,
					47,
					237,
					227,
					49,
					131,
					193,
					136,
					139,
					128,
					56,
					228,
					178,
					107,
					165,
					69,
					52,
					145,
					68,
					97,
					162,
					104,
					122,
					3,
					64,
					54,
					42,
					153,
					248,
					76,
					204,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateProfileKeyCredentialRequestContext(tt.args.serverPublicParams, tt.args.uuid, tt.args.profileKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProfileKeyCredentialRequestContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%#v", got)
		})
	}
}
