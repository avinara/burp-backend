package constants

const (
	InternalServerErrorCode           uint32 = 50001
	DatabaseInitErrorCode             uint32 = 50002
	LoadingConfigurationFileErrorCode uint32 = 50003

	QueryParamUnavailableErrorCode uint32 = 40001
	CookNotFoundErrorCode          uint32 = 40002
	InvalidRequestErrorCode        uint32 = 40003
	UserNotFoundErrorCode          uint32 = 40004
	DatabaseDeletionErrorCode      uint32 = 40006
	DatabaseUpdationErrorCode      uint32 = 40007
	DatabaseInsertionErrorCode     uint32 = 40008
	ScanningRowsErrorCode          uint32 = 40009
	DatabaseQueryErrorCode         uint32 = 40010
)
