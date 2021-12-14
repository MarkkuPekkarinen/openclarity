package config

import (
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

const (
	Verbose                 = "VERBOSE"
	ListeningPort           = "LISTENING_PORT"
	GrypeAddress            = "GRYPE_ADDR"
	KlarTrace               = "KLAR_TRACE"
	KlarResultServicePath   = "KLAR_RESULT_SERVICE_PATH"
	DockleImageName         = "DOCKLE_IMAGE_NAME"
	DockleTimeoutSec        = "DOCKLE_TIMEOUT_SEC"
	DockleResultServicePath = "DOCKLE_RESULT_SERVICE_PATH"
	ResultListenPort        = "RESULT_LISTEN_PORT"
	ScannerHttpsProxy       = "SCANNER_HTTPS_PROXY"
	ScannerHttpProxy        = "SCANNER_HTTP_PROXY"
	CredsSecretNamespace    = "CREDS_SECRET_NAMESPACE"
)

type Config struct {
	Verbose                 bool
	WebappPort              string
	GrypeAddress            string
	KlarTrace               bool
	KlarResultServicePath   string
	DockleImageName         string
	DockleTimeoutSec        string
	DockleResultServicePath string
	ResultListenPort        string
	ResultServiceAddress    string
	ScannerHttpsProxy       string
	ScannerHttpProxy        string
	CredsSecretNamespace    string
}

func setConfigDefaults() {
	viper.SetDefault(Verbose, "false")
	viper.SetDefault(ListeningPort, "8080")
	viper.SetDefault(KlarResultServicePath, "http://kubei.kubei:8081/result/")
	viper.SetDefault(DockleTimeoutSec, "300")
	viper.SetDefault(DockleResultServicePath, "http://kubei.kubei:8081/dockerfileScanResult/")
	viper.SetDefault(ResultListenPort, "8081")
	viper.SetDefault(KlarTrace, "false") // Run Klar in more verbose mode
	viper.SetDefault(GrypeAddress, "grype-server.kubei:9991")
	viper.SetDefault(ScannerHttpsProxy, "")
	viper.SetDefault(ScannerHttpProxy, "")
	viper.SetDefault(CredsSecretNamespace, "kubei")

	viper.AutomaticEnv()
}

// Extracts service hostname from the full url.
// example: `http://kubei.kubei:8081/result/` should return `kubei.kubei`
func getServiceAddress(serviceFullPath string) string {
	u, err := url.Parse(serviceFullPath)
	if err != nil {
		panic(err)
	}

	return u.Hostname()
}

// Add default http scheme in case scheme is missing.
func getServiceFullPath(serviceFullPath string) string {
	if !strings.Contains(serviceFullPath, "://") {
		serviceFullPath = "http://" + serviceFullPath
	}

	return serviceFullPath
}

func LoadConfig() *Config {
	setConfigDefaults()

	klarResultServicePath := getServiceFullPath(viper.GetString(KlarResultServicePath))
	dockleResultServicePath := getServiceFullPath(viper.GetString(DockleResultServicePath))

	config := &Config{
		Verbose:                 viper.GetBool(Verbose),
		WebappPort:              viper.GetString(ListeningPort),
		GrypeAddress:            viper.GetString(GrypeAddress),
		KlarTrace:               viper.GetBool(KlarTrace),
		KlarResultServicePath:   klarResultServicePath,
		DockleImageName:         viper.GetString(DockleImageName),
		DockleTimeoutSec:        viper.GetString(DockleTimeoutSec),
		DockleResultServicePath: dockleResultServicePath,
		ResultListenPort:        viper.GetString(ResultListenPort),
		ResultServiceAddress:    getServiceAddress(klarResultServicePath),
		ScannerHttpsProxy:       viper.GetString(ScannerHttpsProxy),
		ScannerHttpProxy:        viper.GetString(ScannerHttpProxy),
		CredsSecretNamespace:    viper.GetString(CredsSecretNamespace),
	}

	return config
}
