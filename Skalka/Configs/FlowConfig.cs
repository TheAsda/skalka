using System.Collections.Generic;

namespace Skalka.Configs
{
    public record FlowConfig
    {
        public int Version { get; set; }
        public List<string> Env { get; set; }
        public List<Requirement> Requirements { get; set; }
        public Dictionary<string, Job> Jobs { get; set; }
    }
}