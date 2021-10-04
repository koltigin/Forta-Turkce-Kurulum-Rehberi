package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"github.com/fatih/color"
	"github.com/forta-network/forta-node/config"
	"github.com/forta-network/forta-node/store"
)

const (
	contractAddressCacheExpiry = time.Hour
)

// useEnsDefaults gets and uses ENS defaults if needed.
func useEnsDefaults() error {
	if cfg.Registry.ContractAddress != "" && cfg.Query.PublishTo.ContractAddress != "" {
		return nil
	}

	return ensureLatestContractAddresses()
}

// useEnsAgentReg finds the agent registry from a contract.
func useEnsAgentReg() error {
	return ensureLatestContractAddresses()
}

func ensureLatestContractAddresses() error {
	now := time.Now().UTC()

	cache, ok := getContractAddressCache()
	if ok && now.Before(cache.ExpiresAt) {
		setContractAddressesFromCache(cache)
		return nil
	}

	whiteBold("Refreshing contract address cache...\n")

	ens, err := store.DialENSStore(getRegRpcUrl())
	if err != nil {
		return err
	}

	names := config.GetENSNames()
	cache.Dispatch, err = findContractAddress(ens, names.Dispatch)
	if err != nil {
		return err
	}
	cache.Alerts, err = findContractAddress(ens, names.Alerts)
	if err != nil {
		return err
	}
	cache.Agents, err = findContractAddress(ens, names.Agents)
	if err != nil {
		return err
	}
	cache.ExpiresAt = time.Now().UTC().Add(contractAddressCacheExpiry)

	b, err := json.MarshalIndent(&cache, "", "  ") // indent by two spaces
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(cfg.FortaDir, "contracts.json"), b, 0644); err != nil {
		return err
	}

	setContractAddressesFromCache(cache)

	return nil
}

func findContractAddress(ens store.ENS, input string) (string, error) {
	addr, err := ens.Resolve(input)
	if err != nil {
		return "", err
	}

	fmt.Printf("%s: %s\n", input, color.New(color.FgYellow).Sprintf(addr.String()))

	return addr.String(), nil
}

// sets only if not overridden
func setContractAddressesFromCache(cache contractAddressCache) {
	if cfg.Registry.ContractAddress == "" {
		cfg.Registry.ContractAddress = cache.Dispatch
	}
	if cfg.Query.PublishTo.ContractAddress == "" {
		cfg.Query.PublishTo.ContractAddress = cache.Alerts
	}
	cfg.AgentRegistryContractAddress = cache.Agents
}

type contractAddressCache struct {
	Dispatch  string    `json:"dispatch"`
	Alerts    string    `json:"alerts"`
	Agents    string    `json:"agents"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func getContractAddressCache() (cache contractAddressCache, ok bool) {
	b, err := ioutil.ReadFile(path.Join(cfg.FortaDir, "contracts.json"))
	if err != nil {
		return
	}

	if err := json.Unmarshal(b, &cache); err != nil {
		return
	}

	ok = true
	return
}

func getRegRpcUrl() string {
	if cfg.Registry.Ethereum.JsonRpcUrl != "" {
		return cfg.Registry.Ethereum.JsonRpcUrl
	}
	return cfg.Registry.Ethereum.WebsocketUrl
}
