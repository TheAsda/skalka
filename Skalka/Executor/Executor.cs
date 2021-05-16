using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Skalka.Configs;
using Skalka.Queue;

namespace Skalka.Executor
{
    public class Executor : IExecutor
    {
        private readonly IQueueProcessorFactory _factory;
        private List<IQueueProcessor> _processors;

        public Executor(IQueueProcessorFactory factory)
        {
            _factory = factory;
            _processors = new List<IQueueProcessor>();
        }

        public void Fill(List<string> env, Dictionary<string, Job> jobs)
        {
            foreach (var (key, value) in jobs)
            {
                var processor = _factory.Create();
                processor.Fill(key, env, value);
                _processors.Add(processor);
            }
        }

        public void Start()
        {
            _processors.ForEach(processor => { processor.Start(); });
        }

        public void Wait()
        {
            Task.WaitAll(_processors.Select(processor => processor.Wait()).ToArray());
        }
    }
}