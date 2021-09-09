package ngapConvert

import (
	"github.com/free5gc/aper"
	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/openapi/models"
)

func UpIntegrToModels(
	ngapIntegrProtectIndi ngapType.IntegrityProtectionIndication) models.UpIntegrity {
	switch ngapIntegrProtectIndi.Value {
	case ngapType.IntegrityProtectionIndicationPresentRequired:
		return models.UpIntegrity_REQUIRED
	case ngapType.IntegrityProtectionIndicationPresentPreferred:
		return models.UpIntegrity_PREFERRED
	}
	return models.UpIntegrity_NOT_NEEDED
}

func UpIntegrToNgap(
	modelsUpIntegr models.UpIntegrity) aper.Enumerated {
	switch modelsUpIntegr {
	case models.UpIntegrity_REQUIRED:
		return ngapType.IntegrityProtectionIndicationPresentRequired
	case models.UpIntegrity_PREFERRED:
		return ngapType.IntegrityProtectionIndicationPresentPreferred
	}
	return ngapType.IntegrityProtectionIndicationPresentNotNeeded
}

func UpConfidToModels(
	ngapConfidProtectIndi ngapType.ConfidentialityProtectionIndication) models.UpConfidentiality {
	switch ngapConfidProtectIndi.Value {
	case ngapType.ConfidentialityProtectionIndicationPresentRequired:
		return models.UpConfidentiality_REQUIRED
	case ngapType.ConfidentialityProtectionIndicationPresentPreferred:
		return models.UpConfidentiality_PREFERRED
	}
	return models.UpConfidentiality_NOT_NEEDED
}

func UpConfidToNgap(
	modelsUpConfid models.UpConfidentiality) aper.Enumerated {
	switch modelsUpConfid {
	case models.UpConfidentiality_REQUIRED:
		return ngapType.ConfidentialityProtectionIndicationPresentRequired
	case models.UpConfidentiality_PREFERRED:
		return ngapType.ConfidentialityProtectionIndicationPresentPreferred
	}
	return ngapType.ConfidentialityProtectionIndicationPresentNotNeeded
}

func UpIntegrMaxProtectDataRateToModels(
	ngapMaxIntegrProtectDataRate ngapType.MaximumIntegrityProtectedDataRate) models.MaxIntegrityProtectedDataRate {
	switch ngapMaxIntegrProtectDataRate.Value {
	case ngapType.MaximumIntegrityProtectedDataRatePresentMaximumUERate:
		return models.MaxIntegrityProtectedDataRate_MAX_UE_RATE
	}
	return models.MaxIntegrityProtectedDataRate__64_KBPS
}

func UpIntegrMaxProtectDataRateToNgap(
	modelsMaxIntegrProtectDataRate models.MaxIntegrityProtectedDataRate) aper.Enumerated {
	switch modelsMaxIntegrProtectDataRate {
	case models.MaxIntegrityProtectedDataRate_MAX_UE_RATE:
		return ngapType.MaximumIntegrityProtectedDataRatePresentMaximumUERate
	}
	return ngapType.MaximumIntegrityProtectedDataRatePresentBitrate64kbs
}
