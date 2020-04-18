package configuration

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Configuration struct {
	RedisHost     string
	RedisPort     uint16
	RedisDB       uint8
	RedisPassword string

	Transport uint8
}

const (
	TransportPubsub uint8 = iota
	TransportStream
	TransportFifo

	TransportPubsubString = "pubsub"
	TransportStreamString = "stream"
	TransportFifoString   = "fifo"

	DefaultConfigFile   = "redis-pubsub.yaml"
	DefaultConfigFolder = "."

	RedisHostKey     = "redis_host"
	RedisPortKey     = "redis_port"
	RedisDBKey       = "redis_db"
	RedisPasswordKey = "redis_password"
	TransportKey     = "transport"

	RedisHostDefaultValue     = "0.0.0.0"
	RedisPortDefaultValue     = 6379
	RedisDBDefaultValue       = 0
	RedisPasswordDefaultValue = ""

	Transport = TransportPubsub
)

// New method create a new configuration object
func New() *Configuration {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	alternativesConfigFolders := []string{
		dir,
		"$HOME",
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("redis_pubsub")

	viper.SetConfigName(DefaultConfigFile)
	viper.SetConfigType("yaml")

	viper.SetDefault(RedisHostKey, RedisHostDefaultValue)
	viper.SetDefault(RedisPortKey, RedisPortDefaultValue)
	viper.SetDefault(RedisDBKey, RedisDBDefaultValue)
	viper.SetDefault(RedisPasswordKey, RedisPasswordDefaultValue)
	viper.SetDefault(TransportKey, Transport)

	for _, alternativeConfigFolder := range alternativesConfigFolders {
		viper.AddConfigPath(alternativeConfigFolder)
	}
	// Set the default config folder as last default option
	viper.AddConfigPath(DefaultConfigFolder)

	// when configuration is created no error is shown if readinconfig files. It will use the defaults
	viper.ReadInConfig()

	return &Configuration{
		RedisHost:     viper.GetString(RedisHostKey),
		RedisPort:     uint16(viper.GetInt(RedisPortKey)),
		RedisDB:       uint8(viper.GetInt(RedisDBKey)),
		RedisPassword: viper.GetString(RedisPasswordKey),
		Transport:     TransportToUint8(viper.GetString(TransportKey)),
	}
}

func TransportToUint8(trasnportString string) uint8 {
	var transport uint8
	// read communication type
	switch trasnportString {

	case TransportPubsubString:
		transport = TransportPubsub
	case TransportStreamString:
		transport = TransportStream
	case TransportFifoString:
		transport = TransportFifo
	default:
		transport = TransportPubsub
	}

	return transport
}

// LoadFromFile method returns a configuration object loaded from a file
func LoadFromFile(file string) (*Configuration, error) {

	//	viper.Reset()
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.New("(Configuration::LoadFromFile) Error loading configuration from file '" + file + "' -> " + err.Error())
	}

	return &Configuration{
		RedisHost:     viper.GetString(RedisHostKey),
		RedisPort:     uint16(viper.GetInt(RedisPortKey)),
		RedisDB:       uint8(viper.GetInt(RedisDBKey)),
		RedisPassword: viper.GetString(RedisPasswordKey),
		Transport:     uint8(viper.GetInt(TransportKey)),
	}, nil

}

// ReloadConfigurationFromFile
func (c *Configuration) ReloadConfigurationFromFile(file string) error {
	newConfig, err := LoadFromFile(file)
	if err != nil {
		return errors.New("(Configuration::ReloadConfigurationFromFile) Error reloading configuration from file '" + file + "' -> " + err.Error())
	}

	*c = *newConfig
	return nil
}

func (c *Configuration) String() string {
	str := ""

	str = fmt.Sprintln()

	str = fmt.Sprintln(str, RedisHostKey, ": ", c.RedisHost)
	str = fmt.Sprintln(str, RedisPortKey, ": ", fmt.Sprint(c.RedisPort))
	str = fmt.Sprintln(str, RedisDBKey, ": ", fmt.Sprint(c.RedisDB))
	str = fmt.Sprintln(str, RedisPasswordKey, ": ", c.RedisPassword)

	return str
}

func (c *Configuration) ToArray() ([][]string, error) {

	if c == nil {
		return nil, errors.New("(Configuration::ToArray) Configuration is nil")
	}

	arrayConfig := [][]string{}
	arrayConfig = append(arrayConfig, []string{RedisHostKey, c.RedisHost})
	arrayConfig = append(arrayConfig, []string{RedisPortKey, fmt.Sprint(c.RedisPort)})
	arrayConfig = append(arrayConfig, []string{RedisDBKey, fmt.Sprint(c.RedisDB)})
	arrayConfig = append(arrayConfig, []string{RedisPasswordKey, c.RedisPassword})

	return arrayConfig, nil
}

func ConfigurationHeaders() []string {
	h := []string{
		"PARAMETER",
		"VALUE",
	}

	return h
}
