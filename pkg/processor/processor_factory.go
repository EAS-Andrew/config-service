package processor

type ProcessorFactory struct{}

func NewProcessorFactory() *ProcessorFactory {
	return &ProcessorFactory{}
}

func (pf *ProcessorFactory) GetProcessors() []ALFolderProcessor {
	processors := []ALFolderProcessor{
		new(NexusRMProcessor),
	}
	return processors
}
