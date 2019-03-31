package lsp

import (
	"context"
	"golang.org/x/tools/internal/lsp/protocol"
	"golang.org/x/tools/internal/lsp/source"
	"golang.org/x/tools/internal/span"
)

func (s *server) symbols(ctx context.Context, query string) ([]protocol.SymbolInformation, error) {
	symbols := source.Symbols(ctx, s.view, s.workspace.Search, query, 100)

	var symbolInfos []protocol.SymbolInformation
	for _, symbol := range symbols {
		symbolInfos = append(symbolInfos, toProtocolSymbolInformation(symbol))
	}
	return symbolInfos, nil
}

func toProtocolSymbolInformation(symbol source.Symbol) protocol.SymbolInformation {
	return protocol.SymbolInformation{
		Name: symbol.Name,
		Kind: toProtocolSymbolKind(symbol.Kind),
		Location: toProtocolLocation(symbol.Span),
		ContainerName: symbol.Detail,
	}
}

func toProtocolLocation(spn span.Span) protocol.Location {
	return protocol.Location{URI: string(spn.URI()), Range: toProtocolRange(spn)}
}