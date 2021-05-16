using System.Collections.Generic;
using System.Threading.Tasks;
using Skalka.Configs;

namespace Skalka.Executor
{
    public interface IExecutor
    {
        void Fill(List<string> env, Dictionary<string, Job> jobs);
        void Start();
        void Wait();
    }
}