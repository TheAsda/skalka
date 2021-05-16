using Skalka.TaskRunner;

namespace Skalka.Queue
{
    public class QueueProcessorFactory : IQueueProcessorFactory
    {
        private readonly ITaskRunner _runner;

        public QueueProcessorFactory(ITaskRunner runner)
        {
            _runner = runner;
        }

        public IQueueProcessor Create()
        {
            return new QueueProcessor(_runner);
        }
    }
}