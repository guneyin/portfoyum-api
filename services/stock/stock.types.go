package stock

type SyncSymbolRequestDTO struct {
	Code string `json:"code"`
	Data []Symbol
}

type SyncSymbolDetailRequestDTO struct {
	Code string `json:"code"`
	Data struct {
		HisseYuzeysel SymbolDetail
	}
}

type GetSymbolRequestDTO struct {
	Code string `json:"code"`
}

type GetSymbolResponseDTO struct {
	Symbols *[]Symbol `json:"symbols"`
}
