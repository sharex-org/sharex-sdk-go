// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deshare

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CountryInfo is an auto generated low-level Go binding around an user-defined struct.
type CountryInfo struct {
	Iso2      [2]byte
	Timestamp uint32
}

// DetailedStatsInfo is an auto generated low-level Go binding around an user-defined struct.
type DetailedStatsInfo struct {
	BasicStats         StatsInfo
	MetricNames        []string
	MetricDescriptions []string
	LastUpdated        *big.Int
	ContractVersion    Version
}

// DeviceInfo is an auto generated low-level Go binding around an user-defined struct.
type DeviceInfo struct {
	Id          uint64
	Timestamp   uint32
	DeviceId    string
	DeviceType  string
	PartnerCode string
	MerchantId  string
}

// DeviceParams is an auto generated low-level Go binding around an user-defined struct.
type DeviceParams struct {
	DeviceId    string
	DeviceType  string
	PartnerCode string
	MerchantId  string
}

// MerchantInfo is an auto generated low-level Go binding around an user-defined struct.
type MerchantInfo struct {
	Id           uint64
	Timestamp    uint32
	Iso2         [2]byte
	MerchantName string
	MerchantId   string
	LocationId   string
	Location     string
	MerchantType string
	Verification string
	Description  string
}

// MerchantParams is an auto generated low-level Go binding around an user-defined struct.
type MerchantParams struct {
	MerchantName string
	MerchantId   string
	Description  string
	Iso2         [2]byte
	LocationId   string
	Location     string
	MerchantType string
	Verification string
}

// PartnerInfo is an auto generated low-level Go binding around an user-defined struct.
type PartnerInfo struct {
	Id           uint64
	Timestamp    uint32
	Iso2         [2]byte
	PartnerCode  string
	Verification string
	PartnerName  string
	Description  string
	BusinessType string
}

// PartnerParams is an auto generated low-level Go binding around an user-defined struct.
type PartnerParams struct {
	PartnerCode  string
	PartnerName  string
	Iso2         [2]byte
	Verification string
	Description  string
	BusinessType string
}

// StatsInfo is an auto generated low-level Go binding around an user-defined struct.
type StatsInfo struct {
	PartnersCount           uint64
	MerchantsCount          uint64
	DevicesCount            uint64
	TransactionBatchesCount uint64
	CountriesCount          uint64
}

// TransactionBatch is an auto generated low-level Go binding around an user-defined struct.
type TransactionBatch struct {
	Id             uint64
	OrderCount     uint32
	BatchTimestamp uint32
	DeviceId       string
	TotalAmount    string
	DateComparable string
}

// UploadBatchParams is an auto generated low-level Go binding around an user-defined struct.
type UploadBatchParams struct {
	DeviceId        string
	DateComparable  string
	OrderCount      uint32
	TotalAmount     string
	TransactionData []byte
}

// Version is an auto generated low-level Go binding around an user-defined struct.
type Version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

// DeshareMetaData contains all meta data concerning the Deshare contract.
var DeshareMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyTransactionData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entityType\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"entityId\",\"type\":\"bytes32\"}],\"name\":\"EntityAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entityType\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"entityId\",\"type\":\"bytes32\"}],\"name\":\"EntityNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxLength\",\"type\":\"uint256\"}],\"name\":\"InvalidStringLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderCountMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyTransactions\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"major\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"minor\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"patch\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structVersion\",\"name\":\"version\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ContractInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CountryRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"deviceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"deviceType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"partnerCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merchantId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DeviceRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"merchantName\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"merchantId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MerchantRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"partnerCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"partnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PartnerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"deviceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"orderCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"totalAmount\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dateComparable\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TransactionBatchUploaded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TransactionDataUploaded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"}],\"name\":\"countryExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"deviceId\",\"type\":\"string\"}],\"name\":\"deviceExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"}],\"name\":\"getCountry\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structCountryInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDetailedStats\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"partnersCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"merchantsCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"devicesCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionBatchesCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"countriesCount\",\"type\":\"uint64\"}],\"internalType\":\"structStatsInfo\",\"name\":\"basicStats\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"metricNames\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"metricDescriptions\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"major\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"minor\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"patch\",\"type\":\"uint8\"}],\"internalType\":\"structVersion\",\"name\":\"contractVersion\",\"type\":\"tuple\"}],\"internalType\":\"structDetailedStatsInfo\",\"name\":\"detailedStats\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deviceId\",\"type\":\"uint256\"}],\"name\":\"getDevice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"deviceId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"deviceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"partnerCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"merchantId\",\"type\":\"string\"}],\"internalType\":\"structDeviceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"deviceId\",\"type\":\"string\"}],\"name\":\"getDeviceById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"deviceId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"deviceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"partnerCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"merchantId\",\"type\":\"string\"}],\"internalType\":\"structDeviceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"merchantId\",\"type\":\"uint256\"}],\"name\":\"getMerchant\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"},{\"internalType\":\"string\",\"name\":\"merchantName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"merchantId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"locationId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"merchantType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verification\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structMerchantInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"merchantId\",\"type\":\"string\"}],\"name\":\"getMerchantById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"},{\"internalType\":\"string\",\"name\":\"merchantName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"merchantId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"locationId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"merchantType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verification\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structMerchantInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"},{\"internalType\":\"string\",\"name\":\"locationId\",\"type\":\"string\"}],\"name\":\"getMerchantCountByRegion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"}],\"name\":\"getPartner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"},{\"internalType\":\"string\",\"name\":\"partnerCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verification\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"partnerName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"businessType\",\"type\":\"string\"}],\"internalType\":\"structPartnerInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"partnerCode\",\"type\":\"string\"}],\"name\":\"getPartnerByCode\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"},{\"internalType\":\"string\",\"name\":\"partnerCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verification\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"partnerName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"businessType\",\"type\":\"string\"}],\"internalType\":\"structPartnerInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"businessType\",\"type\":\"string\"}],\"name\":\"getPartnerCountByBusinessType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStats\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"partnersCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"merchantsCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"devicesCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transactionBatchesCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"countriesCount\",\"type\":\"uint64\"}],\"internalType\":\"structStatsInfo\",\"name\":\"stats\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"getTransactionBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"orderCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"batchTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"deviceId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"totalAmount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateComparable\",\"type\":\"string\"}],\"internalType\":\"structTransactionBatch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"getTransactionData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"jsonData\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"major\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"minor\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"patch\",\"type\":\"uint8\"}],\"internalType\":\"structVersion\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"merchantId\",\"type\":\"string\"}],\"name\":\"merchantExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"partnerCode\",\"type\":\"string\"}],\"name\":\"partnerExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"}],\"name\":\"registerCountry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"deviceId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"deviceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"partnerCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"merchantId\",\"type\":\"string\"}],\"internalType\":\"structDeviceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"registerDevice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"merchantName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"merchantId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"},{\"internalType\":\"string\",\"name\":\"locationId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"merchantType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verification\",\"type\":\"string\"}],\"internalType\":\"structMerchantParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"registerMerchant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"partnerCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"partnerName\",\"type\":\"string\"},{\"internalType\":\"bytes2\",\"name\":\"iso2\",\"type\":\"bytes2\"},{\"internalType\":\"string\",\"name\":\"verification\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"businessType\",\"type\":\"string\"}],\"internalType\":\"structPartnerParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"registerPartner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"deviceId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateComparable\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"orderCount\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"totalAmount\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"transactionData\",\"type\":\"bytes\"}],\"internalType\":\"structUploadBatchParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"uploadTransactionBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DeshareABI is the input ABI used to generate the binding from.
// Deprecated: Use DeshareMetaData.ABI instead.
var DeshareABI = DeshareMetaData.ABI

// Deshare is an auto generated Go binding around an Ethereum contract.
type Deshare struct {
	DeshareCaller     // Read-only binding to the contract
	DeshareTransactor // Write-only binding to the contract
	DeshareFilterer   // Log filterer for contract events
}

// DeshareCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeshareCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeshareTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeshareTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeshareFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeshareFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeshareSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeshareSession struct {
	Contract     *Deshare          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeshareCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeshareCallerSession struct {
	Contract *DeshareCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DeshareTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeshareTransactorSession struct {
	Contract     *DeshareTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DeshareRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeshareRaw struct {
	Contract *Deshare // Generic contract binding to access the raw methods on
}

// DeshareCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeshareCallerRaw struct {
	Contract *DeshareCaller // Generic read-only contract binding to access the raw methods on
}

// DeshareTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeshareTransactorRaw struct {
	Contract *DeshareTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeshare creates a new instance of Deshare, bound to a specific deployed contract.
func NewDeshare(address common.Address, backend bind.ContractBackend) (*Deshare, error) {
	contract, err := bindDeshare(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deshare{DeshareCaller: DeshareCaller{contract: contract}, DeshareTransactor: DeshareTransactor{contract: contract}, DeshareFilterer: DeshareFilterer{contract: contract}}, nil
}

// NewDeshareCaller creates a new read-only instance of Deshare, bound to a specific deployed contract.
func NewDeshareCaller(address common.Address, caller bind.ContractCaller) (*DeshareCaller, error) {
	contract, err := bindDeshare(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeshareCaller{contract: contract}, nil
}

// NewDeshareTransactor creates a new write-only instance of Deshare, bound to a specific deployed contract.
func NewDeshareTransactor(address common.Address, transactor bind.ContractTransactor) (*DeshareTransactor, error) {
	contract, err := bindDeshare(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeshareTransactor{contract: contract}, nil
}

// NewDeshareFilterer creates a new log filterer instance of Deshare, bound to a specific deployed contract.
func NewDeshareFilterer(address common.Address, filterer bind.ContractFilterer) (*DeshareFilterer, error) {
	contract, err := bindDeshare(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeshareFilterer{contract: contract}, nil
}

// bindDeshare binds a generic wrapper to an already deployed contract.
func bindDeshare(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DeshareMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deshare *DeshareRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deshare.Contract.DeshareCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deshare *DeshareRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deshare.Contract.DeshareTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deshare *DeshareRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deshare.Contract.DeshareTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deshare *DeshareCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deshare.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deshare *DeshareTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deshare.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deshare *DeshareTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deshare.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Deshare *DeshareCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Deshare *DeshareSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Deshare.Contract.DEFAULTADMINROLE(&_Deshare.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Deshare *DeshareCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Deshare.Contract.DEFAULTADMINROLE(&_Deshare.CallOpts)
}

// CountryExists is a free data retrieval call binding the contract method 0x9a1d0ed9.
//
// Solidity: function countryExists(bytes2 iso2) view returns(bool exists)
func (_Deshare *DeshareCaller) CountryExists(opts *bind.CallOpts, iso2 [2]byte) (bool, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "countryExists", iso2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CountryExists is a free data retrieval call binding the contract method 0x9a1d0ed9.
//
// Solidity: function countryExists(bytes2 iso2) view returns(bool exists)
func (_Deshare *DeshareSession) CountryExists(iso2 [2]byte) (bool, error) {
	return _Deshare.Contract.CountryExists(&_Deshare.CallOpts, iso2)
}

// CountryExists is a free data retrieval call binding the contract method 0x9a1d0ed9.
//
// Solidity: function countryExists(bytes2 iso2) view returns(bool exists)
func (_Deshare *DeshareCallerSession) CountryExists(iso2 [2]byte) (bool, error) {
	return _Deshare.Contract.CountryExists(&_Deshare.CallOpts, iso2)
}

// DeviceExists is a free data retrieval call binding the contract method 0xaaed589d.
//
// Solidity: function deviceExists(string deviceId) view returns(bool exists)
func (_Deshare *DeshareCaller) DeviceExists(opts *bind.CallOpts, deviceId string) (bool, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "deviceExists", deviceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DeviceExists is a free data retrieval call binding the contract method 0xaaed589d.
//
// Solidity: function deviceExists(string deviceId) view returns(bool exists)
func (_Deshare *DeshareSession) DeviceExists(deviceId string) (bool, error) {
	return _Deshare.Contract.DeviceExists(&_Deshare.CallOpts, deviceId)
}

// DeviceExists is a free data retrieval call binding the contract method 0xaaed589d.
//
// Solidity: function deviceExists(string deviceId) view returns(bool exists)
func (_Deshare *DeshareCallerSession) DeviceExists(deviceId string) (bool, error) {
	return _Deshare.Contract.DeviceExists(&_Deshare.CallOpts, deviceId)
}

// GetCountry is a free data retrieval call binding the contract method 0xb16f9ab0.
//
// Solidity: function getCountry(bytes2 iso2) view returns((bytes2,uint32))
func (_Deshare *DeshareCaller) GetCountry(opts *bind.CallOpts, iso2 [2]byte) (CountryInfo, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getCountry", iso2)

	if err != nil {
		return *new(CountryInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CountryInfo)).(*CountryInfo)

	return out0, err

}

// GetCountry is a free data retrieval call binding the contract method 0xb16f9ab0.
//
// Solidity: function getCountry(bytes2 iso2) view returns((bytes2,uint32))
func (_Deshare *DeshareSession) GetCountry(iso2 [2]byte) (CountryInfo, error) {
	return _Deshare.Contract.GetCountry(&_Deshare.CallOpts, iso2)
}

// GetCountry is a free data retrieval call binding the contract method 0xb16f9ab0.
//
// Solidity: function getCountry(bytes2 iso2) view returns((bytes2,uint32))
func (_Deshare *DeshareCallerSession) GetCountry(iso2 [2]byte) (CountryInfo, error) {
	return _Deshare.Contract.GetCountry(&_Deshare.CallOpts, iso2)
}

// GetDetailedStats is a free data retrieval call binding the contract method 0x44a0cceb.
//
// Solidity: function getDetailedStats() view returns(((uint64,uint64,uint64,uint64,uint64),string[],string[],uint256,(uint8,uint8,uint8)) detailedStats)
func (_Deshare *DeshareCaller) GetDetailedStats(opts *bind.CallOpts) (DetailedStatsInfo, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getDetailedStats")

	if err != nil {
		return *new(DetailedStatsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DetailedStatsInfo)).(*DetailedStatsInfo)

	return out0, err

}

// GetDetailedStats is a free data retrieval call binding the contract method 0x44a0cceb.
//
// Solidity: function getDetailedStats() view returns(((uint64,uint64,uint64,uint64,uint64),string[],string[],uint256,(uint8,uint8,uint8)) detailedStats)
func (_Deshare *DeshareSession) GetDetailedStats() (DetailedStatsInfo, error) {
	return _Deshare.Contract.GetDetailedStats(&_Deshare.CallOpts)
}

// GetDetailedStats is a free data retrieval call binding the contract method 0x44a0cceb.
//
// Solidity: function getDetailedStats() view returns(((uint64,uint64,uint64,uint64,uint64),string[],string[],uint256,(uint8,uint8,uint8)) detailedStats)
func (_Deshare *DeshareCallerSession) GetDetailedStats() (DetailedStatsInfo, error) {
	return _Deshare.Contract.GetDetailedStats(&_Deshare.CallOpts)
}

// GetDevice is a free data retrieval call binding the contract method 0x09f110b1.
//
// Solidity: function getDevice(uint256 deviceId) view returns((uint64,uint32,string,string,string,string))
func (_Deshare *DeshareCaller) GetDevice(opts *bind.CallOpts, deviceId *big.Int) (DeviceInfo, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getDevice", deviceId)

	if err != nil {
		return *new(DeviceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DeviceInfo)).(*DeviceInfo)

	return out0, err

}

// GetDevice is a free data retrieval call binding the contract method 0x09f110b1.
//
// Solidity: function getDevice(uint256 deviceId) view returns((uint64,uint32,string,string,string,string))
func (_Deshare *DeshareSession) GetDevice(deviceId *big.Int) (DeviceInfo, error) {
	return _Deshare.Contract.GetDevice(&_Deshare.CallOpts, deviceId)
}

// GetDevice is a free data retrieval call binding the contract method 0x09f110b1.
//
// Solidity: function getDevice(uint256 deviceId) view returns((uint64,uint32,string,string,string,string))
func (_Deshare *DeshareCallerSession) GetDevice(deviceId *big.Int) (DeviceInfo, error) {
	return _Deshare.Contract.GetDevice(&_Deshare.CallOpts, deviceId)
}

// GetDeviceById is a free data retrieval call binding the contract method 0x12a7ebe6.
//
// Solidity: function getDeviceById(string deviceId) view returns((uint64,uint32,string,string,string,string))
func (_Deshare *DeshareCaller) GetDeviceById(opts *bind.CallOpts, deviceId string) (DeviceInfo, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getDeviceById", deviceId)

	if err != nil {
		return *new(DeviceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DeviceInfo)).(*DeviceInfo)

	return out0, err

}

// GetDeviceById is a free data retrieval call binding the contract method 0x12a7ebe6.
//
// Solidity: function getDeviceById(string deviceId) view returns((uint64,uint32,string,string,string,string))
func (_Deshare *DeshareSession) GetDeviceById(deviceId string) (DeviceInfo, error) {
	return _Deshare.Contract.GetDeviceById(&_Deshare.CallOpts, deviceId)
}

// GetDeviceById is a free data retrieval call binding the contract method 0x12a7ebe6.
//
// Solidity: function getDeviceById(string deviceId) view returns((uint64,uint32,string,string,string,string))
func (_Deshare *DeshareCallerSession) GetDeviceById(deviceId string) (DeviceInfo, error) {
	return _Deshare.Contract.GetDeviceById(&_Deshare.CallOpts, deviceId)
}

// GetMerchant is a free data retrieval call binding the contract method 0x5d32798a.
//
// Solidity: function getMerchant(uint256 merchantId) view returns((uint64,uint32,bytes2,string,string,string,string,string,string,string))
func (_Deshare *DeshareCaller) GetMerchant(opts *bind.CallOpts, merchantId *big.Int) (MerchantInfo, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getMerchant", merchantId)

	if err != nil {
		return *new(MerchantInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(MerchantInfo)).(*MerchantInfo)

	return out0, err

}

// GetMerchant is a free data retrieval call binding the contract method 0x5d32798a.
//
// Solidity: function getMerchant(uint256 merchantId) view returns((uint64,uint32,bytes2,string,string,string,string,string,string,string))
func (_Deshare *DeshareSession) GetMerchant(merchantId *big.Int) (MerchantInfo, error) {
	return _Deshare.Contract.GetMerchant(&_Deshare.CallOpts, merchantId)
}

// GetMerchant is a free data retrieval call binding the contract method 0x5d32798a.
//
// Solidity: function getMerchant(uint256 merchantId) view returns((uint64,uint32,bytes2,string,string,string,string,string,string,string))
func (_Deshare *DeshareCallerSession) GetMerchant(merchantId *big.Int) (MerchantInfo, error) {
	return _Deshare.Contract.GetMerchant(&_Deshare.CallOpts, merchantId)
}

// GetMerchantById is a free data retrieval call binding the contract method 0xc97f8c7f.
//
// Solidity: function getMerchantById(string merchantId) view returns((uint64,uint32,bytes2,string,string,string,string,string,string,string))
func (_Deshare *DeshareCaller) GetMerchantById(opts *bind.CallOpts, merchantId string) (MerchantInfo, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getMerchantById", merchantId)

	if err != nil {
		return *new(MerchantInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(MerchantInfo)).(*MerchantInfo)

	return out0, err

}

// GetMerchantById is a free data retrieval call binding the contract method 0xc97f8c7f.
//
// Solidity: function getMerchantById(string merchantId) view returns((uint64,uint32,bytes2,string,string,string,string,string,string,string))
func (_Deshare *DeshareSession) GetMerchantById(merchantId string) (MerchantInfo, error) {
	return _Deshare.Contract.GetMerchantById(&_Deshare.CallOpts, merchantId)
}

// GetMerchantById is a free data retrieval call binding the contract method 0xc97f8c7f.
//
// Solidity: function getMerchantById(string merchantId) view returns((uint64,uint32,bytes2,string,string,string,string,string,string,string))
func (_Deshare *DeshareCallerSession) GetMerchantById(merchantId string) (MerchantInfo, error) {
	return _Deshare.Contract.GetMerchantById(&_Deshare.CallOpts, merchantId)
}

// GetMerchantCountByRegion is a free data retrieval call binding the contract method 0x675715c2.
//
// Solidity: function getMerchantCountByRegion(bytes2 iso2, string locationId) view returns(uint256 count)
func (_Deshare *DeshareCaller) GetMerchantCountByRegion(opts *bind.CallOpts, iso2 [2]byte, locationId string) (*big.Int, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getMerchantCountByRegion", iso2, locationId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerchantCountByRegion is a free data retrieval call binding the contract method 0x675715c2.
//
// Solidity: function getMerchantCountByRegion(bytes2 iso2, string locationId) view returns(uint256 count)
func (_Deshare *DeshareSession) GetMerchantCountByRegion(iso2 [2]byte, locationId string) (*big.Int, error) {
	return _Deshare.Contract.GetMerchantCountByRegion(&_Deshare.CallOpts, iso2, locationId)
}

// GetMerchantCountByRegion is a free data retrieval call binding the contract method 0x675715c2.
//
// Solidity: function getMerchantCountByRegion(bytes2 iso2, string locationId) view returns(uint256 count)
func (_Deshare *DeshareCallerSession) GetMerchantCountByRegion(iso2 [2]byte, locationId string) (*big.Int, error) {
	return _Deshare.Contract.GetMerchantCountByRegion(&_Deshare.CallOpts, iso2, locationId)
}

// GetPartner is a free data retrieval call binding the contract method 0x7fc96619.
//
// Solidity: function getPartner(uint256 partnerId) view returns((uint64,uint32,bytes2,string,string,string,string,string))
func (_Deshare *DeshareCaller) GetPartner(opts *bind.CallOpts, partnerId *big.Int) (PartnerInfo, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getPartner", partnerId)

	if err != nil {
		return *new(PartnerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PartnerInfo)).(*PartnerInfo)

	return out0, err

}

// GetPartner is a free data retrieval call binding the contract method 0x7fc96619.
//
// Solidity: function getPartner(uint256 partnerId) view returns((uint64,uint32,bytes2,string,string,string,string,string))
func (_Deshare *DeshareSession) GetPartner(partnerId *big.Int) (PartnerInfo, error) {
	return _Deshare.Contract.GetPartner(&_Deshare.CallOpts, partnerId)
}

// GetPartner is a free data retrieval call binding the contract method 0x7fc96619.
//
// Solidity: function getPartner(uint256 partnerId) view returns((uint64,uint32,bytes2,string,string,string,string,string))
func (_Deshare *DeshareCallerSession) GetPartner(partnerId *big.Int) (PartnerInfo, error) {
	return _Deshare.Contract.GetPartner(&_Deshare.CallOpts, partnerId)
}

// GetPartnerByCode is a free data retrieval call binding the contract method 0x18fae941.
//
// Solidity: function getPartnerByCode(string partnerCode) view returns((uint64,uint32,bytes2,string,string,string,string,string))
func (_Deshare *DeshareCaller) GetPartnerByCode(opts *bind.CallOpts, partnerCode string) (PartnerInfo, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getPartnerByCode", partnerCode)

	if err != nil {
		return *new(PartnerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PartnerInfo)).(*PartnerInfo)

	return out0, err

}

// GetPartnerByCode is a free data retrieval call binding the contract method 0x18fae941.
//
// Solidity: function getPartnerByCode(string partnerCode) view returns((uint64,uint32,bytes2,string,string,string,string,string))
func (_Deshare *DeshareSession) GetPartnerByCode(partnerCode string) (PartnerInfo, error) {
	return _Deshare.Contract.GetPartnerByCode(&_Deshare.CallOpts, partnerCode)
}

// GetPartnerByCode is a free data retrieval call binding the contract method 0x18fae941.
//
// Solidity: function getPartnerByCode(string partnerCode) view returns((uint64,uint32,bytes2,string,string,string,string,string))
func (_Deshare *DeshareCallerSession) GetPartnerByCode(partnerCode string) (PartnerInfo, error) {
	return _Deshare.Contract.GetPartnerByCode(&_Deshare.CallOpts, partnerCode)
}

// GetPartnerCountByBusinessType is a free data retrieval call binding the contract method 0x96eeafb2.
//
// Solidity: function getPartnerCountByBusinessType(string businessType) view returns(uint256 count)
func (_Deshare *DeshareCaller) GetPartnerCountByBusinessType(opts *bind.CallOpts, businessType string) (*big.Int, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getPartnerCountByBusinessType", businessType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPartnerCountByBusinessType is a free data retrieval call binding the contract method 0x96eeafb2.
//
// Solidity: function getPartnerCountByBusinessType(string businessType) view returns(uint256 count)
func (_Deshare *DeshareSession) GetPartnerCountByBusinessType(businessType string) (*big.Int, error) {
	return _Deshare.Contract.GetPartnerCountByBusinessType(&_Deshare.CallOpts, businessType)
}

// GetPartnerCountByBusinessType is a free data retrieval call binding the contract method 0x96eeafb2.
//
// Solidity: function getPartnerCountByBusinessType(string businessType) view returns(uint256 count)
func (_Deshare *DeshareCallerSession) GetPartnerCountByBusinessType(businessType string) (*big.Int, error) {
	return _Deshare.Contract.GetPartnerCountByBusinessType(&_Deshare.CallOpts, businessType)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Deshare *DeshareCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Deshare *DeshareSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Deshare.Contract.GetRoleAdmin(&_Deshare.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Deshare *DeshareCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Deshare.Contract.GetRoleAdmin(&_Deshare.CallOpts, role)
}

// GetStats is a free data retrieval call binding the contract method 0xc59d4847.
//
// Solidity: function getStats() view returns((uint64,uint64,uint64,uint64,uint64) stats)
func (_Deshare *DeshareCaller) GetStats(opts *bind.CallOpts) (StatsInfo, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getStats")

	if err != nil {
		return *new(StatsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(StatsInfo)).(*StatsInfo)

	return out0, err

}

// GetStats is a free data retrieval call binding the contract method 0xc59d4847.
//
// Solidity: function getStats() view returns((uint64,uint64,uint64,uint64,uint64) stats)
func (_Deshare *DeshareSession) GetStats() (StatsInfo, error) {
	return _Deshare.Contract.GetStats(&_Deshare.CallOpts)
}

// GetStats is a free data retrieval call binding the contract method 0xc59d4847.
//
// Solidity: function getStats() view returns((uint64,uint64,uint64,uint64,uint64) stats)
func (_Deshare *DeshareCallerSession) GetStats() (StatsInfo, error) {
	return _Deshare.Contract.GetStats(&_Deshare.CallOpts)
}

// GetTransactionBatch is a free data retrieval call binding the contract method 0xe3b41fb0.
//
// Solidity: function getTransactionBatch(uint256 batchId) view returns((uint64,uint32,uint32,string,string,string))
func (_Deshare *DeshareCaller) GetTransactionBatch(opts *bind.CallOpts, batchId *big.Int) (TransactionBatch, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getTransactionBatch", batchId)

	if err != nil {
		return *new(TransactionBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(TransactionBatch)).(*TransactionBatch)

	return out0, err

}

// GetTransactionBatch is a free data retrieval call binding the contract method 0xe3b41fb0.
//
// Solidity: function getTransactionBatch(uint256 batchId) view returns((uint64,uint32,uint32,string,string,string))
func (_Deshare *DeshareSession) GetTransactionBatch(batchId *big.Int) (TransactionBatch, error) {
	return _Deshare.Contract.GetTransactionBatch(&_Deshare.CallOpts, batchId)
}

// GetTransactionBatch is a free data retrieval call binding the contract method 0xe3b41fb0.
//
// Solidity: function getTransactionBatch(uint256 batchId) view returns((uint64,uint32,uint32,string,string,string))
func (_Deshare *DeshareCallerSession) GetTransactionBatch(batchId *big.Int) (TransactionBatch, error) {
	return _Deshare.Contract.GetTransactionBatch(&_Deshare.CallOpts, batchId)
}

// GetTransactionData is a free data retrieval call binding the contract method 0x3b24b547.
//
// Solidity: function getTransactionData(uint256 batchId) view returns(string jsonData)
func (_Deshare *DeshareCaller) GetTransactionData(opts *bind.CallOpts, batchId *big.Int) (string, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getTransactionData", batchId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTransactionData is a free data retrieval call binding the contract method 0x3b24b547.
//
// Solidity: function getTransactionData(uint256 batchId) view returns(string jsonData)
func (_Deshare *DeshareSession) GetTransactionData(batchId *big.Int) (string, error) {
	return _Deshare.Contract.GetTransactionData(&_Deshare.CallOpts, batchId)
}

// GetTransactionData is a free data retrieval call binding the contract method 0x3b24b547.
//
// Solidity: function getTransactionData(uint256 batchId) view returns(string jsonData)
func (_Deshare *DeshareCallerSession) GetTransactionData(batchId *big.Int) (string, error) {
	return _Deshare.Contract.GetTransactionData(&_Deshare.CallOpts, batchId)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns((uint8,uint8,uint8))
func (_Deshare *DeshareCaller) GetVersion(opts *bind.CallOpts) (Version, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(Version), err
	}

	out0 := *abi.ConvertType(out[0], new(Version)).(*Version)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns((uint8,uint8,uint8))
func (_Deshare *DeshareSession) GetVersion() (Version, error) {
	return _Deshare.Contract.GetVersion(&_Deshare.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns((uint8,uint8,uint8))
func (_Deshare *DeshareCallerSession) GetVersion() (Version, error) {
	return _Deshare.Contract.GetVersion(&_Deshare.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Deshare *DeshareCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Deshare *DeshareSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Deshare.Contract.HasRole(&_Deshare.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Deshare *DeshareCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Deshare.Contract.HasRole(&_Deshare.CallOpts, role, account)
}

// MerchantExists is a free data retrieval call binding the contract method 0x7f473266.
//
// Solidity: function merchantExists(string merchantId) view returns(bool exists)
func (_Deshare *DeshareCaller) MerchantExists(opts *bind.CallOpts, merchantId string) (bool, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "merchantExists", merchantId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MerchantExists is a free data retrieval call binding the contract method 0x7f473266.
//
// Solidity: function merchantExists(string merchantId) view returns(bool exists)
func (_Deshare *DeshareSession) MerchantExists(merchantId string) (bool, error) {
	return _Deshare.Contract.MerchantExists(&_Deshare.CallOpts, merchantId)
}

// MerchantExists is a free data retrieval call binding the contract method 0x7f473266.
//
// Solidity: function merchantExists(string merchantId) view returns(bool exists)
func (_Deshare *DeshareCallerSession) MerchantExists(merchantId string) (bool, error) {
	return _Deshare.Contract.MerchantExists(&_Deshare.CallOpts, merchantId)
}

// PartnerExists is a free data retrieval call binding the contract method 0x70036c16.
//
// Solidity: function partnerExists(string partnerCode) view returns(bool exists)
func (_Deshare *DeshareCaller) PartnerExists(opts *bind.CallOpts, partnerCode string) (bool, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "partnerExists", partnerCode)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PartnerExists is a free data retrieval call binding the contract method 0x70036c16.
//
// Solidity: function partnerExists(string partnerCode) view returns(bool exists)
func (_Deshare *DeshareSession) PartnerExists(partnerCode string) (bool, error) {
	return _Deshare.Contract.PartnerExists(&_Deshare.CallOpts, partnerCode)
}

// PartnerExists is a free data retrieval call binding the contract method 0x70036c16.
//
// Solidity: function partnerExists(string partnerCode) view returns(bool exists)
func (_Deshare *DeshareCallerSession) PartnerExists(partnerCode string) (bool, error) {
	return _Deshare.Contract.PartnerExists(&_Deshare.CallOpts, partnerCode)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Deshare *DeshareCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Deshare.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Deshare *DeshareSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Deshare.Contract.SupportsInterface(&_Deshare.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Deshare *DeshareCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Deshare.Contract.SupportsInterface(&_Deshare.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Deshare *DeshareTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deshare.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Deshare *DeshareSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deshare.Contract.GrantRole(&_Deshare.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Deshare *DeshareTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deshare.Contract.GrantRole(&_Deshare.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) payable returns()
func (_Deshare *DeshareTransactor) Initialize(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Deshare.contract.Transact(opts, "initialize", admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) payable returns()
func (_Deshare *DeshareSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _Deshare.Contract.Initialize(&_Deshare.TransactOpts, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) payable returns()
func (_Deshare *DeshareTransactorSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _Deshare.Contract.Initialize(&_Deshare.TransactOpts, admin)
}

// RegisterCountry is a paid mutator transaction binding the contract method 0x290f7c78.
//
// Solidity: function registerCountry(bytes2 iso2) returns()
func (_Deshare *DeshareTransactor) RegisterCountry(opts *bind.TransactOpts, iso2 [2]byte) (*types.Transaction, error) {
	return _Deshare.contract.Transact(opts, "registerCountry", iso2)
}

// RegisterCountry is a paid mutator transaction binding the contract method 0x290f7c78.
//
// Solidity: function registerCountry(bytes2 iso2) returns()
func (_Deshare *DeshareSession) RegisterCountry(iso2 [2]byte) (*types.Transaction, error) {
	return _Deshare.Contract.RegisterCountry(&_Deshare.TransactOpts, iso2)
}

// RegisterCountry is a paid mutator transaction binding the contract method 0x290f7c78.
//
// Solidity: function registerCountry(bytes2 iso2) returns()
func (_Deshare *DeshareTransactorSession) RegisterCountry(iso2 [2]byte) (*types.Transaction, error) {
	return _Deshare.Contract.RegisterCountry(&_Deshare.TransactOpts, iso2)
}

// RegisterDevice is a paid mutator transaction binding the contract method 0xe8ed7dfb.
//
// Solidity: function registerDevice((string,string,string,string) params) returns()
func (_Deshare *DeshareTransactor) RegisterDevice(opts *bind.TransactOpts, params DeviceParams) (*types.Transaction, error) {
	return _Deshare.contract.Transact(opts, "registerDevice", params)
}

// RegisterDevice is a paid mutator transaction binding the contract method 0xe8ed7dfb.
//
// Solidity: function registerDevice((string,string,string,string) params) returns()
func (_Deshare *DeshareSession) RegisterDevice(params DeviceParams) (*types.Transaction, error) {
	return _Deshare.Contract.RegisterDevice(&_Deshare.TransactOpts, params)
}

// RegisterDevice is a paid mutator transaction binding the contract method 0xe8ed7dfb.
//
// Solidity: function registerDevice((string,string,string,string) params) returns()
func (_Deshare *DeshareTransactorSession) RegisterDevice(params DeviceParams) (*types.Transaction, error) {
	return _Deshare.Contract.RegisterDevice(&_Deshare.TransactOpts, params)
}

// RegisterMerchant is a paid mutator transaction binding the contract method 0x3ee5290d.
//
// Solidity: function registerMerchant((string,string,string,bytes2,string,string,string,string) params) returns()
func (_Deshare *DeshareTransactor) RegisterMerchant(opts *bind.TransactOpts, params MerchantParams) (*types.Transaction, error) {
	return _Deshare.contract.Transact(opts, "registerMerchant", params)
}

// RegisterMerchant is a paid mutator transaction binding the contract method 0x3ee5290d.
//
// Solidity: function registerMerchant((string,string,string,bytes2,string,string,string,string) params) returns()
func (_Deshare *DeshareSession) RegisterMerchant(params MerchantParams) (*types.Transaction, error) {
	return _Deshare.Contract.RegisterMerchant(&_Deshare.TransactOpts, params)
}

// RegisterMerchant is a paid mutator transaction binding the contract method 0x3ee5290d.
//
// Solidity: function registerMerchant((string,string,string,bytes2,string,string,string,string) params) returns()
func (_Deshare *DeshareTransactorSession) RegisterMerchant(params MerchantParams) (*types.Transaction, error) {
	return _Deshare.Contract.RegisterMerchant(&_Deshare.TransactOpts, params)
}

// RegisterPartner is a paid mutator transaction binding the contract method 0xaeb0f5e2.
//
// Solidity: function registerPartner((string,string,bytes2,string,string,string) params) returns()
func (_Deshare *DeshareTransactor) RegisterPartner(opts *bind.TransactOpts, params PartnerParams) (*types.Transaction, error) {
	return _Deshare.contract.Transact(opts, "registerPartner", params)
}

// RegisterPartner is a paid mutator transaction binding the contract method 0xaeb0f5e2.
//
// Solidity: function registerPartner((string,string,bytes2,string,string,string) params) returns()
func (_Deshare *DeshareSession) RegisterPartner(params PartnerParams) (*types.Transaction, error) {
	return _Deshare.Contract.RegisterPartner(&_Deshare.TransactOpts, params)
}

// RegisterPartner is a paid mutator transaction binding the contract method 0xaeb0f5e2.
//
// Solidity: function registerPartner((string,string,bytes2,string,string,string) params) returns()
func (_Deshare *DeshareTransactorSession) RegisterPartner(params PartnerParams) (*types.Transaction, error) {
	return _Deshare.Contract.RegisterPartner(&_Deshare.TransactOpts, params)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Deshare *DeshareTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Deshare.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Deshare *DeshareSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Deshare.Contract.RenounceRole(&_Deshare.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Deshare *DeshareTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Deshare.Contract.RenounceRole(&_Deshare.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Deshare *DeshareTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deshare.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Deshare *DeshareSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deshare.Contract.RevokeRole(&_Deshare.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Deshare *DeshareTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deshare.Contract.RevokeRole(&_Deshare.TransactOpts, role, account)
}

// UploadTransactionBatch is a paid mutator transaction binding the contract method 0xce300cab.
//
// Solidity: function uploadTransactionBatch((string,string,uint32,string,bytes) params) returns()
func (_Deshare *DeshareTransactor) UploadTransactionBatch(opts *bind.TransactOpts, params UploadBatchParams) (*types.Transaction, error) {
	return _Deshare.contract.Transact(opts, "uploadTransactionBatch", params)
}

// UploadTransactionBatch is a paid mutator transaction binding the contract method 0xce300cab.
//
// Solidity: function uploadTransactionBatch((string,string,uint32,string,bytes) params) returns()
func (_Deshare *DeshareSession) UploadTransactionBatch(params UploadBatchParams) (*types.Transaction, error) {
	return _Deshare.Contract.UploadTransactionBatch(&_Deshare.TransactOpts, params)
}

// UploadTransactionBatch is a paid mutator transaction binding the contract method 0xce300cab.
//
// Solidity: function uploadTransactionBatch((string,string,uint32,string,bytes) params) returns()
func (_Deshare *DeshareTransactorSession) UploadTransactionBatch(params UploadBatchParams) (*types.Transaction, error) {
	return _Deshare.Contract.UploadTransactionBatch(&_Deshare.TransactOpts, params)
}

// DeshareContractInitializedIterator is returned from FilterContractInitialized and is used to iterate over the raw logs and unpacked data for ContractInitialized events raised by the Deshare contract.
type DeshareContractInitializedIterator struct {
	Event *DeshareContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeshareContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeshareContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeshareContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeshareContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeshareContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeshareContractInitialized represents a ContractInitialized event raised by the Deshare contract.
type DeshareContractInitialized struct {
	Admin     common.Address
	Version   Version
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContractInitialized is a free log retrieval operation binding the contract event 0x4509b6016c9d4ea94d6f3835eee5db5d1a62863febb6d4f83b449b4bff4afa47.
//
// Solidity: event ContractInitialized(address indexed admin, (uint8,uint8,uint8) version, uint256 timestamp)
func (_Deshare *DeshareFilterer) FilterContractInitialized(opts *bind.FilterOpts, admin []common.Address) (*DeshareContractInitializedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Deshare.contract.FilterLogs(opts, "ContractInitialized", adminRule)
	if err != nil {
		return nil, err
	}
	return &DeshareContractInitializedIterator{contract: _Deshare.contract, event: "ContractInitialized", logs: logs, sub: sub}, nil
}

// WatchContractInitialized is a free log subscription operation binding the contract event 0x4509b6016c9d4ea94d6f3835eee5db5d1a62863febb6d4f83b449b4bff4afa47.
//
// Solidity: event ContractInitialized(address indexed admin, (uint8,uint8,uint8) version, uint256 timestamp)
func (_Deshare *DeshareFilterer) WatchContractInitialized(opts *bind.WatchOpts, sink chan<- *DeshareContractInitialized, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Deshare.contract.WatchLogs(opts, "ContractInitialized", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeshareContractInitialized)
				if err := _Deshare.contract.UnpackLog(event, "ContractInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractInitialized is a log parse operation binding the contract event 0x4509b6016c9d4ea94d6f3835eee5db5d1a62863febb6d4f83b449b4bff4afa47.
//
// Solidity: event ContractInitialized(address indexed admin, (uint8,uint8,uint8) version, uint256 timestamp)
func (_Deshare *DeshareFilterer) ParseContractInitialized(log types.Log) (*DeshareContractInitialized, error) {
	event := new(DeshareContractInitialized)
	if err := _Deshare.contract.UnpackLog(event, "ContractInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeshareCountryRegisteredIterator is returned from FilterCountryRegistered and is used to iterate over the raw logs and unpacked data for CountryRegistered events raised by the Deshare contract.
type DeshareCountryRegisteredIterator struct {
	Event *DeshareCountryRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeshareCountryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeshareCountryRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeshareCountryRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeshareCountryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeshareCountryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeshareCountryRegistered represents a CountryRegistered event raised by the Deshare contract.
type DeshareCountryRegistered struct {
	Iso2      [2]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCountryRegistered is a free log retrieval operation binding the contract event 0x69d32a8886b260f80ec184b9fd78e3e308fa49ccd3afbd7edee87b4a2973e0c5.
//
// Solidity: event CountryRegistered(bytes2 indexed iso2, uint256 timestamp)
func (_Deshare *DeshareFilterer) FilterCountryRegistered(opts *bind.FilterOpts, iso2 [][2]byte) (*DeshareCountryRegisteredIterator, error) {

	var iso2Rule []interface{}
	for _, iso2Item := range iso2 {
		iso2Rule = append(iso2Rule, iso2Item)
	}

	logs, sub, err := _Deshare.contract.FilterLogs(opts, "CountryRegistered", iso2Rule)
	if err != nil {
		return nil, err
	}
	return &DeshareCountryRegisteredIterator{contract: _Deshare.contract, event: "CountryRegistered", logs: logs, sub: sub}, nil
}

// WatchCountryRegistered is a free log subscription operation binding the contract event 0x69d32a8886b260f80ec184b9fd78e3e308fa49ccd3afbd7edee87b4a2973e0c5.
//
// Solidity: event CountryRegistered(bytes2 indexed iso2, uint256 timestamp)
func (_Deshare *DeshareFilterer) WatchCountryRegistered(opts *bind.WatchOpts, sink chan<- *DeshareCountryRegistered, iso2 [][2]byte) (event.Subscription, error) {

	var iso2Rule []interface{}
	for _, iso2Item := range iso2 {
		iso2Rule = append(iso2Rule, iso2Item)
	}

	logs, sub, err := _Deshare.contract.WatchLogs(opts, "CountryRegistered", iso2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeshareCountryRegistered)
				if err := _Deshare.contract.UnpackLog(event, "CountryRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCountryRegistered is a log parse operation binding the contract event 0x69d32a8886b260f80ec184b9fd78e3e308fa49ccd3afbd7edee87b4a2973e0c5.
//
// Solidity: event CountryRegistered(bytes2 indexed iso2, uint256 timestamp)
func (_Deshare *DeshareFilterer) ParseCountryRegistered(log types.Log) (*DeshareCountryRegistered, error) {
	event := new(DeshareCountryRegistered)
	if err := _Deshare.contract.UnpackLog(event, "CountryRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeshareDeviceRegisteredIterator is returned from FilterDeviceRegistered and is used to iterate over the raw logs and unpacked data for DeviceRegistered events raised by the Deshare contract.
type DeshareDeviceRegisteredIterator struct {
	Event *DeshareDeviceRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeshareDeviceRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeshareDeviceRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeshareDeviceRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeshareDeviceRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeshareDeviceRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeshareDeviceRegistered represents a DeviceRegistered event raised by the Deshare contract.
type DeshareDeviceRegistered struct {
	Id          *big.Int
	DeviceId    [32]byte
	DeviceType  [32]byte
	PartnerCode [32]byte
	MerchantId  [32]byte
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeviceRegistered is a free log retrieval operation binding the contract event 0x53e4c9b67f9a34da5dfbabe4c5eb8badcf129f3a7f409b966c4b6accc5941395.
//
// Solidity: event DeviceRegistered(uint256 indexed id, bytes32 indexed deviceId, bytes32 deviceType, bytes32 partnerCode, bytes32 merchantId, uint256 timestamp)
func (_Deshare *DeshareFilterer) FilterDeviceRegistered(opts *bind.FilterOpts, id []*big.Int, deviceId [][32]byte) (*DeshareDeviceRegisteredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var deviceIdRule []interface{}
	for _, deviceIdItem := range deviceId {
		deviceIdRule = append(deviceIdRule, deviceIdItem)
	}

	logs, sub, err := _Deshare.contract.FilterLogs(opts, "DeviceRegistered", idRule, deviceIdRule)
	if err != nil {
		return nil, err
	}
	return &DeshareDeviceRegisteredIterator{contract: _Deshare.contract, event: "DeviceRegistered", logs: logs, sub: sub}, nil
}

// WatchDeviceRegistered is a free log subscription operation binding the contract event 0x53e4c9b67f9a34da5dfbabe4c5eb8badcf129f3a7f409b966c4b6accc5941395.
//
// Solidity: event DeviceRegistered(uint256 indexed id, bytes32 indexed deviceId, bytes32 deviceType, bytes32 partnerCode, bytes32 merchantId, uint256 timestamp)
func (_Deshare *DeshareFilterer) WatchDeviceRegistered(opts *bind.WatchOpts, sink chan<- *DeshareDeviceRegistered, id []*big.Int, deviceId [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var deviceIdRule []interface{}
	for _, deviceIdItem := range deviceId {
		deviceIdRule = append(deviceIdRule, deviceIdItem)
	}

	logs, sub, err := _Deshare.contract.WatchLogs(opts, "DeviceRegistered", idRule, deviceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeshareDeviceRegistered)
				if err := _Deshare.contract.UnpackLog(event, "DeviceRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeviceRegistered is a log parse operation binding the contract event 0x53e4c9b67f9a34da5dfbabe4c5eb8badcf129f3a7f409b966c4b6accc5941395.
//
// Solidity: event DeviceRegistered(uint256 indexed id, bytes32 indexed deviceId, bytes32 deviceType, bytes32 partnerCode, bytes32 merchantId, uint256 timestamp)
func (_Deshare *DeshareFilterer) ParseDeviceRegistered(log types.Log) (*DeshareDeviceRegistered, error) {
	event := new(DeshareDeviceRegistered)
	if err := _Deshare.contract.UnpackLog(event, "DeviceRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeshareInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Deshare contract.
type DeshareInitializedIterator struct {
	Event *DeshareInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeshareInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeshareInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeshareInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeshareInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeshareInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeshareInitialized represents a Initialized event raised by the Deshare contract.
type DeshareInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Deshare *DeshareFilterer) FilterInitialized(opts *bind.FilterOpts) (*DeshareInitializedIterator, error) {

	logs, sub, err := _Deshare.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DeshareInitializedIterator{contract: _Deshare.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Deshare *DeshareFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DeshareInitialized) (event.Subscription, error) {

	logs, sub, err := _Deshare.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeshareInitialized)
				if err := _Deshare.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Deshare *DeshareFilterer) ParseInitialized(log types.Log) (*DeshareInitialized, error) {
	event := new(DeshareInitialized)
	if err := _Deshare.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeshareMerchantRegisteredIterator is returned from FilterMerchantRegistered and is used to iterate over the raw logs and unpacked data for MerchantRegistered events raised by the Deshare contract.
type DeshareMerchantRegisteredIterator struct {
	Event *DeshareMerchantRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeshareMerchantRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeshareMerchantRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeshareMerchantRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeshareMerchantRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeshareMerchantRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeshareMerchantRegistered represents a MerchantRegistered event raised by the Deshare contract.
type DeshareMerchantRegistered struct {
	Id           *big.Int
	MerchantName [32]byte
	MerchantId   [32]byte
	Iso2         [2]byte
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMerchantRegistered is a free log retrieval operation binding the contract event 0x62883210ffa0b95c4d9e7cac152eab1bd2abff2d795540f09065ea734a17a449.
//
// Solidity: event MerchantRegistered(uint256 indexed id, bytes32 indexed merchantName, bytes32 indexed merchantId, bytes2 iso2, uint256 timestamp)
func (_Deshare *DeshareFilterer) FilterMerchantRegistered(opts *bind.FilterOpts, id []*big.Int, merchantName [][32]byte, merchantId [][32]byte) (*DeshareMerchantRegisteredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var merchantNameRule []interface{}
	for _, merchantNameItem := range merchantName {
		merchantNameRule = append(merchantNameRule, merchantNameItem)
	}
	var merchantIdRule []interface{}
	for _, merchantIdItem := range merchantId {
		merchantIdRule = append(merchantIdRule, merchantIdItem)
	}

	logs, sub, err := _Deshare.contract.FilterLogs(opts, "MerchantRegistered", idRule, merchantNameRule, merchantIdRule)
	if err != nil {
		return nil, err
	}
	return &DeshareMerchantRegisteredIterator{contract: _Deshare.contract, event: "MerchantRegistered", logs: logs, sub: sub}, nil
}

// WatchMerchantRegistered is a free log subscription operation binding the contract event 0x62883210ffa0b95c4d9e7cac152eab1bd2abff2d795540f09065ea734a17a449.
//
// Solidity: event MerchantRegistered(uint256 indexed id, bytes32 indexed merchantName, bytes32 indexed merchantId, bytes2 iso2, uint256 timestamp)
func (_Deshare *DeshareFilterer) WatchMerchantRegistered(opts *bind.WatchOpts, sink chan<- *DeshareMerchantRegistered, id []*big.Int, merchantName [][32]byte, merchantId [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var merchantNameRule []interface{}
	for _, merchantNameItem := range merchantName {
		merchantNameRule = append(merchantNameRule, merchantNameItem)
	}
	var merchantIdRule []interface{}
	for _, merchantIdItem := range merchantId {
		merchantIdRule = append(merchantIdRule, merchantIdItem)
	}

	logs, sub, err := _Deshare.contract.WatchLogs(opts, "MerchantRegistered", idRule, merchantNameRule, merchantIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeshareMerchantRegistered)
				if err := _Deshare.contract.UnpackLog(event, "MerchantRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMerchantRegistered is a log parse operation binding the contract event 0x62883210ffa0b95c4d9e7cac152eab1bd2abff2d795540f09065ea734a17a449.
//
// Solidity: event MerchantRegistered(uint256 indexed id, bytes32 indexed merchantName, bytes32 indexed merchantId, bytes2 iso2, uint256 timestamp)
func (_Deshare *DeshareFilterer) ParseMerchantRegistered(log types.Log) (*DeshareMerchantRegistered, error) {
	event := new(DeshareMerchantRegistered)
	if err := _Deshare.contract.UnpackLog(event, "MerchantRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DesharePartnerRegisteredIterator is returned from FilterPartnerRegistered and is used to iterate over the raw logs and unpacked data for PartnerRegistered events raised by the Deshare contract.
type DesharePartnerRegisteredIterator struct {
	Event *DesharePartnerRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DesharePartnerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DesharePartnerRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DesharePartnerRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DesharePartnerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DesharePartnerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DesharePartnerRegistered represents a PartnerRegistered event raised by the Deshare contract.
type DesharePartnerRegistered struct {
	Id          *big.Int
	PartnerCode [32]byte
	PartnerName [32]byte
	Iso2        [2]byte
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPartnerRegistered is a free log retrieval operation binding the contract event 0x5eb4a8754c947ee86e5f620753122383f5ff5d23b6fe64ca7822f210a4e615ab.
//
// Solidity: event PartnerRegistered(uint256 indexed id, bytes32 indexed partnerCode, bytes32 partnerName, bytes2 iso2, uint256 timestamp)
func (_Deshare *DeshareFilterer) FilterPartnerRegistered(opts *bind.FilterOpts, id []*big.Int, partnerCode [][32]byte) (*DesharePartnerRegisteredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var partnerCodeRule []interface{}
	for _, partnerCodeItem := range partnerCode {
		partnerCodeRule = append(partnerCodeRule, partnerCodeItem)
	}

	logs, sub, err := _Deshare.contract.FilterLogs(opts, "PartnerRegistered", idRule, partnerCodeRule)
	if err != nil {
		return nil, err
	}
	return &DesharePartnerRegisteredIterator{contract: _Deshare.contract, event: "PartnerRegistered", logs: logs, sub: sub}, nil
}

// WatchPartnerRegistered is a free log subscription operation binding the contract event 0x5eb4a8754c947ee86e5f620753122383f5ff5d23b6fe64ca7822f210a4e615ab.
//
// Solidity: event PartnerRegistered(uint256 indexed id, bytes32 indexed partnerCode, bytes32 partnerName, bytes2 iso2, uint256 timestamp)
func (_Deshare *DeshareFilterer) WatchPartnerRegistered(opts *bind.WatchOpts, sink chan<- *DesharePartnerRegistered, id []*big.Int, partnerCode [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var partnerCodeRule []interface{}
	for _, partnerCodeItem := range partnerCode {
		partnerCodeRule = append(partnerCodeRule, partnerCodeItem)
	}

	logs, sub, err := _Deshare.contract.WatchLogs(opts, "PartnerRegistered", idRule, partnerCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DesharePartnerRegistered)
				if err := _Deshare.contract.UnpackLog(event, "PartnerRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePartnerRegistered is a log parse operation binding the contract event 0x5eb4a8754c947ee86e5f620753122383f5ff5d23b6fe64ca7822f210a4e615ab.
//
// Solidity: event PartnerRegistered(uint256 indexed id, bytes32 indexed partnerCode, bytes32 partnerName, bytes2 iso2, uint256 timestamp)
func (_Deshare *DeshareFilterer) ParsePartnerRegistered(log types.Log) (*DesharePartnerRegistered, error) {
	event := new(DesharePartnerRegistered)
	if err := _Deshare.contract.UnpackLog(event, "PartnerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeshareRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Deshare contract.
type DeshareRoleAdminChangedIterator struct {
	Event *DeshareRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeshareRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeshareRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeshareRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeshareRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeshareRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeshareRoleAdminChanged represents a RoleAdminChanged event raised by the Deshare contract.
type DeshareRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Deshare *DeshareFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DeshareRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Deshare.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DeshareRoleAdminChangedIterator{contract: _Deshare.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Deshare *DeshareFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DeshareRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Deshare.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeshareRoleAdminChanged)
				if err := _Deshare.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Deshare *DeshareFilterer) ParseRoleAdminChanged(log types.Log) (*DeshareRoleAdminChanged, error) {
	event := new(DeshareRoleAdminChanged)
	if err := _Deshare.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeshareRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Deshare contract.
type DeshareRoleGrantedIterator struct {
	Event *DeshareRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeshareRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeshareRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeshareRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeshareRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeshareRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeshareRoleGranted represents a RoleGranted event raised by the Deshare contract.
type DeshareRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deshare *DeshareFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DeshareRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Deshare.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DeshareRoleGrantedIterator{contract: _Deshare.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deshare *DeshareFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DeshareRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Deshare.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeshareRoleGranted)
				if err := _Deshare.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deshare *DeshareFilterer) ParseRoleGranted(log types.Log) (*DeshareRoleGranted, error) {
	event := new(DeshareRoleGranted)
	if err := _Deshare.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeshareRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Deshare contract.
type DeshareRoleRevokedIterator struct {
	Event *DeshareRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeshareRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeshareRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeshareRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeshareRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeshareRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeshareRoleRevoked represents a RoleRevoked event raised by the Deshare contract.
type DeshareRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deshare *DeshareFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DeshareRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Deshare.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DeshareRoleRevokedIterator{contract: _Deshare.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deshare *DeshareFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DeshareRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Deshare.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeshareRoleRevoked)
				if err := _Deshare.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deshare *DeshareFilterer) ParseRoleRevoked(log types.Log) (*DeshareRoleRevoked, error) {
	event := new(DeshareRoleRevoked)
	if err := _Deshare.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeshareTransactionBatchUploadedIterator is returned from FilterTransactionBatchUploaded and is used to iterate over the raw logs and unpacked data for TransactionBatchUploaded events raised by the Deshare contract.
type DeshareTransactionBatchUploadedIterator struct {
	Event *DeshareTransactionBatchUploaded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeshareTransactionBatchUploadedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeshareTransactionBatchUploaded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeshareTransactionBatchUploaded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeshareTransactionBatchUploadedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeshareTransactionBatchUploadedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeshareTransactionBatchUploaded represents a TransactionBatchUploaded event raised by the Deshare contract.
type DeshareTransactionBatchUploaded struct {
	BatchId        *big.Int
	DeviceId       [32]byte
	OrderCount     uint32
	TotalAmount    string
	DateComparable string
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransactionBatchUploaded is a free log retrieval operation binding the contract event 0xf59aead83aa674db1c77b1ce654d3b4f3716ebb8cd49f1753097d7ecd6bccc3f.
//
// Solidity: event TransactionBatchUploaded(uint256 indexed batchId, bytes32 indexed deviceId, uint32 orderCount, string totalAmount, string dateComparable, uint256 timestamp)
func (_Deshare *DeshareFilterer) FilterTransactionBatchUploaded(opts *bind.FilterOpts, batchId []*big.Int, deviceId [][32]byte) (*DeshareTransactionBatchUploadedIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}
	var deviceIdRule []interface{}
	for _, deviceIdItem := range deviceId {
		deviceIdRule = append(deviceIdRule, deviceIdItem)
	}

	logs, sub, err := _Deshare.contract.FilterLogs(opts, "TransactionBatchUploaded", batchIdRule, deviceIdRule)
	if err != nil {
		return nil, err
	}
	return &DeshareTransactionBatchUploadedIterator{contract: _Deshare.contract, event: "TransactionBatchUploaded", logs: logs, sub: sub}, nil
}

// WatchTransactionBatchUploaded is a free log subscription operation binding the contract event 0xf59aead83aa674db1c77b1ce654d3b4f3716ebb8cd49f1753097d7ecd6bccc3f.
//
// Solidity: event TransactionBatchUploaded(uint256 indexed batchId, bytes32 indexed deviceId, uint32 orderCount, string totalAmount, string dateComparable, uint256 timestamp)
func (_Deshare *DeshareFilterer) WatchTransactionBatchUploaded(opts *bind.WatchOpts, sink chan<- *DeshareTransactionBatchUploaded, batchId []*big.Int, deviceId [][32]byte) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}
	var deviceIdRule []interface{}
	for _, deviceIdItem := range deviceId {
		deviceIdRule = append(deviceIdRule, deviceIdItem)
	}

	logs, sub, err := _Deshare.contract.WatchLogs(opts, "TransactionBatchUploaded", batchIdRule, deviceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeshareTransactionBatchUploaded)
				if err := _Deshare.contract.UnpackLog(event, "TransactionBatchUploaded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransactionBatchUploaded is a log parse operation binding the contract event 0xf59aead83aa674db1c77b1ce654d3b4f3716ebb8cd49f1753097d7ecd6bccc3f.
//
// Solidity: event TransactionBatchUploaded(uint256 indexed batchId, bytes32 indexed deviceId, uint32 orderCount, string totalAmount, string dateComparable, uint256 timestamp)
func (_Deshare *DeshareFilterer) ParseTransactionBatchUploaded(log types.Log) (*DeshareTransactionBatchUploaded, error) {
	event := new(DeshareTransactionBatchUploaded)
	if err := _Deshare.contract.UnpackLog(event, "TransactionBatchUploaded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeshareTransactionDataUploadedIterator is returned from FilterTransactionDataUploaded and is used to iterate over the raw logs and unpacked data for TransactionDataUploaded events raised by the Deshare contract.
type DeshareTransactionDataUploadedIterator struct {
	Event *DeshareTransactionDataUploaded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeshareTransactionDataUploadedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeshareTransactionDataUploaded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeshareTransactionDataUploaded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeshareTransactionDataUploadedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeshareTransactionDataUploadedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeshareTransactionDataUploaded represents a TransactionDataUploaded event raised by the Deshare contract.
type DeshareTransactionDataUploaded struct {
	BatchId          *big.Int
	TransactionCount *big.Int
	Timestamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTransactionDataUploaded is a free log retrieval operation binding the contract event 0xd1398856cd83ffdea8eb5839ca5050509783beebd98e4a8b906127d25ff9137e.
//
// Solidity: event TransactionDataUploaded(uint256 indexed batchId, uint256 transactionCount, uint256 timestamp)
func (_Deshare *DeshareFilterer) FilterTransactionDataUploaded(opts *bind.FilterOpts, batchId []*big.Int) (*DeshareTransactionDataUploadedIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Deshare.contract.FilterLogs(opts, "TransactionDataUploaded", batchIdRule)
	if err != nil {
		return nil, err
	}
	return &DeshareTransactionDataUploadedIterator{contract: _Deshare.contract, event: "TransactionDataUploaded", logs: logs, sub: sub}, nil
}

// WatchTransactionDataUploaded is a free log subscription operation binding the contract event 0xd1398856cd83ffdea8eb5839ca5050509783beebd98e4a8b906127d25ff9137e.
//
// Solidity: event TransactionDataUploaded(uint256 indexed batchId, uint256 transactionCount, uint256 timestamp)
func (_Deshare *DeshareFilterer) WatchTransactionDataUploaded(opts *bind.WatchOpts, sink chan<- *DeshareTransactionDataUploaded, batchId []*big.Int) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Deshare.contract.WatchLogs(opts, "TransactionDataUploaded", batchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeshareTransactionDataUploaded)
				if err := _Deshare.contract.UnpackLog(event, "TransactionDataUploaded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransactionDataUploaded is a log parse operation binding the contract event 0xd1398856cd83ffdea8eb5839ca5050509783beebd98e4a8b906127d25ff9137e.
//
// Solidity: event TransactionDataUploaded(uint256 indexed batchId, uint256 transactionCount, uint256 timestamp)
func (_Deshare *DeshareFilterer) ParseTransactionDataUploaded(log types.Log) (*DeshareTransactionDataUploaded, error) {
	event := new(DeshareTransactionDataUploaded)
	if err := _Deshare.contract.UnpackLog(event, "TransactionDataUploaded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
