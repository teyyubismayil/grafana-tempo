package storage

import (
	"flag"

	"github.com/prometheus/common/model"

	"github.com/cortexproject/cortex/pkg/chunk"
	"github.com/cortexproject/cortex/pkg/chunk/storage"
)

// Config is the loki storage configuration
type Config struct {
	storage.Config    `yaml:",inline"`
	MaxChunkBatchSize int `yaml:"max_chunk_batch_size"`
}

// RegisterFlags adds the flags required to configure this flag set.
func (cfg *Config) RegisterFlags(f *flag.FlagSet) {
	cfg.Config.RegisterFlags(f)
	f.IntVar(&cfg.MaxChunkBatchSize, "max-chunk-batch-size", 50, "The maximun of chunks to fetch per batch.")
}

// Store is the Loki chunk store to retrieve and save chunks.
type Store interface {
	chunk.Store
	//	LazyQuery(ctx context.Context) (iter.EntryIterator, error)
}

type store struct {
	chunk.Store
	cfg Config
}

// NewStore creates a new Loki Store using configuration supplied.
func NewStore(cfg Config, storeCfg chunk.StoreConfig, schemaCfg chunk.SchemaConfig, limits storage.StoreLimits) (Store, error) {
	s, err := storage.NewStore(cfg.Config, storeCfg, schemaCfg, limits)
	if err != nil {
		return nil, err
	}
	return &store{
		Store: s,
		cfg:   cfg,
	}, nil
}

func filterChunksByTime(from, through model.Time, chunks []chunk.Chunk) []chunk.Chunk {
	filtered := make([]chunk.Chunk, 0, len(chunks))
	for _, chunk := range chunks {
		if chunk.Through < from || through < chunk.From {
			continue
		}
		filtered = append(filtered, chunk)
	}
	return filtered
}
