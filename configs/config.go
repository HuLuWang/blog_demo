package configs

import "time"

// 与yaml一致
type Server struct {
	System SystemSettingS   `mapstructure:"System" json:"System"`
	App    AppSettingS      `mapstructure:"App" json:"App"`
	DB     DatabaseSettingS `mapstructure:"Database" json:"Database"`
	Logger LoggerSettingS   `mapstructure:"Logger" json:"Logger"`
}

type SystemSettingS struct {
	RunMode      string        `mapstructure:"RunMode" json:"RunMode"`
	HttpPort     int           `mapstructure:"HttpPort" json:"HttpPort"`
	ReadTimeout  time.Duration `mapstructure:"ReadTimeout" json:"ReadTimeout"`
	WriteTimeout time.Duration `mapstructure:"WriteTimeout" json:"WriteTimeout"`
}

type AppSettingS struct {
	DefaultPageSize int `mapstructure:"DefaultPageSize" json:"DefaultPageSize"`
	MaxPageSize     int `mapstructure:"MaxPageSize" json:"MaxPageSize"`
}

type LoggerSettingS struct {
	LogSavePath string `mapstructure:"LogSavePath" json:"LogSavePath"`
	LogFileName string `mapstructure:"LogFileName" json:"LogFileName"`
	LogFileExt  string `mapstructure:"LogFileExt" json:"LogFileExt"`
}

type DatabaseSettingS struct {
	DBType       string `mapstructure:"DBType" json:"DBType"`
	UserName     string `mapstructure:"UserName" json:"UserName"`
	Password     string `mapstructure:"Password" json:"Password"`
	Host         string `mapstructure:"Host" json:"Host"`
	DBName       string `mapstructure:"DBName" json:"DBName"`
	TablePrefix  string `mapstructure:"TablePrefix" json:"TablePrefix"`
	Charset      string `mapstructure:"Charset" json:"Charset"`
	ParseTime    bool   `mapstructure:"ParseTime" json:"ParseTime"`
	MaxIdleConns int    `mapstructure:"MaxIdleConns" json:"MaxIdleConns"`
	MaxOpenConns int    `mapstructure:"MaxOpenConns" json:"MaxOpenConns"`
}
