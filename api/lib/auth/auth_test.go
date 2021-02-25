package auth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const base64PrivatePEM = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKSndJQkFBS0NBZ0VBdzc1Nk5mMFVCUVlSWWJHZ2tKZGdsRk5aYmZHNDFhLzlCd2xjNXZ2R2hGSWtZeXdqCitCcUxNYUU1U3ZQR0Q4ZS9DQ1JDRGx2TW1ZUjBuVnlWNGZzV1JmcFNGYTVZcVc3RkN2TE1kSFlmU3lBNXNrSm8Kb2hhWUlGc2lOd2xQVlRyOXFudVpKd3JsaEw5ekRqMnRmV0I2VHI4S3M3Zm5VN3dXeFZQYVMxSmIyaGo4R1BMNgpZcldzektJaXBaaEo3Z3NtUEREWFdCUEpwVlg5NDQybHRZSHh4L3NjeE5LWFlNeW81elBxeUpXaG03cVdSOEsvCkJEUSs4eXJXbjU4UGRQQnFyT0dmSzlqdU1CSkZGMkZJK0RwNkFCNGxQWjFtSUtoYVdFcFJnSExoNThFaW51dFcKOWZ3cjUvQWwzeU1ueGMzUVdPeitVYVd0b3laS0NTRndmYjF3QXBURkFrVG1qNGVhYW5OVURDWmdnNnNKeDNpdApEc2pVQW5rMVNyeFNkRjd6c0VyQnBMWDNGVHFUL3RtejJzV2xaaVRNZDAwZWpxWm8zdEhTblFKeXdwaHhsMDA0CjE1dWQvSGtYSzFXQnRrR2hNcDBUY3BQTWNZWnFSaHZnbDk3dnFpRmNnWFZONzdyRERuLytRN29XWTlRdVJHVlUKV1BMaTZhREhsY1hzdXlZZTdna3J3bkJScklodU5UOXlKL1NZcnVoUExTRitoUUVJZDFXUjM5SnJJY083Y2d3TApLb1hFcUZVd21CTThiVmtFTURWUTFwV1JWejBabHEyZk9pbnl2bTYxd1kyQnltbUxGUUtrWklWbER4clEyeEpkClZiTzFidmFBamxsZFozRUhnYzRFMW9CeHRsczN4ZkQrem1sOHNpa2NIakxQS25wKzh1RFlXcC9RNjVrQ0F3RUEKQVFLQ0FnQjFTTmdwS1M2cG8rMGVRRFFZN3RycmhOVjh6dTBVL0pIN2VWeTArZjhFb2NNenVPc0VhY01sUlpqeQpsQVlFeG9acjltMnQ5TXN1NFBLT3B6OFhYRDhJUnVpUUhScjZ5bWcrR3lUdVV5aUU5eFlhL1RkOGgxVTNiZU9lClhuR3VlOTRxSEV5ejNBK1I3clNkdjg4SDVKcmtQQXZKaTFPTUZKTUFRVEgzRjYzNWpDYmhQQlZTdDlDRi9GQU8KTUtWN2dDcTB2ZjhKd3pGN1kzN3dyWnF3bXQrb0Y0b3hWSTFuYnJsMWJ1SXF0WjFUbGdVZGtrNnAwalExdDJGSgoyeUNEek5uZUpJSWUzNmc0SFhwUWUvWWc1Y2piOFRRWEt3eGFYUVZsQ1lLMEdDSExueW1EVnplUHhEejZpQ1hwCkFVTi9mY1pzd2ZQUU1CbE9QRTc0RVpVNWdoeHVJYkZKSnBPU1hYc1JubkNCYUh3cXFRV3hqK3VPSitxSTQxOEEKUjhIMTNjVVNyaG1JeFp1dCtMQjVzNVhMNktmVnpIL09FalVTYmJoWjJhd21SaHZjd3BuYUVDOUpoNHdPMUpLYwpOblVDWDlLUml6c0FBSHVNNXlOeUN1L0dUUjZOR0lhU29SdjZQSmVYY2xNMkJRUG5NdVVmRGY0cnEyQmNHemkrClM0bVcvVys4M3BUZEwwK1ZVVFJTKzM1eklEbjRyYTJVUW5DRTV2Vm9lSk9XZGFwR0N1VU02L0VJU29DajFKZFQKSzlSRUl1c2V6RjJiamhNS1lwd2xCc2dQMmFobG5wdWlySXJYN3JGTklveG9pYlpvVnA0TUplU0xpMHJXaTk2bgpmTlhjTEs1YzFiRDFGRW9VVzBpS0I2TkdXaURmejdZb2l4Z052NDBKTlBvNnNhd3VRUUtDQVFFQStTbUorWU1yCjk2TnZpdDBHRllUa2hNcjRhZEdKdmc4L1ZTeHlNV3EralZ2dm5JQ0ZCT2ZhenRqZ25TcktadjdYOWsvdGM2N3YKVk5rUjN3UlZ2K1NPQjVaMjdPRXh3ZE1PVmVLOS9NZXBRNnVOTUVaRllRc0dzM1FhUDMxa2UyRUFzbUc2YmRhSgpTTDlzdXFYWFBXWXp0T0VoS0NTZ1M3RHQ5RDZCbW9pRm1FbThjdEllUnFuWWFOL1VwSjBCL0dwbG5wVERYNTdVCkxHMHJxUTJEVWlZaVhRTER3QithdzdVa1A2RkRpOEVqQ016TVo1Zi9YUWU4MktGME1XeTBwUHhpUm5sdlY1S0QKeldrTitNWURWd01Zb3RLbGRQZkZyV0h6Q0VwOFArbk9aam1ZZEtUTVVlbVBOcSthNnN4enZ0clAzWld3TzA5MwpscEdGM1pRa0NtUU9CUUtDQVFFQXlSMm5wczlyc1F3aWRSWGpGd0I1YnhwUFdPK0I0K0lLcnlnK3YwNUM1ZkNlCjhURjFjR2svM2xtMFFoNkVDRUlyRDFiL0lhenMvQUtsdmNCdWhIMm5WM2dVdE9BNnA2bDFLVVY1L0xFN01nZVUKRW1vUDVBNnplOVhhbldZeENUOFBUQkI2clZwR3RLZlhBZy9aUU83TVI4d1JuZGkrTW8wN0MxaS9jVnQ3b1Q0VwpXUHFyYVRvUXhoTVZSTmpCaDFacHlGRjk2alhOVDJwTTNFcEVnSGt1R1JIREJYKzdFMGo1bElKZkVlY3lzVlV2ClZ2QUw4M2ZYYmE4aVNKVmM5VVo1VVdkRER3Ui81RHAyaFFWekFSRTlqR3dIVkFMTWY3TXptY3FjVm9xcmFQcUwKdTBpTlhqdUlxWi9KRTdqbVRRVXNZR1p1WURpVXExKzNvZk14aG4ySGhRS0NBUUI0S1ZaSEpNRnhmanB1NThUbApYOSt0UXcybVNzMTVBWFZ6cUNteU4xNktZY29lMmNSTTUxd1k0WE1CbVA5ZnlJdHlXSDJWaXNvVlMxSlpFWWdoCk1TSmwrbVNFUE41NE1VYjZtSTB6ZVQ3aVNidWZpbVF2TnRnV2QybXBNTm5pdnBkTmIrQXUxSVlFdFh6RVR2S00KdzlzdjRsclJOMGl1K0RicDBiTkRTWS9VTDh0WVBJa3BYd1BsSC8wM0hoazFHRUxGeGN5ck1yZjBiUG5mWDRyegpkVHloU3BJSTk3VDFxVWcvLzQvSnVHMGk4MUdvckI4VlNJUUVuV2loNVdFQktFNWsybStkOWlUbVRVSFZ0ZmtxClgyM2tLRjV2R0ljVzVPUUdIWlhxWk9HTFh6OXRFWVVRQStsellDRUFGM1hDbDFnajd1cTh6OGhHcnd1MkhwbDcKQVdIWkFvSUJBQjRlb2RnYW9sOCttUDUrQzZlTE82U0hCVEVsbFlkaWVBVXBldFE5elVrUWswMCtBZitXMDZqaQpPRnZhcEIxMGcyeGx0QW9BRXZIZkY2Rm1hMmJPUnJ5VDBFNFNjdFpmUzV1bHV4SThITWh1V0IrMTRMRmYva05pCmtMNFg2dW9lbHBUbXR1aTFaM2R3MTRPSWlobnVhWXVySlV4RmhKNmZoaU01NUZuK3dISXlrVGc4T21Xays1UWUKa1lYaERJTFBUMEpmNmdLa2toMzlwb3NyV0QwQmFRVDZJd0gwMGppZUtqcXlsN2hmcnJqZU1Cdk9FWXdoKzVLQwpzeGk1dmRWQis3TlVTcmU2RGNsSmRDeVIxMXBta1pneEtadk1XNElZZlhiZjREQVp6bmdIWmR2ampzQjVIQXVEClVhTHhneFF2M1NpR3pxcjhiaytJSTBGTzRiRHBwZDBDZ2dFQWZaalFFOGxzTmN1VUlLVjRrZE92WVBhRG9RMnUKLzYxQk9jVXNodzBCTFVZc25ydHM1OVZFMGUwc25OWmpTQlVDdUNoM1k3SkoxY0wwalJCelRXTDhROGgxc1dxVgpVcDhCV0VhSjdDWGJHSjlnalBRd1JtVXJPVmMrVittTlphTW5VYXYrbktDMGZlbWZWaml0alFCR0N1Q0RoQVhHCndodTJENlNTSmUwOUljUHpnMlFrdmdZNmNKQVdZUDIvZmFndlg3ektoVU9qOGRrQkJ3UkltUkswVUJobzdtZjgKUDN5TlFkR2wvMEhJM0w1c2RwQ2NMWEpGRlg4bENBakd6M1FWMDc0eTFlaEFldThyOGZEc28wZXJqdnVHcXp3SQp0Q3JGeHd6ck1WRjBZeVJtQXlld1JReU45MGZ2eEJxWG5SeXlzVVhCbGJEc2loYmhLVjRXZzFWNmZBPT0KLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K"

func Test_ParseBase64PrivatePEM(t *testing.T) {
	a := assert.New(t)
	_, err := ParseBase64PrivatePEM(base64PrivatePEM)
	a.NoError(err)
}

func Test_JwtSigner(t *testing.T) {
	a := assert.New(t)
	priv, err := ParseBase64PrivatePEM(base64PrivatePEM)
	a.NoError(err)
	signer := JwtSigner{
		PrivateKey: priv,
	}
	jwt, err := signer.Sign(NewClaims("bob@gmail.com"))
	a.NoError(err)
	fmt.Println(jwt)
}
