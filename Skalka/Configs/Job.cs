using System.Collections.Generic;

namespace Skalka.Configs
{
    public record Job
    {
        public bool FlushOnError { get; set; }
        public List<string> Env { get; set; }
        public List<Step> Steps { get; set; }
    }
}