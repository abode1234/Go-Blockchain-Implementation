package blockchain

import (
    "encoding/json"
    "os"
)

// json config data

type Config struct {
    Difficulty      int         `json: "difficulty"`
    Block_reward    int         `json: "block_reward"`
    Genesis_data    string      `json: "genesis_data"`
    Port            string      `json: "port"`
}

func LoadConfig(path string) (*Config, error) {

    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    config := &Config{}
    decoder := json.NewDecoder(file)
    err = decoder.Decode(config)
    return config, err
}

