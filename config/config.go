package config

import (
    "encoding/json"
    "os"
)

type Inbound struct {
    Type   string `json:"type"`
    Listen string `json:"listen"`
    Method string `json:"method,omitempty"`
    Password string `json:"password,omitempty"`
}

type Outbound struct {
    Type string `json:"type"`
}

type Config struct {
    Inbounds  []Inbound  `json:"inbounds"`
    Outbounds []Outbound `json:"outbounds"`
}

func Load(path string) (*Config, error) {
    var cfg Config
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    err = json.Unmarshal(data, &cfg)
    return &cfg, err
}
