using System.Collections.Generic;
using System.Linq;
using Skalka.Configs;
using Skalka.TaskRunner;
using Task = System.Threading.Tasks.Task;

namespace Skalka.Queue
{
    public class QueueProcessor : IQueueProcessor
    {
        private Queue _queue;
        private readonly ITaskRunner _runner;
        private Task _task;

        public QueueProcessor(ITaskRunner runner)
        {
            _runner = runner;
        }

        public void Fill(string name, List<string> env, Job job)
        {
            _queue = new Queue(name);
            var allEnv = (env ?? new List<string>()).Concat(job.Env ?? new List<string>()).ToList();
            _queue.Fill(allEnv, job.Steps);
        }

        public void Start()
        {
            _task = Task.Run(() =>
            {
                while (!_queue.IsEmpty())
                {
                    var t = _queue.GetNext();
                    _runner.Run(t);
                }
            });
        }

        public async Task Wait()
        {
            await _task;
        }
    }
}