using System.Collections.Generic;
using System.Threading.Tasks;
using Skalka.Configs;

namespace Skalka.Queue
{
    public interface IQueueProcessor
    {
        void Fill(string name, List<string> env, Job job);
        void Start();
        Task Wait();
    }
}