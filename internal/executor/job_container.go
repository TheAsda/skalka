package executor

import (
	"github.com/TheAsda/skalka/internal/queue_processor"
	"github.com/TheAsda/skalka/internal/transaction_manager"
	"github.com/TheAsda/skalka/internal/variables_store"
	"github.com/TheAsda/skalka/pkg/config"
)

type JobContainer struct {
	job       config.Job
	processor *queue_processor.QueueProcessor
	tm        *transaction_manager.TransactionManager
	store     *variables_store.VariablesStore
}
